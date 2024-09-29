package main

import (
	"fmt"
	"github.com/obaydullahmhs/stats"
)

func main() {
	statsCard := stats.NewCard().
		WithUsername("obaydullahmhs").
		WithFilename("obaydullahmhs-github-stats").
		WithPrivateCount().
		WithIcons()

	err := statsCard.Generate()
	if err != nil {
		fmt.Println("Failed to generate statscard")
	}

	// gistCard := gists.NewCard().
	// 	WithId("d8ba36175dba10bc139f488d5c7c1f04").
	// 	WithFilename("publisher-consumer-pattern")

	// err = gistCard.Generate()
	// if err != nil {
	// 	fmt.Println("Faile to generate statscard")
	// }
}
