package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	palette "github.com/thebromo/bromium/color"
)

func main() {
	colors := palette.GeneratePalette()

	//open file
	template, err := os.ReadFile("./template.json")

	if err != nil {
		panic(err)
	}

	content := string(template)

	//replace for each color in file
	for _, color := range colors {
		re := regexp.MustCompile(color.TemplateName)
		content = re.ReplaceAllString(content, color.Hex)
	}

	//save file
	f, err := os.Create("./../bromium-color-theme.json.generated")

	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f)

	n4, err := w.WriteString(content)

	if err != nil {
		panic(err)
	}

	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}
