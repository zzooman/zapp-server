package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/zzooman/zapp-server/utils"
)

func createRandomUser(t *testing.T) User {
	password := utils.RandomString(6)
	hashedPassword, err := utils.HashPassword(password)
	require.NoError(t, err)

	arg := CreateUserParams{
		Username: utils.RandomString(6),
		Password: hashedPassword,
		Email:    pgtype.Text{String: utils.RandomString(12), Valid: true},
		Phone:    pgtype.Text{String: "1234567890", Valid: true},		
	}

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Phone, user.Phone)	

	return user
}

func TestCreateUser(t *testing.T) {
	user := createRandomUser(t)
	require.NotEmpty(t, user)
	require.NotZero(t, user.Username)
	require.NotZero(t, user.CreatedAt)	
}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)
	err := testStore.DeleteUser(context.Background(), user.Username)
	
	require.NoError(t, err)
}


func TestGetUser(t *testing.T) {
	createdUser := createRandomUser(t)
	gettedUser, err :=  testStore.GetUser(context.Background(), createdUser.Username)

	require.NoError(t, err)
	require.NotEmpty(t, gettedUser)
	require.Equal(t, createdUser.Username, gettedUser.Username)
	require.Equal(t, createdUser.Username, gettedUser.Username)
	require.Equal(t, createdUser.Email, gettedUser.Email)
	require.Equal(t, createdUser.Phone, gettedUser.Phone)	
	require.Equal(t, createdUser.CreatedAt, gettedUser.CreatedAt)
}
func TestUpdateUser(t *testing.T) {
	user := createRandomUser(t)

	// Define the input parameters for the UpdateUser method
	params := UpdateUserParams{		
		Username: user.Username,
		Password: "newpassword",
		Phone:    pgtype.Text{String: "9876543210", Valid: true},
		Email:    pgtype.Text{String: utils.RandomString(12), Valid: true},
		Profile: pgtype.Text{String: "profile", Valid: true},
	}

	// Call the UpdateUser method
	_, err := testStore.UpdateUser(context.Background(), params)
	require.NoError(t, err)

	// Retrieve the updated user
	updatedUser, err := testStore.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)

	// Verify the updated user's fields
	require.Equal(t, params.Username, updatedUser.Username)
	require.Equal(t, params.Email, updatedUser.Email)
	require.Equal(t, params.Phone, updatedUser.Phone)	
	require.Equal(t, user.CreatedAt, updatedUser.CreatedAt)
}