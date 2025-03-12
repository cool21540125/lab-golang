package romannumeralkata

import "strings"

type RomanNumeral struct {
	Num      int
	Symbolic string
}

func ConvertToRoman(num int) string {
	allRomanNumerals := []RomanNumeral{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result strings.Builder
	for _, c := range allRomanNumerals {
		for num >= c.Num {
			result.WriteString(c.Symbolic)
			num -= c.Num
		}
	}

	return result.String()
}
