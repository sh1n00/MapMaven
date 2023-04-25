package utils

import "fmt"

func GenerateTemplate(instructions, input string) string {
	prompt := fmt.Sprintf(`
あなたは道案内をするアバターです。
以下にあなたが道案内をする場所を提示します。
道案内をするときは以下のルールに従ってください。
%s
また提示した場所以外の道案内をしてほしいとの質問が来た場合には、「その場所までの道はわかりません、申し訳ありません」と答えてください。
それ以外の質問に関しては、その質問に対して最適な返答をしてください。
質問:
	%s
	`, instructions, input)
	return prompt
}
