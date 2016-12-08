package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	result := ""
	for scanner.Scan() {
		text := scanner.Text()
		text = text + " "

		re := regexp.MustCompile("-[[:space:]]+$")
		text = re.ReplaceAllString(text, "")

		re = regexp.MustCompile("[[:space:]]+")
		text = re.ReplaceAllString(text, " ")

		re = regexp.MustCompile("^[[:space:]]+")
		text = re.ReplaceAllString(text, "")

		result += text
	}
	fmt.Println(strings.TrimSpace(result))
}
