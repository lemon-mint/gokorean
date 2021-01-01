package vspeach

import "github.com/lemon-mint/gokorean/hangul"

var defaultPronunciation = map[rune]struct {
	ChoSung  IPA
	JungSung IPA
	JongSung IPA
}{}

func analyze(word []rune) {
	kword := make([]hangul.KChar, len(word))
	for i := 0; i < len(word); i++ {
		kword[i], _ = hangul.CharSplit(word[i])

	}
}
