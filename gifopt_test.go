package gifopt

import (
	"fmt"
	"image"
	_ "image/gif"
	"os"
	"testing"
)

func TestResize(t *testing.T) {

	tests := []struct {
		O, N string
		W    int
	}{
		{"blob.gif", "blob_resized_250.gif", 250},
		{"blob.gif", "blob_resized_200.gif", 200},
	}

	for _, te := range tests {
		err := Resize(te.O, te.N, te.W)
		if err != nil {
			t.Errorf("image was not resized properly: %v", err)
		}

		newWidth, newHeight := getImageDimension(te.N)
		if newWidth != te.W || newHeight != te.W {
			t.Errorf("new image dimensions were incorrect : %v, %v", newWidth, newHeight)
		}
	}
}

func getImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return image.Width, image.Height
}
