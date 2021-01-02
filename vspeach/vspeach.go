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

func analyze(word []rune) {
	kword := make([]hangul.KChar, len(word))
	for i := 0; i < len(word); i++ {
		kword[i], _ = hangul.CharSplit(word[i])

	}
}
