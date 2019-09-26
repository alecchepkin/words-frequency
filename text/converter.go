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
	for _, w := range strings.Fields(string(data)) {
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
