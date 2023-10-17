package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	palette "github.com/thebromo/bromium/icon"
)

func main() {
	colors := palette.GenerateIconPalette()

	folder, err := os.ReadDir("../icons")

	if err != nil {
		panic(err)
	}

	os.MkdirAll("results/icons", os.ModeAppend)

	for _, v := range folder {
		if strings.Contains(v.Name(), ".svg") {

			//open file
			template, err := os.ReadFile("../icons/" + v.Name())

			if err != nil {
				panic(err)
			}

			content := string(template)

			//replace for each color in file
			for _, color := range colors {
				re := regexp.MustCompile(color.CATHEX)
				content = re.ReplaceAllString(content, color.Hex)
			}

			//save file
			f, err := os.Create("./results/icons/" + v.Name())

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
	}

}
