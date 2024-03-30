package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func BuildGameGen() {
	var MySQLDSN = "root:Bo5zcd*2ozsvtjss0@tcp(127.0.0.1:3306)/x_live?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:      "xdb/game_query",
		ModelPkgPath: "game_model",
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
	})
	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}

func BuildOrderGen() {
	var MySQLDSN = "root:Bo5zcd*2ozsvtjss0@tcp(127.0.0.1:3306)/x_live?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:      "xdb/order_query",
		ModelPkgPath: "order_model",
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
	})
	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}

func main() {
	BuildGameGen()
	BuildOrderGen()
}
