package client

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func checkFile(urlFile string, shaCheck string) (err error) {

	response, err := http.Get(urlFile)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	fmt.Println("+++CheckFile+++")
	fmt.Println("Status:", response.Status)
	fmt.Printf("TLS:%+v \n", response.TLS)
	//	fmt.Println("TLS:", response.TLS)
	fmt.Printf("Header:%+v \n", response.Header)
	fmt.Println("ContentLength:", response.ContentLength)

	hash := sha256.New()
	if _, err := io.Copy(hash, response.Body); err != nil {
		return err
	}
	calculatedHash := hash.Sum(nil)
	calculatedHashString := hex.EncodeToString(calculatedHash)

	if calculatedHashString != shaCheck {
		fmt.Printf("конфиг хеш - %v \n факт хеш - %v \n", shaCheck, calculatedHashString)
		return errors.New("Хеш файла не совпадает с ожидаемым значением")
	}

	return
}
