package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {
	token, err := getEnviron("SLACK_API_KEY")
	if err != nil {
		log.Fatal(err.Error())
	}
	ch, err := getEnviron("CHANNEL_ID")
	if err != nil {
		log.Fatal(err.Error())
	}
	chgr := []string{ch}

	if err != nil {
		log.Fatal(err.Error())
		return
	}
	files := []string{"hello.txt"}
	api := slack.New(token)

	for _, file := range files {
		params := slack.FileUploadParameters{
			Channels: chgr,
			File:     file,
		}

		f, err := api.UploadFile(params)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Printf("File: %s, Sise: %d", f.Name, f.Size)
	}

}

func getEnviron(key string) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return "", nil
	}

	value := os.Getenv(key)

	return value, nil
}
