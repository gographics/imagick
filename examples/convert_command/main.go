package main

import (
	"fmt"

	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	ret, err := imagick.ConvertImageCommand([]string{
		"convert", "logo:", "-resize", "100x100", "/tmp/out.png",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Metadata:\n%s\n", ret.Meta)
}
