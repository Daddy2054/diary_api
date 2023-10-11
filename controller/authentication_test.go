package controller

import (
        "diary_api/model"
        "encoding/json"
        "net/http"
        "testing"

        "github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
        newUser := model.AuthenticationInput{
                Username: "yemiwebby",
                Password: "test",
        }
        writer := makeRequest("POST", "/auth/register", newUser, false)
        assert.Equal(t, http.StatusCreated, writer.Code)
}

func TestLogin(t *testing.T) {
        user := model.AuthenticationInput{
                Username: "yemiwebby",
                Password: "test",
        }

        writer := makeRequest("POST", "/auth/login", user, false)

        assert.Equal(t, http.StatusOK, writer.Code)

        var response map[string]string
        json.Unmarshal(writer.Body.Bytes(), &response)
        _, exists := response["jwt"]

        assert.Equal(t, true, exists)
}
