package main

// import (
// 	"fmt"
// 	"log"
// 	"time"
// )

// const totalLicenses = 5

// var usedLicenses = 0

// type ConnectionMaker interface {
// 	connect()
// }

// type User struct {
// 	name string
// }

// type Connection struct {
// 	user User
// }

// func (c Connection) connect() {
// 	fmt.Printf("\n%s connecting!\n", c.user.name)
// }

// type ProxyConnectionMaker struct {
// 	connectionMaker ConnectionMaker
// 	user            User
// }

// func (p *ProxyConnectionMaker) connect() {
// 	if usedLicenses < totalLicenses {
// 		usedLicenses++
// 		p.connectionMaker.connect()
// 	} else {
// 		log.Println("\nToo many users, please try again later.")
// 	}
// }

// func main() {

// 	time.AfterFunc(1*time.Second, func() { usedLicenses = 0 })

// 	user1 := User{name: "Tohan"}
// 	user2 := User{name: "Cemhan"}
// 	user3 := User{name: "Gökçe"}
// 	user4 := User{name: "Batu"}
// 	user5 := User{name: "Ezgi"}
// 	user6 := User{name: "Emrehan"}
// 	user7 := User{name: "Özüm"}

// 	connectionMaker := &[]ProxyConnectionMaker{
// 		{&Connection{user1}, user1},
// 		{&Connection{user2}, user2},
// 		{&Connection{user3}, user3},
// 		{&Connection{user4}, user4},
// 		{&Connection{user5}, user5},
// 		{&Connection{user6}, user6},
// 		{&Connection{user7}, user7},
// 	}

// 	for _, connectionMaker := range *connectionMaker {
// 		connectionMaker.connect()
// 	}

// 	connectionMaker2 := &[]ProxyConnectionMaker{
// 		{&Connection{user1}, user1},
// 		{&Connection{user2}, user2},
// 		{&Connection{user3}, user3},
// 		{&Connection{user4}, user4},
// 		{&Connection{user5}, user5},
// 	}

// 	for _, connectionMaker := range *connectionMaker2 {
// 		connectionMaker.connect()
// 	}

// 	time.Sleep(2 * time.Second)

// 	connectionMaker3 := &[]ProxyConnectionMaker{
// 		{&Connection{user1}, user1},
// 		{&Connection{user2}, user2},
// 		{&Connection{user3}, user3},
// 		{&Connection{user4}, user4},
// 		{&Connection{user5}, user5},
// 		{&Connection{user6}, user6},
// 		{&Connection{user7}, user7},
// 	}

// 	for _, connectionMaker := range *connectionMaker3 {
// 		connectionMaker.connect()
// 	}

// }
