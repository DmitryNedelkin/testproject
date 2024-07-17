package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Message struct {
	Index    int    `json:"index"`
	Id       string `json:"id"`
	Guid     string `json:"guid"`
	IsActive bool   `json:"isActive"`
}

type Messages []Message

var (
	AllMessages = []Messages{}
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pgpwd4habr"
	dbname   = "postgres"
)

func CreateDatabase() {
	log.Println("CreateDatabase start")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connected!")

	createTable(db)
}

func createTable(db *sql.DB) {
	_, err := db.Exec(`
        CREATE TABLE users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            email VARCHAR(255) UNIQUE NOT NULL,
            password VARCHAR(255) NOT NULL
        )
    `)

	if err != nil {
		panic(err)
	}

	fmt.Println("table created")
}

func GetMessageFromDatabase(index int) *Message {
	for i, message := range AllMessages {
		if message[i].Index == index {
			return &message[i]
		}
	}

	return nil
}

func GetAllMessages() []Messages {
	return AllMessages
}

func PutMessagesToDatabase(messages []Message) {
	AllMessages = append(AllMessages, messages)
}
