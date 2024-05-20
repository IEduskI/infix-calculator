package evaluate

import (
	"testing"
)

func TestPostfix(t *testing.T) {
	postfix := "1 3 4 * +"

	wantResult := 13.0

	got, _ := Postfix(postfix)

	if got != wantResult {
		t.Errorf("Evalute postfix got = %v, want %v", got, wantResult)
	}
}
