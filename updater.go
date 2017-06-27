package main

import (
	cr "github.com/korinoken/twitch-crawler/app"
	"log"
	"encoding/json"
	"io/ioutil"
	_ "github.com/ftrvxmtrx/tga"
)
func main() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	crawler := cr.TwitchCrawler{}
	json.Unmarshal(data, &crawler)
	log.Printf("Loaded config: %v", crawler)
	saveImagesAndUpdate(&crawler)


}
func saveImagesAndUpdate (crawler *cr.TwitchCrawler) (err error) {
	list, err := crawler.GetImageList()
	err = crawler.SaveImages(list)
	return err
}
