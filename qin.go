package qin

import (
	"regexp"
	"strings"
)

const (
	CUT_SET = " \t\n"
)

func Parse(addr string) (r []string) {
	array := strings.Split(addr, `,`)
	sz := len(array)
	if sz == 0 {
		return
	}
	if sz == 1 {
		return parseAddr(addr)
		r = append(r, trimP(array[0]))
		return
	}
	r = append(r, trimP(array[0]))
	r = append(r, trimC(array[1]))
	return
}

func parseAddr(addr string) (r []string) {
	re := regexp.MustCompile(`(` +
		provincesRegexp + `)(` +
		provincesSuffix + `)?\s*((` +
		citiesRegexp + `)(` +
		citiesSuffix + `)?)?`)

	ss := re.FindStringSubmatch(addr)
	if ss[1] != "" {
		r = append(r, ss[1])
	}
	if ss[4] != "" {
		r = append(r, ss[4])
	}
	return r
}

func trimP(str string) (r string) {
	r = trim(str, `省`)
	return
}

func trimC(str string) (r string) {
	r = trim(str, `市`)
	return
}

func trim(str string, postfix string) (r string) {
	r = strings.Trim(str, CUT_SET+postfix)
	return
}
