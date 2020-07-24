package universe

import "testing"

func TestShouldInitializeUniverse(t *testing.T) {
	var universe = initialUniverse()

	if universe == nil {
		t.Fail()
	}
}
