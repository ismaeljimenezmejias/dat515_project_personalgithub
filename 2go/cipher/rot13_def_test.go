package cipher

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func testRot13(t *testing.T, scoreDec func()) {
	b := make([]byte, 1024)
	for _, test := range rot13Tests {
		r := rot13Reader{strings.NewReader(test.in)}
		n, err := r.Read(b)
		if err != nil {
			t.Errorf("rot13.Read(%q): got %v, expected EOF", test.in, err)
			scoreDec()
			continue
		}
		got := string(b[:n])
		if diff := cmp.Diff(test.want, got); diff != "" {
			t.Errorf("rot13.Read(%q): (-want +got):\n%s", test.in, diff)
			scoreDec()
		}
	}
}

var rot13Tests = []struct {
	in, want string
}{
	{
		"Go programming is fun.",
		"Tb cebtenzzvat vf sha.",
	},
	{
		"Tb cebtenzzvat vf sha.",
		"Go programming is fun.",
	},
	{
		"There are two hard things in computer science: cache invalidation, naming things, and off-by-one errors.",
		"Gurer ner gjb uneq guvatf va pbzchgre fpvrapr: pnpur vainyvqngvba, anzvat guvatf, naq bss-ol-bar reebef.",
	},
	{
		"Gurer ner gjb uneq guvatf va pbzchgre fpvrapr: pnpur vainyvqngvba, anzvat guvatf, naq bss-ol-bar reebef.",
		"There are two hard things in computer science: cache invalidation, naming things, and off-by-one errors.",
	},
	{
		"All problems in computer science can be solved by another level of indirection, except of course for the problem of too many indirections.",
		"Nyy ceboyrzf va pbzchgre fpvrapr pna or fbyirq ol nabgure yriry bs vaqverpgvba, rkprcg bs pbhefr sbe gur ceboyrz bs gbb znal vaqverpgvbaf.",
	},
	{
		"Nyy ceboyrzf va pbzchgre fpvrapr pna or fbyirq ol nabgure yriry bs vaqverpgvba, rkprcg bs pbhefr sbe gur ceboyrz bs gbb znal vaqverpgvbaf.",
		"All problems in computer science can be solved by another level of indirection, except of course for the problem of too many indirections.",
	},
}
