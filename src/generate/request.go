// 
// 
// 

package generate

import (
	"time"
	"strings"

	"github.com/aamsur/goshort/datastore/model"

	"git.qasico.com/cuxs/validation"
	"git.qasico.com/cuxs/env"
)

// createRequest data struct that stored request data when requesting an create user process.
// All data must be provided and must be match with specification validation below.
// handler function should be bind this with context to matches incoming request
// data keys to the defined json tag.
type createRequest struct {
	LongUrl    string `json:"long_url" valid:"required"`
	Host       string `json:"host"`
	CustomHash string `json:"custom_hash"`

	ShortUrl string `json:"-"`
}

// Validate implement validation.Requests interfaces.
func (r *createRequest) Validate() *validation.Output {
	o := &validation.Output{Valid: true}

	// init default host
	host := r.Host
	if host == "" {
		host = env.GetString("DEFAULT_HOST", "locahost")
	}

	host = strings.TrimRight(host, "/")
	r.Host = host

	if r.CustomHash != "" {
		// validasi custom hash
		su := &model.Link{Host: r.Host}
		su.SetShortUrl(r.CustomHash)

		su.Read("ShortUrl")

		if su.ID != 0 {
			o.Failure("custom_hash", "custom_hash is exist")
		}

		r.ShortUrl = su.ShortUrl
	}

	return o
}

// Messages implement validation.Requests interfaces
// return custom messages when validation fails.
func (r *createRequest) Messages() map[string]string {
	return map[string]string{}
}

// Transform transforming request into model.
func (r *createRequest) Transform() *model.Link {
	l := &model.Link{
		LongUrl:    r.LongUrl,
		Host:       r.Host,
		CustomHash: r.CustomHash,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return l
}
