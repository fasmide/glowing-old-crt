package main

import (
	"bytes"
	"http/client"
	"log"
	"os/exec"
)

type Video struct {
	Url     string
	DataUrl string
	Data    []byte
	Title   string
}

func NewVideo(youtubeUrl string) *Video {
	return &Video{Url: youtubeUrl}
}

func (v *Video) Download() error {

	err := v.GetDataUrl()

	if err != nil {
		return err
	}

	return nil
}

func (v *Video) GetDataUrl() error {
	cmd := exec.Command("youtube-dl", "-g", "-f", "135", v.Url)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		return err
	}
	v.DataUrl = out.String()
	return nil
}
