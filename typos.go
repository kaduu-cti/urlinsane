// Copyright © 2018 rangertaha rangertaha@gmail.com
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package urlinsane

import (
	"fmt"
	"strings"
)

// The registry for typos functions
var TREGISTRY = make(map[string][]Typo)

var missingDot = Typo{
	Code:        "MD",
	Name:        "Missing Dot",
	Description: "Missing Dot is created by omitting a dot from the domain.",
	Exec:        missingDotFunc,
}
var missingDashes = Typo{
	Code:        "MDS",
	Name:        "Missing Dashes",
	Description: "Missing Dashes is created by omitting a dash from the domain.",
	Exec:        missingDashFunc,
}
var stripDashes = Typo{
	Code:        "SD",
	Name:        "Strip Dashes",
	Description: "Strip Dashes is created by omitting a dot from the domain",
	Exec:        stripDashesFunc,
}
var characterOmission = Typo{
	Code:        "CO",
	Name:        "Character Omission",
	Description: "Character Omission Omitting a character from the domain.",
	Exec:        characterOmissionFunc,
}
var characterSwap = Typo{
	Code:        "CS",
	Name:        "Character Swap",
	Description: "Character Swap Swapping two consecutive characters in a domain",
	Exec:        characterSwapFunc,
}
var adjacentCharacterSubstitution = Typo{
	Code:        "ACS",
	Name:        "Adjacent Character Substitution",
	Description: "Adjacent Character Substitution replaces adjacent characters",
	Exec:        adjacentCharacterSubstitutionFunc,
}
var adjacentCharacterInsertion = Typo{
	Code:        "ACI",
	Name:        "Adjacent Character Insertion",
	Description: "Adjacent Character Insertion inserts adjacent character ",
	Exec:        adjacentCharacterInsertionFunc,
}
var homoglyphs = Typo{
	Code:        "HG",
	Name:        "Homoglyphs",
	Description: "Homoglyphs replaces characters with characters that look similar",
	Exec:        homoglyphFunc,
}
var singularPluralise = Typo{
	Code:        "SP",
	Name:        "Singular Pluralise",
	Description: "Singular Pluralise creates a singular domain plural and vice versa",
	Exec:        singularPluraliseFunc,
}

var characterRepeat = Typo{
	Code:        "CR",
	Name:        "Character Repeat",
	Description: "Character Repeat Repeats a character of the domain name twice",
	Exec:        characterRepeatFunc,
}

var doubleCharacterReplacement = Typo{
	Code:        "DCR",
	Name:        "Double Character Replacement",
	Description: "Double Character Replacement repeats a character twice.",
	Exec:        doubleCharacterReplacementFunc,
}
var commonMisspellings = Typo{
	Code:        "CM",
	Name:        "Common Misspellings",
	Description: "Common Misspellings are created from common misspellings",
	Exec:        commonMisspellingsFunc,
}
var homophones = Typo{
	Code:        "HP",
	Name:        "Homophones",
	Description: "Homophones Typos are created from sets of words that sound the same",
	Exec:        homophonesFunc,
}

var vowelSwapping = Typo{
	Code:        "VS",
	Name:        "Vowel Swapping",
	Description: "Vowel Swapping is created by swaps vowels",
	Exec:        vowelSwappingFunc,
}

var bitFlipping = Typo{
	Code:        "BF",
	Name:        "Bit Flipping",
	Description: "Bitsquatting relies on random bit-errors to redirect connections",
	Exec:        bitsquattingFunc,
}

var wrongTopLevelDomain = Typo{
	Code:        "WTLD",
	Name:        "Wrong TLD",
	Description: "Wrong Top Level Domain",
	Exec:        wrongTopLevelDomainFunc,
}

var wrongSecondLevelDomain = Typo{
	Code:        "WSLD",
	Name:        "Wrong SLD",
	Description: "Wrong Second Level Domain",
	Exec:        wrongSecondLevelDomainFunc,
}

var numeralSwap = Typo{
	Code:        "NS",
	Name:        "Numeral Swap",
	Description: "Numeral Swap numbers, words and vice versa",
	Exec:        numeralSwapFunc,
}

func init() {
	TRegister("MD", missingDot)
	TRegister("MDS", missingDashes)
	TRegister("CO", characterOmission)
	TRegister("CS", characterSwap)
	TRegister("ACS", adjacentCharacterSubstitution)
	TRegister("ACI", adjacentCharacterInsertion)
	TRegister("CR", characterRepeat)
	TRegister("DCR", doubleCharacterReplacement)

	TRegister("SD", stripDashes)
	TRegister("SP", singularPluralise)
	TRegister("CM", commonMisspellings)
	TRegister("VS", vowelSwapping)
	TRegister("HG", homoglyphs)
	TRegister("WTLD", wrongTopLevelDomain)
	TRegister("WSLD", wrongSecondLevelDomain)
	TRegister("HP", homophones)
	TRegister("BF", bitFlipping)
	TRegister("NS", numeralSwap)

	TRegister("ALL",
		missingDot,
		missingDashes,
		characterOmission,
		characterSwap,
		adjacentCharacterSubstitution,
		adjacentCharacterInsertion,
		characterRepeat,
		doubleCharacterReplacement,
		missingDashes,
		stripDashes,
		singularPluralise,
		commonMisspellings,
		vowelSwapping,
		homoglyphs,
		wrongTopLevelDomain,
		wrongSecondLevelDomain,
		homophones,
		bitFlipping,
		numeralSwap,
	)

}

// missingDotFunc typos are created by omitting a dot from the domain. For example, wwwgoogle.com and www.googlecom
func missingDotFunc(tc TypoConfig) (results []TypoConfig) {
	for _, str := range missingCharFunc(tc.Domain.Domain, ".") {
		dm := Domain{tc.Domain.Subdomain, str, tc.Domain.Suffix}
		results = append(results, TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})
	}
	return results
}

// missingDashFunc typos are created by omitting a dash from the domain. For example, www.a-bc.com and www.ab-c.com
func missingDashFunc(tc TypoConfig) (results []TypoConfig) {
	for _, str := range missingCharFunc(tc.Domain.Domain, "-") {
		if tc.Domain.Domain != str {
			dm := Domain{tc.Domain.Subdomain, str, tc.Domain.Suffix}
			results = append(results, TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})
		}
	}
	return results
}

// characterOmissionFunc typos are when one character in the original domain name is omitted.
// For example: www.exmple.com
func characterOmissionFunc(tc TypoConfig) (results []TypoConfig) {
	for i := range tc.Domain.Domain {
		if i <= len(tc.Domain.Domain)-1 {
			domain := fmt.Sprint(
				tc.Domain.Domain[:i],
				tc.Domain.Domain[i+1:],
			)
			if tc.Domain.Domain != domain {
				dm := Domain{tc.Domain.Subdomain, domain, tc.Domain.Suffix}
				results = append(results,
					TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})
			}
		}
	}
	return results
}

// characterSwapFunc typos are when two consecutive characters are swapped in the original domain name.
// Example: www.examlpe.com
func characterSwapFunc(tc TypoConfig) (results []TypoConfig) {
	for i := range tc.Domain.Domain {
		if i <= len(tc.Domain.Domain)-2 {
			domain := fmt.Sprint(
				tc.Domain.Domain[:i],
				string(tc.Domain.Domain[i+1]),
				string(tc.Domain.Domain[i]),
				tc.Domain.Domain[i+2:],
			)
			if tc.Domain.Domain != domain {
				dm := Domain{tc.Domain.Subdomain, domain, tc.Domain.Suffix}
				results = append(results,
					TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})
			}
		}
	}
	return results
}

// adjacentCharacterSubstitutionFunc typos are when characters are replaced in the original domain name by their
// adjacent ones on a specific keyboard layout, e.g., www.ezample.com, where “x” was replaced by the adjacent
// character “z” in a the QWERTY keyboard layout.
func adjacentCharacterSubstitutionFunc(tc TypoConfig) (results []TypoConfig) {
	for _, keyboard := range tc.Keyboards {
		for i, char := range tc.Domain.Domain {
			for _, key := range keyboard.Adjacent(string(char)) {
				domain := tc.Domain.Domain[:i] + string(key) + tc.Domain.Domain[i+1:]
				dm := Domain{tc.Domain.Subdomain, domain, tc.Domain.Suffix}
				results = append(results,
					TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})
			}
		}
	}
	return
}

// adjacentCharacterInsertionFunc are created by inserting letters adjacent of each letter. For example, www.googhle.com
// and www.goopgle.com
func adjacentCharacterInsertionFunc(tc TypoConfig) (results []TypoConfig) {
	for _, keyboard := range tc.Keyboards {
		for i, char := range tc.Domain.Domain {
			for _, key := range keyboard.Adjacent(string(char)) {
				d1 := tc.Domain.Domain[:i] + string(key) + string(char) + tc.Domain.Domain[i+1:]
				dm1 := Domain{tc.Domain.Subdomain, d1, tc.Domain.Suffix}
				results = append(results,
					TypoConfig{dm1, tc.Keyboards, tc.Languages, tc.Typo})

				d2 := tc.Domain.Domain[:i] + string(char) + string(key) + tc.Domain.Domain[i+1:]
				dm2 := Domain{tc.Domain.Subdomain, d2, tc.Domain.Suffix}
				results = append(results,
					TypoConfig{dm2, tc.Keyboards, tc.Languages, tc.Typo})
			}
		}
	}
	return
}

// characterRepeatFunc are created by repeating a letter of the domain name.
// Example, www.ggoogle.com and www.gooogle.com
func characterRepeatFunc(tc TypoConfig) (results []TypoConfig) {
	for i := range tc.Domain.Domain {
		if i <= len(tc.Domain.Domain) {
			domain := fmt.Sprint(
				tc.Domain.Domain[:i],
				string(tc.Domain.Domain[i]),
				string(tc.Domain.Domain[i]),
				tc.Domain.Domain[i+1:],
			)
			if tc.Domain.Domain != domain {
				dm := Domain{tc.Domain.Subdomain, domain, tc.Domain.Suffix}
				results = append(results,
					TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})
			}
		}
	}
	return results
}

// doubleCharacterReplacementFunc are created by replacing identical, consecutive
// letters of the domain name with adjacent letters on the keyboard.
// For example, www.gppgle.com and www.giigle.com
func doubleCharacterReplacementFunc(tc TypoConfig) (results []TypoConfig) {
	for _, keyboard := range tc.Keyboards {
		for i, char := range tc.Domain.Domain {
			if i < len(tc.Domain.Domain)-1 {
				if tc.Domain.Domain[i] == tc.Domain.Domain[i+1] {
					for _, key := range keyboard.Adjacent(string(char)) {
						domain := tc.Domain.Domain[:i] + string(key) + string(key) + tc.Domain.Domain[i+2:]
						dm := Domain{tc.Domain.Subdomain, domain, tc.Domain.Suffix}
						results = append(results,
							TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})
					}
				}
			}
		}
	}
	return
}

// stripDashesFunc typos are created by omitting a dot from the domain.
// For example, www.a-b-c.com becomes www.abc.com
func stripDashesFunc(tc TypoConfig) (results []TypoConfig) {
	for _, str := range replaceCharFunc(tc.Domain.Domain, "-", "") {
		if tc.Domain.Domain != str {
			dm := Domain{tc.Domain.Subdomain, str, tc.Domain.Suffix}
			results = append(results, TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})
		}
	}
	return
}

// singularPluraliseFunc are created by making a singular domain plural and
// vice versa. For example, www.google.com becomes www.googles.com and
// www.games.co.nz becomes www.game.co.nz
func singularPluraliseFunc(tc TypoConfig) (results []TypoConfig) {
	var domain string
	if strings.HasSuffix(tc.Domain.Domain, "s") {
		domain = strings.TrimSuffix(tc.Domain.Domain, "s")
	} else {
		domain = tc.Domain.Domain + "s"
	}
	dm := Domain{tc.Domain.Subdomain, domain, tc.Domain.Suffix}
	results = append(results, TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})
	return
}

// CcommonMisspellingsFunc are created with common misspellings in the given
// language. For example, www.youtube.com becomes www.youtub.com and
// www.abseil.com becomes www.absail.com
func commonMisspellingsFunc(tc TypoConfig) (results []TypoConfig) {
	for _, keyboard := range tc.Keyboards {
		for _, word := range keyboard.Language.SimilarSpellings(tc.Domain.Domain) {
			dm := Domain{tc.Domain.Subdomain, word, tc.Domain.Suffix}
			results = append(results, TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})

		}
	}
	return
}

// vowelSwappingFunc swaps vowels within the domain name except for the first letter.
// For example, www.google.com becomes www.gaagle.com.
func vowelSwappingFunc(tc TypoConfig) (results []TypoConfig) {
	for _, keyboard := range tc.Keyboards {
		for _, vchar := range keyboard.Language.Vowels {
			if strings.Contains(tc.Domain.Domain, vchar) {
				for _, vvchar := range keyboard.Language.Vowels {
					new := strings.Replace(tc.Domain.Domain, vchar, vvchar, -1)
					if new != tc.Domain.Domain {
						dm := Domain{tc.Domain.Subdomain, new, tc.Domain.Suffix}
						results = append(results, TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})
					}

				}

			}
		}
	}
	return
}

// homophonesFunc are created from sets of words that sound the same when spoken.
// For example, www.base.com becomes www .bass.com.
func homophonesFunc(tc TypoConfig) (results []TypoConfig) {
	for _, keyboard := range tc.Keyboards {
		for _, word := range keyboard.Language.SimilarSounds(tc.Domain.Domain) {
			dm := Domain{tc.Domain.Subdomain, word, tc.Domain.Suffix}
			results = append(results, TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})

		}
	}
	return
}

// homoglyphFunc when one or more characters that look similar to another
// character but are different are called homogylphs. An example is that the
// lower case l looks similar to the numeral one, e.g. l vs 1. For example,
// google.com becomes goog1e.com.
func homoglyphFunc(tc TypoConfig) (results []TypoConfig) {
	for i, char := range tc.Domain.Domain {
		// Check the alphabet of the language associated with the keyboard for
		// homoglyphs
		for _, keyboard := range tc.Keyboards {
			for _, kchar := range keyboard.Language.SimilarChars(string(char)) {
				domain := fmt.Sprint(tc.Domain.Domain[:i], kchar, tc.Domain.Domain[i+1:])
				if tc.Domain.Domain != domain {
					dm := Domain{tc.Domain.Subdomain, domain, tc.Domain.Suffix}
					results = append(results,
						TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})
				}
			}
		}
		// Check languages given with the (-l --language) CLI options for homoglyphs.
		for _, language := range tc.Languages {
			for _, lchar := range language.SimilarChars(string(char)) {
				domain := fmt.Sprint(tc.Domain.Domain[:i], lchar, tc.Domain.Domain[i+1:])
				if tc.Domain.Domain != domain {
					dm := Domain{tc.Domain.Subdomain, domain, tc.Domain.Suffix}
					results = append(results,
						TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})
				}
			}
		}

	}
	return results
}

// wrongTopLevelDomain for example, www.trademe.co.nz becomes www.trademe.co.nz
// and www.google.com becomes www.google.org. uses the 19 most common top level
// domains.
func wrongTopLevelDomainFunc(tc TypoConfig) (results []TypoConfig) {
	tc.Domain.Domain = tc.Domain.Domain + ":S"
	results = append(results, tc)
	return results
}

// wrongSecondLevelDomain uses an alternate, valid second level domain for the
// top level domain. For example, www.trademe.co.nz becomes www.trademe.ac.nz
// and www.trademe.iwi.nz
func wrongSecondLevelDomainFunc(tc TypoConfig) (results []TypoConfig) {
	tc.Domain.Domain = tc.Domain.Domain + ":S"
	results = append(results, tc)
	return results
}

// bitsquattingFunc relies on random bit- errors to redirect connections
// intended for popular domains
func bitsquattingFunc(tc TypoConfig) (results []TypoConfig) {
	// TOOO: need to improve.
	masks := []int{1, 2, 4, 8, 16, 32, 64, 128}
	charset := make(map[string][]string)
	for _, board := range tc.Keyboards {
		for _, alpha := range board.Language.Graphemes {
			for _, mask := range masks {
				new := int([]rune(alpha)[0]) ^ mask
				for _, a := range board.Language.Graphemes {
					if string(a) == string(new) {
						charset[string(alpha)] = append(charset[string(alpha)], string(new))
					}
				}

			}
		}
	}

	for d, dchar := range tc.Domain.Domain {
		for _, char := range charset[string(dchar)] {

			dnew := tc.Domain.Domain[:d] + string(char) + tc.Domain.Domain[d+1:]
			dm := Domain{tc.Domain.Subdomain, dnew, tc.Domain.Suffix}
			results = append(results, TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})
		}
	}
	return
}

// missingCharFunc removes a character one at a time from the string.
// For example, wwwgoogle.com and www.googlecom
func missingCharFunc(str, character string) (results []string) {
	for i, char := range str {
		if character == string(char) {
			results = append(results, str[:i]+str[i+1:])
		}
	}
	return
}

// replaceCharFunc omits a character from the entire string.
// For example, www.a-b-c.com becomes www.abc.com
func replaceCharFunc(str, old, new string) (results []string) {
	results = append(results, strings.Replace(str, old, new, -1))
	return
}

// numeralSwapFunc are created by swapping numbers and corresponding words
func numeralSwapFunc(tc TypoConfig) (results []TypoConfig) {
	return
}

// TRegister
func TRegister(name string, typo ...Typo) {
	_, registered := TREGISTRY[strings.ToUpper(name)]
	if !registered {
		TREGISTRY[strings.ToUpper(name)] = typo
	}
}

// TRetrieve
func TRetrieve(strs ...string) (results []Typo) {
	for _, f := range strs {
		value, ok := TREGISTRY[strings.ToUpper(f)]
		if ok {
			results = append(results, value...)
		}
	}
	if len(strs) == 0 {
		for _, value := range TREGISTRY {
			results = append(results, value...)
		}
	}
	return
}