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
