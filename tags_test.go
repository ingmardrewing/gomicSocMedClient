package main

import "testing"

func TestGetTagsForTwitter(t *testing.T) {
	expected := 17
	actual := len(GetTagsForTwitter(""))

	if actual != expected {
		t.Error("Expected to get", expected, "tags, but got", actual)
	}

	expected = 2
	actual = len(GetTagsForTwitter("What the f**king hell is going on around here? How much text do we actually need to decrease the number of tag returned?"))

	if actual != expected {
		t.Error("Expected to get", expected, "tags, but got", actual)
	}
}
