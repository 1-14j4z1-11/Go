package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var formatFlag = flag.String("format", "jpeg", "todo")

func main() {
	flag.Parse()

	if err := convert(os.Stdin, os.Stdout, *formatFlag); err != nil {
		fmt.Fprintf(os.Stderr, "image: %v\n", err)
		os.Exit(1)
	}
}

func convert(in io.Reader, out io.Writer, format string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stderr, "Input format = ", kind)

	switch format {
	case "jpeg":
		fallthrough
	case "jpg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, &gif.Options{})
	default:
		return fmt.Errorf("Unsupported format : %s", format)
	}
}
