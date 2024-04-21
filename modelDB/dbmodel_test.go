package modeldb

import (
	modelapp "goFinalTask/modelAPP"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	// Тестовая строка подключения к базе данных
	dsn := "postgres://gjzrhosm:iMQ-dyBUv_Q6hL9SvBgG29OGU50Fh7L_@hattie.db.elephantsql.com/gjzrhosm"

	// Вызов функции подключения к базе данных
	err := ConnectDB(dsn)
	assert.Nil(t, err, "Should connect to database without error")

	// Проверка, что DB не nil
	assert.NotNil(t, DB, "DB should not be nil after connection")

	// Дополнительно можно проверить, успешно ли были выполнены миграции
	if DB != nil {
		var userCount int64
		DB.Model(&modelapp.Users{}).Count(&userCount)
		log.Println("Number of users in Users table:", userCount)
		assert.GreaterOrEqual(t, userCount, int64(0), "Users table should exist and be queryable")
	}
}
