package rword

import (
	"encoding/json"
	"github.com/cxnky/rword-go/lib"
	"io/ioutil"
	"strconv"
	"strings"
)

/**

 * Created by cxnky on 08/05/2019 at 13:47
 * rword_go
 * https://github.com/cxnky/

**/

type WordList struct {
	Words	[]string
}

type GenerateOptions struct {
	Length 		string
	Capitalise  string
}

type RWord struct {
	globalPool []string
	Options GenerateOptions
}

var (
	words = make([]string, 0)
)

func (r *RWord) Generate(count int) []string {
	pool := make([]string, 0)
	length, err := strconv.Atoi(r.Options.Length)

	if err != nil {
		panic(err)
	}

	capitalise := strings.ToLower(r.Options.Capitalise)

	if capitalise != "all" && capitalise != "first" && capitalise != "none" {
		panic("invalid capitalise option, valid options are all/first/none")
	}

	// As Go does not have any lambda functions; we will have to iterate over the array manually
	for _, w := range words {
		if len(w) == length {
			pool = append(pool, w)
		}
	}

	if len(pool) == 0 {
		return nil
	}

	indexes := lib.GenerateIndexes(len(pool), count)
	temp := make([]string, 0)

	for _, i := range indexes {
		temp = append(temp, pool[i])
	}

	pool = temp

	if capitalise == "all" {
		for i, w := range pool {
			pool[i] = strings.ToUpper(w)
		}
	} else if capitalise == "first" {
		for i, w := range pool {
			pool[i] = strings.ToUpper(string(w[0])) + w[1:]
		}
	}

	return pool

}

func (r *RWord) Shuffle() {
	lib.ShuffleWords(words)
}

func (r *RWord) Load(wordList string) {

	if wordList != "big" && wordList != "small" {
		panic("invalid word list. valid options are big/small")
	}

	bytes, err := ioutil.ReadFile("words/" + wordList + ".json")

	if err != nil {
		panic(err)
	}

	var wordsList []string
	err = json.Unmarshal(bytes, &wordsList)

	if err != nil {
		panic(err)
	}

	words = wordsList
	r.Shuffle()
}