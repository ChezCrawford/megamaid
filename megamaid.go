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
	url := "https://traffic.libsyn.com/secure/anjunadeep/The_Anjunadeep_Edition_380_with_Daniel_Curpen_FINAL_MIX.mp3"
	name := "380 with Daniel Curpen.mp3"
	DownloadFile(url, name)
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
