package main

import (
	"database/sql"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"regexp"
	"strings"
	"time"
)

// Check for errors and raise a panic if they exist
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Takes sprite tag text and returns appropriate text
// Just uses the sprite name in title case
func spriter(bytes []byte) []byte {
	s := string(bytes)
	re := regexp.MustCompile("grail|lantern|heart|forge|moth|winter|edge")
	m := re.FindString(s)
	if m == "" {
		fmt.Printf("Error getting sprite tag %s\n", bytes)
		return bytes
	}
	cased := strings.Title(strings.ToLower(m))
	return []byte(cased)
}

// Trims tags
func trimtag(bytes []byte) []byte {
	s := string(bytes)
	trimmed := s[3 : len(s)-4]
	return []byte(trimmed)
}

// Adds appropriate formatting by fixing escaped quotes and the HTML-style tags.
// The hex values are because we are using this Unicode set for formatted text:
// https://en.wikipedia.org/wiki/Mathematical_Alphanumeric_Symbols#Tables_of_styled_letters_and_digits
func format(entry string) string {
	re := regexp.MustCompile("''")
	quoted := re.ReplaceAllString(entry, "'")

	re = regexp.MustCompile("<br>")
	newlined := re.ReplaceAllString(quoted, "\n")

	re = regexp.MustCompile("<b>.*</b>")
	offset := re.ReplaceAllFunc([]byte(newlined), trimtag)

	re = regexp.MustCompile("<i>.*</i>")
	italiced := re.ReplaceAllFunc([]byte(offset), trimtag)

	re = regexp.MustCompile("<sprite.*?>")
	sprited := re.ReplaceAllFunc([]byte(italiced), spriter)

	return string(sprited)
}

// Tweet an entry string with the API, printing any errors
func tweet(api *anaconda.TwitterApi, entry string) {
	tweet, err := api.PostTweet(entry, nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("-- %s\n%s\n", time.Now().String(), tweet.Text)
	}
}

// Returns a random entry from the SQLite DB
// Limited to grabbing < 280 character entries that haven't been posted
// Will then increment the entry's postcount.
// This will eventually panic, once everything has been posted.
func getRandomEntry() string {
	var id int
	var entry string
	db, err := sql.Open("sqlite3", "./data/core.db")
	check(err)
	defer db.Close()
	row := db.QueryRow(
		`SELECT id, entry
		FROM entries
		WHERE LENGTH(entry)<=280
		ORDER BY RANDOM() LIMIT 1;`)
	err = row.Scan(&id, &entry)
	check(err)
	_, err = db.Exec(
		`UPDATE entries
		SET postcount = postcount + 1
		WHERE id = $1;`, id)
	check(err)
	return entry
}

// Load secret tokens/keys from the OS environment & create the API
// Then loop forever, grabbing an entry and tweeting it periodically
func main() {
	godotenv.Load()
	s := fmt.Sprintf("%sm", os.Getenv("TWEET_TIME_INTERVAL_MINUTES"))
	seconds, err := time.ParseDuration(s)
	check(err)
	api := anaconda.NewTwitterApiWithCredentials(
		os.Getenv("ACCESS_TOKEN"),
		os.Getenv("ACCESS_TOKEN_SECRET"),
		os.Getenv("CONSUMER_KEY"),
		os.Getenv("CONSUMER_KEY_SECRET"))
	for {
		entry := getRandomEntry()
		entry = format(entry)
		tweet(api, entry)
		time.Sleep(seconds)
	}
}
