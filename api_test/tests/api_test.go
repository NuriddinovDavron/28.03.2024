package tests

import (
	"api_exam/api_test/handlers"
	"api_exam/api_test/storage"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApi(t *testing.T) {
	gin.SetMode(gin.TestMode)
	require.NoError(t, SetupMinimumInstance(""))
	buffer, err := OpenFile("user.json")

	require.NoError(t, err)
	// User Create
	req := NewRequest(http.MethodPost, "/users", buffer)
	res := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/users", handlers.CreateUser)
	r.ServeHTTP(res, req)
	//assert.Equal(t, http.StatusOK, res.Code)

	var user storage.User
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &user))

	require.Equal(t, "nuriddinovdavron2003@gmail.com", user.Email)
	require.Equal(t, "Davron", user.FirstName)
	require.Equal(t, "Nuriddinov", user.LastName)
	require.Equal(t, "qwertyui123546", user.Password)
	require.NotNil(t, user.Id)

	// GetUser
	getReq := NewRequest(http.MethodGet, "/users/:id", nil)
	q := getReq.URL.Query()
	q.Add("id", user.Id)
	getReq.URL.RawQuery = q.Encode()
	getRes := httptest.NewRecorder()
	r = gin.Default()
	r.GET("/users/:id", handlers.GetUser)
	r.ServeHTTP(getRes, getReq)
	assert.Equal(t, http.StatusOK, getRes.Code)
	var getUserResp storage.User
	bodyBytes, err := io.ReadAll(getRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &getUserResp))
	assert.Equal(t, user.Id, getUserResp.Id)
	assert.Equal(t, user.FirstName, getUserResp.FirstName)
	assert.Equal(t, user.LastName, getUserResp.LastName)
	assert.Equal(t, user.Password, getUserResp.Password)
	assert.Equal(t, user.Email, getUserResp.Email)

	// User List
	listReq := NewRequest(http.MethodGet, "/users", nil)
	listRes := httptest.NewRecorder()
	r = gin.Default()
	r.GET("/users", handlers.ListUsers)
	r.ServeHTTP(listRes, listReq)
	assert.Equal(t, http.StatusOK, listRes.Code)
	bodyBytes, err = io.ReadAll(listRes.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bodyBytes)

	// User Delete
	delReq := NewRequest(http.MethodDelete, "/user/:id"+user.Id, nil)
	delRes := httptest.NewRecorder()
	r.DELETE("/user/:id", handlers.DeleteUser)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var respm storage.Message
	bodyBytes, _ = io.ReadAll(delRes.Body)
	require.NoError(t, json.Unmarshal(bodyBytes, &respm))
	require.Equal(t, "user was deleted successfully", respm.Message)

	// User Register
	regReq := NewRequest(http.MethodPost, "/register", buffer)
	regRes := httptest.NewRecorder()
	r.POST("/register", handlers.RegisterUser)
	r.ServeHTTP(regRes, regReq)
	assert.Equal(t, http.StatusOK, regRes.Code)
	var resp storage.Message
	bodyBytes, err = io.ReadAll(regRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &resp))
	require.NotNil(t, resp.Message)
	require.Equal(t, "One time password sent to your email", resp.Message)

	// User Verify with correct code
	verURLCorrect := "/user/verification"
	verReqCorrect := NewRequest(http.MethodGet, verURLCorrect, buffer)
	verResCorrect := httptest.NewRecorder()
	r = gin.Default()
	r.GET("/user/verification", handlers.Verify)
	r.ServeHTTP(verResCorrect, verReqCorrect)

	assert.Equal(t, http.StatusOK, verResCorrect.Code)
	var responseCorrect storage.Message
	bodyBytesCorrect, err := io.ReadAll(verResCorrect.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytesCorrect, &responseCorrect))
	require.Equal(t, "Success", responseCorrect.Message)

	// User Verify with incorrect code
	verURLIncorrect := "/user/verification"
	verReqIncorrect := NewRequest(http.MethodGet, verURLIncorrect, buffer)
	verResIncorrect := httptest.NewRecorder()
	r.ServeHTTP(verResIncorrect, verReqIncorrect)

	assert.Equal(t, http.StatusBadRequest, verResIncorrect.Code)
	var responseIncorrect storage.Message
	bodyBytesIncorrect, err := io.ReadAll(verResIncorrect.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytesIncorrect, &responseIncorrect))
	require.Equal(t, "Incorrect code", responseIncorrect.Message)

	//PRODUCT TEST

	gin.SetMode(gin.TestMode)
	require.NoError(t, SetupMinimumInstance(""))
	buffer, err = OpenFile("product.json")

	require.NoError(t, err)

	// Product create
	req = NewRequest(http.MethodPost, "/products", buffer)
	res = httptest.NewRecorder()
	r = gin.Default()
	r.POST("/products", handlers.CreateProduct)
	r.ServeHTTP(res, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.Code)
	var product storage.Product
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &product))
	require.Equal(t, product.Description, "chevrolet")
	require.Equal(t, product.Name, "car")
	require.Equal(t, product.Price, int64(45))

	// Get Product
	getReq = NewRequest(http.MethodGet, "/products/by/:id", buffer)
	q = getReq.URL.Query()
	q.Add("id", product.Id)
	getReq.URL.RawQuery = q.Encode()
	getRes = httptest.NewRecorder()
	r = gin.Default()
	r.GET("/products/by/:id", handlers.GetProduct)
	r.ServeHTTP(getRes, getReq)
	assert.Equal(t, http.StatusOK, getRes.Code)
	var getProduct storage.Product
	bodyBytes, err = io.ReadAll(getRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &getProduct))
	require.Equal(t, product.Id, getProduct.Id)
	require.Equal(t, product.Description, getProduct.Description)
	require.Equal(t, product.Name, getProduct.Name)
	require.Equal(t, product.Price, getProduct.Price)

	// List Products
	listReqq := NewRequest(http.MethodGet, "/products", buffer)
	listRess := httptest.NewRecorder()
	r = gin.Default()
	r.GET("/products", handlers.ListProducts)
	r.ServeHTTP(listRess, listReqq)
	assert.Equal(t, http.StatusOK, listRess.Code)
	bodyBytes, err = io.ReadAll(listRess.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bodyBytes)

	// Delete Product
	delReq = NewRequest(http.MethodDelete, "/products/:id"+product.Id, buffer)
	delRes = httptest.NewRecorder()
	r.DELETE("/products/:id", handlers.DeleteProduct)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var respmessage storage.Message
	bodyBytes, _ = io.ReadAll(delRes.Body)
	require.NoError(t, json.Unmarshal(bodyBytes, &respmessage))
	require.Equal(t, "product was deleted successfully", respmessage.Message)

}
