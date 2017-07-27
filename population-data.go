package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os"
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

	// Rows -- количество строк, которое хочу считать
	Rows = 500

	// ErrorMessageArgs -- строка об ошибке!
	ErrorMessageArgs = "Что-то не так с аргументом комендной строки -data"
)

func main() {
	wordPtr := flag.String("data", "both", "which file do you need?")
	flag.Parse()

	var path string
	if *wordPtr == "diff" {
		path = PathDiff
	} else if *wordPtr == "both" {
		path = PathBoth
	} else {
		fmt.Println(ErrorMessageArgs)
		return
	}

	fmt.Println("Выбран файл: " + path)
	rand.Seed(time.Now().Unix())

	lines := file2lines(path)

	str := make(chan string)
	quit := make(chan int)
	go launchServer(str, quit)
	linesPusher(lines, str, quit)

}

func launchServer(row chan string, quit chan int) {
	ln, _ := net.Listen("tcp", Port)
	conn, _ := ln.Accept()
	defer conn.Close()

	fmt.Println("в го рутине")
	for {
		select {
		case msg := <-row:
			conn.Write([]byte(msg + "\n"))
		case <-quit:
			return
		}
	}
}

func linesPusher(lines []string, row chan string, quit chan int) {
	counter := 1
	count := len(lines)
	fmt.Printf("В мейн. %d отправил из %d\n", counter, count)
	for {
		nextStep := counter + rand.Intn(Rows)
		for i := counter; i < count && i < nextStep; i++ {
			newmessage := lines[i]
			row <- newmessage
		}
		fmt.Printf("%d lines were sended of %d\n", nextStep, count)

		sleepTime := time.Duration(rand.Intn(Rows*2)) * time.Millisecond
		time.Sleep(sleepTime)

		counter = nextStep
		if counter >= count {
			break
		}
	}
	quit <- 0
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
