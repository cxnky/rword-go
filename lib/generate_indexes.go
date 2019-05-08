package lib

import (
	"github.com/heapwolf/go-indexof"
	"math/rand"
)

/**

 * Created by cxnky on 08/05/2019 at 13:37
 * lib
 * https://github.com/cxnky/

**/

// GenerateIndexes randomly generates the numbers for indexes within the words array
func GenerateIndexes(length, count int) []int {
	indexes := make([]int, 0)

	for {
		index := rand.Intn(length - 0) + 0
		if indexof.IndexOf(index, indexes) == -1 {
			indexes = append(indexes, index)
		}
		if len(indexes) == count {
			break
		} else if length < count && len(indexes) == length {
			break
		}
	}

	return indexes
}