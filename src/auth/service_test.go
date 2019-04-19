// 
// 
// 

package auth_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aamsur/goshort/datastore/model"

	"git.qasico.com/cuxs/cuxs"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"os"
	"github.com/aamsur/goshort/test"
	"github.com/aamsur/goshort/src/auth"
)

func TestMain(m *testing.M) {
	test.Setup()

	// run tests
	res := m.Run()

	// cleanup
	test.DataCleanUp()

	os.Exit(res)
}

func LoginTokenTest(token string) (*auth.SessionData, error) {
	validAuth := "Bearer " + token

	req, _ := http.NewRequest(echo.GET, "/", nil)
	req.Header.Set(echo.HeaderAuthorization, validAuth)
	res := httptest.NewRecorder()

	e := cuxs.New()
	c := e.NewContext(req, res)
	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}

	h := cuxs.Authorized()(handler)
	if h(c) == nil {
		ctx := cuxs.NewContext(c)

		sd, e := auth.UserSession(ctx)

		return sd, e
	}

	return nil, errors.New("tidak dapat melakukan login")
}

func TestLogin(t *testing.T) {
	user := &model.User{Username: "testing_login"}
	user.Save()

	// mengecek lastlogin
	sd, e := auth.Login(user)
	assert.NoError(t, e)
	assert.NotEqual(t, user.LastLogin, sd.User.LastLogin)
}

func TestToken(t *testing.T) {
	user := &model.User{Username: "testing_token"}
	user.Save()

	// get token from username
	sd, e := auth.Login(user)

	// mencoba untuk login menggunakan token
	// dan mendapatkan user session dari context
	r, e := LoginTokenTest(sd.Token)
	assert.NoError(t, e)
	assert.NotEmpty(t, r.Token, "Seharusnya token ada")
}

func TestGetUsername(t *testing.T) {

	user := &model.User{Username: "testing_username"}
	user.Save()

	m, e := auth.GetUsername(user.Username)
	assert.NoError(t, e, "seharusnya tidak error")
	assert.Equal(t, user.Username, m.Username)
}
