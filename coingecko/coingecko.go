package coingecko

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	m "govulnapi/models"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//go:embed coin_data.zip
var coinDataZip []byte

type jsonCoin struct {
	Prices       [][]float64 `json:"prices"`
	MarketCaps   [][]float64 `json:"market_caps"`
	TotalVolumes [][]float64 `json:"total_volumes"`
}

type Coingecko struct {
	router        *mux.Router // CWE-1104: Use of Unmaintained Third Party Components
	coins         map[string][]m.Coin
	listenAddress string
}

func New(listenAddress string) *Coingecko {
	coingecko := Coingecko{
		router:        mux.NewRouter(),
		coins:         map[string][]m.Coin{},
		listenAddress: listenAddress,
	}

	return &coingecko
}

func (c *Coingecko) Run() {
	c.loadZipCoinData()
	c.setupRoutes()
	log.Println("Starting virtual coingecko ...")
	log.Fatalln(http.ListenAndServe(c.listenAddress, c.router))
}

func (c *Coingecko) loadZipCoinData() {
	// Unzip market data
	r, err := zip.NewReader(bytes.NewReader(coinDataZip), int64(len(coinDataZip)))
	if err != nil {
		log.Fatalln(err)
	}

	// Iterate over each coin file
	for _, f := range r.File {
		coinName := strings.Split(f.Name, ".")[0]
		rc, err := f.Open()
		if err != nil {
			log.Fatalln(err)
		}
		defer rc.Close()

		var coinData jsonCoin
		json.NewDecoder(rc).Decode(&coinData)

		// Parse individual coin fields
		for _, v := range coinData.Prices {
			date := fmt.Sprintf("%v", int(v[0]))
			price := v[1]

			coin := m.Coin{
				Id:    coinName,
				Price: price,
			}

			c.coins[date] = append(c.coins[date], coin)
		}
	}
}

func (c *Coingecko) setupRoutes() {
	c.router.HandleFunc("/coins", c.getCoins)
	c.router.HandleFunc("/coins/{date}", c.getCoinsOnDate)
}

func (c *Coingecko) getCoinsOnDate(w http.ResponseWriter, r *http.Request) {
	date := mux.Vars(r)["date"]
	json.NewEncoder(w).Encode(c.coins[date])
}

func (c *Coingecko) getCoins(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(c.coins)
}
