package vspeach

//Caution All IPA markings are based on X-SAMPA standards.

//Muted : Muted sound
const Muted = 0

//BilabialPlosive : Bilabial Plosive like (p, b)
const BilabialPlosive = 1

//BilabialNasal : Bilabial Nasal like (m)
const BilabialNasal = 2

//BilabialTrill : Bilabial Trill like (B\)
const BilabialTrill = 3

//BilabialFricative : Bilabial Fricative like (p\, B)
const BilabialFricative = 4

//DiacriticNone : Diacritic None
const DiacriticNone = -1

//DiacriticAspirated : Diacritic Aspirated [_h]
const DiacriticAspirated = 3

//DiacriticBreathyVoiced : Diacritic Breathy voiced [_t]
const DiacriticBreathyVoiced = 13

//ToneNone : Tone None
const ToneNone = -1

//AccentNone : Accent None
const AccentNone = -1

//SuprasegmentalNone : Suprasegmental None
const SuprasegmentalNone = -1

//IPA : International Phonetic Alphabet
type IPA struct {
	MainClassification int
	IsVowel            bool
	Diacritic          int
	Tone               int
	Accent             int
	Suprasegmental     int
	Base               string
	Notation           string
}

//NoSound No sound is produced, but a virtually added sound for program operation
var NoSound = IPA{
	MainClassification: Muted,
	IsVowel:            false,
	Diacritic:          DiacriticNone,
	Tone:               ToneNone,
	Accent:             AccentNone,
	Base:               "",
	Notation:           "",
}
