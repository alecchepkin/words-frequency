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

// defNum default number of frequent words
const defNum = 20

func main() {
	start := time.Now()
	defer func() { fmt.Printf("time execution:%v", time.Now().Sub(start)) }()

	if len(os.Args) == 1 {
		panic(errors.New("filename was not received"))
	}
	name := os.Args[1]

	num := defNum
	if len(os.Args) > 2 {
		var err error
		num, err = strconv.Atoi(os.Args[2])
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

	trie := list.NewTrie(num)
	converter := text.NewConverter()
	for scanner.Scan() {
		data := scanner.Bytes()
		words := converter.Split(data)

		for _, w := range words {
			trie.Insert(w)
		}
	}

	for _, node := range trie.GetMostFrequent() {
		fmt.Printf("%d %s\n", node.Count, node.Word)
	}
}
