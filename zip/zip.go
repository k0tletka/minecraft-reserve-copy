package zip

import (
    "context"
    "os/exec"
    "io"
    "text/template"
    "bytes"
    "path/filepath"
)

type ZipProcess struct {
    WorldPath           string
    ArchiveNameTemplate string
    TimeTemplate        string
    ResultOutputWriter  io.Writer

    cmdZipProcess *exec.Cmd
}

func (z *ZipProcess) RunZip(ctx context.Context) (io.ReadCloser, error) {
    // Create template and generate archive name
    archiveNameTemplate, err := template.New("archive-name-template").Parse(ArchiveNameTemplate)
    if err != nil {
        return nil, err
    }

    archiveTemplateData := &types.archiveTemplateData{
        Datetime: time.Now().Format(z.TimeTemplate),
        WorldSaveName: filepath.Base(z.WorldPath),
    }

    archiveNameBuffer := &bytes.Buffer{}
    if err := archiveNameTemplate.Execute(archiveNameBuffer, archiveTemplateData); err != nil {
        return nil, err
    }

    // Create zip process
    z.cmdZipProcess = exec.CommandContext(
        ctx,
        "zip",
        "-r",
        "-", // Write archive to process stdout
        WorldPath,
    )

    // Get stdout pipe reader and start process
    stdoutReader, err := z.cmdZipProcess.StdoutPipe()
    if err != nil {
        return nil, err
    }

    return stdoutReader, z.cmdZipProcess.Start()
}
