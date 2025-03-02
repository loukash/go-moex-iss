package main

import (
	"github.com/Ruvad39/go-moex-iss"
	"log/slog"
)

func main() {

	// создание клиента
	client, err := iss.NewClient()
	if err != nil {
		slog.Error("main", "NewClient", err.Error())
	}
	iss.SetLogLevel(slog.LevelDebug)

	//Sec, err := client.GetBondsInfo("tqob") // ОФЗ
	//Sec, err := client.GetBondsInfo("TQCB") // Корпоративные
	// выборка по boardGroup = все облигации
	Sec, err := client.GetBondsInfo("", 58)

	if err != nil {
		slog.Error("main", "ошибка GetBondsInfo", err.Error())
	}

	slog.Info("GetBondsInfo", slog.Int("всего len(Sec)", len(Sec)))
	for n, sec := range Sec {
		slog.Info("bonds",
			"row", n,
			"SecID", sec.SecID,
			"SecName", sec.SecName,
			"board", sec.BoardID,
			"CouponDate", sec.NextCoupon,
			"CouponValue", sec.CouponValue,
			"PrevPrice", sec.PrevPrice,
		)
	}

	// текущие рыночные данные по облигации
	//SecData, err := client.GetBondsData("tqob", "SU26218RMFS6")
	//SecData, err := client.GetBondsData("tqob", "")
	//SecData, err := client.GetBondsData("", "", 58)
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
	//SecHistory, err := client.GetBondHistory("SU26218RMFS6", "2025-02-01", "2025-02-21")
	////SecHistory, err := client.GetBondsHistoryDate("2025-02-21")
	//if err != nil {
	//	slog.Error("main", "ошибка GetBondHistory", err.Error())
	//}
	//
	//slog.Info("GetBondHistory", slog.Int("всего len(Sec)", len(SecHistory)))
	//for row, sec := range SecHistory {
	//	slog.Info(strconv.Itoa(row),
	//		"SecHistory", sec,
	//	)
	//}

}
