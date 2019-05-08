package lib

import (
	"math/rand"
	"time"
)

/**

 * Created by cxnky on 08/05/2019 at 13:43
 * lib
 * https://github.com/cxnky/

**/

// ShuffleWords shuffles the words array in place
func ShuffleWords(words []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func (i, j int) {
		words[i], words[j] = words[j], words[i]
	})
}