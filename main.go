package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	c := make(chan *[]time.Time)

	count := 5

	for i := 0; i < count; i++ {
		go stamp(c)
	}

	f, err := os.Create("stamp.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	for i := 0; i < count; i++ {
		writeFile(f, string(65+i), <-c)
	}

}

func stamp(c chan *[]time.Time) {
	var t []time.Time
	for i := 0; i < 10; i++ {
		t = append(t, time.Now())
	}
	c <- &t
}

func writeFile(f *os.File, s string, t *[]time.Time) {

	w := bufio.NewWriter(f)
	for i, v := range *t {
		w.WriteString(fmt.Sprintf("%s, %d, %d\n", s, i, v.UnixNano()))
	}
	w.Flush()

}
