package dbhelper

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

type postGresConnector struct {
	host     string `yaml:"host"`
	port     int    `yaml:"port"`
	user     string `yaml:"user"`
	password string `yaml:"password"`
	dbname   string `yaml:"dbname"`
}
type Connection struct {
	postgresDescription postGresConnector `yaml:"postgres"`
	conn                *sql.DB
}

func (c *Connection) LoadConfig(path string) {
	// Read the file
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a map to hold the YAML data
	var config map[string]interface{}

	// Unmarshal the YAML data into the map
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the data
	fmt.Println(config)
	config = config["postgres"].(map[string]interface{})
	c.postgresDescription = postGresConnector{host: config["host"].(string), port: config["port"].(int), user: config["user"].(string), password: config["password"].(string), dbname: config["dbname"].(string)}
	fmt.Println(c)

}
func (c *Connection) Connect() {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.postgresDescription.user, c.postgresDescription.password, c.postgresDescription.host, c.postgresDescription.port, c.postgresDescription.dbname)
	log.Default().Println("Trying to Connect")

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	c.conn = conn
	// defer c.conn.Close()

	// rows, err := c.conn.Query("INSE")
	if err = c.conn.Ping(); err != nil {
		panic(err)
	}

}

func (c *Connection) InsertQuery(tableName string, values string, query string) (*sql.Rows, error) {
	q := fmt.Sprintf("INSERT INTO %s %s %s", tableName, values, query)
	return c.conn.Query(q)
}
func (c *Connection) GetAll(tableName string, query string) (*sql.Rows, error) {
	return c.conn.Query(fmt.Sprintf("SELECT * FROM %s %s", tableName, query))

}
func (c *Connection) GetCondition(tableName string, where string, query string) (*sql.Rows, error) {
	// fmt.Println(fmt.Sprintf("SELECT * FROM %s WHERE %s %s", tableName, where, query))
	return c.conn.Query(fmt.Sprintf("SELECT * FROM %s WHERE %s %s", tableName, where, query))

}
func (c *Connection) CloseConnection() {
	c.conn.Close()
}
