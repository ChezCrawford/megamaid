package main

import "fmt"

func main() {
	fmt.Println("Starting MegaMaid...")
	refreshFeed()
	loadFeed()
	// downloadEpisode()
	fmt.Println("Shutting down MegaMaid...")
}

func downloadEpisode() {
	itemIndex := 0
	itemsToDownload := 1
	feed, _ := LoadFeed("the-anjunadeep-edition.xml")

	for itemIndex < itemsToDownload {
		item, _ := feed.GetItemAtIndex(itemIndex)
		name := item.GetFileName(feed)
		url := item.Enclosure.Url
		DownloadFile(url, name)
		itemIndex++
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
