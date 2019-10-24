package main

import (
	"testing"
)

var flagtests = []struct {
	to     string
	from   string
	offset int64
	limit  int64
	count  int
}{
	{"folder/1/example.txt", "folder/2/example1.txt", 10, 100, 100},
	{"folder/1/example.txt", "folder/2/example2.txt", 100, 1000, 370},
	{"folder/1/example.txt", "folder/2/example2.txt", 480, 1000, 0},
}

func TestCopy(t *testing.T) {
	for _, tt := range flagtests {
		t.Run("simple", func(t *testing.T) {
			s := copy(tt.to, tt.from, tt.offset, tt.limit)
			if s != tt.count {
				t.Errorf("got %d, want %d", s, tt.count)
			}
		})
	}

}
