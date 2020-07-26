package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/blmayer/gortune/fortunes"
)

//go:generate go run gen/main.go

func main() {
	source := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Print a random fortune
	n := source.Intn(len(fortunes.Categories))
	k := source.Intn(len(fortunes.List[fortunes.Categories[n]]))
	fmt.Println(fortunes.List[fortunes.Categories[n]][k])
}
