package main

import (
	"fmt"
	"io"
	"math"
	"os"

	"github.com/cheggaaa/pb/v3"
	flag "github.com/spf13/pflag"
)

var from string
var to string
var offset int64
var limit int64

func init() {
	flag.StringVar(&from, "from", "", "absolut path from copy")
	flag.StringVar(&to, "to", "", "absolut path to copy")
	flag.Int64VarP(&offset, "offset", "o", 0, "offset bytes")
	flag.Int64VarP(&limit, "limit", "l", 0, "limit bytes")
}
func main() {
	flag.Parse()
	if from == "" || to == "" {
		fmt.Println("No such requery params path from and to")
		return
	}
	bytes := copy(from, to, offset, limit)
	fmt.Printf("%d bytes\n", bytes)
}
func copy(from string, to string, offset int64, limit int64) int {
	var bufSize int64 = 8
	var count int
	fileFrom, err := os.Open(from)
	if err != nil {
		fmt.Printf("Can't open file: %v %s", err, from)
		return 0
	}
	fi, err := fileFrom.Stat()
	if err != nil {
		fmt.Printf("Could not obtain stat: %v %s", err, from)
		return 0
	}
	if offset > fi.Size() {
		fmt.Printf("offset more then file size")
		return 0
	}
	println(fi.Size())
	if offset+limit > fi.Size() {
		limit = fi.Size() - offset
	}
	bar := pb.StartNew(int(math.Round(float64(limit)/float64(bufSize) + 0.49)))
	//fileTo, err := os.OpenFile(to, os.O_RDWR|os.O_CREATE, 0666)
	fileTo, err := os.Create(to)
	if err != nil {
		fmt.Printf("Can't create/open file: %v %s", err, to)
		return 0
	}
	var i int64
	for i = 0; i < limit; i = i + bufSize {
		bar.Increment()
		o3, err := fileFrom.Seek(int64(offset), 0)

		if err != nil {
			fmt.Printf("failed to seek: %v %d", err, o3)
			return 0
		}
		if i+bufSize > limit {
			bufSize = limit - i
		}
		b3 := make([]byte, bufSize)
		n3, err := io.ReadAtLeast(fileFrom, b3, int(bufSize))
		count = count + n3
		if err == io.EOF {
			bufSize = limit
			continue
		}
		a, err := fileTo.Write(b3)

		if err != nil {
			fmt.Printf("failed to write: %v %d", err, a)
			return count
		}
		offset = offset + bufSize
	}
	fileFrom.Close()
	fileTo.Close()
	bar.Finish()

	return count
}
