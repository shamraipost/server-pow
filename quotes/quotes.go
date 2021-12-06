package quotes

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	list = make([]string, 0)
)

func init() {
	readFromFile()
}

func GetRandom() string {
	randSource := rand.NewSource(time.Now().UnixNano())
	randNew := rand.New(randSource)
	return list[randNew.Intn(len(list))]
}

func readFromFile() {
	file, err := os.OpenFile("wordofwisdom.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		list = append(list, sc.Text())
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
}
