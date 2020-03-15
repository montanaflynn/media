package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/montanaflynn/media"
)

func main() {
	giphy := "https://media1.giphy.com/media/l0ErxFClZX9L3bgBi/giphy.gif"
	r, err := http.Get(giphy)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	m, err := media.Parse(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("media dimensions: %v\n", m.Size())
	// media dimensions: {480 270}

	fmt.Printf("content type: %q\n", m.Type())
	// content type: "image/gif"

	j, err := json.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", j)
	// {"ContentType":"image/gif","Dimensions":{"Width":480,"Height":270}}
}
