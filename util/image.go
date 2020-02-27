package util

import (
	"io"
	"image"
	"os"
)

// Decode reads an image from r.
func DecodeImageConfig(r io.Reader) (image.Config, error) {
	cfg, _, err := image.DecodeConfig(r)
	if err != nil {
		return image.Config{}, err
	}
	return cfg,nil
}

// Open loads an image from file
func ImageConfig(filename string) (image.Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return image.Config{}, err
	}
	defer file.Close()
	return DecodeImageConfig(file)
}

