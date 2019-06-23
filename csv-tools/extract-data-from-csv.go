package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var (
	mFlag = flag.String("m", "2019-01", "Set this month.")
)

func transformEncoding(rawReader io.Reader, trans transform.Transformer) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(rawReader, trans))
	if err == nil {
		return string(ret), nil
	} else {
		return "", err
	}
}

func ShiftJIStoUTF8(str string) (string, error) {
	return transformEncoding(strings.NewReader(str), japanese.ShiftJIS.NewDecoder())
}

func main() {

	flag.Parse()

	var fp *os.File
	if len(os.Args) < 2 {
		fp = os.Stdin
	} else {
		var err error
		fp, err = os.Open(os.Args[3])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	reader := csv.NewReader(fp)
	reader.Comma = ','
	reader.LazyQuotes = true
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		var csvRecord string
		if strings.HasPrefix(record[1], *mFlag) {
			for i, v := range record {
				if i == 0 {
					csvRecord = v
				} else {
					csvRecord = csvRecord + "," + v
				}
			}
			result, err := ShiftJIStoUTF8(csvRecord)
			if err != nil {
				panic(err)
			}
			fmt.Println(result)
		}
	}
}
