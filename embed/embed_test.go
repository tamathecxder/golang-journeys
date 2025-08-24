package embed

import (
	"embed"
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

//go:embed texts/1.txt
//go:embed texts/2.txt
//go:embed texts/3.txt
var texts embed.FS

func TestEmbedMultiFiles(t *testing.T) {
	first, _ := fs.ReadFile(texts, "texts/1.txt")
	second, _ := fs.ReadFile(texts, "texts/2.txt")
	third, _ := fs.ReadFile(texts, "texts/3.txt")

	fmt.Println(string(first))
	fmt.Println(string(second))
	fmt.Println(string(third))
}

//go:embed texts/*.txt
var dir embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := dir.ReadDir("texts")

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println("Filename:", entry.Name())
			text, _ := dir.ReadFile("texts/" + entry.Name())
			fmt.Println(string(text))
		}
	}
}
