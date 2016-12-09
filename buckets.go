package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {

	result, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(result.Body)
	defer result.Body.Close()
	scanner.Split(bufio.ScanWords)
	buckets := make([]int, 10)

	for scanner.Scan() {
		n := hashBucket(scanner.Text(), 10)
		buckets[n]++
	}
	fmt.Println(buckets)
}

func hashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
