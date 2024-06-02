package main

// 次の仕様のデータ構造を考えてみてください
//	とあるゲームの得点を集計をするプログラム
//	ゲームの結果は0点から100点まで1点刻みで点数が付けられる
//	集計は複数回のゲームの結果をもとにユーザごとに行う
//	どういうデータ構造で1回のゲーム結果を表現すべきか
//	適切だと思うユーザ定義型を定義してください

import (
	"fmt"

	"github.com/google/uuid"
)

type Score struct {
	UserID string
	GameId int
	Point  int
}

func main() {
	score := Score{
		UserID: uuid.New().String(),
		GameId: 1,
		Point:  100,
	}
	fmt.Println(score)
}
