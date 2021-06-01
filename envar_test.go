package envar

import (
	"reflect"
	"testing"
	"time"
)

func TestGetDefs(t *testing.T) {
	expected := map[string]interface{}{
		"k1": "v1",
		"k2": 123,
		"k3": true,
		"k4": 4.56,
		"k5": 3 * time.Second,
	}

	for k, v := range expected {
		SetDef(k, v)
	}

	got := GetDefs()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected:\n%v\ngot:\n%v\n", expected, got)
	}
}
