package main

import (
	"github.com/ataberkcanitez/araqr/cmd"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	if err := cmd.RunRootCmd(); err != nil {
		panic(err)
	}
}

func main2() {
	// Create the barcode
	qrCode, _ := qr.Encode("https://dev.randevumu.com/5-ataberk-hair-shop", qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
}
