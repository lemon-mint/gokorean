package vspeach

//BilabialPlosive : Bilabial Plosive like (p, b)
const BilabialPlosive = 0

//BilabialNasal : Bilabial Nasal like (m)
const BilabialNasal = 1

//BilabialTrill : Bilabial Trill like (Β)
const BilabialTrill = 2

//DiacriticNone : Diacritic None
const DiacriticNone = -1

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
	Notation           string
}
