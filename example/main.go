package example

import (
	"fmt"
	"github.com/cxnky/rword-go"
	"strconv"
)

/**

 * Created by cxnky on 08/05/2019 at 14:46
 * example
 * https://github.com/cxnky/

**/

func main() {
	wordGen := rword.RWord{Options: rword.GenerateOptions{Length: "8", Capitalise: "none"}}
	wordGen.Load("big")

	words := wordGen.Generate(5000)

	for i, w := range words {
		humanPosition := strconv.Itoa(i + 1)
		fmt.Println("Word at position #" + humanPosition + " is " + w)
	}
}