package gifopt

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Resize(oldf, newf string, width int) error {

	args := []string{
		oldf,
		"-O2",
		"--resize",
		fmt.Sprintf("%vx%v", width, width),
	}

	gifsicleCmd, err := exec.LookPath("gifsicle")
	if err != nil {
		return err
	}
	cmd := exec.Command(gifsicleCmd, args...)
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	outfile, err := os.Create(newf)
	if err != nil {
		return err
	}
	defer outfile.Close()

	writer := bufio.NewWriter(outfile)

	err = cmd.Start()
	if err != nil {
		return err
	}

	go io.Copy(writer, stdoutPipe)
	cmd.Wait()

	return nil
}
