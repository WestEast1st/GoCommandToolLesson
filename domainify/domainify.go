package main

/*
ドメインに利用できない文字を変換するコード
末尾にトップレベルドメインを追加することが可能
*/
import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

// top level domeins
var tlds = []string{"com", "cat"}

// 使用可能な文字列
const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func main() {
	var flagTlds = flag.String("tlds", "com, cat", "トップドメインの追加")
	flag.Parse()
	slice := strings.Split(*flagTlds, ",")
	for _, str := range slice {
		tlds = append(tlds, str)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune
		for _, r := range text {
			if unicode.IsSpace(r) {
				r = '-'
			}
			if !strings.ContainsRune(allowedChars, r) {
				continue
			}
			newText = append(newText, r)
		}
		fmt.Println(string(newText) + "." + tlds[rand.Intn(len(tlds))])
	}
}
