package utils

import (
	"crypto/md5"
	"encoding/hex"
	// _ tells go that we want to import so we can use the drivers without ever referencing the library directly in code
	_ "github.com/lib/pq"
)

const serverPort = 3333

// The types of
type ComicBook struct {
	Title string
	year  int
	id    int
}

type Series struct {
	Title       string
	Description string
	Id          int
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func ExecuteMigration() {
	//privateAPIKey := "e754f147e1f5268c1c1443eb9ce7383d30a6e5ec"
	//publicAPIKey := "72871c7e8329864eea4b8fa177884e51"
	//
	//now := time.Now().Unix()
	//
	//fmt.Println(now)
	//
	//nowString := strconv.FormatInt(now, 16)
	//
	//hash_this := nowString + privateAPIKey + publicAPIKey
	//
	//hash := GetMD5Hash(hash_this)

	//fmt.Println(hash)

	//url := fmt.Sprintf("http://gateway.marvel.com/v1/public/stories/1810/events?limit=10&ts=%s&apikey=%s&hash=%s", nowString, publicAPIKey, hash)
	//
	//res, err := http.Get(url)
	//
	//if err != nil {
	//	fmt.Printf("error making http request: %s\n", err)
	//	os.Exit(1)
	//}

	//res_body, _ := ioutil.ReadAll(res.Body)

	//fmt.Printf("client: got response!\n")
	//fmt.Printf("client: status code: %d\n", res.StatusCode)
	//
	//fmt.Printf(string(res_body))

	// Start db access here
	DbConnect()
	querySeries := "select title, description, id from series;"
	SelectRows(querySeries)

}
