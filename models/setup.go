package models

import (
	"fmt"
	"log"

	"github.com/awesome-sphere/as-seating/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/sharding"
)

var DB *gorm.DB

func InitDatabase() {
	dbUser := utils.GetenvOr("POSTGRES_USER", "pkinwza")
	dbPassword := utils.GetenvOr("POSTGRES_PASSWORD", "securepassword")
	dbHost := utils.GetenvOr("POSTGRES_HOST", "localhost")
	dbPort := utils.GetenvOr("POSTGRES_PORT", "5432")
	dbName := utils.GetenvOr("POSTGRES_DB", "as-cinema")

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&SeatType{})
	db.AutoMigrate(&Seat{}, &Theater{})
	db.AutoMigrate(&TimeSlots_1{}, &TimeSlots_2{}, &TimeSlots_3{}, &TimeSlots_4{}, &TimeSlots_5{})
	db.AutoMigrate(&SeatInfos_1{}, &SeatInfos_2{}, &SeatInfos_3{}, &SeatInfos_4{}, &SeatInfos_5{})
	db.Use(sharding.Register(sharding.Config{
		ShardingKey:         "theater_id",
		NumberOfShards:      6,
		PrimaryKeyGenerator: sharding.PKPGSequence,
	}, "time_slots", "seat_infos"))
	DB = db
	// testSharding := UseSharding()
	// var test []SeatInfo
	// DB.Model(&SeatInfo{}).Where("theater_id", 1).Find(&test)
	// fmt.Printf("%#v\n", test)
}
