package main

import (
    "flag"
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "path/filepath"
    "context"
    "io"
    "time"

    "minecraft-reverse-copy/config"
    "minecraft-reverse-copy/zip"
    "minecraft-reverse-copy/tmpl"
    "minecraft-reverse-copy/common/types"

    "github.com/k0tletka/gowebdav"
)

const (
    // Exit codes
    GeneralIssue = 1
    ConfigurationReadError = 2
    ZipError = 3
    WebdavError = 4
)

// Flag initialization
var (
    optConfig = flag.String("c", "", "Configuration file to load")
)

func main() {
    var (
        conf        *config.Configuration
        ctx         context.Context
        logFd       io.WriteCloser
        davClient   *gowebdav.Client
        zipFile     *os.File
        archiveName string
        err         error
    )

    flag.Parse()

    if *optConfig == "" {
        // Load default configuration
        conf = config.GetDefaultConfiguration()
    } else {
        if conf, err = config.ReadConfiguration(*optConfig); err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(ConfigurationReadError)
        }
    }

    // Setup context
    ctx, _ = signal.NotifyContext(context.Background(),
        syscall.SIGINT,
        syscall.SIGKILL,
        syscall.SIGTERM,
    )


    if conf.LogFile != "" {
        logFd, err = os.OpenFile(conf.LogFile, os.O_WRONLY | os.O_APPEND, 0755)

        if err != nil {
            fmt.Fprintln(os.Stdout, "Warning: Can't open file for log, using stdout instead")
            logFd = os.Stdout
        }
    } else {
        logFd = os.Stdout
    }

    defer logFd.Close()

    // Signal close handler
    go func() {
        <-ctx.Done()
        fmt.Fprintln(logFd, "Got signal, exiting...")
        os.Exit(0)
    }()

    // Initialize webdav connection
    if conf.Webdav.UseAuth {
        davClient = gowebdav.NewClientBasicAuth(
            conf.Webdav.WebdavHost,
            conf.Webdav.WebdavAuthConfiguration.Username,
            conf.Webdav.WebdavAuthConfiguration.Password,
        )
    } else {
        davClient = gowebdav.NewClient(conf.Webdav.WebdavHost, "", "")
    }

    if err = davClient.Connect(); err != nil {
        fmt.Fprintf(logFd, "Error: can't connect to webdav: %s\n", err)
        os.Exit(WebdavError)
    }

    // Generate archive name
    archiveTemplateData := &types.ArchiveTemplateData{
        Datetime: time.Now().Format(conf.TimeTemplate),
        WorldSaveName: filepath.Base(conf.WorldPath),
    }

    if archiveName, err = tmpl.GenerateArchiveName(conf.ArchiveNameTemplate, archiveTemplateData); err != nil {
        fmt.Fprintf(logFd, "Error: can't generate archive name: %s\n", err)
        os.Exit(GeneralIssue)
    }

    // Start zip process
    zipProcess := &zip.ZipProcess{
        WorldPath: conf.WorldPath,
    }

    if zipFile, err = zipProcess.RunZip(ctx); err != nil {
        fmt.Fprintf(logFd, "Error: can't start zip process: %s\n", err)
        os.Exit(ZipError)
    }

    defer func() {
        zipFile.Close()
        os.Remove("/tmp/" + zipFile.Name())
    }()

    // Upload file
    if err = davClient.WriteStream(filepath.Join(conf.Webdav.WebdavSavePath, archiveName), zipFile, 0); err != nil {
        fmt.Fprintf(logFd, "Error: can't send archive to server: %s\n", err)
        os.Exit(WebdavError)
    }
}
