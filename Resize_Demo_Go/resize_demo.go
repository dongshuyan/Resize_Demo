package main

import (
	"os"
	"fmt"
	"log"
	"image/jpeg"
	"github.com/nfnt/resize"
)
func main() {
	filename := "resize_demo.JPEG"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Load picture failed", filename)
		log.Fatal(err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println("Decode picture failed", filename)
		log.Fatal(err)
	}
	file.Close()

	m := resize.Resize(256, 256, img, resize.Lanczos3)
	
	out, err := os.Create("resize_demo_resized.JPEG")
	if err != nil {
		fmt.Println("Write picture failed", filename)
		log.Fatal(err)
	}
	defer out.Close()

	jpeg.Encode(out, m, nil)
}