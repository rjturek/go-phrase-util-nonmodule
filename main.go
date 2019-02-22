package main

import (
	"fmt"

	"github.com/rjturek/go-phrase-util-nonmodule/rjtphrasenonmod"
)

var phrase1 = "But when night comes round, oh gosh, oh gee"

func GetPhrase() string {
	return phrase1
}

func main() {
	fmt.Println(GetPhrase())
	fmt.Println(rjtphrasenonmod.GetPhrase())
}
