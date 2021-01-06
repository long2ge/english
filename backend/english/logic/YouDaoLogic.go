package logic

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const (
	WORD_URL    = "http://www.youdao.com/w/{word}"
	EXAMPLE_URL = "http://dict.youdao.com/example/{word}"
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

func FindTest() {
	text := `
	<h2 class="wordbook-js">
	<span class="keyword">study</span>
						<div class="baav">
						<span class="pronounce">英
								<span class="phonetic">[ˈstʌdi]</span>
												<a href="#" title="真人发音" class="sp dictvoice voice-js log-js" data-rel="study&type=1" data-4log="dict.basic.ec.uk.voice"></a>
							</span>
								  <span class="pronounce">美
								<span class="phonetic">[ˈstʌdi]</span>
												<a href="#" title="真人发音" class="sp dictvoice voice-js log-js" data-rel="study&type=2" data-4log="dict.basic.ec.us.voice"></a>
							</span>
							  </div>
		</h2>`

	cityListReg := `<span class="phonetic">([\s\S]*?)</span>[\s\S]*?<span class="phonetic">([\s\S]*?)</span>`

	// match, _ := regexp.MatchString(cityListReg, text)
	//true
	// fmt.Println(match)

	compile := regexp.MustCompile(cityListReg)

	submatch := compile.FindAllStringSubmatch(text, 1)

	fmt.Println(submatch[0])

}

// 获取相关单词和翻译
func FindTranslatesByText(text string) {
	cityListReg := `[\s\S]*?<div class="trans-container">[\s\S]*?<ul>([\s\S]*?)</ul>[\s\S]*?<p class="additional">([\s\S]*?)</p>[\s\S]*?</div>`

	// match, _ := regexp.MatchString(cityListReg, text)
	//true
	// fmt.Println(match)

	compile := regexp.MustCompile(cityListReg)

	submatch := compile.FindAllStringSubmatch(text, -1)

	translates := submatch[0][1]
	relatedWords := submatch[0][2]

	fmt.Println(translates)
	fmt.Println(relatedWords)

}

// 匹配词组
func FindAbc() {
	text := `
	<div id="ads" class="ads">
  <div id="dict-inter" class="dict-inter">
          <p class="hd">词组短语</p>
      <div id="wordGroup2" class="trans-container tab-content hide more-collapse">
                                                                                                        <p class="wordGroup">
                                            <span class="contentTitle"><a class="search-js" href="/w/eng/case_study/#keyfrom=dict.basic.wordgroup">case study</a></span>
                                                                                                        个案研究；案例研究
                                                                  </p>
                                                                                                            <p class="wordGroup collapse">
                                            <span class="contentTitle"><a class="search-js" href="/w/eng/study_group/#keyfrom=dict.basic.wordgroup">study group</a></span>
                                                                                                        学习集团，学习小组；学习研讨会
                                                                  </p>
                                                                                                            <p class="wordGroup collapse">
                                            <span class="contentTitle"><a class="search-js" href="/w/eng/in_a_brown_study/#keyfrom=dict.basic.wordgroup">in a brown study</a></span>
                                                                                                        沉思；幻想；呆想
                                                                  </p>
                                                                                                            <p class="wordGroup collapse">
                                            <span class="contentTitle"><a class="search-js" href="/w/eng/field_of_study/#keyfrom=dict.basic.wordgroup">field of study</a></span>
                                                                                                        研究领域
                                                                  </p>
                                      <div class="more"><a href="#" class="sp more_sp">&nbsp;</a>
        <span class="show_more">更多</span><span
        class="show_less">收起</span>词组短语
        </div>
              </div>
		</div>`

	cityListReg := `<div id="wordGroup2" class="trans-container tab-content hide more-collapse">([\s\S]*?<p class="wordGroup collapse">[\s\S]*?>([\s\S]*?)</a>[\s\S]*?)</div>`

	compile := regexp.MustCompile(cityListReg)

	submatch := compile.FindAllStringSubmatch(text, -1)

	fmt.Println(submatch[0])
}
