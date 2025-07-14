package main

import (
	"github.com/mdpe-ir/md-goinsta/internal/config"
	"github.com/mdpe-ir/md-goinsta/pkg/postgen"
)

// "fmt"
// "os"

// "github.com/mdpe-ir/md-goinsta/pkg/ai"

func main() {

	cfg := config.Load()
	_ = postgen.PostGen(postgen.PostGenConfig{
		InstagramUsername: cfg.Instagram.InstagramUsername,
		Content:           "Test",
		SlidesCount:       10,
		Index:             1,
	})

	// client := ai.NewLLMClient()
	// subject := "Golang programming language"
	// result, err := ai.GenerateInstagramContent(client, subject, "persian")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }
	// fmt.Println("Slides:")
	// for i, slide := range result.Slides {
	// 	fmt.Printf("%d. %s\n", i+1, slide)
	// }
	// fmt.Println("\nCaption:")
	// fmt.Println(result.Caption)
}
