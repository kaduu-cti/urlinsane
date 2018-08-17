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

var MissingDot = Typo{
	Code:        "md",
	Name:        "Missing Dot",
	Description: "Omitting a dot from the domain. For example, wwwgoogle.com and www.googlecom",
	Exec:        MissingDotFunc,
}
var MissingDashes = Typo{
	Code: "mds",
	Name: "Missing Dashes",
	Description: "Omitting a dash from the domain. For example, www.a-b-c.com becomes " +
		  "www.ab-c.com and www.ab-c.com",
	Exec: MissingDashFunc,
}
var StripDashes = Typo{
	Code:        "sd",
	Name:        "Strip Dashes",
	Description: "Omitting a dot from the domain. For example, www.a-b-c.com becomes www.abc.com",
	Exec:        StripDashesFunc,
}
var CharacterOmission = Typo{
	Code:        "co",
	Name:        "Character Omission",
	Description: "Omitting a character from the domain. For example, www.gogle.com and googl.com",
	Exec:        CharacterOmissionFunc,
}
var CharacterSwap = Typo{
	Code:        "cp",
	Name:        "Character Swap",
	Description: "Swapping two consecutive characters in a domain. For example, www.ogogle.com",
	Exec:        CharacterSwapFunc,
}
var AdjacentCharacterSubstitution = Typo{
	Code:        "acs",
	Name:        "Adjacent Character Substitution",
	Description: "Replace characters with adjacent characters on the keyboard For example, www.foogle.com",
	Exec:        AdjacentCharacterSubstitutionFunc,
}
var AdjacentCharacterInsertion = Typo{
	Code:        "aci",
	Name:        "Adjacent Character Insertion",
	Description: "Created by inserting letters adjacent of each letter. Example: www.googhle.com, www.goopgle.com",
	Exec:        AdjacentCharacterInsertionFunc,
}
var Homoglyphs = Typo{
	Code:        "hg",
	Name:        "Homoglyphs",
	Description: "Replace characters with characters that look similar. For example, www.göögle.com",
	Exec:        HomoglyphFunc,
}
var SingularPluralise = Typo{
	Code:        "hg",
	Name:        "Singular Pluralise",
	Description: "Making a singular domain plural and vice versa. Example googles.com",
	Exec:        SingularPluraliseFunc,
}

var CharacterRepeat = Typo{
	Code:        "cr",
	Name:        "Character Repeat",
	Description: "Repeat a letter of the domain name twice. Example, www.ggoogle.com and www.gooogle.com",
	Exec:        CharacterRepeatFunc,
}

var DoubleCharacterReplacement = Typo{
	Code:        "dcr",
	Name:        "Double Character Replacement",
	Description: "Repeat a letter of the domain name twice. Example, www.ggoogle.com and www.gooogle.com",
	Exec:        DoubleCharacterReplacementFunc,
}
var CommonMisspellings = Typo{
	Code:        "cm",
	Name:        "Common Misspellings",
	Description: "Typos created from common misspellings in the given languages.",
	Exec:        CommonMisspellingsFunc,
}
var Homophones = Typo{
	Code:        "hp",
	Name:        "Homophones",
	Description: "Typos created from sets of words that sound the same when spoken.",
	Exec:        HomophonesFunc,
}

var VowelSwapping = Typo{
	Code:        "vs",
	Name:        "Vowel Swapping",
	Description: "Created by swaps vowels within the domain name except for the first letter.",
	Exec:        VowelSwappingFunc,
}

func init() {
	Register("md", MissingDot)
	Register("mds", MissingDashes)
	Register("co", CharacterOmission)
	Register("cp", CharacterSwap)
	Register("acs", AdjacentCharacterSubstitution)
	Register("aci", AdjacentCharacterInsertion)
	Register("cr", CharacterRepeat)
	Register("dcr", DoubleCharacterReplacement)

	Register("sd", StripDashes)
	Register("sp", SingularPluralise)
	Register("cm", CommonMisspellings)
	Register("vs", VowelSwapping)
	//Register("bf", BitFlipping)
	Register("hg", Homoglyphs)
	//Register("wtld", WrongTopLevelDomain)
	//Register("wsld", WrongSecondLevelDomain)
	Register("hp", Homophones)
	//Register("bs", Bitsquatting)

	Register("all",
		MissingDot,
		MissingDashes,
		CharacterOmission,
		CharacterSwap,
		AdjacentCharacterSubstitution,
		AdjacentCharacterInsertion,
		CharacterRepeat,
		DoubleCharacterReplacement,
		MissingDashes,
		StripDashes,
		SingularPluralise,
		CommonMisspellings,
		VowelSwapping,
		//BitFlipping,
		Homoglyphs,
		//WrongTopLevelDomain,
		//WrongSecondLevelDomain,
		Homophones,
		//Bitsquatting,
	)

}

// MissingDotFunc typos are created by omitting a dot from the domain. For example, wwwgoogle.com and www.googlecom
func MissingDotFunc(tc TypoConfig) (results []TypoConfig) {
	for _, str := range MissingCharFunc(tc.Domain.Domain, ".") {
		dm := Domain{tc.Domain.Subdomain, str, tc.Domain.Suffix}
		results = append(results, TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})
	}
	return results
}

// MissingDashFunc typos are created by omitting a dash from the domain. For example, www.a-bc.com and www.ab-c.com
func MissingDashFunc(tc TypoConfig) (results []TypoConfig) {
	for _, str := range MissingCharFunc(tc.Domain.Domain, "-") {
		if tc.Domain.Domain != str {
			dm := Domain{tc.Domain.Subdomain, str, tc.Domain.Suffix}
			results = append(results, TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})
		}
	}
	return results
}

// CharacterOmissionFunc typos are when one character in the original domain name is omitted.
// For example: www.exmple.com
func CharacterOmissionFunc(tc TypoConfig) (results []TypoConfig) {
	for i := range tc.Domain.Domain {
		if i <= len(tc.Domain.Domain) - 1 {
			domain := fmt.Sprint(
				tc.Domain.Domain[:i],
				tc.Domain.Domain[i + 1:],
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

// CharacterSwapFunc typos are when two consecutive characters are swapped in the original domain name.
// Example: www.examlpe.com
func CharacterSwapFunc(tc TypoConfig) (results []TypoConfig) {
	for i := range tc.Domain.Domain {
		if i <= len(tc.Domain.Domain) - 2 {
			domain := fmt.Sprint(
				tc.Domain.Domain[:i],
				string(tc.Domain.Domain[i + 1]),
				string(tc.Domain.Domain[i]),
				tc.Domain.Domain[i + 2:],
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

// AdjacentCharacterSubstitutionFunc typos are when characters are replaced in the original domain name by their
// adjacent ones on a specific keyboard layout, e.g., www.ezample.com, where “x” was replaced by the adjacent
// character “z” in a the QWERTY keyboard layout.
func AdjacentCharacterSubstitutionFunc(tc TypoConfig) (results []TypoConfig) {
	for _, keyboard := range tc.Keyboards {
		for i, char := range tc.Domain.Domain {
			for _, key := range keyboard.Adjacent(string(char)) {
				domain := tc.Domain.Domain[:i] + string(key) + tc.Domain.Domain[i + 1:]
				dm := Domain{tc.Domain.Subdomain, domain, tc.Domain.Suffix}
				results = append(results,
					TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})
			}
		}
	}
	return
}

// AdjacentCharacterInsertionFunc are created by inserting letters adjacent of each letter. For example, www.googhle.com
// and www.goopgle.com
func AdjacentCharacterInsertionFunc(tc TypoConfig) (results []TypoConfig) {
	for _, keyboard := range tc.Keyboards {
		for i, char := range tc.Domain.Domain {
			for _, key := range keyboard.Adjacent(string(char)) {
				d1 := tc.Domain.Domain[:i] + string(key) + string(char) + tc.Domain.Domain[i + 1:]
				dm1 := Domain{tc.Domain.Subdomain, d1, tc.Domain.Suffix}
				results = append(results,
					TypoConfig{dm1, tc.Keyboards, tc.Languages, tc.Typo})

				d2 := tc.Domain.Domain[:i] + string(char) + string(key) + tc.Domain.Domain[i + 1:]
				dm2 := Domain{tc.Domain.Subdomain, d2, tc.Domain.Suffix}
				results = append(results,
					TypoConfig{dm2, tc.Keyboards, tc.Languages, tc.Typo})
			}
		}
	}
	return
}

// CharacterRepeatFunc are created by repeating a letter of the domain name. Example, www.ggoogle.com and www.gooogle.com
func CharacterRepeatFunc(tc TypoConfig) (results []TypoConfig) {
	for i := range tc.Domain.Domain {
		if i <= len(tc.Domain.Domain) {
			domain := fmt.Sprint(
				tc.Domain.Domain[:i],
				string(tc.Domain.Domain[i]),
				string(tc.Domain.Domain[i]),
				tc.Domain.Domain[i + 1:],
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

// DoubleCharacterReplacementFunc are created by replacing identical, consecutive letters of the domain name with
// adjacent letters on the keyboard. For example, www.gppgle.com and www.giigle.com
func DoubleCharacterReplacementFunc(tc TypoConfig) (results []TypoConfig) {
	for _, keyboard := range tc.Keyboards {
		for i, char := range tc.Domain.Domain {
			if i < len(tc.Domain.Domain) - 1 {
				if tc.Domain.Domain[i] == tc.Domain.Domain[i + 1] {
					for _, key := range keyboard.Adjacent(string(char)) {
						domain := tc.Domain.Domain[:i] + string(key) + string(key) + tc.Domain.Domain[i + 2:]
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

// StripDashesFunc typos are created by omitting a dot from the domain. For example, www.a-b-c.com becomes www.abc.com
func StripDashesFunc(tc TypoConfig) (results []TypoConfig) {
	for _, str := range ReplaceCharFunc(tc.Domain.Domain, "-", "") {
		if tc.Domain.Domain != str {
			dm := Domain{tc.Domain.Subdomain, str, tc.Domain.Suffix}
			results = append(results, TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})
		}
	}
	return
}

// SingularPluraliseFunc are created by making a singular domain plural and vice versa. For example, www.google.com
// becomes www.googles.com and www.games.co.nz becomes www.game.co.nz
func SingularPluraliseFunc(tc TypoConfig) (results []TypoConfig) {
	var domain string
	if strings.HasSuffix(tc.Domain.Domain, "s"){
		domain = strings.TrimSuffix(tc.Domain.Domain, "s")
	} else {
		domain = tc.Domain.Domain + "s"
	}
	dm := Domain{tc.Domain.Subdomain, domain, tc.Domain.Suffix}
	results = append(results, TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})
	return
}

// CommonMisspellingsFunc are created with common misspellings in the given language.
// For example, www.youtube.com becomes www.youtub.com and www.abseil.com becomes www.absail.com
func CommonMisspellingsFunc(tc TypoConfig) (results []TypoConfig) {
	for _, keyboard := range tc.Keyboards {
		for _, word := range keyboard.Language.SimilarSpellings(tc.Domain.Domain) {
			dm := Domain{tc.Domain.Subdomain, word, tc.Domain.Suffix}
			results = append(results, TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})

		}
	}
	return
}


// VowelSwappingFunc swaps vowels within the domain name except for the first letter.
// For example, www.google.com becomes www.gaagle.com.
func VowelSwappingFunc(tc TypoConfig) (results []TypoConfig) {
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


// HomophonesFunc are created from sets of words that sound the same when spoken.
// For example, www.base.com becomes www .bass.com.
func HomophonesFunc(tc TypoConfig) (results []TypoConfig) {
	for _, keyboard := range tc.Keyboards {
		for _, word := range keyboard.Language.SimilarSounds(tc.Domain.Domain) {
			dm := Domain{tc.Domain.Subdomain, word, tc.Domain.Suffix}
			results = append(results, TypoConfig{dm, tc.Keyboards, tc.Languages, tc.Typo})

		}
	}
	return
}

//
//// BitFlipping each letter in a domain name is an 8bit character. The character is substituted with the set of valid
//// characters that can be made after a single bit flip. For example, facebook.com becomes bacebook.com, dacebook.com,
//// faaebook.com,fabebook.com,facabook.com, etc.
//func BitFlipping(tc TypoConfig) (results []TypoConfig) {
//	tc.Domain.Domain = tc.Domain.Domain + ":S"
//	results = append(results, tc)
//	return results
//}

// HomoglyphFunc when one or more characters that look similar to another character but are different are called
// homogylphs. An example is that the lower case l looks similar to the numeral one, e.g. l vs 1.
// For example, google.com becomes goog1e.com.
func HomoglyphFunc(tc TypoConfig) (results []TypoConfig) {
	for i, char := range tc.Domain.Domain {
		// Check the alphabet of the language associated with the keyboard for homoglyphs
		for _, keyboard := range tc.Keyboards {
			for _, kchar := range keyboard.Language.SimilarChars(string(char)) {
				domain := fmt.Sprint(tc.Domain.Domain[:i], kchar, tc.Domain.Domain[i + 1:])
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
				domain := fmt.Sprint(tc.Domain.Domain[:i], lchar, tc.Domain.Domain[i + 1:])
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

//// WrongTopLevelDomain for example, www.trademe.co.nz becomes www.trademe.co.nz and www.google.com becomes
//// www.google.org. uses the 19 most common top level domains.
//func WrongTopLevelDomain(tc TypoConfig) (results []TypoConfig) {
//	tc.Domain.Domain = tc.Domain.Domain + ":S"
//	results = append(results, tc)
//	return results
//}
//
//
//// WrongSecondLevelDomain uses an alternate, valid second level domain for the top level domain.
//// For example, www.trademe.co.nz becomes www.trademe.ac.nz and www.trademe.iwi.nz
//func WrongSecondLevelDomain(tc TypoConfig) (results []TypoConfig) {
//	tc.Domain.Domain = tc.Domain.Domain + ":S"
//	results = append(results, tc)
//	return results
//}
//
//// Bitsquatting relies on random bit- errors to redirect connections intended for popular domains
//func Bitsquatting(tc TypoConfig) (results []TypoConfig) {
//	tc.Domain.Domain = tc.Domain.Domain + ":S"
//	results = append(results, tc)
//	return results
//}

// MissingCharFunc removes a character one at a time from the string. For example, wwwgoogle.com and www.googlecom
func MissingCharFunc(str, character string) (results []string) {
	for i, char := range str {
		if character == string(char) {
			results = append(results, str[:i] + str[i + 1:])
		}
	}
	return
}

// ReplaceCharFunc omits a character from the entire string. For example, www.a-b-c.com becomes www.abc.com
func ReplaceCharFunc(str, old, new string) (results []string) {
	results = append(results, strings.Replace(str, old, new, -1))
	return
}
