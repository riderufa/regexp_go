package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	var filename string
	fmt.Print("Введите название входного файла: ")
	_, err := fmt.Scanln(&filename)
	f, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("Такого файла не существует")
		return
	}
	defer f.Close()

	fileReader := bufio.NewReader(f)

	fmt.Print("Введите название выходного файла: ")
	_, err = fmt.Scanln(&filename)
	_ = os.Remove(filename)
	output, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	for {
		line, _, err := fileReader.ReadLine()
		if err != nil {
			break
		}
		re := regexp.MustCompile(`([0-9]+)([+-]{1})([0-9]+)=+`)
		submatch := re.FindStringSubmatch(string(line))
		if len(submatch) < 1 {
			continue
		}
		a, err := strconv.Atoi(submatch[1])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(submatch[3])
		if err != nil {
			panic(err)
		}
		switch submatch[2] {
		case "+":
			_, _ = output.WriteString(submatch[0])
			_, _ = output.WriteString(fmt.Sprintf("%d\n", a+b))
		case "-":
			_, _ = output.WriteString(submatch[0])
			_, _ = output.WriteString(fmt.Sprintf("%d\n", a-b))
		default:
			fmt.Println("Unsupported operator")
		}
	}
}
