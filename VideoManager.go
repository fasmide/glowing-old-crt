package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type VideoManager struct {
	AvailableVideos []VideoDescription
}

type VideoDescription struct {
	Url   string
	Title string
}

func NewVideoManager() VideoManager {

	rand.Seed(time.Now().UTC().UnixNano())

	m := VideoManager{AvailableVideos: make([]VideoDescription, 0)}

	doc, err := goquery.NewDocument("https://www.youtube.com/channel/UC-9-kyTW8ZkZNDHQJ6FgpwQ")

	if err != nil {
		fmt.Println(err)
	}

	doc.Find(".yt-uix-tile-link").Each(func(i int, s *goquery.Selection) {

		link, exists := s.Attr("href")

		if !exists {
			return
		}

		m.AvailableVideos = append(m.AvailableVideos, VideoDescription{Url: "https://youtube.com" + link, Title: s.Text()})

	})

	return m
}

func (m *VideoManager) RandomVideo() VideoDescription {
	var selected VideoDescription
	for _, i := range rand.Perm(len(m.AvailableVideos)) {
		selected = m.AvailableVideos[i]
	}

	return selected

}
