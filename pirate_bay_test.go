package petsounds_scrapers

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const validProxy string = "http://tpb.unblocked.co"

func TestSanity(t *testing.T) {
	pirate := NewPirateBay(validProxy)

	if pirate.ProxyUrl != validProxy {
		t.Errorf("ProxyUrl was wrong %s", pirate.ProxyUrl)
	}
}

func TestMagnet2Torrent(t *testing.T) {
	mag := "magnet:?xt=urn:btih:2df72f01370886102cabde3c96e1e7d3930f7865&dn=Daft+Punk+-+Random+Access+Memories+%5BMP3+192%5D&tr=udp%3A%2F%2Ftracker.openbittorrent.com%3A80&tr=udp%3A%2F%2Ftracker.publicbt.com%3A80&tr=udp%3A%2F%2Ftracker.istole.it%3A6969&tr=udp%3A%2F%2Ftracker.ccc.de%3A80&tr=udp%3A%2F%2Fopen.demonii.com%3A1337"
	MagnetToTorrent(mag, "./")

	if files, _ := filepath.Glob("./*.torrent"); len(files) == 0 {
		t.Errorf("File does not exist")
	} else {
		os.Remove("output.torrent")
	}
}

func TestSearchForTorrent(t *testing.T) {
	pirate := NewPirateBay(validProxy)
	term := "Daft Punk Random Access Memories"

	result := pirate.Search(term)

	if len(result) == 0 && !strings.Contains(result, "magnet:") {
		t.Errorf("No results were found for '%s', which is weird. Result: %v", term, result)
	}
}
