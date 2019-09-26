package main

import (
	"bufio"
	"errors"
	"fmt"
	"frequency/list"
	"frequency/text"
	"log"
	"os"
	"strconv"
	"time"
)

const defCount = 20

func main() {
	start := time.Now()
	defer func() { fmt.Printf("time execution:%v", time.Now().Sub(start)) }()

	if len(os.Args) == 1 {
		panic(errors.New("filename was not received"))
	}
	name := os.Args[1]

	count := defCount
	if len(os.Args) > 2 {
		var err error
		count, err = strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
	}

	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	trie := list.NewTrie()
	converter := text.NewConverter()
	for scanner.Scan() {
		data := scanner.Bytes()
		words := converter.Split(data)

		for _, w := range words {
			trie.Insert(w)
		}
	}

	mostFrequent := trie.GetMostFrequent(count)

	fmt.Printf("%d most frequent\n:", count)
	for _, node := range mostFrequent {
		fmt.Printf("%d %s\n", node.Count, node.Count)
	}
}
