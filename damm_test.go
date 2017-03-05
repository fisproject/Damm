package damm

import (
	"testing"
)

func TestAppend(t *testing.T) {
	candidates := []struct {
		input    string
		expected string
	}{
		{"22456433232", "224564332323"},
		{"54352543234", "543525432346"},
		{"83532323322", "835323233226"},
		{"88498347454", "884983474543"},
		{"19347355682", "193473556821"},
		{"75932958575", "759329585754"},
		{"33584758462", "335847584629"},
		{"08989435403", "089894354035"},
		{"10493839530", "104938395305"},
		{"54994384990", "549943849904"},
	}

	for _, c := range candidates {
		result, err := Append(c.input)

		if err != nil {
			t.Errorf("expected %v, but %v", nil, err)
		}

		if result != c.expected {
			t.Errorf("input is expected %v, but %v", c.input, c.expected, result)
		}
	}
}

func TestVerify(t *testing.T) {
	candidates := []struct {
		input    string
		expected bool
	}{
		{"224564332321", false},
		{"543525432340", false},
		{"835323233227", false},
		{"884983474541", false},
		{"193473556822", false},
		{"759329585754", true},
		{"335847584629", true},
		{"089894354035", true},
		{"104938395305", true},
		{"549943849904", true},
	}

	for _, c := range candidates {
		result, err := Verify(c.input)

		if err != nil {
			t.Errorf("expected %v, but %v", nil, err)
		}

		if result != c.expected {
			t.Errorf("%v is expected %v, but %v", c.input, c.expected, result)
		}
	}
}
