package handler

import (
	"bytes"
	"encoding/json"
	"dts-fga-swswg/product-middleware-ut/dto"
	"dts-fga-swswg/product-middleware-ut/pkg/errs"
	"dts-fga-swswg/product-middleware-ut/service/service_mocks"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserHandler_Register_Success(t *testing.T) {
	userService := service_mocks.NewUserServiceMock()

	userHandler := NewUserHandler(userService)

	payload := dto.NewUserRequest{
		Email:    "test@mail.com",
		Password: "123456",
	}

	bs, err := json.Marshal(payload)

	require.Nil(t, err)

	service_mocks.CreateNewUser = func(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
		result := dto.NewUserResponse{
			Result:     "success",
			Message:    "registered successfully",
			StatusCode: http.StatusCreated,
		}

		return &result, nil
	}

	request, err := http.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(bs))

	require.Nil(t, err)

	w := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)

	route := gin.Default()

	route.POST("/users/register", userHandler.Register)

	route.ServeHTTP(w, request)

	response := w.Result()

	responseBody, err := ioutil.ReadAll(response.Body)

	require.Nil(t, err)

	defer response.Body.Close()

	var userResponse dto.NewUserResponse

	err = json.Unmarshal(responseBody, &userResponse)

	require.Nil(t, nil)

	assert.Equal(t, http.StatusCreated, userResponse.StatusCode)

	assert.Equal(t, "success", userResponse.Result)
}

func TestUserHandler_Register_InvalidJsonError(t *testing.T) {
	userHandler := NewUserHandler(nil)

	payload :=
		`
		{
			"email "test@mail.com",
			"password": "123456"
		}
	`

	request, err := http.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer([]byte(payload)))

	require.Nil(t, err)

	w := httptest.NewRecorder()

	gin.SetMode(gin.TestMode)

	route := gin.Default()

	route.POST("/users/register", userHandler.Register)

	route.ServeHTTP(w, request)

	response := w.Result()

	responseBody, err := ioutil.ReadAll(response.Body)

	require.Nil(t, err)

	defer response.Body.Close()

	var errMessage errs.MessageErrData

	err = json.Unmarshal(responseBody, &errMessage)

	require.Nil(t, err)

	assert.Equal(t, http.StatusUnprocessableEntity, errMessage.Status())

	assert.Equal(t, "invalid request body", errMessage.Message())

	assert.Equal(t, "INVALID_REQUEST_BODY", errMessage.Error())
}
