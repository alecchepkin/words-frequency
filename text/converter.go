package text

import (
	"log"
	"regexp"
	"strings"
)

type Converter struct {
}

//Split should split row by the words
func (converter Converter) Split(data []byte) []string {
	res := make([]string, 0)

	for _, w := range strings.Fields(string(data)) {
		w = strings.ToLower(w)
		r, err := regexp.Compile("[^a-zA-Z]")
		if err != nil {
			log.Fatal(err)
		}
		w = r.ReplaceAllString(w, "")
		if len(w) == 0 {
			continue
		}
		res = append(res, w)
	}
	return res
}

func NewConverter() *Converter {
	return &Converter{}
}
