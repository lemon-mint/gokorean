package vspeach

import "github.com/lemon-mint/gokorean/hangul"

var defaultPronunciation = map[rune]struct {
	First  IPA
	Middle IPA
	End    IPA
}{
	'ㅂ': {
		IPA{
			MainClassification: BilabialPlosive,
			IsVowel:            false,
			Diacritic:          DiacriticNone,
			Suprasegmental:     SuprasegmentalNone,
			Tone:               ToneNone,
			Accent:             AccentNone,
			Notation:           "p",
		},
		IPA{
			MainClassification: BilabialPlosive,
			IsVowel:            false,
			Diacritic:          DiacriticNone,
			Suprasegmental:     SuprasegmentalNone,
			Tone:               ToneNone,
			Accent:             AccentNone,
			Notation:           "b",
		},
		IPA{
			MainClassification: BilabialPlosive,
			IsVowel:            false,
			Diacritic:          DiacriticNone,
			Suprasegmental:     SuprasegmentalNone,
			Tone:               ToneNone,
			Accent:             AccentNone,
			Notation:           "p",
		},
	},
	'ㅍ': {
		IPA{
			MainClassification: BilabialPlosive,
			IsVowel:            false,
			Diacritic:          DiacriticBreathyVoiced,
			Suprasegmental:     SuprasegmentalNone,
			Tone:               ToneNone,
			Accent:             AccentNone,
			Notation:           "p_t",
		},
		IPA{
			MainClassification: BilabialPlosive,
			IsVowel:            false,
			Diacritic:          DiacriticBreathyVoiced,
			Suprasegmental:     SuprasegmentalNone,
			Tone:               ToneNone,
			Accent:             AccentNone,
			Notation:           "p_t",
		},
		IPA{},
	},
}

func analyze(word []rune) {
	kword := make([]hangul.KChar, len(word))
	for i := 0; i < len(word); i++ {
		kword[i], _ = hangul.CharSplit(word[i])

	}
}
