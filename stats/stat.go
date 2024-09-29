package stats

import (
	"bytes"
	"fmt"
	"github.com/go-resty/resty/v2"
	"io"
	"os"
)

var (
	defaultTheme    = "github_dark"
	generatorAPI    = "https://github-readme-stats-lime-nu-90.vercel.app/api"
	defaultFileName = "stat"
	defaultSvgDir   = "stats"
)

type Card struct {
	Username     string
	Theme        string
	FileName     string
	FileDir      string
	PrivateCount bool
	Icons        bool
}

func NewCard() *Card {
	return &Card{
		Username:     "",
		Theme:        defaultTheme,
		FileName:     defaultFileName,
		FileDir:      defaultSvgDir,
		PrivateCount: false,
		Icons:        false,
	}
}

func (g *Card) WithUsername(id string) *Card {
	g.Username = id
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

func (g *Card) WithPrivateCount() *Card {
	g.PrivateCount = true
	return g
}

func (g *Card) WithIcons() *Card {
	g.Icons = true
	return g
}

func (g *Card) Generate() error {

	reqURL := fmt.Sprintf("%s?username=%s&show_icons=%t&theme=%s", generatorAPI, g.Username, g.Icons, g.Theme)
	fmt.Println(reqURL)

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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Failed to close file")
			return
		}
	}(file)

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
