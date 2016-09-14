package nodetwitterapi

import "testing"

func TestAPI(t *testing.T) {
	_, err := getSession()
	if err != nil {
		t.Fail()
	}

	var analytic Analytic
	analytic, err := getAnalytic()
	if err != nil {
		t.Fail()
	}
}
