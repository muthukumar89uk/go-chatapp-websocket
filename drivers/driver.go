package drivers

import (
	"chatApp/helpers"
	"chatApp/models"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func DbConnection() {

	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	userName := viper.GetString(`DB.userName`)
	password := viper.GetString(`DB.password`)
	host := viper.GetString(`DB.host`)
	port := viper.GetString(`DB.port`)
	dbName := viper.GetString(`DB.dbName`)

	connectionURI := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", userName, password, host, port, dbName)

	//connecting to postgres-SQL
	helpers.DB, err = gorm.Open(postgres.Open(connectionURI), &gorm.Config{})
	if err != nil {
		panic("could not connect to DB!")
	}
	
	fmt.Println("DataBase Connected Successfully!!!")
}

func Migration() {
	err = helpers.DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Failed to create Tables:", err)
		return
	}
	fmt.Println("Tables created successfully")
}
