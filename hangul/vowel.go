package hangul

const highVowel uint16 = 0
const midVowel uint16 = 1
const lowVowel uint16 = 2

const roundedVowel uint16 = 10
const flatVowel uint16 = 20

const frontTongueVowel uint16 = 100
const backTongueVowel uint16 = 200

//KVowel Classification of vowels for pronunciation analysis
type KVowel struct {
	IsDoubleVowel bool
	TonguePos     uint16
	TongueShape   uint16
	VowelHeight   uint16
}

var vowels = map[rune]KVowel{
	'ㅣ': {
		IsDoubleVowel: false,
		TonguePos:     frontTongueVowel,
		TongueShape:   flatVowel,
		VowelHeight:   highVowel,
	},
	'ㅔ': {
		IsDoubleVowel: false,
		TonguePos:     frontTongueVowel,
		TongueShape:   flatVowel,
		VowelHeight:   midVowel,
	},
	'ㅐ': {
		IsDoubleVowel: false,
		TonguePos:     frontTongueVowel,
		TongueShape:   flatVowel,
		VowelHeight:   lowVowel,
	},
	'ㅟ': {
		IsDoubleVowel: false,
		TonguePos:     frontTongueVowel,
		TongueShape:   roundedVowel,
		VowelHeight:   highVowel,
	},
	'ㅚ': {
		IsDoubleVowel: false,
		TonguePos:     frontTongueVowel,
		TongueShape:   roundedVowel,
		VowelHeight:   midVowel,
	},
	'ㅡ': {
		IsDoubleVowel: false,
		TonguePos:     backTongueVowel,
		TongueShape:   flatVowel,
		VowelHeight:   highVowel,
	},
	'ㅓ': {
		IsDoubleVowel: false,
		TonguePos:     backTongueVowel,
		TongueShape:   flatVowel,
		VowelHeight:   midVowel,
	},
	'ㅏ': {
		IsDoubleVowel: false,
		TonguePos:     backTongueVowel,
		TongueShape:   flatVowel,
		VowelHeight:   lowVowel,
	},
	'ㅜ': {
		IsDoubleVowel: false,
		TonguePos:     backTongueVowel,
		TongueShape:   roundedVowel,
		VowelHeight:   highVowel,
	},
	'ㅗ': {
		IsDoubleVowel: false,
		TonguePos:     backTongueVowel,
		TongueShape:   roundedVowel,
		VowelHeight:   midVowel,
	},
	'ㅑ': {
		IsDoubleVowel: true,
	},
	'ㅒ': {
		IsDoubleVowel: true,
	},
	'ㅕ': {
		IsDoubleVowel: true,
	},
	'ㅖ': {
		IsDoubleVowel: true,
	},
	'ㅘ': {
		IsDoubleVowel: true,
	},
	'ㅙ': {
		IsDoubleVowel: true,
	},
	'ㅛ': {
		IsDoubleVowel: true,
	},
	'ㅝ': {
		IsDoubleVowel: true,
	},
	'ㅞ': {
		IsDoubleVowel: true,
	},
	'ㅠ': {
		IsDoubleVowel: true,
	},
	'ㅢ': {
		IsDoubleVowel: true,
	},
}
