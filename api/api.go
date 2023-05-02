package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"govulnapi/api/database"
	m "govulnapi/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type Api struct {
	db               *database.DB
	router           *chi.Mux
	coins            []m.Coin
	currentDate      time.Time
	dayDuration      time.Duration
	coingeckoBaseUrl string
	listenAddress    string
	jwtAuth          *jwtauth.JWTAuth
}

func New(listenAddress string, coingeckoBaseUrl string) *Api {
	db := database.Init("api.db")
	virtualTime := time.Date(2014, time.January, 1, 0, 0, 0, 0, time.UTC)
	priceRefreshInterval := time.Minute
	coins, err := db.GetCoins()
	if err != nil {
		log.Fatalln(err)
	}

	api := Api{
		db:               db,
		router:           chi.NewRouter(),
		currentDate:      virtualTime,
		dayDuration:      priceRefreshInterval,
		coins:            coins,
		coingeckoBaseUrl: coingeckoBaseUrl,
		listenAddress:    listenAddress,
		// CWE-547: Use of Hard-coded, Security-relevant Constants
		jwtAuth: jwtauth.New("HS256", []byte("safe-secret"), nil),
	}

	return &api
}

func (a *Api) Run() {
	go a.managePrices()
	a.setupRoutes()
	log.Println("Starting API ...")
	// CWE-319: Cleartext Transmission of Sensitive Information
	log.Fatalln(http.ListenAndServe(a.listenAddress, a.router))
}

func (a *Api) Shutdown() {
	a.db.Close()
}

func (a *Api) managePrices() {
	log.Println("Starting price management daemon ...")
	for {
		a.refreshCoins()
		a.currentDate = a.currentDate.Add(time.Hour * 24)
		time.Sleep(a.dayDuration)
	}
}

func (a *Api) refreshCoins() {
	var (
		coins []m.Coin
		r     *http.Response
		err   error
	)

	url := fmt.Sprintf("%s/coins/%v", a.coingeckoBaseUrl, a.currentDate.UnixMilli())

	for {
		r, err = http.Get(url)
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}

	json.NewDecoder(r.Body).Decode(&coins)
	a.coins = coins
}

func (a *Api) getCoin(coin_id string) (m.Coin, error) {
	for _, coin := range a.coins {
		if coin.Id == coin_id {
			return coin, nil
		}
	}
	return m.Coin{}, errors.New("Requested coin doesn't exist!")
}
