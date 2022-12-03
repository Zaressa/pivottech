package main


import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	publicKey := os.Getenv("MARVEL_PUBLIC_KEY")
	privateKey := os.Getenv("MARVEL_PRIVATE_KEY")

	client := marvelClient{
		publicKey:  publicKey,
		privateKey: privateKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	

type Character struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type marvelClient struct {
	publicKey  string
	privateKey string
	httpClient *http.Client
	BaseURL    string
}

type CharacterResponse struct {
	Code      int         `json:"code"`
	Status    string      `json:"status"`
	Etag      string      `json:"etag"`
	AttributionHTML string `json:"attributionHTML"`
	AttributionText string `json:"attributionText"`
	Data      struct {
		Offset      int         `json:"offset"`
		Limit       int         `json:"limit"`
		Total       int         `json:"total"`
		Count       int         `json:"count"`
		Results     []Character `json:"results"`
	} `json:"data"`
	
}
func (c *marvelClient) md5Hash(ts int64) string {
	hasher := md5.New()
	hasher.Write([]byte(ts + c.privateKey + c.publicKey))
	return hex.EncodeToString(hasher.Sum(nil))
}	

func (c *marvelClient) md5Hash (ts, string) string {
	ts := strconv.Itoa(ts)
	hash := md5.Sum([]byte(ts + c.privateKey + c.publicKey))
	return hex.EncodeToString(hash[:])

}
func (c *MarvelClient) signURL(url string) string {
	ts := time.Now().Unix()
	hash := c.md5Hash(ts)
	return fmt.Sprintf("%s&ts=%d&apikey=%s&hash=%s", url, ts, c.PubKey, hash)
}
func (c *marvelClient) GetCharacters(limit int) ([]Character, error) {
	limit := stronv.Itoa(limit)
	url := c.BaseURL + "/characters?limit=" + limit
	url = c.signURL(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	var CharacterResponse CharacterResponse
	err = json.NewDecoder(req.Body).Decode(&CharacterResponse)
	if err != nil {
		return nil, err
	}
	return CharacterResponse.Data.Results, nil
}


// Path: marvel-client/marvel.go