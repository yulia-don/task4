package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var url *string
var result *string
var dataUrl []string

func init() {
	url = flag.String("url", "C://Users/Стажер/Desktop/task/4/adres.txt", "a string")
	result = flag.String("result", "C://Users/Стажер/Desktop/task/4/result/", "a string")
}

func main() {
	flag.Parse()
	file, err := os.Open(*url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 128)

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		dataUrl = strings.Fields(string(data[:n]))
	}

	for i := 0; i < len(dataUrl); i++ {
		go getsH(i)
	}
	fmt.Scanln()
}

func getsH(i int) {
	resp, err := http.Get("" + dataUrl[i] + "")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	tmp, errr := os.Create(*result + strconv.Itoa(i) + ".html")
	if errr != nil {
		fmt.Println("Unable to create file:", errr)
		os.Exit(1)
	}
	defer tmp.Close()
	io.Copy(tmp, resp.Body)
	fmt.Println(strconv.Itoa(i) + " all")
}
