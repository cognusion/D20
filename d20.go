package main

import (
	u "github.com/cognusion/go-unique"

	"bytes"
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func getChars(charSet string) string {
	switch charSet {
	case "list":
		return "all bytes alphanumeric alphanumeric-nosim numeric alphabet binary hexadecimal"
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
func randString(size int, charset string) string {
	bytes := make([]byte, size)
	setLen := byte(len(charset))

	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = charset[v%setLen]
	}
	return string(bytes)
}

func randBytes(size int) []byte {
	bytes := make([]byte, size)

	rand.Read(bytes)
	return bytes
}

func blockstring(s string, n int) string {
	var buffer bytes.Buffer
	var n1 = n - 1
	var l1 = len(s) - 1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n1 && i != l1 {
			buffer.WriteString("\n")
		}
	}
	return buffer.String()
}

type generateConfig struct {
	chars     string
	length    int
	blocksize int
	b64       bool
	mangle    string
}

func generateString(conf generateConfig) string {
	var s string
	if conf.chars != "" {
		// Strings!
		s = randString(conf.length, conf.chars)

		if conf.b64 {
			s = base64.StdEncoding.EncodeToString([]byte(s))
		}

	} else {
		// Bytes!
		b := randBytes(conf.length)

		if conf.b64 {
			s = base64.StdEncoding.EncodeToString(b)
		} else {
			s = string(b)
		}
	}
	// POST: s is a populated string of something

	switch strings.ToLower(conf.mangle) {
	case "uc":
		s = strings.ToUpper(s)
	case "lc":
		s = strings.ToLower(s)
	}

	if conf.blocksize > 0 {
		s = blockstring(s, conf.blocksize)
	}

	return s
}

func main() {

	var (
		charset      string
		stringlength int
		stringcount  int
		mangle       string
		b64          bool
		block        bool
		keyblock     bool
		pin          int
		blocksize    int
		unique       bool
		separator    string
		custom       string

		chars string
		uniq  = u.New()
	)

	flag.StringVar(&charset, "chars", "all", "Characters to use ("+getChars("list")+")")
	flag.IntVar(&stringlength, "length", 20, "Length of string")
	flag.IntVar(&stringcount, "count", 20, "Number of strings")
	flag.StringVar(&mangle, "mangle", "", "Mangle the output (WARN: Decreases cardinality, should not be used with --base64) (UC LC)")
	flag.BoolVar(&b64, "base64", false, "Base64 encode the output")
	flag.BoolVar(&block, "block", false, "Block the output to 65 character lines")
	flag.BoolVar(&keyblock, "keyblock", false, "Shortcut to '--char bytes --base64 --block --blocksize 65' (HINT: --length 741, perhaps?)")
	flag.IntVar(&pin, "pin", 0, "Shortcut to '--char numeric --length int'")
	flag.IntVar(&blocksize, "blocksize", 65, "Slight misnomer: if --block is used, sets the line length to int")
	flag.BoolVar(&unique, "unique", false, "Ensure generated strings are unique. Lame. Also may result in < count number of returned results")
	flag.StringVar(&separator, "separator", "\n", "What character or string should each value be separated with?")
	flag.StringVar(&custom, "custom", "", "A list of characters you want to use in lieu of '--chars' (repeat for prevalence)")
	flag.Parse()

	// Sanity
	if stringcount > 10000 {
		fmt.Printf("Requested count of %d is too damn high!\n", stringcount)
		return
	}

	if charset == "list" {
		charset = "all"
	}
	if keyblock {
		charset = "bytes"
		b64 = true
		block = true
		blocksize = 65
	}
	if pin > 0 {
		charset = "numeric"
		stringlength = pin
	}

	// String Quoting Madness
	if separator == "\n" {
		separator = `"\n"`
	} else {
		separator = `"` + separator + `"`
	}
	separator, err := strconv.Unquote(separator)
	if err != nil {
		fmt.Printf("Separator error: %s\n", err.Error())
		return
	}

	// Yes, globals suck, and there are "better" ways to do this.
	// No, in this instance it doesn't matter. Short-lived program, and it spares
	//   us a ton of computation by doing this.
	if custom != "" {
		chars = custom
	} else if charset != "bytes" {
		chars = getChars(charset)
	}

	if !block {
		// Safety, before calling generateString
		blocksize = 0
	}

	perms := math.Pow(float64(len(chars)), float64(stringlength))
	if float64(stringcount) >= perms {
		fmt.Printf("**WARNING** Desired number of %d is >= the possible permutations (%.2f) given the characters and length requested!\n", stringcount, perms)
		if unique {
			fmt.Printf("\tSetting the count from %d to %d because uniqueness is required\n", stringcount, int(perms))
			stringcount = int(perms)
		}
	}

	stringChan := make(chan string, 100)
	configChan := make(chan generateConfig, stringcount)

	// Print All The Strings!

	// Build the config
	gConfig := generateConfig{
		chars:     chars,
		length:    stringlength,
		blocksize: blocksize,
		b64:       b64,
		mangle:    mangle,
	}

	// Spawn off the workers
	go func() {
		for gc := range configChan {
			go func(gc generateConfig) {
				s := generateString(gc)
				if unique && !uniq.IsUnique(s) {
					configChan <- gc // do it again
					return
				}
				stringChan <- s
			}(gc)
		}
	}()

	// queue the work
	for i := 0; i < stringcount; i++ {
		configChan <- gConfig
	}

	var c int
	for s := range stringChan {
		c++
		if c >= stringcount {
			close(configChan)
			close(stringChan)
		}

		fmt.Printf("%s%s", s, separator)
	}
}
