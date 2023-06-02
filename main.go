package main

import (
	"encoding/json"
	"fmt"
	svg "github.com/ajstarks/svgo"
	"github.com/joho/godotenv"
	"io"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func checkError(err error, httpCode int, httpRes http.ResponseWriter) bool {
	if err != nil {
		httpRes.WriteHeader(httpCode)
		fmt.Println(" [ERROR] - " + err.Error())
		return true
	}
	return false
}

func stringToInt64(s string) int64 {
	seed := int64(1)
	for _, c := range s {
		seed *= int64(c - '0')
	}
	return seed
}

func getWHUrlParams(req *http.Request, w *int, h *int, maxW int, maxH int) (err error) {
	if !req.URL.Query().Has("w") && !req.URL.Query().Has("h") {
		return nil
	}
	if req.URL.Query().Has("w") {
		wUrl, err := strconv.Atoi(req.URL.Query().Get("w"))
		if err != nil {
			return err
		}
		if wUrl < 1 || wUrl > maxW {
			return err
		}
		*w, *h = wUrl, wUrl
	}
	if req.URL.Query().Has("h") {
		hUrl, err := strconv.Atoi(req.URL.Query().Get("h"))
		if err != nil {
			return err
		}
		if hUrl < 1 || hUrl > maxH {
			return err
		}
		*h = hUrl
	}
	return nil
}

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fileData, err := os.ReadFile("index.html")
		if checkError(err, http.StatusInternalServerError, res) {
			return
		}
		res.Header().Set("Content-Type", "text/html")
		_, _ = res.Write(fileData)
	})

	// seed - required
	// w 	- width (if not defined 7) 								in range [1..100]
	// h 	- height (if not defined 7 or same as defined width)	in range [1..100]
	// <host>/render?seed=any&w=number&h=number
	http.HandleFunc("/hashmap", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("\n[hashMap]")
		if req.Method != "GET" {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		res.Header().Set("Content-Type", "image/svg+xml")

		// setting random seed
		seed := time.Now().UnixNano()
		if req.URL.Query().Has("seed") {
			seed = stringToInt64(req.URL.Query().Get("seed"))
		}
		rand.Seed(seed)
		fmt.Println("seed: ", seed)

		// getting width and height with default 7x7 and in range of [1..100]
		w, h := 7, 7
		checkError(
			getWHUrlParams(req, &w, &h, 100, 100),
			http.StatusBadRequest,
			res,
		)

		colors := []string{"#000", "#fff"}
		img := svg.New(res)
		img.Start(w, h)
		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				currentColor := colors[rand.Intn(2)]
				img.Rect(i, j, 1, 1, "fill:"+currentColor+";stroke:none")
			}
		}
		img.End()
	})

	http.HandleFunc("/picsum", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("\n[picsum]")
		if req.Method != "GET" {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		// setting random seed
		seed := strconv.Itoa(int(time.Now().UnixNano()))
		if req.URL.Query().Has("seed") {
			seed = req.URL.Query().Get("seed")
		}
		fmt.Println(" seed: " + seed)

		// fetching info from picsum/.../info and getting donwload url
		fetchedData, err := http.Get("https://picsum.photos/seed/" + seed + "/info")
		if checkError(err, http.StatusInternalServerError, res) {
			return
		}

		data, err := io.ReadAll(fetchedData.Body)
		if checkError(err, http.StatusInternalServerError, res) {
			return
		}

		type downloadUrl struct {
			Url string `json:"download_url"`
		}
		var url downloadUrl
		err = json.Unmarshal(data, &url)
		if checkError(err, http.StatusInternalServerError, res) {
			return
		}

		// processing download url
		w, h := 64, 64
		checkError(
			getWHUrlParams(req, &w, &h, math.MaxInt, math.MaxInt),
			http.StatusBadRequest,
			res,
		)

		isSecondSlash := false
		for i := len(url.Url) - 1; i >= -1; i-- {
			if url.Url[i] != '/' {
				continue
			}
			if !isSecondSlash {
				isSecondSlash = true
				continue
			}
			url.Url = url.Url[0 : i+1]
			break
		}
		url.Url = url.Url + strconv.Itoa(w) + "/" + strconv.Itoa(h)

		// reading image
		fetchedData, err = http.Get(url.Url)
		if checkError(err, http.StatusInternalServerError, res) {
			return
		}

		data, err = io.ReadAll(fetchedData.Body)
		if checkError(err, http.StatusInternalServerError, res) {
			return
		}

		// writing response
		res.Header().Set("Content-Type", req.Header.Get("Content-Type"))
		_, _ = res.Write(data)
		return
	})

	err := godotenv.Load(".env")
	fmt.Println("[SERVER] - starting on :" + os.Getenv("PORT") + " ...")
	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
