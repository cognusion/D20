package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"strings"
)

var CharSet string


func getChars(charSet string) string {
	switch charSet {
	case "list":
		return "all alphanumeric alphanumeric-nosim numeric alphabet binary hexadecimal"

	case "numeric":
		return "0123456789"
	case "bin":
		fallthrough
	case "binary":
		return "01"
	case "hex":
		fallthrough
	case "hexadecimal":
		return "0123456789ABCDEF"
	case "alphanumeric-nosim":
		fallthrough
	case "alpha-nosim":
		return "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghjkmnopqrstuvwxyz"
	case "alphabet":
		return "ABCDEFGHJKLMNPQRSTUVWXYZabcdefghjkmnopqrstuvwxyz"
	case "alphanumeric":
		fallthrough
	case "alpha":
		return "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	default:
		// Everything
		return "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_-!$%^&();:.,<>/?"
	}
}

// Return a randomish string of the specified size,
// using the global CharSet
func randString(size int) string {

	bytes := make([]byte, size)
	setLen := byte(len(CharSet))

	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = CharSet[v%setLen]
	}
	return string(bytes)
}

func main() {

	var (
		charset      string
		stringlength int
		stringcount  int
		mangle       string
	)

	flag.StringVar(&charset, "chars", "all", "Characters to use ("+getChars("list")+")")
	flag.IntVar(&stringlength, "length", 20, "Length of string")
	flag.IntVar(&stringcount, "count", 20, "Number of strings")
	flag.StringVar(&mangle, "mangle", "", "Mangle the output (Decreases cardinality) (UC LC)")
	flag.Parse()

	// Sanity
	if charset == "list" {
		charset = "all"
	}

	// Yes, globals suck, and there are "better" ways to do this.
	// No, in this instance it doesn't matter. Short-lived program, and it spares
	//   us a ton of computation by doing this.
	CharSet = getChars(charset)

	// Print All The Strings!
	for i := 0; i < stringcount; i++ {
		s := randString(stringlength)
		
		switch strings.ToLower(mangle) {
		case "uc":
			s = strings.ToUpper(s)
		case "lc":
			s = strings.ToLower(s)
		}
		
		fmt.Println(s)
	}
}
