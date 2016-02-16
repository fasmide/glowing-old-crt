package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type Video struct {
	Url     string
	DataUrl string
	Data    []byte
	Title   string
}

type DownloadProgress struct {
	Bytes uint64
	Total uint64
}

func (wc *DownloadProgress) Write(p []byte) (int, error) {
	n := len(p)
	wc.Bytes += uint64(n)
	log.Printf("Read %d bytes for a total of %d\n", wc.Bytes, wc.Total)
	return n, nil
}

func NewVideo(youtubeUrl string) *Video {
	return &Video{Url: youtubeUrl}
}

func (v *Video) Download() error {

	err := v.GetDataUrl()

	if err != nil {
		return err
	}

	log.Printf("Today we will be downloading %+v", v.DataUrl)

	resp, err := http.Get(v.DataUrl)

	if err != nil {
		// handle error
		return err
	}

	defer resp.Body.Close()
	counter := &DownloadProgress{Total: uint64(resp.ContentLength)}

	// Instrument with our counter.
	src := io.TeeReader(resp.Body, counter)

	v.Data, err = ioutil.ReadAll(src)

	if err != nil {
		return err
	}

	return nil
}

func (v *Video) GetDataUrl() error {
	//asking youtube-dl for links is somewhat ugly but i cannot find any api returning
	//video urls in youtubes official docs
	cmd := exec.Command("youtube-dl", "-g", "-f", "135", v.Url)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		return err
	}
	v.DataUrl = strings.Trim(out.String(), "\n")
	return nil
}
