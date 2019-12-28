package Connections

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Databases struct {
	Databases []Database `json:"databases"`
}

// User struct which contains a name
// a type and a list of social links
type Database struct {
	DriverName string `json:"driverName"`
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	Host       string `json:"host"`
	DbName     string `json:"dbName" default:"test"`
	Port       int    `json:"port"`
}

func getConfig() Databases {

	// Open our jsonFile
	jsonFile, err := os.Open(DB_FILE_NAME)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var databases Databases

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'databases' which we defined above
	json.Unmarshal(byteValue, &databases)

	return databases;
}

func (Databases) EstablishConnection(database Database) (*sql.DB, error) {

	//set defaults..
	if database.DbName == "" && database.DriverName == "" {
		database.setDefaults()
	}

	databasesList := getConfig();

	// we iterate through every user within our databases array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(databasesList.Databases); i++ {

		if databasesList.Databases[i].DbName == database.DbName {
			dbInfo := databasesList.Databases[i]
			query := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbInfo.UserName, dbInfo.Password, dbInfo.Host, dbInfo.Port, dbInfo.DbName)
			db, err := sql.Open(dbInfo.DriverName, query)
			return db, err;
		}
	}

	panic("DB connection went wrong, please check the connection settings!!!")
}

func (d *Database) setDefaults() {
	d.DriverName = "mysql"
	d.DbName = TEXT_DEFAULT_CONNECTION
}
