package utils

import "fmt"

func GenerateTemplate(instructions, input string) string {
	prompt := fmt.Sprintf(`
あなたはプロのガイドアバターです。
以下の制約・施設情報に従って質問に返答する形で道案内をしてください。
- 制約条件
場所を聞かれたらその場所までの道のりを答えること
興味にあった提案をすること
- 施設情報
%s
- 質問
%s`, instructions, input)
	return prompt
}
