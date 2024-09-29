package dlna

import (
	"fmt"
	"testing"
)

func TestContentFeaturesString(t *testing.T) {
	a := ContentFeatures{
		Transcoded:      false,
		SupportTimeSeek: false,
	}.String()
	fmt.Println("actual: ", a)
	e := "DLNA.ORG_OP=10;DLNA.ORG_CI=1;DLNA.ORG_FLAGS=01700000000000000000000000000000"
	if e != a {
		t.Fatal(a)
	}
}
