package text

import (
	"regexp"
	"strings"
)

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

//Split parse string and split row by the words
func (Converter) Split(data []byte) []string {
	res := make([]string, 0)
	re, _ := regexp.Compile("[.*'/-]+")
	str := re.ReplaceAllString(string(data), " ")

	for _, w := range strings.Fields(str) {
		w = strings.ToLower(w)
		re := regexp.MustCompile("[a-zA-Z]+")
		match := re.FindStringSubmatch(w)
		if len(match) == 0 {
			continue
		}
		w = match[0]
		res = append(res, w)
	}
	return res
}
