package main

import "testing"

func TestEncode(t *testing.T) {

	user := Input{
		text:      "leonardo",
		shift:     7,
		direction: "encode",
	}

	populateAlphabet()
	result := cipher(user.text, user.shift, user.direction)

	if result != "slvuhykv" {
		t.Errorf("Expected = %s; got %s", user.text, result)
	}
}

func TestDecode(t *testing.T) {

	user := Input{
		text:      "slvuhykv",
		shift:     7,
		direction: "decode",
	}

	populateAlphabet()
	result := cipher(user.text, user.shift, user.direction)

	if result != "leonardo" {
		t.Errorf("Expected = %s; got %s", user.text, result)
	}
}
