package main

/*
Web向けの語句をいくつか生成、空いているドメインが見つかる確率が上がります。
 - さまざまな文字列の変換方法を配列として表現する。元の語が現れる場所に特別な文字を使った文字列によって変換方法を示す。
 - bfuio パッケージを使って標準入力から語を読み込み、変換された語をfmt.Printlnで標準出力に書き出します
 - math/rand パッケージを使い変換方法をランダムに選択する。
*/
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord,
	otherWord,
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get" + otherWord,
	"go" + otherWord,
	"lest" + otherWord,
}

// 標準入力を永遠に受け付けてランダムに文字列を追加する
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}
