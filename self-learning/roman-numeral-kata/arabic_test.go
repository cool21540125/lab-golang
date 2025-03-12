package romannumeralkata

import (
	"fmt"
	"testing"
)

func TestConvertToRoman(t *testing.T) {
	cases := []RomanNumeral{
		{Num: 1, Symbolic: "I"},
		{Num: 4, Symbolic: "IV"},
		{Num: 5, Symbolic: "V"},
		{Num: 6, Symbolic: "VI"},
		{Num: 9, Symbolic: "IX"},
		{Num: 10, Symbolic: "X"},
		{Num: 11, Symbolic: "XI"},
		{Num: 13, Symbolic: "XIII"},
		{Num: 15, Symbolic: "XV"},
		{Num: 19, Symbolic: "XIX"},
		{Num: 40, Symbolic: "XL"},
		{Num: 47, Symbolic: "XLVII"},
		{Num: 49, Symbolic: "XLIX"},
		{Num: 50, Symbolic: "L"},
		{Num: 100, Symbolic: "C"},
		{Num: 400, Symbolic: "CD"},
		{Num: 500, Symbolic: "D"},
		{Num: 798, Symbolic: "DCCXCVIII"},
		{Num: 900, Symbolic: "CM"},
		{Num: 999, Symbolic: "CMXCIX"},
		{Num: 1000, Symbolic: "M"},
		{Num: 1006, Symbolic: "MVI"},
		{Num: 1050, Symbolic: "ML"},
		{Num: 1080, Symbolic: "MLXXX"},
		{Num: 1090, Symbolic: "MXC"},
		{Num: 2014, Symbolic: "MMXIV"},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q -> %d", test.Symbolic, test.Num), func(t *testing.T) {
			got := ConvertToRoman(test.Num)
			if test.Symbolic != got {
				t.Errorf("%d expected %q, but got %q", test.Num, test.Symbolic, got)
			}
		})
	}
}
