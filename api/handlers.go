package api

import (
	"encoding/json"
	"net/http"

	m "govulnapi/models"
)

// @Summary		  Coin data
// @Description	Get data for coins
// @Tags			  Coins
// @Produce		  json
// @Success	   	200	"ok"
// @Failure	    500	"internal server error"
// @Router			/coins [get]
func (s *Api) getCoins(w http.ResponseWriter, r *http.Request) {
	var response []byte
	coins, err := json.Marshal(s.coins)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response = []byte(err.Error())
	} else {
		w.Header().Set("Content-Type", "application/json")
		response = coins
	}

	w.Write(response)
}

// @Summary		  Get coin balances
// @Description	Fetches coin balances
// @Tags		    Trading
// @Produce	    json
// @Success	    200	"ok"
// @Failure	    401	"unauthorized"
// @Router			/balances/coin [get]
// @Security		Bearer
func (s *Api) getCoinBalances(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(m.User)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.CoinBalances)
}

// @Summary		  Get usd balances
// @Description	Fetches usd balances
// @Tags		    Trading
// @Produce	    json
// @Success	    200	"ok"
// @Failure	    401	"unauthorized"
// @Router			/balances/usd [get]
// @Security		Bearer
func (s *Api) getUsdBalances(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(m.User)

	w.Header().Set("Content-Type", "application/json")
	usdBalances := map[string]float64{
		"UsdBalance":         user.UsdBalance,
		"UsdStartingBalance": user.UsdStartingBalance,
	}
	json.NewEncoder(w).Encode(usdBalances)
}

// @Summary		  Get past orders
// @Description	Fetches past orders
// @Tags		    Trading
// @Produce	    json
// @Success	    200	"ok"
// @Failure	    401	"unauthorized"
// @Router			/orders [get]
// @Security		Bearer
func (s *Api) getOrders(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(m.User)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.Orders)
}

// @Summary		  Buy/sell coins
// @Description	Creates new buy/sell order
// @Tags		    Trading
// @Accept	    json
// @Produce	    plain
// @Param		    order	body		m.Order	true	"New order"
// @Success	    200	"order went through"
// @Failure	    401	"unauthorized"
// @Failure	    404	"requested coin not found"
// @Failure	    500	"internal server error"
// @Router			/orders [post]
// @Security		Bearer
func (s *Api) addOrder(w http.ResponseWriter, r *http.Request) {
	var (
		response = "Order successfully made!"
		user     = r.Context().Value("user").(m.User)
	)

	var order m.Order
	// CWE-472: External Control of Assumed-Immutable Web Parameter
	// CWE-639: Authorization Bypass Through User-Controlled Key
	// CWE-915: Improperly Controlled Modification of Dynamically-Determined Object Attributes
	// Any logged in user can set the hidden "userId" field in POST
	// request, enabling the user to make orders for another users
	// without knowing their credentials
	order.UserId = user.Id

	// CWE-20: Improper Input Validation
	json.NewDecoder(r.Body).Decode(&order)

	coin, err := s.getCoin(order.CoinId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response = err.Error()
	} else if err = s.db.AddOrder(order.UserId, coin.Id, coin.Price, order.IsBuy, order.Qty); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response = err.Error()
	}

	w.Write([]byte(response))
}

// @Summary		  Get past transactions
// @Description	Fetches past transactions
// @Tags		    Transactions
// @Produce	    json
// @Success	    200	"ok"
// @Failure	    401	"unauthorized"
// @Router			/transactions [get]
// @Security		Bearer
func (s *Api) getTransactions(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(m.User)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.Transactions)
}

// @Summary		  Send coins
// @Description	Creates new transaction
// @Tags		    Transactions
// @Accept	    json
// @Produce	    plain
// @Param		    transaction	body m.Transaction	true	"New transaction"
// @Success	    200	"transaction went through"
// @Failure	    400	"bad request"
// @Failure	    401	"unauthorized"
// @Failure	    412	"not enough coin"
// @Router			/transactions [post]
// @Security		Bearer
func (a *Api) addTransaction(w http.ResponseWriter, r *http.Request) {
	var (
		response = "Transaction successfully made!"
		user     = r.Context().Value("user").(m.User)
	)

	var transaction m.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = err.Error()
	} else {
		err := a.db.AddTransaction(user.Id, transaction.CoinId, transaction.Address, transaction.Qty, transaction.Note)
		if err != nil {
			w.WriteHeader(http.StatusPreconditionFailed)
			response = err.Error()
		}
	}

	w.Write([]byte(response))
}
