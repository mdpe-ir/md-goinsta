package main

import (
	"fmt"
	"os"

	"github.com/mdpe-ir/md-goinsta/pkg/ai"
)

func main() {
	client := ai.NewLLMClient()
	subject := "Golang programming language"
	result, err := ai.GenerateInstagramContent(client, subject, "persian")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Slides:")
	for i, slide := range result.Slides {
		fmt.Printf("%d. %s\n", i+1, slide)
	}
	fmt.Println("\nCaption:")
	fmt.Println(result.Caption)
}
