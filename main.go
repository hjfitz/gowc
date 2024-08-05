package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	//dbytes := flag.Bool("c", false, "The number of bytes in each input is written to the standard output.")
	dlines := flag.Bool("l", false, "The number of lines in each input file is written to the standard output.")
	//dchars := flag.Bool("m", false, "The number of characters in each input file is written to the standard output. If the current locale does not support multbyte characters, this is equivalent to the -c option.")
	//dwords := flag.Bool("w", false, "The number of words in each input file is written to the standard output.")
	flag.Parse()
	fname := flag.Args()[0]

	fmt.Printf("    ")

	if *dlines {
		lcount := glines(fname)
		fmt.Printf("%d ", lcount)
	}

	fmt.Printf("%s\n", fname)
}

func glines(fname string) int64 {
	var lines int64
	lines = 0
	var wg sync.WaitGroup
	wg.Add(1)
	stream(fname, func(chunk string, err error) {
		if err == io.EOF {
			wg.Done()
			return
		}
		// by default, bufio will scan to each newline, so we just sum the lines
		lines += int64(1)
	})
	wg.Wait()

	return lines
}

func stream(fname string, cb func(c string, e error)) {
	file, err := os.Open(fname)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 1024*1024)
	scanner.Buffer(buf, 10*1024*1024)

	for scanner.Scan() {
		chk := scanner.Text()
		cb(chk, nil)
	}
	cb("", io.EOF)
}
