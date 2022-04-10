package zip

import (
    "context"
    "os"
    "os/exec"
)

type ZipProcess struct {
    WorldPath           string

    cmdZipProcess *exec.Cmd
}

func (z *ZipProcess) RunZip(ctx context.Context) (*os.File, error) {
    // Create zip process
    z.cmdZipProcess = exec.CommandContext(
        ctx,
        "zip",
        "-r",
        "-", // Write archive to process stdout
        z.WorldPath,
    )

    // Create temporary file for archive storing
    archiveTemp, err := os.CreateTemp("/tmp", "*.zip")
    if err != nil {
        return nil, err
    }

    z.cmdZipProcess.Stdout = archiveTemp
    return archiveTemp, z.cmdZipProcess.Run()
}
