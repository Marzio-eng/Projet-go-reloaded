package main

import "testing"

func TestProcess(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"marqueurs simples hex et bin", "Simply add 42 (hex) and 10 (bin)", "Simply add 66 and 2"},
		{"marqueur up sur ponctuation", "Ready, set, go (up) !", "Ready, set, GO!"},
		{"article a -> an devant voyelle", "There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"cas de reference 1", "If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?", "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"},
		{"cas de reference 2", "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure", "I have to pack 5 outfits. Packed 26 just to be sure"},
		{"cas de reference 4", "harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '", "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"},
		{"marqueur en debut de ligne", "(up) hello", "hello"},
		{"n plus grand que le nombre de mots", "go now (up, 9)", "GO NOW"},
		{"capitalize ne casse pas les accents", "élan (cap)", "Élan"},
		{"conversion hex invalide", "zz (hex)", "zz"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Process(c.in)
			if got != c.want {
				t.Errorf("Process(%q)\n got  %q\n want %q", c.in, got, c.want)
			}
		})
	}
}

func TestConvertBase(t *testing.T) {
	cases := []struct {
		in   string
		base int
		want string
	}{
		{"ff", 16, "255"},
		{"1E", 16, "30"},
		{"101", 2, "5"},
		{"zz", 16, "zz"},
	}

	for _, c := range cases {
		got := convertBase(c.in, c.base)
		if got != c.want {
			t.Errorf("convertBase(%q, %d) = %q, want %q", c.in, c.base, got, c.want)
		}
	}
}

func TestCapitalize(t *testing.T) {
	cases := []struct{ in, want string }{
		{"bridge", "Bridge"},
		{"bRIDGE", "Bridge"},
		{"élan", "Élan"},
		{"", ""},
	}

	for _, c := range cases {
		got := capitalize(c.in)
		if got != c.want {
			t.Errorf("capitalize(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
