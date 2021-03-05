package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math"
	"os"
	"strings"

	"github.com/alecthomas/kingpin"
)

var (
	app = kingpin.New("ficonv", "A tool to parse files as an image and vice versa")

	appIn  = app.Arg("in", "The file you want to convert").Required().String()
	appOut = app.Arg("out", "The output file").Required().String()

	appTrim    = app.Flag("trim", "Trims trailing NULL bytes. Useful for txt files, can damage other types. Only used together with \"reverse\"").Default("false").Short('t').Bool()
	appReverse = app.Flag("reverse", "Reverses an image to a file. By default the given file will be converted to an image.").Default("false").Short('r').Bool()
)

func main() {

	// Parse args
	kingpin.MustParse(app.Parse(os.Args[1:]))

	// Convert accordingly
	if *appReverse {
		revertFromImage(*appIn, *appOut)
	} else {
		convertToImage(*appIn, *appOut)
	}
}

func revertFromImage(filepath, output string) {

	// Read the file
	f, err := os.Open(filepath)
	tmpImg, _, err := image.Decode(f)
	img := tmpImg.(*image.NRGBA)
	defer f.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	// Get Image width and height
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	// Iterate each pixel and build the file
	fileBytes := make([]byte, width*4*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rgba := img.NRGBAAt(x, y)

			fileBytes[y*width*4+x*4] = rgba.R
			fileBytes[y*width*4+x*4+1] = rgba.G
			fileBytes[y*width*4+x*4+2] = rgba.B
			fileBytes[y*width*4+x*4+3] = rgba.A

			// fmt.Println("r:", byte(r), "g:", byte(g), "b:", byte(b), "a:", byte(a))
			// fmt.Println("Written to ", (y*width + x), "-", (y*width + x + 3))
		}
	}

	// Remove trailing NULL bytes
	if *appReverse {
		fileBytes = []byte(strings.TrimRight(string(fileBytes), "\x00"))
	}

	// Create file and write bytes
	err = ioutil.WriteFile(output, fileBytes, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func convertToImage(filepath string, output string) {

	// Read the file
	file, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Println(err.Error())
	}

	// Get the needed size
	var bytes int
	bytes = int(len(file))
	_ = bytes

	// Set Image Size
	imageSqrt := math.Sqrt(float64(bytes))

	width := int(imageSqrt)
	height := int(imageSqrt) + 1

	// Adjust width and height so the width is dividable by 4
	dif := differenceToBiggestDivisor(width, 4)
	width -= dif
	height += dif + 1

	// Create the base image
	img := image.NewNRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width / 4, height}})

	// Set color for each pixel.
	for y := 0; y < height; y++ {
		for x := 0; x < width; x += 4 {
			startByte := y*width + x

			// Skip empty bytes
			if startByte >= bytes {
				continue
			}

			// The bytes represented by R,G,B,A
			byte1 := file[startByte]
			byte2 := byte(0)
			byte3 := byte(0)
			byte4 := byte(0)

			// For bytes 2-4 check if there is any data left -> add if yes
			if startByte+1 < bytes {
				byte2 = file[startByte+1]
			}
			if startByte+2 < bytes {
				byte3 = file[startByte+2]
			}
			if startByte+3 < bytes {
				byte4 = file[startByte+3]
			}

			// Set pixel into img
			pixel := color.NRGBA{R: byte1, G: byte2, B: byte3, A: byte4}
			img.Set(x/4, y, pixel)
		}
	}

	// Encode as PNG.
	if !strings.HasSuffix(output, ".png") {
		output += ".png"
	}
	f, _ := os.Create(output)
	png.Encode(f, img)
}

func differenceToBiggestDivisor(num int, divisor int) (rest int) {
	for num%divisor != 0 {
		rest++
		num--
	}
	return rest
}
