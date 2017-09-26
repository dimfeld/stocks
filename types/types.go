package types

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

// VendorSpecific holds information that isn't common to the supported platforms,
// and isn't vital, but might be interesting to use when it's present.
type VendorSpecific struct {
	Data map[string]string `json:"data"`
	// Keys defines a preferred order to print the keys.
	Keys []string `json:"keys"`
}

type Account struct {
	Id          string `json:"id"`
	Broker      string `json:"broker"`
	Description string `json:"description"`
}

func (a Account) String() string {
	return fmt.Sprintf("%s:%s - %s", a.Broker, a.Id, a.Description)
}

type ConnectionStatus struct {
	Connected bool
	Error     error
}

type Tristate int

const (
	Yes Tristate = iota
	No
	Maybe
)

type Quote struct {
	High   float64 `json:"high,omitempty"`
	Low    float64 `json:"low,omitempty"`
	Open   float64 `json:"open,omitempty"`
	Close  float64 `json:"close,omitempty"`
	Mark   float64 `json:"mark,omitempty"`
	Volume int64   `json:"volume,omitempty"`

	Bid     float64 `json:"bid,omitempty"`
	BidSize int64   `json:"bid_size,omitempty"`
	BidExch string  `json:"bid_exch,omitempty"`

	Ask     float64 `json:"ask,omitempty"`
	AskSize int64   `json:"ask_size,omitempty"`
	AskExch string  `json:"ask_exch,omitempty"`

	LastTime time.Time `json:"last_time,omitempty"`
	Last     float64   `json:"last,omitempty"`
	LastSize int64     `json:"last_size,omitempty"`
	LastExch string    `json:"last_exch,omitempty"`

	OptionHistoricalVolatility float64 `json:"option_hv,omitempty"`
	OptionImpliedVolatility    float64 `json:"option_iv,omitempty"`
	OptionCallOpenInt          int64   `json:"option_call_open_int,omitempty"`
	OptionCallVolume           int64   `json:"option_call_vol,omitempty"`
	OptionPutOpenInt           int64   `json:"option_put_open_int,omitempty"`
	OptionPutVolume            int64   `json:"option_put_vol,omitempty"`

	// Shortable Tristate `json:"shortable"`

	AvgVol float64 `json:"avg_vol,omitempty"` // Not supported by all brokers

	// YearHigh float64 `json:"year_low"`
	// YearLow  float64 `json:"year_high"`

	Time       time.Time `json:"time"`
	Incomplete bool      `json:"incomplete,omitempty"`
}

type OptionQuote struct {
	Quote
	Strike        float64   `json:"strike"`
	Underlying    string    `json:"underlying"`
	Expiration    string    `json:"expiration"`
	Type          PutOrCall `json:"type"`
	MinPriceDelta float64   `json:"min_price_delta"`
	FullSymbol    string    `json:"full_symbol"` // The full symbol for the option
	OpenInterest  int64     `json:"open_interest"`

	ModelPrice float64 `json:"model_price"`
	Delta      float64 `json:"delta"`
	Gamma      float64 `json:"gamma"`
	Theta      float64 `json:"theta"`
	Vega       float64 `json:"vega"`
	Rho        float64 `json:"rho"` // Not always supported
}

type SymbolType int

const (
	SymbolEquity SymbolType = iota
	SymbolOption
)

// func (t SymbolType) String() string {
// 	switch t {
// 	case SymbolEquity:
// 		return "Equity"
// 	case SymbolOption:
// 		return "Option"
// 	default:
// 		return "Unknown"
// 	}
// }

type SymbolDetails struct {
	Symbol      string
	Description string
	Vendor      VendorSpecific
}

type PutOrCall string

const (
	Put  PutOrCall = "PUT"
	Call           = "CALL"
)

type OptionChain struct {
	Underlying  string    `json:"underlying"`
	Multiplier  string    `json:"multiplier,omitempty"`
	Exchanges   []string  `json:"exchanges,omitempty"`
	Strikes     []float64 `json:"strikes"`
	Expirations []string  `json:"expirations"`
}

type Option struct {
	Underlying string
	Strike     float64
	Expiration string
	Type       PutOrCall
}

type OptionCombo struct {
	Legs []Option
}

type ErrorWithCode struct {
	error
	code int
}

func (ec ErrorWithCode) Code() int {
	return ec.code
}

var (
	ErrSymbolNotFound = ErrorWithCode{errors.New("symbol not found"), http.StatusNotFound}
	ErrDisconnected   = errors.New("broker disconnected")
)
