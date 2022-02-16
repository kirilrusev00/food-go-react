package qrdecoderserver

import (
	"image"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

// Remove the unnecessary from the received bytes from the buffer
func clen(n []byte) int {
	for i := 0; i < len(n); i++ {
		if n[i] < 32 {
			return i
		}
	}
	return len(n)
}

/*
	decodeQrCode returns the decoded QR code as string or "Could not decode QR code" if there was an error
*/
func decodeQrCode(localQrPath string) string {
	result, err := getQrCodeContent(localQrPath)

	if err != nil {
		return "Could not decode QR code"
	}

	return result.String()
}

/*
	getQrCodeContent returns the decoded QR code as gozxing.Result or error
*/
func getQrCodeContent(localQrPath string) (*gozxing.Result, error) {
	file, err := os.Open(localQrPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	src := gozxing.NewLuminanceSourceFromImage(img)

	bmp, err := gozxing.NewBinaryBitmap(gozxing.NewGlobalHistgramBinarizer(src))
	if err != nil {
		return nil, err
	}

	qrReader := qrcode.NewQRCodeReader()
	return qrReader.Decode(bmp, nil)
}
