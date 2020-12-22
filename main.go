package main

import (
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"log"
	"strings"

	"github.com/kyo1/go-cartesian-product"
	"gopkg.in/alecthomas/kingpin.v2"
)

var chars = []interface{}{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

var (
	app = kingpin.New("pow", "A command-line Proof-of-Work solver")

	algorithm = kingpin.Flag("algorithm", "md5, sha1, sha224, sha256, sha384, sha512").Required().Short('a').String()

	inputPrefix  = kingpin.Flag("input-prefix", "input prefix").String()
	inputSuffix  = kingpin.Flag("input-suffix", "input suffix").String()
	outputPrefix = kingpin.Flag("output-prefix", "output prefix").String()
	outputSuffix = kingpin.Flag("output-suffix", "output suffix").String()

	dump = kingpin.Flag("dump", "dump a hash digest").Bool()
)

func hashStr(algorithm, input string) (string, error) {
	var h hash.Hash

	switch algorithm {
	case "md5", "MD5":
		h = md5.New()
	case "sha1", "SHA1":
		h = sha1.New()
	case "sha224", "SHA224":
		h = sha256.New224()
	case "sha256", "SHA256":
		h = sha256.New()
	case "sha384", "SHA384":
		h = sha512.New384()
	case "sha512", "SHA512":
		h = sha512.New()
	default:
		return "", fmt.Errorf("%s is not suppurted", algorithm)
	}

	h.Write([]byte(input))
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func interfacesToString(v []interface{}) string {
	var sb strings.Builder
	for _, x := range v {
		sb.WriteString(x.(string))
	}
	return sb.String()
}

func solve(algorithm, inputPrefix, inputSuffix, outputPrefix, outputSuffix string) (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for s := range cartesian.All(ctx, chars) {
		candidate := fmt.Sprintf("%s%s%s", inputPrefix, interfacesToString(s), inputSuffix)

		hexDigest, err := hashStr(algorithm, candidate)
		if err != nil {
			return "", err
		}

		if strings.HasPrefix(hexDigest, outputPrefix) && strings.HasSuffix(hexDigest, outputSuffix) {
			return candidate, nil
		}
	}
	return "", nil
}

func main() {
	kingpin.Parse()

	res, err := solve(*algorithm, *inputPrefix, *inputSuffix, *outputPrefix, *outputSuffix)
	if err != nil {
		log.Fatal(err)
	}

	if *dump {
		hexDigest, _ := hashStr(*algorithm, res)
		fmt.Printf("%s(%s) = %s\n", *algorithm, res, hexDigest)
	} else {
		fmt.Printf("%s\n", res)
	}
}
