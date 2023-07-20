package hash

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type KV struct {
	Key   string
	Value int
}

func Handle(path string, n int, hFunc HFunc) *KV {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	channels := make(map[int]chan string)
	for i := 0; i < n; i++ {
		num := i
		ch := make(chan string)
		go writeFile(fmt.Sprintf("%d", num), ch)
		channels[num] = ch
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num := hFunc(scanner.Text()) % n
		ch := channels[num]
		ch <- scanner.Text()
	}
	for _, v := range channels {
		close(v)
	}
	resCh := make(chan *KV, n)
	for i := 0; i < n; i++ {
		num := i
		go readFile(fmt.Sprintf("%d", num), resCh)
	}
	results := make([]*KV, n)
	for i := 0; i < n; i++ {
		v := <-resCh
		results[i] = v
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].Value > results[j].Value
	})
	return results[0]
}

func writeFile(path string, ch <-chan string) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	defer f.Close()
	for v := range ch {
		fmt.Fprintln(f, v)
	}
}

func readFile(path string, result chan<- *KV) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(f)
	res := make(map[string]int)
	for scanner.Scan() {
		res[scanner.Text()] += 1
	}
	maxKey := ""
	max := 0
	for k, v := range res {
		if v > max {
			maxKey = k
			max = v
		}
	}
	result <- &KV{Key: maxKey, Value: max}
}
