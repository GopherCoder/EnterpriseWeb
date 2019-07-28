package model

import "EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/database"

func ListLottery(ownerID int64) ([]Lottery, error) {
	var results []Lottery

	if dbError := database.Engine.ID(ownerID).Find(&results); dbError != nil {
		return results, dbError
	}
	return results, nil
}

func OneLottery(id int64) (Lottery, error) {
	var result Lottery
	if has, err := database.Engine.ID(id).Get(&result); !has || err != nil {
		return result, err
	}
	return result, nil
}

func InvolvementsLottery(adminID int64) ([]Lottery, error) {
	var results []Lottery
	var adminTakePartIn AdminTakePart
	if has, dbError := database.Engine.ID(adminID).Get(&adminTakePartIn); dbError != nil || !has {
		return results, dbError
	}

	if dbError := database.Engine.In("lottery_ids", adminTakePartIn.LotteryIds).Find(&results); dbError != nil {
		return results, dbError
	}
	return results, nil
}

func CreateLottery(params CreateLotteryParams) (*LotterySerializer, error) {
	var lottery Lottery
	var levels []Level

	tx := database.Engine.NewSession()
	tx.Begin()

	var winnerLottery WinnerLottery
	if has, dbError := tx.ID(params.WinnerLotteryCondition).Get(&winnerLottery); !has || dbError != nil {
		return nil, dbError
	}

	var admin Admin
	if has, dbError := tx.Id(params.AdminId).Get(&admin); !has || dbError != nil {
		return nil, dbError
	}

	for _, i := range params.Levels {
		if err := i.Valid(); err != nil {
			return nil, err
		}
		var one Level
		one = Level{
			Name:     i.Name,
			ImageURL: i.ImageURL,
			Number:   i.Number,
			Class:    i.Class,
		}
		if _, dbError := tx.InsertOne(&one); dbError != nil {
			tx.Rollback()
			return nil, dbError
		}
		levels = append(levels, one)
	}

	lottery = Lottery{
		Deadline:        params.Deadline,
		Levels:          levels,
		WinnerLotteryId: winnerLottery.Id,
		Class:           params.LotteryClass,
		Limit:           params.Limit,
		AdminID:         admin.Id,
	}
	if _, dbError := tx.InsertOne(&lottery); dbError != nil {
		tx.Rollback()
		return nil, dbError
	}
	tx.Commit()
	var result *LotterySerializer
	result = lottery.Serializer()
	return result, nil
}
