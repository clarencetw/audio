package main

import (
	"fmt"
	"github.com/clarencetw/audio"
	"io/ioutil"
	"os"
)

func main() {
	rf, err := ioutil.ReadFile("./output.g711")
	if err != nil {
		panic(err)
	}

	linear := make([]byte, len(rf)*2)
	for i, j := 0, 0; i < len(rf); i, j = i+1, j+2 {
		lt := audio.Alaw2linear(rf[i])
		linear[j] = byte(lt)
		linear[j+1] = byte(lt >> 8)
	}

	wf, err := os.OpenFile("./output.raw", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer wf.Close()

	wn, err := wf.Write(linear)
	if err != nil {
		panic(err)
	}
	fmt.Println("Wrote byte = ", wn, "file name = ./output.raw")
}
