package opening

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

type OpeningObj struct {
	Eco  string
	Name string
	Pgn  string
}

type OpeningMessage struct {
	Key   string
	Value *OpeningObj
}

var (
	OpeningStore = make(map[string]*OpeningObj)
	OpeningChan  = make(chan OpeningMessage, 50)
	wg           sync.WaitGroup
)

func OpeningWriter() {
	for opening := range OpeningChan {
		OpeningStore[opening.Key] = opening.Value
	}
}

// later will add workers to read all files concurrently
func ReadTsv(path string, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("failed to open file")
		return
	}

	defer file.Close()

	tsvReader := csv.NewReader(file)
	tsvReader.Comma = '\t'

	// 2 skip first line read once trick btw
	_, err = tsvReader.Read()
	if err != nil {
		return
	}

	for {
		record, err := tsvReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("error while reading", err.Error())
		}

		if len(record) < 3 {
			fmt.Println("invalid record length", len(record))
			continue
		}

		data := OpeningObj{
			Eco:  record[0],
			Name: record[1],
			Pgn:  record[2],
		}
		_, result := NormalizePgn(record[2])
		OpeningChan <- OpeningMessage{
			Key:   result,
			Value: &data,
		}
	}
}

func NormalizePgn(pgn string) (error, string) {
	field := strings.Fields(pgn)

	temp := make([]string, 0, len(field))

	for _, value := range field {
		if strings.HasSuffix(value, ".") {
			continue
		}
		temp = append(temp, value)
	}

	normalized := strings.Join(temp, "|")
	return nil, normalized
}

func ReadAllFiles() {
	files := []string{
		"Opening/a.tsv",
		"Opening/b.tsv",
		"Opening/c.tsv",
		"Opening/d.tsv",
		"Opening/e.tsv",
	}

	go OpeningWriter()
	for _, value := range files {
		wg.Add(1)
		go ReadTsv(value, &wg)
	}
	wg.Wait()
	close(OpeningChan)
}
