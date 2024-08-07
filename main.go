package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type Flags struct {
	dbytes *bool
	dlines *bool
	dchars *bool
	dwords *bool
}

func main() {
	flags := Flags{}
	flag.BoolVar(flags.dbytes, "c", false, "The number of bytes in each input is written to the standard output.")
	flag.BoolVar(flags.dlines, "l", false, "The number of lines in each input file is written to the standard output.")
	flag.BoolVar(flags.dchars, "m", false, "The number of characters in each input file is written to the standard output. If the current locale does not support multbyte characters, this is equivalent to the -c option.")
	flag.BoolVar(flags.dwords, "w", false, "The number of words in each input file is written to the standard output.")
	flag.Parse()

	for _, fname := range flag.Args() {
		wc(&flags, &fname)
	}
}

func wc(flags *Flags, fname *string) {
	scn, file := stream_lines(fname)
	defer file.Close()

	tlines := int64(0)
	twords := int64(0)
	tbytes := int64(0)
	tchars := int64(0)

	for scn.Scan() {
		chk := scn.Text()

		words := strings.Fields(chk)
		twords += int64(len(words))

		tlines += int64(1)

		tbytes += int64(len(chk))

		tchars += int64(utf8.RuneCountInString(chk))

	}

	// account to chunks being newlines
	tbytes += tlines
	tchars += tlines

	fmt.Printf("   ")
	if *flags.dbytes {
		fmt.Printf("%d    ", tbytes)
	}

	if *flags.dlines {
		fmt.Printf("%d    ", tlines)
	}

	if *flags.dchars {
		fmt.Printf("%d    ", tchars)
	}

	if *flags.dwords {
		fmt.Printf("%d ", twords)
	}

	fmt.Printf("%s\n", *fname)
}

func stream_lines(fname *string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(*fname)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 1024*1024)
	scanner.Buffer(buf, 10*1024*1024)

	return scanner, file

}
