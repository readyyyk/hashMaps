package main

import (
	svg "github.com/ajstarks/svgo"
	"github.com/joho/godotenv"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// seed - required
	// w 	- width (if not defined 7) 								in range [1..100]
	// h 	- height (if not defined 7 or same as defined width)	in range [1..100]
	// <host>/render?seed=any&w=number&h=number
	http.HandleFunc("/render", func(res http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		res.Header().Set("Content-Type", "image/svg+xml")

		// setting random seed
		seedStr := []rune(req.URL.Query().Get("seed"))
		seed := int64(1)
		for c := range seedStr {
			seed *= int64(c - '0')
		}
		rand.Seed(seed)

		// getting width and height with default 7x7 and in range of [1..100]
		w, h := 7, 7
		if req.URL.Query().Has("w") {
			wUrl, err := strconv.Atoi(req.URL.Query().Get("w"))
			if err != nil {
				res.WriteHeader(http.StatusBadRequest)
				return
			}
			if wUrl < 1 || wUrl > 100 {
				res.WriteHeader(http.StatusBadRequest)
				return
			}
			w, h = wUrl, wUrl
		}
		if req.URL.Query().Has("w") {
			hUrl, err := strconv.Atoi(req.URL.Query().Get("h"))
			if err != nil {
				res.WriteHeader(http.StatusBadRequest)
				return
			}
			if hUrl < 1 || hUrl > 100 {
				res.WriteHeader(http.StatusBadRequest)
				return
			}
			h = hUrl
		}

		colors := []string{"#000", "#fff"}
		img := svg.New(res)
		img.Start(w, h)
		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				rand.Shuffle(2, func(i, j int) {
					colors[i], colors[j] = colors[j], colors[i]
				})
				currentColor := colors[0]
				img.Rect(i, j, 1, 1, "fill:"+currentColor+";stroke:none")
			}
		}
		img.End()
	})
	err := godotenv.Load(".env")
	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
