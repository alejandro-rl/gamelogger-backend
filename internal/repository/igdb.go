package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func RequestImageHash(igdb_id int) string {

	type Response struct {
		ID       int    `json:"id"`
		Image_id string `json:"image_id"`
	}

	url := "https://api.igdb.com/v4/covers"
	method := "POST"

	query := fmt.Sprintf(`
	fields image_id;
	where game = %d;
	`, igdb_id)

	payload := strings.NewReader(query)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Client-ID", os.Getenv("IGDBID"))
	req.Header.Add("Authorization", "Bearer "+os.Getenv("IGDBAUTH"))
	req.Header.Add("Content-Type", "text/plain")

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var response []Response

	json.Unmarshal(body, &response)

	return response[0].Image_id

}

// func prettyPrint(i interface{}) string {
// 	s, _ := json.MarshalIndent(i, "", "\t")
// 	return string(s)
// }
