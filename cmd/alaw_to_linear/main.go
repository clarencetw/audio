package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/clarencetw/audio/pkg/g711"
	"github.com/clarencetw/audio/pkg/resample"
)

func main() {
	rf, err := ioutil.ReadFile("./output.g711")
	if err != nil {
		panic(err)
	}

	linear := make([]byte, len(rf)*2)
	for i, j := 0, 0; i < len(rf); i, j = i+1, j+2 {
		lt := g711.Alaw2linear(rf[i])
		linear[j] = byte(lt)
		linear[j+1] = byte(lt >> 8)
	}

	wf, err := os.Create("./output.raw")
	if err != nil {
		panic(err)
	}
	defer wf.Close()

	resampleLinear := resample.Resample(linear, 8000, 16000)

	fmt.Println("linear len: ", len(linear), " resample len: ", len(resampleLinear))
	wn, err := wf.Write(resampleLinear)
	if err != nil {
		panic(err)
	}

	fmt.Println("Wrote byte = ", wn, "file name = ./output.raw")
}
