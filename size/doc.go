/*
Package size is for determining dimensions
of media without decoding the entire file.

Example Usage:

	giphy := "https://media1.giphy.com/media/l0ErxFClZX9L3bgBi/giphy.gif"
	res, err := http.Get(giphy)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	size, err := size.Parse(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonBytes, err := json.Marshal(size)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", jsonBytes)
	// {"Width":480,"Height":270,"MediaType":"GIF"}

MIT License Copyright (c) 2020 Montana Flynn (https://montanaflynn.com)
*/
package size
