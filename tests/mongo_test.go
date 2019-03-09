package tests

import (
	"log"
	"testing"

	Entities "../entities"
	Infra "../infra"
	Services "../services"
)

const (
	dbName             = "test_db"
	userCollectionName = "user"
)

func Test_UserService(t *testing.T) {
	t.Run("CreateUser", createUser_should_insert_user_into_mongo)
}

func createUser_should_insert_user_into_mongo(t *testing.T) {
	//Arrange
	config := Infra.GetConfigurations()

	if config.DbConnection == "" {
		log.Fatalf("Unable to connect to mongo: connection string is empty")
	}

	session, err := Infra.NewSession(config.DbConnection, dbName)
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	defer session.Close()
	userService := Services.NewUserService(session.Copy())

	testUsername := "integration_test_user"
	testPassword := "integration_test_password"
	user := Entities.User{
		Username: testUsername,
		Password: testPassword}

	//Act
	err = userService.Create(&user)

	//Assert
	if err != nil {
		t.Errorf("Unable to create user: %s", err)
	}
	var results []Entities.User
	session.GetCollection(dbName, userCollectionName).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Error("Incorrect number of results. Expected `1`, got: `%i`", count)
	}
	if results[0].Username != user.Username {
		t.Errorf("Incorrect Username. Expected `%s`, Got: `%s`", testUsername, results[0].Username)
	}
}
