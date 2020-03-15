/*
Package media is for determining information
about media without decoding the entire file.

Example Usage:

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

MIT License Copyright (c) 2020 Montana Flynn (https://montanaflynn.com)
*/
package media
