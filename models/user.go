package models

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	db *gorm.DB
)

type User struct {
	ID               int
	Username         string
	Password         string
	Is_authenticated bool
}

func Connect() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	fmt.Println(".env file was loaded successfully")

	dbuser := os.Getenv("POSTGRES_USER")
	dbpasswd := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v", host, dbuser, dbpasswd, dbname, port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("There was an issue with while connecting to PostgreSQL")
	}

	fmt.Println("Connections was established successfully")

	db.AutoMigrate(&User{})

}

func All() {

	query := `select * from users;`
	var res []User

	db.Raw(query).Scan(&res)

	fmt.Printf("Here is users => \n\n")
	for i := range res {
		fmt.Printf("ID: %v; Username: %v; Password: %v; Is_authenticated: %v\n", res[i].ID, res[i].Username, res[i].Password, res[i].Is_authenticated)
	}
	fmt.Println()

}

func New(username, password string) {

	user := User{
		Username: username,
		Password: password,
	}

	db.Create(&user)

}

func ValidateUser(username, password string) bool {

	Connect()
	user := User{}

	err := db.Table("users").Where("username = ?", username).First(&user)

	fmt.Printf("\nPASSWORD = %v\n\n", user.Password)
	if err != nil {
		fmt.Printf("\n err = %v\n\n", err)
	}

	return user.Password == password
}

func CreateUser(username, password string) {

	Connect()
	New(username, password)

}
