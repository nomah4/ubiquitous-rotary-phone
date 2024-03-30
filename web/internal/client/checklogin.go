package client

import (
	"fmt"
	"net/http"
)

func checkLogin(urlLogin string) (err error) {

	response, err := http.Get(urlLogin)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	fmt.Println("+++CheckLogin+++")
	fmt.Println("Status:", response.Status)
	fmt.Printf("TLS:%+v \n", response.TLS)
	//	fmt.Println("TLS:", response.TLS)
	fmt.Printf("Header:%+v \n", response.Header)
	fmt.Println("ContentLength:", response.ContentLength)

	return
}
