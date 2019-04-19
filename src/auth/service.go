// 
// 
// 

package auth

import (
	"errors"
	"time"
	"github.com/aamsur/goshort/datastore/model"

	"git.qasico.com/cuxs/orm"
	"git.qasico.com/cuxs/cuxs"
	"github.com/dgrijalva/jwt-go"
)

// SessionData structur data current user logged in.
type SessionData struct {
	Token string      `json:"token"`
	User  *model.User `json:"user"`
}

// Login mendapatkan session data dari model user
// user diasumsikan sudah valid untuk login.
// jadi disini tidak ada validasi untuk login, hanya
// untuk mendapatkan session data
func Login(user *model.User) (sd *SessionData, e error) {
	//application menu
	if sd, e = StartSession(user.ID); e == nil {
		// update last login dari user tersebut
		user.LastLogin = time.Now()
		user.Save("LastLogin")
		return sd, nil
	}
	return nil, e
}

// StartSession mendapatkan data user entity dengan token
// untuk menandakan session user yang sedang login.
func StartSession(userID int64, token ...string) (sd *SessionData, e error) {
	sd = new(SessionData)

	// buat token baru atau menggunakan yang sebelumnya
	if len(token) == 0 {
		sd.Token = cuxs.JwtToken("id", userID)
	} else {
		sd.Token = token[0]
	}

	// membaca data user terlebih dahulu untuk
	sd.User = &model.User{ID: userID}
	sd.User.Read()

	return
}

// UserSession mendapatkan session data dari user yang mengirimkan request.
func UserSession(ctx *cuxs.Context) (*SessionData, error) {
	if u := ctx.Get("user"); u != nil {
		c := u.(*jwt.Token).Claims.(jwt.MapClaims)
		var userID int64

		// id adalah user id
		if c["id"] != nil {
			userID = int64(c["id"].(float64))
		}

		// memakai token sebelumnya
		token := ctx.Get("user").(*jwt.Token).Raw

		sd, _ := StartSession(userID, token)

		if c.VerifyIssuedAt(sd.User.LastLogoutAt.Unix(), true) {
			return nil, errors.New("expired jwt token")
		}

		// failed to get username, it's mean the user was gone / something change
		// please use another token
		if sd.User.Username == "" {
			return nil, errors.New("expired jwt token")
		}

		return sd, nil
	}

	return nil, errors.New("invalid jwt token")
}

// Logout
func Logout(ctx *cuxs.Context) (*model.User, error) {
	if u := ctx.Get("user"); u != nil {
		c := u.(*jwt.Token).Claims.(jwt.MapClaims)
		var userID int64

		// id adalah user id
		if c["id"] != nil {
			userID = int64(c["id"].(float64))
		}

		u := &model.User{ID: userID}
		u.LastLogoutAt = time.Now()
		u.Save("LastLogoutAt")

		return u, nil
	}

	return nil, errors.New("invalid jwt token")
}

// GetUsername get data user by username
func GetUsername(username string) (m *model.User, e error) {

	o := orm.NewOrm()
	if e = o.Raw("select * from user where username = ?", username).QueryRow(&m); e == nil {
		return m, nil
	}

	return nil, e
}
