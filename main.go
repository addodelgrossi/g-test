package main

import (
	"fmt"
)

func main() {

	interval := 1
	listenKey := fmt.Sprintf("%s@kline_%dm", "BTCUSDT", interval)
	fmt.Printf("using listen key %s\n", listenKey)

	wsHandler := func(message []byte) {
		fmt.Println(string(message))
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}

	// <symbol>@kline_<interval>
	doneC, _, err := WsFutureUserDataServe(listenKey, wsHandler, errHandler, &WsConfig{
		Endpoint: "wss://stream.binancefuture.com/ws",
	})

	if err != nil {
		panic(err)
	}
	<-doneC

}
