package parser

import (
	"bufio"
	"fmt"
	"html"
	"log"
	"os"
	"strings"
)

const boilerPlate string = " <!DOCTYPE html>\n<html>\n<head>\n <title>Checkbox List</title>\n </head>\n <body>	"

var http strings.Builder

func Lexer(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	http.WriteString(boilerPlate)
	defer http.Reset()

	sc := bufio.NewScanner(file)
	count := 1
	for sc.Scan() {
		line := sc.Text()
		if strings.HasPrefix(line, H3) {
			line := strings.TrimPrefix(line, H3)
			http.WriteString("\n<h3>" + html.EscapeString(line) + "</h3>\n")
		} else if strings.HasPrefix(line, H2) {
			line := strings.TrimPrefix(line, H2)
			http.WriteString("\n<h2>" + html.EscapeString(line) + "</h2>\n")
		} else if strings.HasPrefix(line, H1) {
			line := strings.TrimPrefix(line, H1)
			http.WriteString("\n<h1>" + html.EscapeString(line) + "</h1>\n")
		} else if strings.HasPrefix(line, CHECKBOX) {
			http.WriteString(fmt.Sprintf("\n<label><input type=\"checkbox\" name=\"option\"=%d, value=%d>", count, count))
			line := strings.TrimPrefix(line, CHECKBOX)
			http.WriteString(line + CheckTokenEnd)
			count++
		} else {
			http.WriteString("\n" + line + "\n")
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	http.WriteString("\n</body>\n</html>")
	return http.String()
}
