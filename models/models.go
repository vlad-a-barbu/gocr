package models

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode/utf8"
)

type ImageDescriptor struct {
	RowExpr string
	ColExpr string
}

var imageDescriptors = map[ImageDescriptor]rune{
	{"1+2+1+2+", "1*2+[2-3]{2,}1+"}: 'a',
	{"1{3,}2{2,}1*", "1+2{2,}1+"}:   'b',
	{"1+2*1{3,}2*1+", "1+2{3,}"}:    'c',
	{"1{3,}2{6,}", "1+2+1+"}:        'd',
	{"1+2+1{3,}2*1", "13{2,}2*"}:    'e',
	{"2{4,}1+", "1{4,}"}:            'u',
	{"1{6,}", "1{2,}"}:              'l',
	{"1{6,}2+1*", "1{3,}"}:          'J',

	{"2{6,}", "12+1"}:               '0',
	{"9", ""}:                       '1',
	{"1+2{2,}1{5,}", "1*2+3{3,}2+"}: '2',
	{"9", ""}:                       '3',
	{"9", ""}:                       '4',
	{"9", ""}:                       '5',
	{"9", ""}:                       '6',
	{"9", ""}:                       '7',
	{"9", ""}:                       '8',
	{"9", ""}:                       '9',
}

func MatchHists(rows []int, cols []int) []rune {

	return nil
}

func ReadLabels(dir string) []rune {

	subdirs, err := ioutil.ReadDir(fmt.Sprintf("%s/", dir))
	if err != nil {
		log.Fatal(err)
	}

	var runes []rune
	for _, subdir := range subdirs {

		f, err := os.Open(fmt.Sprintf("%s/%s/label.txt", dir, subdir.Name()))
		if err != nil {
			return nil
		}
		b := make([]byte, 1)

		_, err = f.Read(b)
		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Println(err)
			return nil
		}

		f.Close()
		r, _ := utf8.DecodeRune(b)

		runes = append(runes, r)
	}

	return runes
}

func Match(rows []int, cols []int) []rune {

	/*labels := ReadLabels("hists")

	fmt.Println("Matching image against the following labeled runes:")
	for _, label := range labels {
		fmt.Printf("%c ", label)
	}*/

	var candidates []rune
	for m := range imageDescriptors {
		rres, _ := regexp.MatchString(m.RowExpr, ToString(rows))
		cres, _ := regexp.MatchString(m.ColExpr, ToString(cols))
		if rres && cres {
			candidates = append(candidates, imageDescriptors[m])
		}
	}

	return candidates
}

func ToString(xs []int) string {
	var result string
	for _, x := range xs {
		result += strconv.Itoa(x)
	}
	return result
}
