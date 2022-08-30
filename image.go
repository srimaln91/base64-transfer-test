package encode

import (
	"bytes"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"strings"

	"encoding/base64"

	"github.com/google/uuid"
)

func TransferImage(b64data string) error {
	// ctx := context.Background()

	// b64data := image.Images[strings.IndexByte(image.Images, ',')+1:]
	//decode encoded image url
	fileBytes, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		return err
	}

	// check content type, detectcontenttype only needs the first 512 bytes
	ImageType := http.DetectContentType(fileBytes)

	r := bytes.NewReader(fileBytes)

	im, err := jpeg.Decode(r)
	if err != nil {
		return err
	}

	fileName := strings.ToLower(ImageType) + "_" + uuid.New().String() + ".jpg"
	ff, err := os.OpenFile("./"+fileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}

	return jpeg.Encode(ff, im, nil)
}

func TransferImageOptimized(b64data []byte) error {

	reader := base64.NewDecoder(base64.StdEncoding, bytes.NewReader(b64data))

	fileName := strings.ToLower("image/jpeg") + "_" + uuid.New().String() + ".jpg"
	ff, err := os.OpenFile("./"+fileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	_, err = io.Copy(ff, reader)
	if err != nil {
		return err
	}

	return nil
}
