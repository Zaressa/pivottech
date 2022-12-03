package main

// marvel.NewClient()
//client.GetCharacters()
import (
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
	publicKey := os.Getenv("marvel_public_key")
	privateKey := os.Getenv("marvel_private_key")

	client := marvelClient{
		publicKey:  publicKey,
		privateKey: privateKey,
		httpClient: &http.Client{
			Timeout: 20 * time.Second,
		},
	}
	charc, err := client.getCharacters() // this is where I want to call the function
	if err != nil {
		log.Println("action failed", err)
		return

	}
	log.Println(charc)
}

type marvelClient struct {
	publicKey  string
	privateKey string
	httpClient *http.Client
}

func (c *marvelClient) getCharacters() ([]Character, error) {
	var characters []Character
	req, err := http.NewRequest("GET", "https://gateway.marvel.com:443/v1/public/characters", nil) //contains the endpoint
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("apikey", c.publicKey)
	q.Add("ts", "1")
	q.Add("hash", "d41d8cd98f00b204e9800998ecf8427e")
	req.URL.RawQuery = q.Encode()
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	res, err := c.httpClient.Get("https://gateway.marvel.com/v1/public/characters") //contians the endpoint
	if err != nil {
		log.Println("action failed", err)
		return nil, err
	}

	defer res.Body.Close()

	var characterResponse CharacterResponse
	if err := json.NewDecoder(res.Body).Decode(&characterResponse); err != nil {
		return nil, err

	}
	return nil, err

	}
	defer resp.Body.Close()
	
