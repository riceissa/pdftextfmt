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
	var joinNextLine = false
	for scanner.Scan() {
		text := scanner.Text()
		if !joinNextLine {
			text = text + " "
		}

		// If the whole line is just a hyphen, we want to interpret
		// this as part of a word so join it with the following line.
		// An example would be something like ["mathe", "-",
		// "matician"].
		matched, _ := regexp.MatchString("^[[:space:]]*-[[:space:]]*$", text)
		if matched {
			result = strings.TrimSpace(result)
			joinNextLine = true
			continue
		} else {
			joinNextLine = false
		}

		re := regexp.MustCompile("-[[:space:]]+$")
		text = re.ReplaceAllString(text, "")

		re = regexp.MustCompile("[[:space:]]+")
		text = re.ReplaceAllString(text, " ")

		re = regexp.MustCompile("^[[:space:]]+")
		text = re.ReplaceAllString(text, "")

		result += text
	}
	fmt.Print(strings.TrimSpace(result))
}
