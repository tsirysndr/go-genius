package main

import (
	"encoding/json"
	"fmt"

	"github.com/tsirysndr/go-genius"
)

func main() {
	token := "u--H9zhSY3zB4a2sF_7tEgxtETKGsCzqs_KKUJ7q8Ez8NXOG1lp6jjj0uLKrbBBG"
	client := genius.NewClient(token)
	res, _ := client.Artists.Get("16775")
	artist, _ := json.Marshal(res)
	fmt.Println(string(artist))
}
