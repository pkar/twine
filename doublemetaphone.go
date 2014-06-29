// Package twine ...
// see python the implementation of
// double metaphone here http://www.drdobbs.com/the-double-metaphone-search-algorithm/184401251
// or python https://github.com/oubiwann/metaphone
// or c++ http://aspell.net/metaphone/dmetaph.cpp
// Returns a code indicating the "sound" of the supplied word.
// This is intended for use in a spell checker or similar to
// find words that sound the same or similar.
//
// For example:
//	swordfish = SRTF,XRTF
//	Gammon = KMN
//	gamon = KMN
//	gamin = KMN
//	Cameron = KMRN
//	bruise = PRS
//	Bruce = PRS
package twine

import (
	"bytes"
	"fmt"
	"strings"
)

type R []rune

var (
	vowels         = R("AEIOUY")
	silentStarters = []string{"GN", "KN", "PN", "WR", "PS"}
)

// DM holds the meta information about the current string to process.
type DM struct {
	original   string
	length     int
	upper      string
	upperRune  R
	lengthRune int

	position      int
	maxCodeLength int

	primary   bytes.Buffer
	alternate bytes.Buffer
}

// NewDM initializes a double metaphone string parser.
func NewDM(s string, codeLength int) *DM {
	if codeLength == 0 {
		codeLength = 4
	}
	u := strings.ToUpper(s)
	ru := R(u)
	dm := &DM{
		original:      s,
		upper:         u,
		upperRune:     ru,
		length:        len(u),
		lengthRune:    len(ru),
		position:      0,
		maxCodeLength: codeLength,
	}
	return dm
}

// isSlavoGermanic check if string contains slavo germanic characters.
func (dm *DM) isSlavoGermanic() bool {
	return strings.HasPrefix(dm.upper, "SCH") ||
		strings.HasPrefix(dm.upper, "SW") ||
		strings.HasPrefix(dm.upper, "J") ||
		strings.HasPrefix(dm.upper, "W") ||
		strings.HasPrefix(dm.upper, "CZ") ||
		strings.HasPrefix(dm.upper, "K") ||
		strings.HasPrefix(dm.upper, "W") ||
		strings.HasPrefix(dm.upper, "WITZ")
}

// isSilentStarter checks for silent letters at the start of a word.
func (dm *DM) isSilentStarter() bool {
	for _, starter := range silentStarters {
		if strings.HasPrefix(dm.upper, starter) {
			return true
		}
	}
	return false
}

// contains checks if the input rune from start to len sub is a match
// for the given sub rune. It does bounds checking to prevent panics.
func (dm *DM) contains(start int, sub ...R) bool {
	if len(dm.upperRune) == 0 || len(sub) == 0 {
		return false
	}
	if start < 0 {
		return false
	}
	for _, subRune := range sub {
		isIn := true
		end := len(subRune) + start
		if len(dm.upperRune) < end {
			continue
		}
		j := 0
		for i := start; i < end; i++ {
			if dm.upperRune[i] != subRune[j] {
				isIn = false
				break
			}
			j++
		}
		if isIn {
			return true
		}
	}
	return false
}

// write appends to the primary and alternate buffers.
func (dm *DM) write(primary, alternate string) {
	if primary != "" {
		dm.primary.WriteString(primary)
	}
	if alternate != "" {
		dm.alternate.WriteString(alternate)
	}
}

// isVowel checks to see if a rune is a, e i o, u, y...
func (dm *DM) isVowel(pos int) bool {
	if pos < 0 || pos >= dm.lengthRune {
		return false
	}
	switch dm.upperRune[pos] {
	case 'A', 'E', 'I', 'O', 'U', 'Y', 'Á', 'Â', 'Ã', 'Ä', 'Å',
		'Æ', 'È', 'É', 'Ê', 'Ë', 'Ì', 'Í', 'Î', 'Ï', 'Ò', 'Ó', 'Ô',
		'Õ', 'Ö', '', 'Ø', 'Ù', 'Ú', 'Û', 'Ü', 'Ý', '':
		return true
	default:
		return false
	}
}

// parse starts parsing the string and returns the
// double metaphone representation of it.
func (dm *DM) parse() [2]string {

	current := 0
	end := dm.lengthRune - 1

	// skip this at beginning of word
	if dm.isSilentStarter() {
		current++
	}

	isSlavo := dm.isSlavoGermanic()

	// Initial 'X' is pronounced 'Z' e.g. 'Xavier'
	if dm.upperRune[0] == 'X' {
		dm.write("S", "S")
		current++
	}

	for dm.primary.Len() < dm.maxCodeLength || dm.alternate.Len() < dm.maxCodeLength {
		if current > end {
			break
		}

		switch dm.upperRune[current] {
		case 'A', 'E', 'I', 'O', 'U', 'Y':
			if current == 0 {
				dm.write("A", "A")
			}
		case 'B':
			// '-mb', e.g. "dumb", already skipped over ...
			dm.write("P", "P")
			if dm.contains(current+1, R("B")) {
				current += 2
				continue
			}
		case 'Ç':
			dm.write("S", "S")
		case 'C':
			switch {
			case current > 1 &&
				!dm.isVowel(current-2) &&
				dm.contains(current-1, R("ACH")) &&
				!dm.contains(current+2, R("I")) &&
				(!dm.contains(current+2, R("E")) || dm.contains(current-2, R("BACHER"), R("MACHER"))):
				// various gremanic
				dm.write("K", "K")
				current += 2
				continue
			case current == 0 && dm.contains(current, R("CAESAR")):
				// special case 'CAESAR'
				dm.write("S", "S")
				current += 2
				continue
			case dm.contains(current, R("CHIA")):
				// italian 'chianti'
				dm.write("K", "K")
				current += 2
				continue
			case dm.contains(current, R("CH")):
				switch {
				case current > 0 && dm.contains(current, R("CHAE")):
					//find 'michael'
					dm.write("K", "X")
				case current == 0 &&
					(dm.contains(current+1, R("HARAC"), R("HARIS")) || dm.contains(current+1, R("HOR"), R("HYM"), R("HIA"), R("HEM"))) &&
					!dm.contains(0, R("CHORE")):
					// greek roots e.g. 'chemistry', 'chorus'
					dm.write("K", "K")
				case dm.contains(0, R("VAN "), R("VON "), R("SCH")) ||
					dm.contains(current-2, R("ORCHES"), R("ARCHIT"), R("ORCHID")) || // 'architect' but not 'arch', orchestra', 'orchid'
					dm.contains(current+2, R("T"), R("S")) ||
					((dm.contains(current-1, R("A"), R("O"), R("U"), R("E")) || current == 0) &&
						dm.contains(current+2, R("L"), R("R"), R("N"), R("M"), R("B"), R("H"), R("F"), R("V"), R("W"), R(" "))):
					// e.g. 'wachtler', 'weschsler', but not 'tichner'
					//germanic, greek, or otherwise 'ch' for 'kh' sound
					dm.write("K", "K")
				default:
					switch {
					case current > 0:
						if dm.contains(0, R("MC")) {
							// e.g. 'McHugh'
							dm.write("K", "K")
						} else {
							dm.write("X", "K")
						}
					default:
						dm.write("X", "X")
					}
				}

				current += 2
				continue
			case dm.contains(current, R("CZ")) && !dm.contains(current-2, R("WICZ")):
				// e.g. 'czerny'
				dm.write("S", "X")
				current += 2
				continue
			case dm.contains(current+1, R("CIA")):
				// e.g. 'focaccia'
				dm.write("X", "X")
				current += 3
				continue
			case dm.contains(current, R("CC")) && !(current == 1 && dm.contains(0, R("M"))):
				// double 'C', but not McClellan'
				// 'bellocchio' but not 'bacchus'
				if dm.contains(current+2, R("I"), R("E"), R("H")) && !dm.contains(current+2, R("HU")) {
					if (current == 1 && dm.contains(current-1, R("A"))) || dm.contains(current-1, R("UCCEE"), R("UCCES")) {
						// 'accident', 'accede', 'succeed'
						dm.write("KS", "KS")
					} else {
						// 'bacci', 'bertucci', other italian
						dm.write("X", "X")
					}
					current += 3
					continue
				} else {
					// Pierce's rule
					dm.write("K", "K")
				}
				current += 2
				continue
			case dm.contains(current, R("CK"), R("CG"), R("CQ")):
				dm.write("K", "K")
				current += 2
				continue
			case dm.contains(current, R("CI"), R("CE"), R("CY")):
				// italian vs. english
				if dm.contains(current, R("CIO"), R("CIE"), R("CIA")) {
					dm.write("S", "X")
				} else {
					dm.write("S", "S")
				}
				current += 2
				continue
			}

			dm.write("K", "K")
			// name sent in 'mac caffrey', 'mac gregor'
			switch {
			case dm.contains(current+1, R(" C"), R(" Q"), R(" G")):
				current += 3
				continue
			case dm.contains(current+1, R("C"), R("K"), R("Q")) && !dm.contains(current+1, R("CE"), R("CI")):
				current += 2
				continue
			}
		case 'D':
			switch {
			case dm.contains(current, R("DG")):
				// e.g. 'edge'
				if dm.contains(current+2, R("I"), R("E"), R("Y")) {
					dm.write("J", "J")
					current += 3
					continue
				} else {
					// e.g. 'edgar'
					dm.write("TK", "TK")
					current += 2
					continue
				}
			case dm.contains(current, R("DT"), R("DD")):
				dm.write("T", "T")
				current += 2
				continue
			default:
				dm.write("T", "T")
			}
		case 'F':
			dm.write("F", "F")
			if dm.contains(current+1, R("F")) {
				current += 2
				continue
			}
		case 'G':
			switch {
			case dm.contains(current+1, R("H")):
				switch {
				case current > 0 && !dm.isVowel(current-1):
					dm.write("K", "K")
				case current < 3:
					// 'ghislane', ghiradelli
					if current == 0 {
						if dm.contains(current+2, R("I")) {
							dm.write("J", "J")
						} else {
							dm.write("K", "K")
						}
					}
				case ((current > 1) && dm.contains(current-2, R("B"), R("H"), R("D"))) ||
					(current > 2 && dm.contains(current-3, R("B"), R("H"), R("D"))) || // e.g. 'bough'
					(current > 3 && dm.contains(current-4, R("B"), R("H"))): // e.g. 'broughton'
					// Parker's rule (with some further refinements) - e.g., 'hugh'
				case current > 2 && dm.contains(current-1, R("U")) &&
					dm.contains(current-3, R("C"), R("G"), R("L"), R("R"), R("T")):
					//e.g., 'laugh', 'McLaughlin', 'cough', 'gough', 'rough', 'tough'
					dm.write("F", "F")
				case current > 0 && !dm.contains(current-1, R("I")):
					dm.write("K", "K")
				}
				current += 2
				continue
			case dm.contains(current+1, R("N")):
				switch {
				case current == 1 && dm.isVowel(0) && !isSlavo:
					dm.write("KN", "N")
				case !dm.contains(current+2, R("EY")) && !dm.contains(current+1, R("Y")) && !isSlavo:
					//not e.g. 'cagney'
					dm.write("N", "KN")
				default:
					dm.write("KN", "KN")
				}
				current += 2
				continue
			case dm.contains(current+1, R("LI")) && !isSlavo:
				// 'tagliaro'
				dm.write("KL", "L")
				current += 2
				continue
			case current == 0 && (dm.contains(current+1, R("Y")) ||
				dm.contains(current+1, R("ES"), R("EP"), R("EB"), R("EL"), R("EY"), R("IB"), R("IL"), R("IN"), R("IE"), R("EI"), R("ER"))):
				// -ges-, -gep-, -gel- at beginning
				dm.write("K", "J")
				current += 2
				continue
			case (dm.contains(current+1, R("ER")) || dm.contains(current+1, R("Y"))) &&
				!dm.contains(0, R("DANGER"), R("RANGER"), R("MANGER")) &&
				!dm.contains(current-1, R("E"), R("I"), R("RGY"), R("OGY")):
				// -ger-, -gy-
				dm.write("K", "J")
				current += 2
				continue
			case dm.contains(current+1, R("E"), R("I"), R("Y")) || dm.contains(current-1, R("AGGI"), R("OGGI")):
				// italian e.g. 'biaggi'
				switch {
				case dm.contains(0, R("VAN "), R("VON "), R("SCH")) || dm.contains(current+1, R("ET")):
					// obvious germanic
					dm.write("K", "K")
				case dm.contains(current+1, R("IER ")):
					// always soft if french ending
					dm.write("J", "J")
				default:
					dm.write("J", "K")
				}
				current += 2
				continue
			case dm.contains(current+1, R("G")):
				dm.write("K", "K")
				current += 2
				continue

			}

			dm.write("K", "K")
			if dm.contains(current+1, R("G")) {
				current += 2
				continue
			}
		case 'H':
			// only keep if first & before vowel or between 2 vowels
			if current == 0 || dm.isVowel(current+1) && dm.isVowel(current-1) {
				dm.write("H", "H")
				current += 2
				continue
			}
		case 'J':
			// TODO
			// obvious spanish, 'jose', 'san jacinto'
			switch {
			case dm.contains(current, R("JOSE")) || dm.contains(0, R("SAN ")):
				if (current == 0 && dm.contains(current+4, R(" "))) || dm.contains(0, R("SAN ")) {
					dm.write("H", "H")
				} else {
					dm.write("J", "H")
				}
			case current == 0 && !dm.contains(current, R("JOSE")):
				dm.write("J", "A")
			default:
				switch {
				case dm.isVowel(current-1) && !isSlavo && dm.contains(current+1, R("A"), R("O")):
					dm.write("J", "H")
				default:
					switch {
					case current == end:
						dm.write("J", "")
					case !dm.contains(current+1, R("L"), R("T"), R("K"), R("S"), R("N"), R("M"), R("B"), R("Z")) &&
						!dm.contains(current-1, R("S"), R("K"), R("L")):
						dm.write("J", "J")
					}
				}
			}

			if dm.contains(current+1, R("J")) {
				current += 2
				continue
			}
		case 'K':
			dm.write("K", "K")
			if dm.contains(current+1, R("K")) {
				current += 2
				continue
			}
		case 'L':
			if dm.contains(current+1, R("L")) {
				switch {
				// spanish e.g. 'cabrillo', 'gallegos'
				case (current == end-2 && dm.contains(current-1, R("ILLO"), R("ILLA"), R("ALLE"))) ||
					((dm.contains(end-1, R("AS"), R("OS")) || dm.contains(end, R("A"), R("O"))) && dm.contains(current-1, R("ALLE"))):
					dm.write("L", "")
				default:
					dm.write("L", "L")
				}
				current += 2
				continue
			}
			dm.write("L", "L")
		case 'M':
			dm.write("M", "M")
			if dm.contains(current+1, R("M")) ||
				// 'dumb', 'thumb'
				(dm.contains(current-1, R("UMB")) && current+1 == end) ||
				dm.contains(current+2, R("ER")) {
				current += 2
				continue
			}
		case 'N':
			dm.write("N", "N")
			if dm.contains(current+1, R("N")) {
				current += 2
			}
		case 'Ñ':
			dm.write("N", "N")
		case 'P':
			if dm.contains(current+1, R("H")) {
				dm.write("F", "F")
				current += 2
				continue
			}
			dm.write("P", "P")
			// also account for "campbell" and "raspberry"
			if dm.contains(current+1, R("P"), R("B")) {
				current += 2
				continue
			}
		case 'Q':
			dm.write("K", "K")
			if dm.contains(current+1, R("Q")) {
				current += 2
				continue
			}
		case 'R':
			// french e.g. 'rogier', but exclude 'hochmeier'
			if current == end && !isSlavo &&
				dm.contains(current-2, R("IE")) &&
				!dm.contains(current-4, R("ME"), R("MA")) {
				dm.write("", "R")
			} else {
				dm.write("R", "R")
			}

			if dm.contains(current+1, R("R")) {
				current += 2
				continue
			}
		case 'S':
			switch {
			case dm.contains(current-1, R("ISL"), R("YSL")):
				// special cases 'island', 'isle', 'carlisle', 'carlysle'
			case current == 0 && dm.contains(current, R("SUGAR")):
				// special case 'sugar-'
				dm.write("X", "S")
				current += 2
				continue
			case dm.contains(current, R("SH")):
				// germanic
				if dm.contains(current+1, R("HEIM"), R("HOEK"), R("HOLM"), R("HOLZ")) {
					dm.write("S", "S")
				} else {
					dm.write("X", "X")
				}
				current += 2
				continue
			case dm.contains(current, R("SIO"), R("SIA"), R("SIAN")):
				// italian & armenian
				if !isSlavo {
					dm.write("S", "X")
				} else {
					dm.write("S", "S")
				}
				current += 3
				continue
			case (current == 0 && dm.contains(current+1, R("M"), R("N"), R("L"), R("W")) ||
				dm.contains(current+1, R("Z"))):
				// german & anglicisations, e.g. 'smith' match 'schmidt', 'snider' match 'schneider'
				// also, -sz- in slavic language altho in hungarian it is pronounced 's'
				dm.write("S", "X")
				if dm.contains(current+1, R("Z")) {
					current += 2
					continue
				}
			case dm.contains(current, R("SC")):
				// Schlesinger's rule
				switch {
				case dm.contains(current+2, R("H")):
					// dutch origin, e.g. 'school', 'schooner'
					if dm.contains(current+3, R("OO"), R("ER"), R("EN"), R("UY"), R("ED"), R("EM")) {
						// 'schermerhorn', 'schenker'
						if dm.contains(current+3, R("ER"), R("EN")) {
							dm.write("X", "SK")
						} else {
							dm.write("SK", "SK")
						}
					} else {
						if current == 0 && !dm.isVowel(3) && !dm.contains(3, R("W")) {
							dm.write("X", "S")
						} else {
							dm.write("X", "X")
						}
					}
				case dm.contains(current+2, R("I"), R("E"), R("Y")):
					dm.write("S", "S")
				default:
					dm.write("X", "X")
				}
				current += 3
				continue
			default:
				if current == end && dm.contains(current-2, R("AI"), R("OI")) {
					// french e.g. 'resnais', 'artois'
					dm.write("", "S")
					current += 2
					continue
				} else {
					dm.write("S", "S")
				}
				if dm.contains(current+1, R("S"), R("Z")) {
					current += 2
					continue
				}
			}
		case 'T':
			switch {
			case dm.contains(current, R("TIA"), R("TCH"), R("TION")):
				dm.write("X", "X")
				current += 3
				continue
			case dm.contains(current, R("TH"), R("TTH")):
				// special case 'thomas', 'thames' or germanic
				if dm.contains(current+2, R("OM"), R("AM")) || dm.contains(0, R("VAN "), R("VON "), R("SCH")) {
					dm.write("T", "T")
				} else {
					dm.write("0", "T")
				}
				current += 2
				continue
			}
			dm.write("T", "T")
			if dm.contains(current+1, R("T"), R("D")) {
				current += 2
				continue
			}
		case 'V':
			dm.write("F", "F")
			if dm.contains(current+1, R("V")) {
				current += 2
				continue
			}
		case 'W':
			// can also be in middle of word
			if dm.contains(current, R("WR")) {
				dm.write("R", "R")
				current += 2
				continue
			}

			if (current == 0) && (dm.isVowel(current+1) || dm.contains(current, R("WH"))) {
				// Wasserman should match Vasserman
				if dm.isVowel(current + 1) {
					dm.write("A", "F")
				} else {
					// need Uomo to match Womo
					dm.write("A", "A")
				}
			}

			// Arnow should match Arnoff
			if (current == end && dm.isVowel(current-1)) ||
				dm.contains(current-1, R("EWSKI"), R("EWSKY"), R("OWSKI"), R("OWSKY")) ||
				dm.contains(0, R("SCH")) {
				dm.write("", "F")
			}

			// polish e.g. 'filipowicz'
			if dm.contains(current, R("WICZ"), R("WITZ")) {
				dm.write("TS", "FX")
				current += 4
				continue
			}
		case 'X':
			switch {
			// french e.g. breaux
			case current == 0:
				dm.write("S", "S")
			case !(current == end && (dm.contains(current-3, R("IAU"), R("EAU")) ||
				dm.contains(current-2, R("AU"), R("OU")))):
				dm.write("KS", "KS")
			}
			if dm.contains(current+1, R("C"), R("X")) {
				current += 2
				continue
			}
		case 'Z':
			switch {
			// chinese pinyin e.g. 'zhao'
			case dm.contains(current+1, R("H")):
				dm.write("J", "J")
			default:
				if dm.contains(current+1, R("ZO"), R("ZI"), R("ZA")) ||
					(isSlavo && (current > 0 && !dm.contains(current-1, R("T")))) {
					dm.write("S", "TS")
				} else {
					dm.write("S", "S")
				}
			}
			if dm.contains(current+1, R("Z")) {
				current += 2
				continue
			}
		}
		current++
	}

	p := dm.primary.String()
	a := dm.alternate.String()
	if p == a {
		a = ""
	}
	if len(p) > dm.maxCodeLength {
		p = p[:dm.maxCodeLength]
	}
	if len(a) > dm.maxCodeLength {
		a = a[:dm.maxCodeLength]
	}

	return [2]string{p, a}
}

// DoubleMetaphone ...
func DoubleMetaphone(s string, codeLength int) ([2]string, error) {
	if len(s) < 1 {
		return [2]string{}, fmt.Errorf("string length 0")
	}
	dm := NewDM(s, 4)
	result := dm.parse()
	return result, nil
}
