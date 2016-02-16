package main

import (
	"bytes"
	"log"
	"os/exec"
)

type MPlayer struct {
}

func NewMPlayer() *MPlayer {
	return &MPlayer{}
}

func (m *MPlayer) Play(data *[]byte) error {

	cmd := exec.Command("mplayer", "-")

	reader := bytes.NewReader(*data)

	cmd.Stdin = reader

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	log.Printf("Mplayer finished:\n %s\n", out.String())

	return err
}
