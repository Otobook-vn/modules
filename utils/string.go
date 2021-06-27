package utils

import (
	"encoding/json"
	"regexp"
	"strings"
	"unicode"

	"github.com/gosimple/slug"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"gorm.io/datatypes"
)

// Code ...
const (
	ReferralCodeCharacters = "abcdefghjklmnpqrstuvwxyz123456789" // No i, o, and 0
	ReferralCodeLength     = 7

	OTPCharacters = "123456789"
	OTPLength     = 6
)

// RemoveDiacritics ...
func RemoveDiacritics(s string) string {
	if s != "" {
		s = strings.ToLower(s)
		s = replaceStringWithRegex(s, `đ`, "d")
		t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
		result, _, _ := transform.String(t, s)
		result = replaceStringWithRegex(result, `[^a-zA-Z0-9\s]`, "")

		return result
	}
	return ""
}

// replaceStringWithRegex ...
func replaceStringWithRegex(src string, regex string, replaceText string) string {
	reg := regexp.MustCompile(regex)
	return reg.ReplaceAllString(src, replaceText)
}

// TransformKeywordToSearchString ...
func TransformKeywordToSearchString(keyword string) string {
	s := strings.TrimSpace(keyword)
	s = RemoveDiacritics(s)
	s = strings.ReplaceAll(s, " ", "&")
	return s + ":*" // For prefix search
}

// TransformKeywordToSearchILike ...
func TransformKeywordToSearchILike(keyword string) string {
	s := strings.TrimSpace(keyword)
	return "%" + strings.ToLower(s) + "%"
}

// FormatPhoneFull ...
func FormatPhoneFull(countryCode, number string) string {
	// Country code always format as: +{country_code}, ex: +84
	// Convert number to 9 chars, without "0" at start
	numberLength := len(number)

	// Remove char 0
	if numberLength == 10 && string(number[0]) == "0" {
		number = number[1:]
	}

	// Join and return
	return countryCode + number
}

// RandomReferralCode ...
func RandomReferralCode() string {
	var code string
	var length = len(ReferralCodeCharacters)
	for i := 0; i < ReferralCodeLength; i++ {
		// Random character index
		randIndex := RandomIntBetweenRange(0, length)
		code += string(ReferralCodeCharacters[randIndex])
	}
	return code
}

// ToLowercase ...
func ToLowercase(s string) string {
	return strings.ToLower(s)
}

// RandomOTPCode ...
func RandomOTPCode() string {
	var code string
	var length = len(OTPCharacters)
	for i := 0; i < OTPLength; i++ {
		// Random character index
		randIndex := RandomIntBetweenRange(0, length)
		code += string(OTPCharacters[randIndex])
	}
	return code
}

// ToSlug ...
func ToSlug(s string) string {
	return slug.Make(s)
}

// GetNameStringWithLang ...
func GetNameStringWithLang(name datatypes.JSON, lang string) string {
	byteData, err := name.MarshalJSON()
	if err != nil {
		return ""
	}
	var mapData map[string]string
	if err = json.Unmarshal(byteData, &mapData); err != nil {
		return ""
	}
	return mapData[lang]
}

// JoinStringWithUnderscore ...
func JoinStringWithUnderscore(str string) string {
	// Lowercase first, then split string with "space char" and join with "_"
	// E.g: input = "Customer care" -> output = "customer_care"

	parts := strings.SplitN(strings.ToLower(str), " ", -1)
	return strings.Join(parts, "_")
}
