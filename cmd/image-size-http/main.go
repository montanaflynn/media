package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/montanaflynn/media/size"
)

func main() {
	giphy := "https://media1.giphy.com/media/l0ErxFClZX9L3bgBi/giphy.gif"
	res, err := http.Get(giphy)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	size, err := mediasize.Parse(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonBytes, err := json.Marshal(size)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", jsonBytes)
	// {"Width":480,"Height":270,"ImageType":"GIF"}
}
