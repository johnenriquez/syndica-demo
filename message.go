package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Message struct {
	ID      int
	User    string
	From    string
	Message string
	Time    string
}

func AddMessage(m Message) error {
	insert, err := GetDB().Query("INSERT INTO messages (user, `from`, message) VALUES (?,?,?)", m.User, m.From, m.Message)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	insert.Close()
	return nil
}

func GetMessages(email string) (responses []Message) {
	rows, err := GetDB().Query("SELECT user, `from`, message, time FROM messages WHERE user=? OR `from`=? ORDER BY time desc", email, email)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var m Message
		err := rows.Scan(&m.User, &m.From, &m.Message, &m.Time)
		if err != nil {
			log.Println(err.Error())
			break
		}
		responses = append(responses, m)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return responses
}
