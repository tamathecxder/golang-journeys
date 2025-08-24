package embed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

//go:embed dummy.jpg
var image []byte

func TestEmbed(t *testing.T) {
	fmt.Println(version)
}

func TestEmbedImage(t *testing.T) {
	err := ioutil.WriteFile("dummy_copy.jpg", image, fs.ModePerm)

	if err != nil {
		panic(err)
	}
}
