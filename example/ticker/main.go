package main

import (
	iss "github.com/Ruvad39/go-moex-iss"
	"log/slog"
	"strconv"
)

func main() {
	// создание клиента
	// получить данные по стакану можно только с авторизацией

	//url := iss.NewIssRequest().Target("securities").Json().MetaData(false).OnlySecurities().URL()
	//slog.Info("url", slog.Any("url", url))
	user := ""
	pwd := ""
	client, err := iss.NewClient(iss.WithUser(user), iss.WithPwd(pwd))
	if err != nil {
		slog.Error("main", "NewClient", err.Error())
	}
	//iss.SetLogLevel(slog.LevelDebug)

	Sec, err := client.GetTickersAll()
	slog.Info("GetTickersAll", slog.Int("всего len(Sec)", len(Sec)))
	for row, sec := range Sec {
		slog.Info(strconv.Itoa(row),
			"sec", sec,
		)
	}
	//https://iss.moex.com/iss/securities.json
	/*
		user := ""
		pwd := ""
		client, err := iss.NewClient(iss.WithUser(user), iss.WithPwd(pwd))
		if err != nil {
			slog.Error("main", "NewClient", err.Error())
		}
		//iss.SetLogLevel(slog.LevelDebug)

		// создание (поиск) тикера
		ticker, err := client.GetTicker("SBER")
		//ticker, err := client.GetTicker("RTS-9.24")
		if err != nil {
			slog.Error("main", "ошибка NewTicker", err.Error())
		}
		slog.Info("NewTicker", slog.Any("Ticker", ticker))

		// параметры инструмента
		info, err := ticker.Info()
		if err != nil {
			slog.Error("main", "ошибка ticker.Info", err.Error())
		}
		slog.Info("ticker.Info", slog.Any("t_info", info))

		// текущие рыночные данные
		data, err := ticker.Data()
		if err != nil {
			slog.Error("main", "ошибка ticker.data", err.Error())
		}
		slog.Info("ticker.Info", slog.Any("t_data", data))

		// свечи
		candles, err := ticker.Candles(iss.Interval_D1, "2025-01-01", "2025-02-01")
		if err != nil {
			slog.Error("main", "ошибка Candles", err.Error())
			return
		}
		slog.Info("Candles",
			"всего len(candles)", candles.Len(),
			"mindate", candles.First().Begin,
			"maxdate", candles.Last(0).Begin,
			"Symbol", candles.Symbol,
			"Interval", candles.Interval,
		)

		// цикл по списку свечей
		for row, candle := range candles.Data {
			slog.Info(strconv.Itoa(row),
				"time", candle.Time(),
				"candleData", candle,
			)
		}

		// стакан. Нужна авторизация
		orderBook, err := ticker.OrderBook()
		if err != nil {
			slog.Error("main", "ticker.OrderBook", err.Error())
			return
		}
		bid, _ := orderBook.BestBid()
		ask, _ := orderBook.BestAsk()
		bidVolume := orderBook.Bids.SumDepth()
		askVolume := orderBook.Asks.SumDepth()

		slog.Info("orderBook", "BestAsk", ask.Price, "BestBid", bid.Price, "объем асков", askVolume, "объем бидов", bidVolume)
		fmt.Println(orderBook.String())
	*/

}
