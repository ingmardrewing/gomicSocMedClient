package main

import (
	"fmt"
	"os"
	"testing"

	shared "github.com/ingmardrewing/gomicSocMedShared"
)

func TestCreateJson(t *testing.T) {
	link := "linkvalue"
	imgUrl := "imgUrlValue"
	title := "titleValue"
	tagsCsv := "tagsCsvValue"
	description := "descriptionValue"

	expected := fmt.Sprintf(JSON_FORMAT, link, imgUrl, title, tagsCsv, description)
	actual := createJson(title, description, link, imgUrl, tagsCsv)

	if actual != expected {
		t.Error("Expected json was", expected, "but got", actual)
	}
}

func TestCreateCredentials(t *testing.T) {
	os.Setenv(shared.GOMIC_BASIC_AUTH_USER, "user")
	os.Setenv(shared.GOMIC_BASIC_AUTH_PASS, "asdfasdf")

	expected := "'user:asdfasdf'"
	actual := createCredentials()

	if actual != expected {
		t.Error("Expected credentials were", expected, "but got", actual)
	}
}

func TestCreateTargetUrl(t *testing.T) {
	os.Setenv(shared.GOMIC_SOCMED_PROD_URL, "http://example.com")
	os.Setenv(shared.GOMIC_SOCMED_PROD_PORT, "1234")

	expected := "http://example.com:1234/0.1/gomic/socmed/echo"
	actual := createTargetUrl()

	if actual != expected {
		t.Error("Expected url was", expected, "but got", actual)
	}
}

func TestCreateCurl(t *testing.T) {
	credentials := "user:'asdfasdf'"
	target := "http://example.com:0234/0.1/gomic/socmed/echo"
	json := `{"Link":"linkvalue","ImgUrl":"imgUrlValue","Title":"titleValue","TagsCsvString":"tagsCsvValue","Description":"descriptionValue"}`

	expected := `curl -X POST -H "Content-Type application/json; charset=utf-8" -d '` + json + `' -u ` + credentials + " " + target
	actual := createCurl(json, credentials, target)

	if actual != expected {
		t.Error("Expected curl command was\n", expected, "\nbut got\n", actual)
	}
}
