package rjtphrasenonmod

import "github.com/golang/example/stringutil"

var phrase = "And they only use suspenders in the fall"

func GetPhrase() string {
	return stringutil.Reverse(stringutil.Reverse(phrase))
}
