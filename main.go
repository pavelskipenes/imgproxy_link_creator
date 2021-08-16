package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/alexflint/go-arg"
)

// usage: imgproxy_converter "http://10.10.10.10/site/camera01.jpg"
// output: http://imgproxyserver:imgproxyport/randomStuff/site/camera01.jpg

func main() {
	var args struct {
		Path      string `arg:"required"` // Source path
		Server    string `arg:"env"`      // imgproxy server path
		Key       string `arg:"env"`
		Salt      string `arg:"env"`
		Resize    string `default:"fill"`
		Width     int    `default:"300"`
		Height    int    `default:"300"`
		Gravity   string `default:"no"`
		Enlarge   int    `default:"1"`
		Extension string `arg:"env"`
	}
	var err error
	err = arg.Parse(&args)
	if err != nil {
		log.Fatal("Server: ", args.Server)
		log.Fatal("Path: ", args.Path)
		log.Fatal("Key:", args.Key)
		log.Fatal("Salt:", args.Salt)
		log.Fatal("Resize:", args.Resize)
		log.Fatal("Width:", args.Width)
		log.Fatal("Height:", args.Height)
		log.Fatal("Gravity:", args.Gravity)
		log.Fatal("Enlarge:", args.Enlarge)
		log.Fatal("Extension:", args.Extension)
		log.Fatal("Error while parsing arguments")
	}

	var keyBin, saltBin []byte

	if keyBin, err = hex.DecodeString(args.Key); err != nil {
		log.Fatal("Key expected to be hex-encoded string")
	}

	if saltBin, err = hex.DecodeString(args.Salt); err != nil {
		log.Fatal("Salt expected to be hex-encoded string")
	}

	encodedURL := base64.RawURLEncoding.EncodeToString([]byte(args.Path))
	path := fmt.Sprintf("/%s/%d/%d/%s/%d/%s.%s", args.Resize, args.Width, args.Height, args.Gravity, args.Enlarge, encodedURL, args.Extension)

	mac := hmac.New(sha256.New, keyBin)
	mac.Write(saltBin)
	mac.Write([]byte(path))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	fmt.Printf("%s/%s%s\n", args.Server, signature, path)
}
