package main

import (
	"crypto/tls"
	_ "embed"
	"fmt"
	"log"
	"net/http"

	handler "github.com/20ritiksingh/sharego/handlers"
	"github.com/skip2/go-qrcode"
)

//go:embed cert/cert.pem
var certPEM []byte

//go:embed cert/key.pem
var keyPEM []byte

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /upload", handler.UploadHandler)
	mux.HandleFunc("GET /file/", handler.FileServerHandler)
	mux.HandleFunc("GET /list/", handler.FileListHandler)
	mux.HandleFunc("GET /", handler.HomeHandler)
	cert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		println("failed to load tls certificate: ", err)
		fmt.Scanln("enter any key to exit!!")
		return
	}
	server := http.Server{
		Addr:    ":8088",
		Handler: mux,
		TLSConfig: &tls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
		},
	}
	addrs, err := getInterfaceAddrs()
	if err != nil {
		log.Println("error getting network interface: ", err)
		fmt.Scanln("enter any key to exit!!")
		return
	}
	for i, addr := range addrs {

		qr, err := qrcode.New("https://"+addr+":8088", qrcode.Medium)
		if err != nil {
			log.Println("failed to generate qr for: ", addr)
			continue
		}
		fmt.Printf("Interface %d: https://%s:8088", i, addr)
		fmt.Print(qr.ToString(true))
	}
	// log.Printf("app is up n running")
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		log.Println(err)
		fmt.Scanln("enter any key to exit!!")
		return
	}

}
