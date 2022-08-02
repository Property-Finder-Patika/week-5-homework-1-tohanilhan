package main

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

// Total number of licenses. It should be constant for the whole program.
const totalLicenses = 6

// usedLicenses is a global variable that counts the number of connections that are in use.
var usedLicenses = 0

// ConnectionMaker is an interface that defines the connect method.
type ConnectionMaker interface {
	connect()
}

// User is a struct that holds the name of the user.
type User struct {
	name string
}

// Connection is a struct that holds the user.
type Connection struct {
	user User
}

// connect method of Connection struct that prints the connected user name.
func (c Connection) connect() {
	log.Infof("Welcome aboard %s ! You are connected!", c.user.name)
}

// ProxyConnectionMaker is a struct that holds the connection maker and the user.
type ProxyConnectionMaker struct {
	connectionMaker ConnectionMaker
	user            User
}

// connect method of ProxyConnectionMaker struct that checks the number of licenses and prints the connected user name.
func (p *ProxyConnectionMaker) connect() {
	if usedLicenses < totalLicenses {
		usedLicenses++
		p.connectionMaker.connect()
	} else {
		log.Warnf("All connections are in use at the moment, please try again later dear %s.", p.user.name)
	}
}

func main() {

	log.Info("Hello. This is a sample program using proxy design pattern.")
	log.Warn("If you want to exit the program, please press Ctrl+C.")

	for {
		// After some time, usedLicenses will be reset to 0.
		time.AfterFunc(3*time.Second, func() {
			// Reset the usedLicenses with a random number between 0 and totalLicenses
			rand.Seed(time.Now().UnixNano())
			usedLicenses = rand.Intn(totalLicenses)
		})

		// connect concurrently
		connectConcurrently()

		// Next line for better visualization.
		fmt.Println("\n ")

		// Sleep for a while to reset usedLicenses.
		time.Sleep(5 * time.Second)
	}

}

func connectConcurrently() {
	// Initialize the users.
	user1 := User{name: "User1"}
	user2 := User{name: "User2"}
	user3 := User{name: "User3"}
	user4 := User{name: "User4"}
	user5 := User{name: "User5"}
	user6 := User{name: "User6"}

	// Go routines to connect the users.

	go func() {
		connectionMaker := &ProxyConnectionMaker{&Connection{user: user1}, user1}
		connectionMaker.connect()
	}()
	go func() {
		connectionMaker := &ProxyConnectionMaker{&Connection{user: user2}, user2}
		connectionMaker.connect()
	}()
	go func() {
		connectionMaker := &ProxyConnectionMaker{&Connection{user: user3}, user3}
		connectionMaker.connect()
	}()
	go func() {
		connectionMaker := &ProxyConnectionMaker{&Connection{user: user4}, user4}
		connectionMaker.connect()
	}()
	go func() {
		connectionMaker := &ProxyConnectionMaker{&Connection{user: user5}, user5}
		connectionMaker.connect()
	}()
	go func() {
		connectionMaker := &ProxyConnectionMaker{&Connection{user: user6}, user6}
		connectionMaker.connect()
	}()
}
