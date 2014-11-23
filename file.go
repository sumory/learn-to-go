package main

import (
	"fmt"
	//"io"
	"io/ioutil"
	"os"
	"time"
	//"strings"
)

func main() {
	fmt.Println("++++++文件操作实例++++++")
	fileName := "/data/go_snippets/test.txt"

	// var buff = make([]byte, 14)
	// for {
	// 	n, err := f.Read(buff)
	// 	if err != nil && err != io.EOF {
	// 		panic(err)
	// 	}

	// 	if n == 0 {
	// 		break
	// 	}

	// 	os.Stdout.Write(buff[:n])

	// }
	//
	count := 0
	for {
		count += 1
		content := "new content 新内容" + getNow()
		fmt.Println("写之前:", count)
		fmt.Println(readFile(fileName))
		fmt.Println("写之后:", count)
		writeFile(fileName, content)
		fmt.Println(readFile(fileName))

		if count > 5 {
			break
		}
	}

}

func writeFile(fileName string, content string) {
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0660)
	if err != nil {
		panic("open file failed")
	}

	defer f.Close()
	c := []byte(content)
	ioutil.WriteFile(fileName, c, os.ModePerm)
}

func readFile(fileName string) string {
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0660)
	if err != nil {
		panic("open file failed")
	}

	defer f.Close()
	c, err := ioutil.ReadFile(fileName)
	if err != nil {
		return ""
	} else {
		return string(c)
	}
}

func getNow() string {
	now := time.Now()
	year, mon, day := now.Date()
	hour, min, sec := now.Clock()
	return fmt.Sprintf("%d-%d-%d %02d:%02d:%02d", year, mon, day, hour, min, sec)
}
