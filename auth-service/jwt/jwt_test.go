package jwt

import (
	"auth-service/config"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateToken(t *testing.T) {
	config.InitializeViper("../")
	config.ReadConfig()
	token, err := CreateToken("test-username", "user")
	require.NoError(t, err)
	require.NotEmpty(t, token)
}

func TestValidateAndExtractToken(t *testing.T) {
	config.InitializeViper("../")
	config.ReadConfig()
	token, err := CreateToken("test-username", "user")
	require.NoError(t, err)
	require.NotEmpty(t, token)

	claims, err := ValidateAndExtractToken(token)
	if err != nil {
		log.Panic(err.Error())
	}

	require.Equal(t, "test-username", claims.Username)
	require.Equal(t, "user", claims.Role)
}
