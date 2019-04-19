// 
// 
// 

package auth

import (
	"github.com/aamsur/goshort/datastore/model"

	"git.qasico.com/cuxs/common"
	"git.qasico.com/cuxs/validation"
)

// SignInRequest data struct that stored request data when requesting an create auth process.
// All data must be provided and must be match with specification validation below.
// handler function should be bind this with context to matches incoming request
// data keys to the defined json tag.
type SignInRequest struct {
	Username string      `valid:"required" json:"username"`
	Password string      `valid:"required" json:"password"`
	User     *model.User `json:"-"`
}

// Validate implement validation.Requests interfaces.
func (r *SignInRequest) Validate() *validation.Output {
	o := &validation.Output{Valid: true}

	// check username di user
	user := new(model.User)
	user.Username = r.Username
	if e := user.Read("Username"); e != nil {
		o.Failure("username", "Invalid credential, please check your username or password.")
	} else {
		// cek password sesuai dengan inputan
		if err := common.PasswordHash(user.Password, r.Password); err != nil {
			o.Failure("username", "Invalid credential, please check your username or password.")
		}
	}

	if o.Valid {
		r.User = user
	}

	return o
}

// Messages implement validation.Requests interfaces
// return custom messages when validation fails.
func (r *SignInRequest) Messages() map[string]string {
	return map[string]string{}
}
