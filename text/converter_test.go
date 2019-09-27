package text

import (
	"reflect"
	"testing"
)

func TestConverter_Split(t *testing.T) {
	tests := []struct {
		name string
		data string
		want []string
	}{
		{"simple slice", "aaa bbb ccc", []string{"aaa", "bbb", "ccc"}},
		{"to lower", "aaa BBB ccc", []string{"aaa", "bbb", "ccc"}},
		{"only words", "(aaa,) BBB! {ccc}.", []string{"aaa", "bbb", "ccc"}},
		{"empty if no letters", "(,) ! {}.", []string{}},
		{"cutting the word, take only suffix with letters", "abc10dsfsa", []string{"abc"}},
		{"cutting the word, take only suffix with letters", "abc10dsfsa", []string{"abc"}},
		{"split by slash", "abc/def", []string{"abc", "def"}},
		{"split by dash", "Jonas-in-the-Whale", []string{"jonas", "in", "the", "whale"}},
		{"remove asterisk", "**The Project", []string{"the", "project"}},
		{"remove asterisk", "***START**THE", []string{"start", "the"}},
		{"remove quotes", "\"The Indian", []string{"the", "indian"}},
		{"remove dashes", "--THE", []string{"the"}},
		{"should separate by ' ", "WEBSTER'S", []string{"webster", "s"}},
		{"should separate . ' ", "GUTINDEX.ALL", []string{"gutindex", "all"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			converter := Converter{}
			if got := converter.Split([]byte(tt.data)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}
