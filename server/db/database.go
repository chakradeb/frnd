package db

import (
	"fmt"
	"strings"

	"github.com/gocql/gocql"

	"github.com/chakradeb/frnd-server/models"
)

type IDBClient interface {
	CreateUser (string, string) error
	GetUser(string) *models.User
	CheckUserAlreadyExists (string) bool
}

type DB struct {
	session *gocql.Session
}

func New(clusterIPs []string, keyspace string) (*DB, error) {
	cluster := gocql.NewCluster(clusterIPs[:]...)
	cluster.Keyspace = keyspace
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("db: connecting to \"%s\"", strings.Join(clusterIPs, ","))
	}

	return &DB{
		session: session,
	}, nil
}

func (d *DB) CreateUser(username string, password string) error {
	return d.session.Query(
		"INSERT INTO user_creds(username, password) VALUES (?, ?)", username, password,
	).Exec()
}

func (d *DB) GetUser(username string) *models.User {
	user := &models.User{}
	m := map[string]interface{}{}
	iter := d.session.Query("SELECT * FROM user_creds Where username = ?", username).Iter()
	for iter.MapScan(m) {
		user.Username = m["username"].(string)
		user.Password = m["password"].(string)
	}
	return user
}

func (d *DB) CheckUserAlreadyExists(username string) bool {
	var user string
	_ = d.session.Query(
		`SELECT username FROM user_creds WHERE username = ? LIMIT 1`, username,
	).Consistency(
		gocql.One,
	).Scan(&user)
	if user != "" {
		return true
	}
	return false
}
