package messages

import (
	"encoding/json"
	"math/rand"
	"time"
)

type Message struct {
	Symbol string
	ID     int32
	Trades []Trade
}

type Trade struct {
	TradeTimestamp time.Time
	Volume         int
	Bid            float32
	Ask            float32
}

func (message *Message) randomMessageAsBytes() ([]byte, error) {
	message.Symbol = getSymbol()
	message.ID = rand.Int31()
	message.Trades = getTrades()

	if data, err := json.Marshal(&message); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func getTrades() []Trade {
	trades := make([]Trade, rand.Intn(2)+1)

	for i := range trades {
		trades[i] = Trade{
			TradeTimestamp: time.Now(),
			Volume:         rand.Intn(1000000),
			Bid:            rand.Float32() + float32(rand.Intn(100)),
			Ask:            rand.Float32() + float32(rand.Intn(100)),
		}
	}

	return trades
}

func getSymbol() string {
	symbols := []string{
		"QQQ",
		"SPY",
		"MSFT",
		"NFLX",
	}

	return symbols[rand.Int()%len(symbols)]
}
