package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

func numDigits(v int) int {
	n := 1
	for ; v >= 10; v /= 10 {
		n++
	}
	return n
}

func maxPowerToTenBelow(v int) int {
	d := 1
	for ; d*10 <= v; d *= 10 {
	}
	return d
}

func numToFilename(v int) string {
	buf := make([]byte, 0, 64)
	w := 1
	p := 1
	for ; p*10 <= v; p *= 10 {
		w++
	}
	for ; p > 0; p /= 10 {
		d := v / p
		buf = append(buf, byte(d)+'0')
		if w%3 == 1 && w >= 3 {
			buf = append(buf, '/')
		}

		v -= d * p
		w--
	}
	buf = append(buf, ".b"...)
	return string(buf)
}

func main() {
	var num int
	flag.IntVar(&num, "num", 10000, "number of files")
	var baseDir string
	flag.StringVar(&baseDir, "base-dir", "out", "base directory")
	flag.Parse()

	for i := 0; i < num; i++ {
		path := filepath.Join(baseDir, numToFilename(i))
		dir := filepath.Dir(path)
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}
}
