package main

const (
	MAX_LENGTH = 140
)

var TAGS = []string{"inked", "inking", "illustration", "drawing", "drawings", "art", "artwork", "draw", "painting", "sketch", "sketchbook", "artist", "comics", "comicart", "comic", "graphicnovel", "design", "concept", "conceptart", "create", "creative", "image", "imagination"}

// Returns tags fitting into a tweet with
// the given text
func GetTagsForTwitter(txt string) []string {
	addedTags := []string{}
	length := len(txt) + 1

	for _, tag := range TAGS {
		future_length := length + len(tag) + 1
		if future_length > MAX_LENGTH {
			return addedTags
		}
		addedTags = append(addedTags, tag)
		length = future_length
	}

	return addedTags
}
