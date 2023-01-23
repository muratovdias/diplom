package repository

import (
	"fmt"

	"github.com/kamalshkeir/klog"
	"github.com/kamalshkeir/korm"
	"github.com/kamalshkeir/pgdriver"
	"github.com/muratovdias/diplom/internal/app/config"
	"github.com/muratovdias/diplom/internal/models"
)

func IniitDB(config *config.GlobalConfig) error {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	fmt.Println("error env")
	// 	return err
	// }
	s := fmt.Sprintf("%s:%s@localhost:%s", config.PsgDB.User, config.PsgDB.Pwd, config.PsgDB.Port)
	pgdriver.Use()
	err = korm.New(korm.POSTGRES, config.PsgDB.Name, s)
	if err != nil {
		fmt.Println("error connect")
		return err
	}
	return nil
}

func CreateTables() error {
	err = korm.AutoMigrate[models.Trainer]("trainer")
	if klog.CheckError(err) {
		fmt.Println("error create tables")
		return err
	}
	return nil
}
