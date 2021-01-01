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
	'ㅣ': KVowel{
		IsDoubleVowel: false,
		TonguePos:     frontTongueVowel,
		TongueShape:   flatVowel,
		VowelHeight:   highVowel,
	},
	'ㅔ': KVowel{
		IsDoubleVowel: false,
		TonguePos:     frontTongueVowel,
		TongueShape:   flatVowel,
		VowelHeight:   midVowel,
	},
	'ㅐ': KVowel{
		IsDoubleVowel: false,
		TonguePos:     frontTongueVowel,
		TongueShape:   flatVowel,
		VowelHeight:   lowVowel,
	},
	'ㅟ': KVowel{
		IsDoubleVowel: false,
		TonguePos:     frontTongueVowel,
		TongueShape:   roundedVowel,
		VowelHeight:   highVowel,
	},
	'ㅚ': KVowel{
		IsDoubleVowel: false,
		TonguePos:     frontTongueVowel,
		TongueShape:   roundedVowel,
		VowelHeight:   midVowel,
	},
	'ㅡ': KVowel{
		IsDoubleVowel: false,
		TonguePos:     backTongueVowel,
		TongueShape:   flatVowel,
		VowelHeight:   highVowel,
	},
	'ㅓ': KVowel{
		IsDoubleVowel: false,
		TonguePos:     backTongueVowel,
		TongueShape:   flatVowel,
		VowelHeight:   midVowel,
	},
	'ㅏ': KVowel{
		IsDoubleVowel: false,
		TonguePos:     backTongueVowel,
		TongueShape:   flatVowel,
		VowelHeight:   lowVowel,
	},
	'ㅜ': KVowel{
		IsDoubleVowel: false,
		TonguePos:     backTongueVowel,
		TongueShape:   roundedVowel,
		VowelHeight:   highVowel,
	},
	'ㅗ': KVowel{
		IsDoubleVowel: false,
		TonguePos:     backTongueVowel,
		TongueShape:   roundedVowel,
		VowelHeight:   midVowel,
	},
	'ㅑ': KVowel{
		IsDoubleVowel: true,
	},
	'ㅒ': KVowel{
		IsDoubleVowel: true,
	},
	'ㅕ': KVowel{
		IsDoubleVowel: true,
	},
	'ㅖ': KVowel{
		IsDoubleVowel: true,
	},
	'ㅘ': KVowel{
		IsDoubleVowel: true,
	},
	'ㅙ': KVowel{
		IsDoubleVowel: true,
	},
	'ㅛ': KVowel{
		IsDoubleVowel: true,
	},
	'ㅝ': KVowel{
		IsDoubleVowel: true,
	},
	'ㅞ': KVowel{
		IsDoubleVowel: true,
	},
	'ㅠ': KVowel{
		IsDoubleVowel: true,
	},
	'ㅢ': KVowel{
		IsDoubleVowel: true,
	},
}
