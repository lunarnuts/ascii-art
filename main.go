package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		os.Exit(0)
	}

	filename, args := GetFlag("--output=", args)
	reverse, args := GetFlag("--reverse=", args)
	color, args := GetFlag("--color=", args)

	bufColor := map[int]string{}
	if color != "" {
		color = ToLower(color)
		bufColor = getColor(color)
	}
	fontStyle := "standard.txt"
	if len(args) > 1 {
		fontStyle = args[1] + ".txt"
	}

	file, err := os.Open(fontStyle)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer func() {
		if err = file.Close(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	ascii := GetASCII(file)

	if reverse != "" {
		fmt.Println(GetReverse(ascii, reverse))
		os.Exit(0)
	}

	buf := make([]string, 8)
	//asciiChar := 0
	str := strings.Split(args[0], "\\n")
	res := ""
	for _, a := range str {
		for ind, n := range a {
			if n < 32 || n > 126 {
				fmt.Println("Error: Message is not valid")
				os.Exit(1)
			}
			//asciiChar = int(rune(n)) - 32 //given template has ascii chars from 32 to 126 (95 chars)
			buffColor := make([]string, 8)
			for i := range buffColor {
				buffColor[i] = bufColor[ind+1]
			}
			buf = AddCh(buf, buffColor)
			buf = AddCh(buf, ascii[n])
		}
		for i := range buf {
			if bufColor[i] != " " {
				res += bufColor[0] + buf[i] + "\n"
			}

		}
		buf = make([]string, 8)
	}
	if filename == "" {
		Prints(res)
	} else {
		ioutil.WriteFile(filename, []byte(res), 0655)
	}
	os.Exit(0)
}

//GetFlag function gets value of flag and deletes it from args
func GetFlag(flag string, args []string) (string, []string) {
	value := ""
	l := len(flag)
	for i := 0; i < len(args); i++ {
		if len(args[i]) > l {
			if args[i][:l] == flag {
				value = args[i][l:]
				args = append(args[:i], args[i+1:]...)
				i--
			}
		}
	}
	return value, args
}

//GetASCII fuction writes all characters to array
func GetASCII(file *os.File) map[rune][]string {
	ascii := make(map[rune][]string, 95)
	scanner := bufio.NewScanner(file)
	buf := make([]string, 8)
	charLine := 0
	asciiChar := 32
	for scanner.Scan() {
		if scanner.Text() == "" {
			charLine = 0
			buf = nil
			continue
		} else {
			buf = append(buf, scanner.Text())
			if asciiChar == 127 { //break the loop when, 96 char is read, as there are only 95 chars
				asciiChar = 0
				break
			}
			if charLine == 7 {
				ascii[rune(asciiChar)] = buf
				buf = nil
				charLine = 0
				asciiChar++
				continue
			}
			charLine++
		}
	}
	return ascii
}

//AddCh function adds characters one by one
func AddCh(buf, new []string) []string {
	for i := range buf {
		buf[i] = buf[i] + new[i]
	}
	return buf
}

//Prints function prints ASCII message by lines
func Prints(buf string) {
	for i := range buf {
		fmt.Print(string(buf[i]))
	}
}
func getColor(s string) map[int]string {
	buf := strings.Split(s, ",")
	color := make(map[int]string, len(buf))
	colorcodes := map[string]string{
		"black":   "\u001b[30m",
		"red":     "\u001b[31m",
		"green":   "\u001b[32m",
		"yellow":  "\u001b[33m",
		"blue":    "\u001b[34m",
		"magenta": "\u001b[35m",
		"cyan":    "\u001b[36m",
		"white":   "\u001b[37m",
		"orange":  "\u001b[38;2;255;165;0m",
		"reset":   "\u001b[0m",
	}
	color[0] = colorcodes["reset"]
	for _, i := range buf {
		b := strings.Split(i, "-")
		if b[0] == "all" {
			color[0] = colorcodes[b[1]]
		}
		p, err := strconv.Atoi(b[0])
		if err != nil {
			color[0] = colorcodes[b[0]]
			continue
		}
		color[p] = colorcodes[b[1]]
	}
	return color
}

//GetReverse function gets result in string from ascii file
func GetReverse(ascii map[rune][]string, filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	lines := GetLines(content)
	l := len(lines)
	if l%8 != 0 {
		log.Fatal("Invalid ASCII art")
		os.Exit(1)
	}

	result := ""
	for resLine := 0; resLine < l/8; resLine++ {
		start := resLine * 8
		end := start + 8
		newLines := lines[start:end]
		currPart := make([]string, 8)
		for i := 0; i < len(newLines[0]); i++ {
			for j, line := range newLines {
				currPart[j] += string(line[i])
			}
			part := ""
			for _, line := range currPart {
				part += line
			}

			for k := 0; k < 95; k++ {
				asciiChar := ""
				for _, r := range ascii[rune(k+32)] {
					asciiChar += r
				}
				if asciiChar == part {
					result += string(k + 32)
					for l := range currPart {
						currPart[l] = ""
					}
					break
				}
			}
		}
		if end != l {
			result += "\n"
		}
	}
	return result
}

//GetLines function separates lines given in file
func GetLines(content []byte) []string {
	lines := []string{}
	currLine := ""
	for i := 0; i < len(content); i++ {
		currLine += string(content[i])
		if content[i] == '\n' {
			lines = append(lines, currLine)
			currLine = ""
		}
	}
	return lines
}

func ToLower(str string) string {
	String := []rune(str)
	for n := range String {
		if (String[n] > 64) && (String[n] < 91) {
			String[n] = String[n] + rune(32)
		}
	}
	return string(String)
}
