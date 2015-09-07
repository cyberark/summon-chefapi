package main

import "testing"

func TestParsePath(t *testing.T) {
	_, _, _, err := parsePath("")
	if err == nil {
		t.Error("Blank path did not return an error")
	}

	_, _, _, err = parsePath("postgres")
	if err == nil {
		t.Error("Path not fully formed")
	}

	_, _, _, err = parsePath("postgres/password")
	if err == nil {
		t.Error("Path not fully formed")
	}

	bagName, bagItem, keyName, err := parsePath("passwords/postgres/value")
	if err != nil {
		t.Error("Parsing return an error when it should not have")
	}
	if bagName != "passwords" {
		t.Errorf("bagName: %s != %s", bagName, "passwords")
	}
	if bagItem != "postgres" {
		t.Errorf("bagItem: %s != %s", bagItem, "postgres")
	}
	if keyName != "value" {
		t.Errorf("keyName: %s != %s", keyName, "value")
	}

}