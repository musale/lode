package token

import (
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	tokens := []Token{
		Token{NUMBER, "2", nil, 1},
		Token{PLUS, "+", nil, 2},
		Token{NUMBER, "2", nil, 3},
		Token{EQUAL, "=", nil, 4},
		Token{NUMBER, "4", nil, 5},
	}
	strTokens := []string{
		"{NUMBER 2 <nil> 1}",
		"{+ + <nil> 2}",
		"{NUMBER 2 <nil> 3}",
		"{= = <nil> 4}",
		"{NUMBER 4 <nil> 5}",
	}
	for i, token := range tokens {
		strToken := fmt.Sprint(token)
		if strToken != strTokens[i] {
			t.Errorf("Expected %s but got %s", strTokens[i], strToken)
		}
	}
}
