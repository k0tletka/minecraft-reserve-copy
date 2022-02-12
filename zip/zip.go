package zip

import (
    "context"
    "os/exec"
    "io"
)

type ZipProcess struct {
    WorldPath           string

    cmdZipProcess *exec.Cmd
}

func (z *ZipProcess) RunZip(ctx context.Context) (io.ReadCloser, error) {
    // Create zip process
    z.cmdZipProcess = exec.CommandContext(
        ctx,
        "zip",
        "-r",
        "-", // Write archive to process stdout
        z.WorldPath,
    )

    // Get stdout pipe reader and start process
    stdoutReader, err := z.cmdZipProcess.StdoutPipe()
    if err != nil {
        return nil, err
    }

    return stdoutReader, z.cmdZipProcess.Start()
}
