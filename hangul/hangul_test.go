package hangul

import (
	"reflect"
	"testing"
)

func TestCharSplit(t *testing.T) {
	type args struct {
		char rune
	}
	tests := []struct {
		name    string
		args    args
		want    KChar
		wantErr bool
	}{
		{
			"split 가",
			args{
				'가',
			},
			KChar{
				Char:        '가',
				IsHangul:    true,
				IsComponent: false,
				IsFull:      true,
				HasChoSung:  true,
				ChoSung:     'ㄱ',
				HasJungSung: true,
				JungSung:    'ㅏ',
				HasJongSung: false,
				JongSung:    ' ',
			},
			false,
		},
		{
			"split 뵓",
			args{
				'뵓',
			},
			KChar{
				Char:        '뵓',
				IsHangul:    true,
				IsComponent: false,
				IsFull:      true,
				HasChoSung:  true,
				ChoSung:     'ㅂ',
				HasJungSung: true,
				JungSung:    'ㅚ',
				HasJongSung: true,
				JongSung:    'ㄼ',
			},
			false,
		},
		{
			"split ㄱ",
			args{
				'ㄱ',
			},
			KChar{
				Char:        'ㄱ',
				IsHangul:    true,
				IsComponent: true,
				IsFull:      false,
				IsVowel:     false,
				IsConsonant: true,
				HasChoSung:  false,
				ChoSung:     0,
				HasJungSung: false,
				JungSung:    0,
				HasJongSung: false,
				JongSung:    0,
			},
			false,
		},
		{
			"split ㅗ",
			args{
				'ㅗ',
			},
			KChar{
				Char:        'ㅗ',
				IsHangul:    true,
				IsComponent: true,
				IsFull:      false,
				IsVowel:     true,
				IsConsonant: false,
				HasChoSung:  false,
				ChoSung:     0,
				HasJungSung: false,
				JungSung:    0,
				HasJongSung: false,
				JongSung:    0,
			},
			false,
		},
		{
			"split 라",
			args{
				'라',
			},
			KChar{
				Char:        '라',
				IsHangul:    true,
				IsComponent: false,
				IsFull:      true,
				HasChoSung:  true,
				ChoSung:     'ㄹ',
				HasJungSung: true,
				JungSung:    'ㅏ',
				HasJongSung: false,
				JongSung:    ' ',
			},
			false,
		},
		{
			"split 로",
			args{
				'로',
			},
			KChar{
				Char:        '로',
				IsHangul:    true,
				IsComponent: false,
				IsFull:      true,
				HasChoSung:  true,
				ChoSung:     'ㄹ',
				HasJungSung: true,
				JungSung:    'ㅗ',
				HasJongSung: false,
				JongSung:    ' ',
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CharSplit(tt.args.char)
			if (err != nil) != tt.wantErr {
				t.Errorf("CharSplit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CharSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
