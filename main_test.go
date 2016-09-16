package nodetwitterapi

import "testing"

func TestAPI(t *testing.T) {
	_, err := getSession()
	if err != nil {
        t.Errorf("Sorry! Failed to get the session: %s", err)
	}
}
