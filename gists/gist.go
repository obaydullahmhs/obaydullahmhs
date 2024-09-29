package gists

import (
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"os"
)

var (
	defaultTheme    = "dark"
	generatorAPI    = "https://gists-readme.yizack.com/api/pin"
	defaultFileName = "gist"
	defaultSvgDir   = "gists"
)

type Card struct {
	Id       string
	Theme    string
	FileName string
	FileDir  string
}

func NewCard() *Card {
	return &Card{
		Id:       "",
		Theme:    defaultTheme,
		FileName: defaultFileName,
		FileDir:  defaultSvgDir,
	}
}

func (g *Card) WithId(id string) *Card {
	g.Id = id
	return g
}

func (g *Card) WithTheme(theme string) *Card {
	g.Theme = theme
	return g
}

func (g *Card) WithFilename(name string) *Card {
	g.FileName = name
	return g
}
func (g *Card) WithFileDir(dir string) *Card {
	g.FileDir = dir
	return g
}

func (g *Card) Generate() error {

	reqURL := fmt.Sprintf("%s?user=&id=%s&theme=%s", generatorAPI, g.Id, g.Theme)

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get(reqURL)

	if err != nil {
		fmt.Println(err)
		return err
	}

	// Create a file to save the SVG
	file, err := os.Create(fmt.Sprintf("%s/%s.svg", g.FileDir, g.FileName))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	//fmt.Println(string(resp.Body()))

	responseReader := bytes.NewReader(resp.Body())

	// Write the response body to the file
	_, err = io.Copy(file, responseReader)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Println("SVG file saved successfully!")
	return nil
}
