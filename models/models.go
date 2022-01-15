package models

import (
	"regexp"
	"strconv"
)

type ImageDescriptor struct {
	RowExpr string
	ColExpr string
}

var models = map[ImageDescriptor]rune{
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

func Match(rows []int, cols []int) []rune {
	var candidates []rune
	for m := range models {
		rres, _ := regexp.MatchString(m.RowExpr, ToString(rows))
		cres, _ := regexp.MatchString(m.ColExpr, ToString(cols))
		if rres && cres {
			candidates = append(candidates, models[m])
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
