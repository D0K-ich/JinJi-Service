package JinJi_Service

import (
	"fmt"
	"testing"

	"github.com/gocql/gocql"
)

func TestConnect(t  *testing.T) {
	// Создаем кластер и сессию для подключения к базе данных
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "test_keyspace"
	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Printf("failed to create session: %v", err)
		return
	}
	defer session.Close()

	fmt.Println("Connected to database!")
	// Ваш код для работы с базой данных
}
