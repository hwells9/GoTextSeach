package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const serverPort = 3333

// The types of
type ComicBook struct {
	Title string
	year  int
	id    int
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func main() {
	now := time.Now().Unix()

	fmt.Println(now)

	nowString := strconv.FormatInt(now, 16)

	hash_this := nowString + privateAPIKey + publicAPIKey

	fmt.Println(hash_this)

	hash := GetMD5Hash(hash_this)

	fmt.Println(hash)

	// url := fmt.Sprintf("https://gateway.marvel.com:443/v1/public/comics?limit=10&ts=%s&apikey=%s&hash=%s", nowString, publicAPIKey, hash)

	url := fmt.Sprintf("http://gateway.marvel.com/v1/public/stories/1810/events?limit=10&ts=%s&apikey=%s&hash=%s", nowString, publicAPIKey, hash)

	res, err := http.Get(url)

	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	res_body, _ := ioutil.ReadAll(res.Body)

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	fmt.Printf(string(res_body))
}
