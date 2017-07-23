package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

const (
	// Period -- количество милисекунд, через которые
	// новые данные отправляются на сервер
	Period = 450

	// Port - сокет, который слушаем
	Port = ":7777"

	// PathBoth - путь до файла
	PathBoth = "data/unsd-citypopulation-year-both.csv"

	// PathDiff - путь до файла
	PathDiff = "data/unsd-citypopulation-year-fm.csv"
)

func main() {
	rand.Seed(32)

	str := readCsv(PathBoth)
	lines := stringToLines(str)
	count := len(lines)

	ln, _ := net.Listen("tcp", Port)
	conn, _ := ln.Accept()

	c := time.Tick(Period * time.Millisecond)
	counter := 1
	for range c {
		nextStep := counter + rand.Intn(5)
		for i := counter; i < count && i < nextStep; i++ {
			newmessage := lines[i]
			// send new string back to client
			conn.Write([]byte(newmessage + "\n"))
		}
		counter = nextStep
		if counter > count {
			break
		}
	}

	conn.Close()
}

func readCsv(path string) string {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'

	return str
}

func stringToLines(s string) []string {
	var lines []string

	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return lines
}

func file2lines(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return lines
}
