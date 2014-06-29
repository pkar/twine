package twine

import "testing"

var doubleMetaphoneTests = []struct {
	in  string
	out [2]string
}{
	{"aubrey", [2]string{"APR", ""}},
	{"richard", [2]string{"RXRT", "RKRT"}},
	{"Jose", [2]string{"JS", "HS"}},
	{"cambrillo", [2]string{"KMPR", ""}},
	{"otto", [2]string{"AT", ""}},
	{"aubrey", [2]string{"APR", ""}},
	{"maurice", [2]string{"MRS", ""}},
	{"auto", [2]string{"AT", ""}},
	{"maisey", [2]string{"MS", ""}},
	{"catherine", [2]string{"K0RN", "KTRN"}},
	{"katherine", [2]string{"K0RN", "KTRN"}},
	{"geoff", [2]string{"JF", "KF"}},
	{"Chile", [2]string{"XL", ""}},
	{"steven", [2]string{"STFN", ""}},
	{"zhang", [2]string{"JNK", ""}},
	{"bob", [2]string{"PP", ""}},
	{"ray", [2]string{"R", ""}},
	{"Tux", [2]string{"TKS", ""}},
	{"bryan", [2]string{"PRN", ""}},
	{"bryce", [2]string{"PRS", ""}},
	{"Rapelje", [2]string{"RPL", ""}},
	{"solilijs", [2]string{"SLLS", ""}},
	{"Dallas", [2]string{"TLS", ""}},
	{"Schwein", [2]string{"XN", "XFN"}},
	{"dave", [2]string{"TF", ""}},
	{"eric", [2]string{"ARK", ""}},
	{"Parachute", [2]string{"PRKT", ""}},
	{"brian", [2]string{"PRN", ""}},
	{"randy", [2]string{"RNT", ""}},
	{"Through", [2]string{"0R", "TR"}},
	{"Nowhere", [2]string{"NR", ""}},
	{"heidi", [2]string{"HT", ""}},
	{"Arnow", [2]string{"ARN", "ARNF"}},
	{"Thumbail", [2]string{"0MPL", "TMPL"}},
	{"andestādītu", [2]string{"ANTS", ""}},
	{"français", [2]string{"FRNS", ""}},
	{"garçon", [2]string{"KRSN", ""}},
	{"leçon", [2]string{"LSN", ""}},
	{"ach", [2]string{"AX", "AK"}},
	{"bacher", [2]string{"PKR", ""}},
	{"macher", [2]string{"MKR", ""}},
	{"bacci", [2]string{"PX", ""}},
	{"bertucci", [2]string{"PRTX", ""}},
	{"bellocchio", [2]string{"PLX", ""}},
	{"bacchus", [2]string{"PKS", ""}},
	{"focaccia", [2]string{"FKX", ""}},
	{"chianti", [2]string{"KNT", ""}},
	{"tagliaro", [2]string{"TKLR", "TLR"}},
	{"biaggi", [2]string{"PJ", "PK"}},
	{"bajador", [2]string{"PJTR", "PHTR"}},
	{"cabrillo", [2]string{"KPRL", "KPR"}},
	{"gallegos", [2]string{"KLKS", "KKS"}},
	{"San Jacinto", [2]string{"SNHS", ""}},
	{"rogier", [2]string{"RJ", "RKR"}},
	{"breaux", [2]string{"PR", ""}},
	{"Wewski", [2]string{"ASK", "FFSK"}},
	{"zhao", [2]string{"J", ""}},
	{"school", [2]string{"SKL", ""}},
	{"schooner", [2]string{"SKNR", ""}},
	{"schermerhorn", [2]string{"XRMR", "SKRM"}},
	{"schenker", [2]string{"XNKR", "SKNK"}},
	{"Charac", [2]string{"KRK", ""}},
	{"Charis", [2]string{"KRS", ""}},
	{"chord", [2]string{"KRT", ""}},
	{"Chym", [2]string{"KM", ""}},
	{"Chia", [2]string{"K", ""}},
	{"chem", [2]string{"KM", ""}},
	{"chore", [2]string{"XR", ""}},
	{"orchestra", [2]string{"ARKS", ""}},
	{"architect", [2]string{"ARKT", ""}},
	{"orchid", [2]string{"ARKT", ""}},
	{"accident", [2]string{"AKST", ""}},
	{"accede", [2]string{"AKST", ""}},
	{"succeed", [2]string{"SKST", ""}},
	{"mac caffrey", [2]string{"MKFR", ""}},
	{"mac gregor", [2]string{"MKRK", ""}},
	{"mc crae", [2]string{"MKR", ""}},
	{"mcclain", [2]string{"MKLN", ""}},
	{"laugh", [2]string{"LF", ""}},
	{"cough", [2]string{"KF", ""}},
	{"rough", [2]string{"RF", ""}},
	{"gya", [2]string{"K", "J"}},
	{"ges", [2]string{"KS", "JS"}},
	{"gep", [2]string{"KP", "JP"}},
	{"geb", [2]string{"KP", "JP"}},
	{"gel", [2]string{"KL", "JL"}},
	{"gey", [2]string{"K", "J"}},
	{"gib", [2]string{"KP", "JP"}},
	{"gil", [2]string{"KL", "JL"}},
	{"gin", [2]string{"KN", "JN"}},
	{"gie", [2]string{"K", "J"}},
	{"gei", [2]string{"K", "J"}},
	{"ger", [2]string{"KR", "JR"}},
	{"danger", [2]string{"TNJR", "TNKR"}},
	{"dangerous", [2]string{"TNJR", "TNKR"}},
	{"manager", [2]string{"MNKR", "MNJR"}},
	{"dowager", [2]string{"TKR", "TJR"}},
	{"Campbell", [2]string{"KMPL", ""}},
	{"raspberry", [2]string{"RSPR", ""}},
	{"Thomas", [2]string{"TMS", ""}},
	{"Thames", [2]string{"TMS", ""}},
	{"xa", [2]string{"S", ""}},
	{"Çb", [2]string{"SP", ""}},
	{"GNu", [2]string{"N", ""}},
	{"dg", [2]string{"TK", ""}},
	{"dgi", [2]string{"J", ""}},
	{"dge", [2]string{"J", ""}},
	{"dgyB", [2]string{"JP", ""}},
	{"dtB", [2]string{"TP", ""}},
	{"ddB", [2]string{"TP", ""}},
	{"doB", [2]string{"TP", ""}},
	{"ffB", [2]string{"FP", ""}},
	{"fB", [2]string{"FP", ""}},
	{"Hab", [2]string{"HP", ""}},
	{"aHab", [2]string{"AHP", ""}},
	{"aHb", [2]string{"AP", ""}},
	{"kkb", [2]string{"KP", ""}},
	{"kb", [2]string{"KP", ""}},
	{"nb", [2]string{"NP", ""}},
	{"anna", [2]string{"AN", ""}},
	{"Ñb", [2]string{"NP", ""}},
	{"Occasionally", [2]string{"AKSN", "AKXN"}},
	{"antidisestablishmentarianism", [2]string{"ANTT", ""}},
	{"appreciated", [2]string{"APRS", "APRX"}},
	{"beginning", [2]string{"PJNN", "PKNN"}},
	{"changing", [2]string{"XNJN", "XNKN"}},
	{"cheat", [2]string{"XT", ""}},
	{"dangerous", [2]string{"TNJR", "TNKR"}},
	{"development", [2]string{"TFLP", ""}},
	{"etiology", [2]string{"ATLJ", "ATLK"}},
	{"existence", [2]string{"AKSS", ""}},
	{"simplicity", [2]string{"SMPL", ""}},
	{"circumstances", [2]string{"SRKM", ""}},
	{"fiery", [2]string{"FR", ""}},
	{"february", [2]string{"FPRR", ""}},
	{"illegitimate", [2]string{"ALJT", "ALKT"}},
	{"immediately", [2]string{"AMTT", ""}},
	{"happily", [2]string{"HPL", ""}},
	{"judgment", [2]string{"JTKM", "ATKM"}},
	{"knowing", [2]string{"NNK", ""}},
	{"kipper", [2]string{"KPR", ""}},
	{"john", [2]string{"JN", "AN"}},
	{"lesion", [2]string{"LSN", "LXN"}},
	{"Xavier", [2]string{"SF", "SFR"}},
	{"dumb", [2]string{"TM", ""}},
	{"caesar", [2]string{"SSR", ""}},
	{"chianti", [2]string{"KNT", ""}},
	{"michael", [2]string{"MKL", "MXL"}},
	{"chemistry", [2]string{"KMST", ""}},
	{"chorus", [2]string{"KRS", ""}},
	{"architect", [2]string{"ARKT", ""}},
	{"arch", [2]string{"ARX", "ARK"}},
	{"orchestra", [2]string{"ARKS", ""}},
	{"orchid", [2]string{"ARKT", ""}},
	{"wachtler", [2]string{"AKTL", "FKTL"}},
	{"wechsler", [2]string{"AKSL", "FKSL"}},
	{"tichner", [2]string{"TXNR", "TKNR"}},
	{"McHugh", [2]string{"MK", ""}},
	{"czerny", [2]string{"SRN", "XRN"}},
	{"focaccia", [2]string{"FKX", ""}},
	{"bellocchio", [2]string{"PLX", ""}},
	{"bacchus", [2]string{"PKS", ""}},
	{"accident", [2]string{"AKST", ""}},
	{"accede", [2]string{"AKST", ""}},
	{"succeed", [2]string{"SKST", ""}},
	{"bacci", [2]string{"PX", ""}},
	{"bertucci", [2]string{"PRTX", ""}},
	{"mac caffrey", [2]string{"MKFR", ""}},
	{"mac gregor", [2]string{"MKRK", ""}},
	{"edge", [2]string{"AJ", ""}},
	{"edgar", [2]string{"ATKR", ""}},
	{"ghislane", [2]string{"JLN", ""}},
	{"ghiradelli", [2]string{"JRTL", ""}},
	{"hugh", [2]string{"H", ""}},
	{"bough", [2]string{"P", ""}},
	{"broughton", [2]string{"PRTN", ""}},
	{"laugh", [2]string{"LF", ""}},
	{"McLaughlin", [2]string{"MKLF", ""}},
	{"cough", [2]string{"KF", ""}},
	{"gough", [2]string{"KF", ""}},
	{"rough", [2]string{"RF", ""}},
	{"tough", [2]string{"TF", ""}},
	{"cagney", [2]string{"KKN", ""}},
	{"tagliaro", [2]string{"TKLR", "TLR"}},
	{"biaggi", [2]string{"PJ", "PK"}},
	{"san jacinto", [2]string{"SNHS", ""}},
	{"Yankelovich", [2]string{"ANKL", ""}},
	{"Jankelowicz", [2]string{"JNKL", "ANKL"}},
	{"bajador", [2]string{"PJTR", "PHTR"}},
	{"cabrillo", [2]string{"KPRL", "KPR"}},
	{"gallegos", [2]string{"KLKS", "KKS"}},
	{"dumb", [2]string{"TM", ""}},
	{"thumb", [2]string{"0M", "TM"}},
	{"campbell", [2]string{"KMPL", ""}},
	{"raspberry", [2]string{"RSPR", ""}},
	{"hochmeier", [2]string{"HKMR", ""}},
	{"island", [2]string{"ALNT", ""}},
	{"isle", [2]string{"AL", ""}},
	{"carlisle", [2]string{"KRLL", ""}},
	{"carlysle", [2]string{"KRLL", ""}},
	{"smith", [2]string{"SM0", "XMT"}},
	{"schmidt", [2]string{"XMT", "SMT"}},
	{"snider", [2]string{"SNTR", "XNTR"}},
	{"schneider", [2]string{"XNTR", "SNTR"}},
	{"school", [2]string{"SKL", ""}},
	{"schooner", [2]string{"SKNR", ""}},
	{"schermerhorn", [2]string{"XRMR", "SKRM"}},
	{"schenker", [2]string{"XNKR", "SKNK"}},
	{"resnais", [2]string{"RSN", "RSNS"}},
	{"artois", [2]string{"ART", "ARTS"}},
	{"thomas", [2]string{"TMS", ""}},
	{"Wasserman", [2]string{"ASRM", "FSRM"}},
	{"Vasserman", [2]string{"FSRM", ""}},
	{"Uomo", [2]string{"AM", ""}},
	{"Womo", [2]string{"AM", "FM"}},
	{"Arnow", [2]string{"ARN", "ARNF"}},
	{"Arnoff", [2]string{"ARNF", ""}},
	{"filipowicz", [2]string{"FLPT", "FLPF"}},
	{"breaux", [2]string{"PR", ""}},
	{"zhao", [2]string{"J", ""}},
	{"thames", [2]string{"TMS", ""}},
}

func TestDoubleMetaphone(t *testing.T) {
	for _, tt := range doubleMetaphoneTests {
		res, err := DoubleMetaphone(tt.in, 4)
		if err != nil {
			t.Error(err)
			continue
		}
		if res[0] != tt.out[0] || res[1] != tt.out[1] {
			t.Errorf("DoubleMetaphone(%s) => %v, want %v", tt.in, res, tt.out)
		}
	}
}

var containsTests = []struct {
	in    string
	start int
	sub   []R
	out   bool
}{
	// single result
	{"abcdefg", 2, []R{R("CD")}, true},
	{"abcdefg", 2, []R{R("DA")}, false},
	{"abcdefg", 3, []R{R("DC"), R("DA")}, false},
	{"abcdefg", 3, []R{R("DC"), R("DA"), R("DE")}, true},
	{"abcdefg", 4, []R{R("DC"), R("EF")}, true},
	{"", 2, []R{R("DA")}, false},
	{"ab", 2, []R{R("DA")}, false},
}

func TestContains(t *testing.T) {
	for _, tt := range containsTests {
		dm := NewDM(tt.in, 4)
		res := dm.contains(tt.start, tt.sub...)
		if res != tt.out {
			t.Errorf("contains(%v, %d, %v) => %t, want %t", R(tt.in), tt.start, tt.sub, res, tt.out)
		}
	}
}

var silentStarterTests = []struct {
	in  string
	out bool
}{
	{"KNOWING", true},
	{"GNU", true},
	{"ABCDEFG", false},
	{"", false},
}

func TestIsSilentStarter(t *testing.T) {
	for _, tt := range silentStarterTests {
		dm := NewDM(tt.in, 4)
		res := dm.isSilentStarter()
		if res != tt.out {
			t.Errorf("isSilentStarter(%s) => %t, want %t", tt.in, res, tt.out)
		}
	}
}

func BenchmarkDoubleMetaphone(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DoubleMetaphone("abcdefghijklmnopqrstuvwxyz", 4)
	}
}

func BenchmarkDoubleMetaphone2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DoubleMetaphone("bertucci", 4)
	}
}
