package generator

import (
	_ "embed"
	"testing"
)

func TestGetRandomWord(t *testing.T) {

	tests := []struct {
		name  string
		words []string
		want  string
	}{
		{
			name:  "empty",
			words: []string{},
			want:  "",
		},
		{
			name:  "one",
			words: []string{"one"},
			want:  "One",
		},
		{
			name:  "two",
			words: []string{"", ""},
			want:  "",
		},
		{
			name:  "three",
			words: []string{"three", "three", "three"},
			want:  "Three",
		},
		{
			name:  "nil",
			words: nil,
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRandomWord(tt.words); got != tt.want {
				t.Errorf("GetRandomWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
