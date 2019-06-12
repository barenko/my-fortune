package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Load gets a random quote from a fortune file
func Load(filename string) string {
	content, err := ioutil.ReadFile(filename)
	check(err)

	re := regexp.MustCompile(`(\n%\n|%$)`)
	data := re.Split(string(content), -1)

	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(data))

	return data[idx]
}

func hsize(text string, size int) string {
	var s []string
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		lnSize := len(line)
		if lnSize < size {
			pad := strings.Repeat(" ", size-len([]rune(line)))
			s = append(s, line+pad)
		} else {
			i := 0
			for i < lnSize {
				if i+size <= lnSize {
					pos := strings.LastIndex(line[i:i+size], " ")
					idx := 1
					if pos == -1 {
						pos = strings.Index(line[i+size:], " ")
						pos += size
						idx = 0
					}
					if pos == -1 || i+pos > lnSize {
						txt := line[i:]
						pad := strings.Repeat(" ", size-len([]rune(txt)))
						s = append(s, txt+pad)
						i = lnSize
					} else {
						txt := line[i : i+pos]
						pad := strings.Repeat(" ", size-len([]rune(txt)))
						s = append(s, txt+pad)
						i += pos + idx
					}
				} else {
					txt := line[i:]
					pad := strings.Repeat(" ", size-len([]rune(txt)))
					s = append(s, txt+pad)
					i = lnSize
				}
			}
		}
	}

	return strings.Join(s, "\n")
}

func main() {
	filename := flag.String("f", "./my-fortune", "a fortune file")
	size := flag.Int("l", 120, "Width text size")
	flag.Parse()

	text := Load(*filename)

	text = hsize(text, *size)

	fmt.Println(text)
}
