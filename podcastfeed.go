package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName       xml.Name `xml:"channel"`
	Title         string   `xml:"title"`
	Description   string   `xml:"description"`
	LastBuildDate string   `xml:"lastBuildDate"`
	Items         []Item   `xml:"item"`
}

type Item struct {
	Id        string    `xml:"guid"`
	Title     string    `xml:"title"`
	Published string    `xml:"pubDate"`
	Enclosure Enclosure `xml:"enclosure"`
}

type Enclosure struct {
	Url    string `xml:"url,attr"`
	Length int    `xml:"length,attr"`
}

func (i Item) String() string {
	return fmt.Sprintf("Id: %v, Title: %v, Url: %v", i.Id, i.Title, i.Enclosure.Url)
}

func LoadFeed(name string) (*Channel, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	return toFeed(data)
}

func toFeed(data []byte) (*Channel, error) {
	var r Rss
	err := xml.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}

	c := r.Channel

	fmt.Printf("Title: %v, LastBuildDate: %v\n", c.Title, c.LastBuildDate)
	fmt.Printf("Items: %v\n", len(c.Items))
	c.PrintLatestItems()

	return &c, nil
}

func (feed *Channel) GetItemAtIndex(index int) (*Item, error) {
	if len(feed.Items) <= index {
		return nil, errors.New("index out of bounds")
	}

	item := feed.Items[index]
	return &item, nil
}

func (feed *Channel) PrintLatestItems() {
	latestItems := feed.Items[0:5]
	fmt.Println("Latest Items...")
	for index, item := range latestItems {
		fmt.Printf("%v - %v\n", index, item)
	}
}

func (item *Item) GetFileName(channel *Channel) string {
	name := strings.TrimPrefix(item.Title, channel.Title)
	name = strings.TrimSpace(name) + ".mp3"
	fmt.Println(name)
	return name
}
