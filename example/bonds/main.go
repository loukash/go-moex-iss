package main

import (
	"github.com/Ruvad39/go-moex-iss"
	"log/slog"
	"strconv"
)

func main() {

	// создание клиента
	client, err := iss.NewClient()
	if err != nil {
		slog.Error("main", "NewClient", err.Error())
	}
	iss.SetLogLevel(slog.LevelDebug)

	// получить список бондов
	// ОФЗ: board = tqob
	// Корпоративные: board = TQIR
	/*
		Sec, err := client.GetBondsInfo("tqob")
		//Sec, err := client.GetBondsInfo("TQCB")

		if err != nil {
			slog.Error("main", "ошибка GetBondsInfo", err.Error())
		}

		slog.Info("GetBondsInfo", slog.Int("всего len(Sec)", len(Sec)))
		for row, sec := range Sec {

			slog.Info(strconv.Itoa(row),
				"sec", sec,
			)
			//slog.Info("sec", slog.Any("sec", sec))
		}
	*/

	// текущие рыночные данные по облигации
	//SecData, err := client.GetBondsData("tqob", "SU26218RMFS6")
	////SecData, err := client.GetBondsData("tqob", "")
	//if err != nil {
	//	slog.Error("main", "ошибка GetBondsData", err.Error())
	//}
	//
	//slog.Info("GetBondsData", slog.Int("всего len(Sec)", len(SecData)))
	//for row, sec := range SecData {
	//	slog.Info(strconv.Itoa(row),
	//		"SecData", sec,
	//	)
	//}

	// Исторические данные
	SecHistory, err := client.GetBondHistory("SU26218RMFS6", "2025-02-01", "2025-02-21")
	//SecHistory, err := client.GetBondsHistoryDate("2025-02-21")
	if err != nil {
		slog.Error("main", "ошибка GetBondHistory", err.Error())
	}

	slog.Info("GetBondHistory", slog.Int("всего len(Sec)", len(SecHistory)))
	for row, sec := range SecHistory {
		slog.Info(strconv.Itoa(row),
			"SecHistory", sec,
		)
	}

}
