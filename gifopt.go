// Page gifopt provides optimization tools for working with gifs in go
package gifopt

import (
	"bufio"
	"fmt"
	"image/gif"
	"image/jpeg"
	"io"
	"os"
	"os/exec"
)

// Allow to convert gif images to jpeg
// using the first frame
func ToJpg(oldf, newf string) error {
	ob, err := os.Open(oldf)
	defer ob.Close()
	if err != nil {
		return err
	}

	oi, err := gif.Decode(ob)
	if err != nil {
		return err
	}

	outfile, err := os.Create(newf)
	defer outfile.Close()
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(outfile)

	jpeg.Encode(writer, oi, &jpeg.Options{90})

	return nil
}

// Resize a gif using gifsicle
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
	defer outfile.Close()
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(outfile)

	err = cmd.Start()
	if err != nil {
		return err
	}

	go io.Copy(writer, stdoutPipe)
	cmd.Wait()

	return nil
}
