package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Starting MegaMaid...")

	cmd := os.Args[1]
	switch cmd {
	case "refresh":
		refreshFeed()
	case "view":
		loadFeed()
	case "download":
		indexInput := os.Args[2]
		itemIndex, _ := strconv.Atoi(indexInput)
		downloadEpisode(itemIndex)
	default:
		fmt.Printf("I do not know how to '%s' yet\n", cmd)
	}

	fmt.Println("Shutting down MegaMaid...")
}

func downloadEpisode(itemIndex int) {
	itemsToDownload := 1
	feed, _ := LoadFeed("the-anjunadeep-edition.xml")

	for i := itemIndex; i < (itemIndex + itemsToDownload); i++ {
		item, _ := feed.GetItemAtIndex(i)
		name := item.GetFileName(feed)
		url := item.Enclosure.Url
		DownloadFile(url, name)
	}
}

func loadFeed() {
	_, err := LoadFeed("the-anjunadeep-edition.xml")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func refreshFeed() {
	url := "http://static.anjunadeep.com/edition/podcast.xml"
	name := "the-anjunadeep-edition.xml"
	DownloadFile(url, name)
}
