package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shanghuiyang/image-recognizer/oauth"
	"github.com/shanghuiyang/image-recognizer/recognizer"
)

const (
	// replace your_app_key with your app key
	appKey = "your_app_key"
	// replace your_secret_key with your secret key
	secretKey = "your_secret_key"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("error: invalid args")
		fmt.Println("usage: image-recognizer test.jpg")
		os.Exit(1)
	}
	imgf := os.Args[1]

	auth := oauth.New(appKey, secretKey, oauth.NewCacheMan())
	r := recognizer.New(auth)
	text, err := r.Recognize(imgf)
	if err != nil {
		log.Printf("failed to recognize the image, error: %v", err)
		os.Exit(1)
	}
	fmt.Println(text)
}
