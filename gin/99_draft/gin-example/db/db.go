// Define the package interacting with the database
package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

// DBConnection defines the connection structure
type DBConnection struct {
	session *mgo.Session
}

// NewConnection handles connecting to a mongo database
func NewConnection(host string, dbName string) (conn *DBConnection) {
	fmt.Println("host: " + host)
	info := &mgo.DialInfo{
		// Address if its a local db then the value host=localhost
		Addrs: []string{host},
		// Timeout when a failure to connect to db
		Timeout: 60 * time.Second,
		// Database name
		Database: "admin",
		// Database credentials if your db is protected
		//Username: os.Getenv("DB_USER"),
		//Password: os.Getenv("DB_PWD"),
		Username: "MyRoot",
		Password: "MyRootPassw0rd",
	}
	fmt.Println(info.Username)
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	conn = &DBConnection{session}
	return conn
}

// Use handles connect to a certain collection
func (conn *DBConnection) Use(dbName, tableName string) (collection *mgo.Collection) {
	// This returns method that interacts with a specific collection and table
	return conn.session.DB(dbName).C(tableName)
}

// Close handles closing a database connection
func (conn *DBConnection) Close() {

	// This closes the connection
	conn.session.Close()
	return
}
