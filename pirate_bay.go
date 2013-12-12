package petsounds_scrapers

import (
	// "github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Scraper interface {
	Search() []string
	BestMatch() []string
}

func BuildTorrentFilenameFromMagnet(dest string, magnet string) string {
	u, err := url.Parse(magnet)

	if err != nil {
		log.Fatal(err)
	}

	name := strings.Replace(u.Query()["xt"][0], "urn:btih:", "", -1)

	return dest + name + ".torrent"
}

func MagnetToTorrent(magnet string, destination string) {
	resp, err := http.PostForm("http://magnet2torrent.com/upload/", url.Values{"magnet": {magnet}})
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create(BuildTorrentFilenameFromMagnet(destination, magnet))
	defer out.Close()

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(out, resp.Body)
}

type PirateBay struct {
	ProxyUrl string
}

func NewPirateBay(proxyUrl string) *PirateBay {
	return &PirateBay{ProxyUrl: proxyUrl}
}

func (pb PirateBay) Search(term string) []string {
	//doc, _ := goquery.NewDocument(pb.ProxyUrl + "/0/7/0")
	return []string{}
}
