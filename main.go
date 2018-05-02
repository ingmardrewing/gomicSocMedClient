package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	shared "github.com/ingmardrewing/gomicSocMedShared"
)

const (
	JSON_FORMAT = `{"Link":"%s","ImgUrl":"%s","Title":"%s","TagsCsvString":"%s","Description":"%s"}`
	CURL_FORMAT = `curl -X POST -H "Content-Type application/json; charset=utf-8" -d '%s' -u %s %s`
)

type Content struct {
	shared.Content
}

func main() {
	title := askUser("Title?")
	description := askUser("Description?")
	link := askUser("Link to post?")
	imgUrl := askUser("Image url?")
	tags := askUser("Tags (comma separated)?")

	if len(tags) == 0 {
		tags = strings.Join(TAGS, ",")
	}

	json := createJson(title, description, link, imgUrl, tags)
	credentials := createCredentials()
	target := createTargetUrl()

	curlCommand := createCurl(json, credentials, target)

	fmt.Println(curlCommand)

	/*
		curl -X POST -H "Content-Type: application/json; charset=utf-8" -d '{"Link":"https://devabo.de/2017/06/17/Airstrike","ImgUrl":"https://s3-us-west-1.amazonaws.com/devabode-us/comicstrips/DevAbode_0088.png","Title":"#88 Airstrike","TagsCsvString":"webcomic,graphicnovel,comic,comicart,comics,sciencefiction,scifi,geek,nerd,art,artist,artwork,blackandwhite,concept,conceptart,create,creative,design,digital,draw,drawing,drawings,dystopy,fantasy,humor,illustration,illustrator,image,imagination,ink,inked,inking,kunst,malen,malerei,narrative,parody,pulp,sketch,sketchbook,tusche,zeichnen,zeichnung","Description":"Alien elite forces taking their first hit at the people of the abode."}' -u ingmar:'YLy3DS$KE54$U0!!V)sKXoePepWvaE8Ypb9/z' http://stellarco.de:8880/0.1/gomic/socmed/facebook/publish
	*/

}

func createJson(title, description, link, imgUrl, tags string) string {
	return fmt.Sprintf(JSON_FORMAT, link, imgUrl, title, tags, description)
}

func createCredentials() string {
	user := shared.Env(shared.GOMIC_BASIC_AUTH_USER)
	pass := shared.Env(shared.GOMIC_BASIC_AUTH_PASS)
	return fmt.Sprintf("'%s:%s'", user, pass)
}

func createTargetUrl() string {
	basicUrl := shared.Env(shared.GOMIC_SOCMED_PROD_URL)
	port := shared.Env(shared.GOMIC_SOCMED_PROD_PORT)
	restVersion := shared.CURRENT_REST_VERSION
	restBasePath := shared.REST_BASE_PATH
	echoPath := shared.REST_PATH_ECHO

	return fmt.Sprintf("%s:%s/%s/%s%s", basicUrl, port, restVersion, restBasePath, echoPath)
}

func createCurl(json, credentials, target string) string {
	return fmt.Sprintf(CURL_FORMAT, json, credentials, target)
}

func askUser(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt + ": ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSuffix(text, "\n")
}
