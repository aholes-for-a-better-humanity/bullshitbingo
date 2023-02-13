package names

import (
	"testing"
	"testing/quick"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Test_syllab(t *testing.T) {
	v := ""
	err := quick.Check(func() bool {
		v = syllab()
		return len(v) < 6
	}, nil)
	if err != nil {
		t.Fatal(v, "too long")
	}
}

func Test_coinFlip(t *testing.T) {
	var seenT, seenF int
	for i := 0; i < 1000; i++ {
		b := coinFlip()
		if b {
			seenT += 1
		} else {
			seenF += 1
		}
	}
	if seenF == 0 {
		t.Error("no false")
	}
	if seenT == 0 {
		t.Error("no true")
	}
	if seenF < 300 || seenF > 700 {
		t.Error("false", seenF)
	} // expected 500
	if seenF < 300 || seenF > 700 {
		t.Error("true", seenT)
	} // expected 500
	// note: obviously this test relies on i being rather big.
	// For our purpose, a coinflip reliable to 70/30 would still be okay.
}

func Test_Word(t *testing.T) {
	t.Run("words are not overly long", func(t *testing.T) {
		v := ""
		err := quick.Check(func() bool {
			v = Word()
			return len(v) < 15
		}, nil)
		if err != nil {
			t.Fatal(v, "too long")
		}
	})

}

func Test_Pseudo(t *testing.T) {
	t.Run("ensure low chance of collision", func(t *testing.T) {
		seen := make(map[string]bool)
		for i := 0; i < 1000; i++ {
			nioo := Pseudo()
			t.Log(nioo)
			if v, ok := seen[nioo]; ok {
				t.Fatal("dup:", v)
			} else {
				seen[nioo] = v
			}
		}
	})
	t.Run("pseudos are Capitalised", func(t *testing.T) {
		p := ""
		f := func() bool {
			p = Pseudo()
			return p == cases.Title(language.AmericanEnglish).String(p)
		}
		err := quick.Check(f, nil)
		if err != nil {
			t.Fatal(p, "is not Capitalized")
		}
	})
}
