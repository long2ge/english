package logic

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	WORD_URL = "http://www.youdao.com/w/{word}/"
)

// 从有道词典中获取单词信息
func findWordDataByYouDao(word string) string {

	//字符串替换, -1表示全部替换, 0表示不替换, 1表示替换第一个, 2表示替换第二个...
	wordUrl := strings.Replace(WORD_URL, "{word}", word, 1)

	resp, err := http.Get(wordUrl)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return string(body)
}

func YouDaoMake() {

	word := "hello"

	fmt.Println(findWordDataByYouDao(word))

}
