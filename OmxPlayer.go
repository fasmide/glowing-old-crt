package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"syscall"
)

type OmxPlayer struct {
	pipePath string
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewOmxPlayer() *OmxPlayer {
	return &OmxPlayer{pipePath: PreparePipe()}
}

func (m *OmxPlayer) Play(data *[]byte) error {

	cmd := exec.Command("omxplayer", "--win", "-107 0 747 480", "-b", m.pipePath)

	go ioutil.WriteFile(m.pipePath, *data, 0666)

	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()

	log.Printf("OmxPlayer finished:\n %s\n", out.String())

	return err
}

func PreparePipe() string {

	path := "/tmp/" + RandStringBytes(10)

	pipeExists := false
	fileInfo, err := os.Stat(path)

	if err == nil {
		if (fileInfo.Mode() & os.ModeNamedPipe) > 0 {
			pipeExists = true
		} else {
			log.Printf("%d != %d\n", os.ModeNamedPipe, fileInfo.Mode())
			panic(path + " exists, but it's not a named pipe (FIFO)")
		}
	}

	// Try to create pipe if needed
	if !pipeExists {
		err := syscall.Mkfifo(path, 0666)
		if err != nil {
			panic(err.Error())
		}
	}
	return path
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
