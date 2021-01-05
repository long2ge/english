package logic

import (
	"fmt"
	"strings"
)

func Test() {
	fmt.Println(123)
}

func Test2() int {
	return 1
}

func TextToWords(text string) []string {

	// http://www.youdao.com/w/hello/

	data := strings.Split(text, " ")

	return data
}

func Make() {
	text := "After you click on the Link if it doesn't Start then it is a Fake Link. Even if it says Paypal in it somewhere it is a Fake Link. The term  should always precede any website address where you enter personal information. The stands for secure. If you don't see you're not in a secure web session, and you should not enter data.(This goes for any payment Processor including your online Bank Accounts)"

	// words := map[string]string{}

	// 空格切割

	// for i := 0; i < len(str); i++ {
	// 	ch := str[i]
	// 	fmt.Println(12345)
	// 	fmt.Println(string(ch))

	// 	words[string(ch)] = string(ch)
	// }
	fmt.Println(len(TextToWords(text)))
	// for _, ch1 := range str {
	// 	fmt.Println(ch1)
	// }
}
