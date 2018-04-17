package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"fmt"
	"bytes"
	"image/jpeg"
	"strconv"
)

var secretGlobal string

func GenerateKey(w http.ResponseWriter, r *http.Request) {

	totp := GenerateQr("issuer", "account", 30)

	secretGlobal = totp.Secret()

	fmt.Print(totp.Secret())

	img, _ := totp.Image(100,100)

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func ValidateUserKey(w http.ResponseWriter, r *http.Request) {

	key := r.URL.Query()["key"]

	fmt.Println("key", key)
	fmt.Println("secretGlobal", secretGlobal)

	result := Validate(key[0], secretGlobal)
	fmt.Println(result)

}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/generatekey", GenerateKey)
	r.HandleFunc("/validate", ValidateUserKey)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
