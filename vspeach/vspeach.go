package vspeach

import "github.com/lemon-mint/gokorean/hangul"

var defaultConsonantPronunciation = map[rune]struct {
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
			Base:               "p",
			Notation:           "p",
		},
		IPA{
			MainClassification: BilabialPlosive,
			IsVowel:            false,
			Diacritic:          DiacriticNone,
			Suprasegmental:     SuprasegmentalNone,
			Tone:               ToneNone,
			Accent:             AccentNone,
			Base:               "b",
			Notation:           "b",
		},
		IPA{
			MainClassification: BilabialPlosive,
			IsVowel:            false,
			Diacritic:          DiacriticNone,
			Suprasegmental:     SuprasegmentalNone,
			Tone:               ToneNone,
			Accent:             AccentNone,
			Base:               "p",
			Notation:           "p",
		},
	},
	'ㅍ': {
		IPA{
			MainClassification: BilabialPlosive,
			IsVowel:            false,
			Diacritic:          DiacriticAspirated,
			Suprasegmental:     SuprasegmentalNone,
			Tone:               ToneNone,
			Accent:             AccentNone,
			Base:               "p",
			Notation:           "p_h",
		},
		IPA{
			MainClassification: BilabialPlosive,
			IsVowel:            false,
			Diacritic:          DiacriticAspirated,
			Suprasegmental:     SuprasegmentalNone,
			Tone:               ToneNone,
			Accent:             AccentNone,
			Base:               "b",
			Notation:           "b_h",
		},
		IPA{
			MainClassification: BilabialPlosive,
			IsVowel:            false,
			Diacritic:          DiacriticNone,
			Suprasegmental:     SuprasegmentalNone,
			Tone:               ToneNone,
			Accent:             AccentNone,
			Base:               "p",
			Notation:           "p",
		},
	},
	'ㅃ': {
		IPA{
			MainClassification: BilabialPlosive,
			IsVowel:            false,
			Diacritic:          DiacriticBreathyVoiced,
			Suprasegmental:     SuprasegmentalNone,
			Tone:               ToneNone,
			Accent:             AccentNone,
			Base:               "p",
			Notation:           "p_t",
		},
		IPA{
			MainClassification: BilabialPlosive,
			IsVowel:            false,
			Diacritic:          DiacriticBreathyVoiced,
			Suprasegmental:     SuprasegmentalNone,
			Tone:               ToneNone,
			Accent:             AccentNone,
			Base:               "p",
			Notation:           "p_t",
		},
		NoSound,
	},
	'ㅁ': {
		IPA{
			MainClassification: BilabialNasal,
			IsVowel:            false,
			Diacritic:          DiacriticNone,
			Suprasegmental:     SuprasegmentalNone,
			Tone:               ToneNone,
			Accent:             AccentNone,
			Base:               "m",
			Notation:           "m",
		},
		IPA{
			MainClassification: BilabialNasal,
			IsVowel:            false,
			Diacritic:          DiacriticNone,
			Suprasegmental:     SuprasegmentalNone,
			Tone:               ToneNone,
			Accent:             AccentNone,
			Base:               "m",
			Notation:           "m",
		},
		IPA{
			MainClassification: BilabialNasal,
			IsVowel:            false,
			Diacritic:          DiacriticNone,
			Suprasegmental:     SuprasegmentalNone,
			Tone:               ToneNone,
			Accent:             AccentNone,
			Base:               "m",
			Notation:           "m",
		},
	},
	'ㅎ': {
		IPA{
			MainClassification: BilabialFricative,
			IsVowel:            false,
			Diacritic:          DiacriticNone,
			Suprasegmental:     SuprasegmentalNone,
			Tone:               ToneNone,
			Accent:             AccentNone,
			Base:               "p\\",
			Notation:           "p\\",
		},
		IPA{
			MainClassification: BilabialFricative,
			IsVowel:            false,
			Diacritic:          DiacriticNone,
			Suprasegmental:     SuprasegmentalNone,
			Tone:               ToneNone,
			Accent:             AccentNone,
			Base:               "B",
			Notation:           "B",
		},
		IPA{
			MainClassification: AlveolarPlosive,
			IsVowel:            false,
			Diacritic:          DiacriticNone,
			Suprasegmental:     SuprasegmentalNone,
			Tone:               ToneNone,
			Accent:             AccentNone,
			Base:               "t",
			Notation:           "t",
		},
	},
}

var defaultVowelPronunciation = map[rune]IPA{
	'ㅏ': IPA{
		MainClassification: Vowel,
		IsVowel:            true,
		Diacritic:          DiacriticNone,
		Suprasegmental:     SuprasegmentalNone,
		Tone:               ToneNone,
		Accent:             AccentNone,
		Base:               "a",
		Notation:           "a",
	},
	'ㅑ': IPA{
		MainClassification: Vowel,
		IsVowel:            true,
		Diacritic:          DiacriticApical,
		Suprasegmental:     SuprasegmentalNone,
		Tone:               ToneNone,
		Accent:             AccentNone,
		Base:               "ja",
		Notation:           "ja_a",
	},
}

//Analyze Pronunciation
func Analyze(word []rune) {
	kword := make([]hangul.KChar, len(word))
	kipa := make([]IPA, 0)
	for i := 0; i < len(word); i++ {
		kword[i], _ = hangul.CharSplit(word[i])
		if kword[i].IsFull {
			if i == 0 {
				if kword[i].HasChoSung {
					ipa, ok := defaultConsonantPronunciation[kword[i].ChoSung]
					if ok {
						kipa = append(kipa, ipa.First)
					}
				}
				if kword[i].HasJungSung {
					ipa, ok := defaultVowelPronunciation[kword[i].JongSung]
					if ok {
						kipa = append(kipa, ipa)
					}
				}
				if kword[i].HasJongSung {
					ipa, ok := defaultConsonantPronunciation[kword[i].JongSung]
					if ok {
						kipa = append(kipa, ipa.End)
					}
				}
			} else {
				if kword[i].HasChoSung {
					ipa, ok := defaultConsonantPronunciation[kword[i].ChoSung]
					if ok {
						kipa = append(kipa, ipa.Middle)
					}
				}
				if kword[i].HasJungSung {
					ipa, ok := defaultVowelPronunciation[kword[i].JongSung]
					if ok {
						kipa = append(kipa, ipa)
					}
				}
				if kword[i].HasJongSung {
					ipa, ok := defaultConsonantPronunciation[kword[i].JongSung]
					if ok {
						kipa = append(kipa, ipa.End)
					}
				}
			}
		}
	}
}
