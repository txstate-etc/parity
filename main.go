package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var before string
var after string

func init() {
	if len(os.Args) < 3 || os.Args[1] == "" || os.Args[2] == "" {
		log.Panic("Please include before and after filenames.")
	}
	before = os.Args[1]
	after = os.Args[2]
}

func count(links map[string]int, path string, cnt int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linkinfo := scanner.Text()
		links[linkinfo] = links[linkinfo] + cnt
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func sortprint(links map[string]int) (keys []string) {
	for k, v := range links {
		if v != 0 {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return
}

func main() {
	var links = make(map[string]int)
	count(links, before, -1)
	count(links, after, 1)
	list := sortprint(links)
	for _, k := range list {
		fmt.Printf("%6d %s\n", links[k], k)
	}
}
