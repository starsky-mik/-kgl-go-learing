package roman_numerals

import (
	"fmt"
	. "github.com/starsky-mik/kgd-go-learing/internal/pkg/testing_tools"
	. "github.com/starsky-mik/kgd-go-learing/lesson/roman-numerals"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{90, "XC"},
	{100, "C"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestConvertToRoman(t *testing.T) {
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("convert %d to %s", testCase.Arabic, testCase.Roman), func(t *testing.T) {
			got := ConvertToRoman(testCase.Arabic)
			want := testCase.Roman

			AssertEquals(t, got, want)
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, testCase := range cases {
		t.Run(fmt.Sprintf("convert %s to %d", testCase.Roman, testCase.Arabic), func(t *testing.T) {
			got := ConvertToArabic(testCase.Roman)
			want := testCase.Arabic

			AssertEquals(t, got, want)
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}

		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
