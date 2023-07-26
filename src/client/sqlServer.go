//This example uses the ORM jet
package client

import (
	"database/sql"
	"github.com/denisenkom/go-mssqldb"

	"fmt"
	"sync"
)

var (
	instanceClient CommandClient
	onceClient     sync.Once
)

type CommandClient interface {
	Connect() error
	DeleteParameter(commandSql string, interfaceSql interface{}) error
	SelectOneParameter(commandSql string, interfaceSql interface{}) (string, error)
	SelectTwoParameter(commandSql string, interfaceSql interface{}, interfaceSql_second interface{}) (string, error)
	
}

type clientSql struct {
	db *sql.DB
}

func GetInstance() CommandClient {
	onceClient.Do(func() {
		instanceClient = &clientSql{}
	})
	return instanceClient
}

func (c *clientSql) Connect() error {
	query := url.Values{}
	query.Add("app name", "MyAppName")
	
	u := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(username, password),
		Host:     fmt.Sprintf("%s:%d", hostname, port),
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	db, err := sql.Open("sqlserver", u.String())
	if err != nil {
		return err
	}

	return nil
}

func (c *clientSql) DeleteParameter(commandSql string, interfaceSql interface{}) error {
	_, err := c.db.Exec(commandSql, interfaceSql)
	if err != nil {
		return err
	}

	return nil
}
func (c *clientSql) SelectOneParameter(commandSql string, interfaceSql interface{}) (string, error) {
	var result string
	row := c.db.QueryRow(commandSql, interfaceSql)
	err := row.Scan(&result)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
			return result, nil
		}
	}
	return result, err
}

func (c *clientSql) SelectTwoParameter(commandSql string, interfaceSql interface{}, interfaceSql_second interface{}) (string, error) {
	var result string
	row := c.db.QueryRow(commandSql, interfaceSql, interfaceSql_second)
	err := row.Scan(&result)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
			return "", nil
		} else {
			return "", err
		}
	}
	return result, nil
}


func (c *clientSql) Insert(TableName string, interfaceSql interface{}, interfaceSql_second interface{}) (string, error) {
	stmt, err := db.Prepare("INSERT INTO @p1 VALUES ()")
	tin, err := time.Parse(time.RFC3339, "2006-01-02T22:04:05.7870015-07:00")   // data containing 7 decimal digits and has time zone awareness
	param := mssql.DateTime1(tin)   // data is stored in mssql.DateTime1 type
	_, err = stmt.Exec(TableName)
}