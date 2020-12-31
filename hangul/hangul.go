package hangul

import "errors"

//KChar Analyzed Korean Character
type KChar struct {
	Char        rune
	IsHangul    bool
	IsComponent bool
	IsFull      bool
	IsVowel     bool
	IsConsonant bool
	HasChoSung  bool
	ChoSung     rune
	HasJungSung bool
	JungSung    rune
	HasJongSung bool
	JongSung    rune //If the letter does not have a JongSung, the space is returned.
}

//ErrIsNotHangul : The requested character is not a Hangul character
var ErrIsNotHangul = errors.New("The requested character is not a Hangul character")

//CharSplit : Split Hangul character to Hangul Component
func CharSplit(char rune) (KChar, error) {
	CharCode := int(char)
	isComponent := ((0x1100 <= CharCode) && (CharCode <= 0x11ff)) || ((0x3130 <= CharCode) && (CharCode <= 0x318f))
	isFull := (0xac00 <= CharCode) && (CharCode <= 0xd7a3)
	IsHangul := false
	if isComponent || isFull {
		IsHangul = true
	} else {
		return KChar{
			Char:        char,
			IsHangul:    IsHangul,
			IsComponent: false,
			IsFull:      false,
			IsVowel:     false,
			IsConsonant: false,
			HasChoSung:  false,
			HasJungSung: false,
			HasJongSung: false,
		}, ErrIsNotHangul
	}

	if isComponent {
		return KChar{
			Char:        char,
			IsHangul:    IsHangul,
			IsComponent: true,
			IsFull:      false,
			IsVowel:     false,
			IsConsonant: false,
			HasChoSung:  false,
			HasJungSung: false,
			HasJongSung: false,
		}, nil
	}

	ChoSungIndex := ((CharCode - 0xac00) / 28) / 21
	JungSungIndex := ((CharCode - 0xac00) / 28) % 21
	JongSungIndex := (CharCode - 0xac00) % 28
	HasJongSung := JongSungIndex == 0

	return KChar{
		Char:        char,
		IsHangul:    IsHangul,
		IsComponent: false,
		IsFull:      true,
		HasChoSung:  true,
		ChoSung:     chosung[ChoSungIndex],
		HasJungSung: true,
		JungSung:    jungsung[JungSungIndex],
		HasJongSung: HasJongSung,
		JongSung:    jongsung[JongSungIndex],
	}, nil
}
