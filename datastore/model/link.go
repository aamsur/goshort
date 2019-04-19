// Copyright 2017 PT. Qasico Teknologi Indonesia. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

import (
	"encoding/json"
	"time"

	"git.qasico.com/cuxs/common"
	"git.qasico.com/cuxs/orm"
)

func init() {
	orm.RegisterModel(new(Link))
}

// Link model for link table.
type Link struct {
	ID            int64     `orm:"column(id);auto" json:"-"`
	LongUrl       string    `orm:"column(long_url)" json:"long_url"`
	ShortUrl      string    `orm:"column(short_url);size(100);null" json:"short_url"`
	Host          string    `orm:"column(host);size(145);null" json:"host"`
	CustomHash    string    `orm:"column(custom_hash);size(145);null" json:"custom_hash"`
	Clicked       int64     `orm:"column(clicked);null" json:"clicked"`
	LastClickedAt time.Time `orm:"column(last_clicked_at);type(timestamp);null" json:"last_clicked_at,omitempty"`
	CreatedAt     time.Time `orm:"column(created_at);type(timestamp)" json:"created_at"`
	UpdatedAt     time.Time `orm:"column(updated_at);type(timestamp);null" json:"updated_at"`
}

// MarshalJSON customized data struct when marshaling data
// into JSON format, all Primary key & Foreign key will be encrypted.
func (m *Link) MarshalJSON() ([]byte, error) {
	type Alias Link

	return json.Marshal(&struct {
		ID string `json:"id"`
		*Alias
	}{
		ID:    common.Encrypt(m.ID),
		Alias: (*Alias)(m),
	})
}

// Save inserting or updating Link struct into link table.
// It will updating if this struct has valid Id
// if not, will inserting a new row to link.
// The field parameter is an field that will be saved, it is
// usefull for partial updating data.
func (m *Link) Save(fields ...string) (err error) {
	o := orm.NewOrm()
	if m.ID > 0 {
		_, err = o.Update(m, fields...)
	} else {
		m.ID, err = o.Insert(m)
	}
	return
}

// Delete permanently deleting link data
// this also will truncated all data from all table
// that have relation with this link.
func (m *Link) Delete() (err error) {
	o := orm.NewOrm()
	if m.ID > 0 {
		var i int64
		if i, err = o.Delete(m); i == 0 && err == nil {
			err = orm.ErrNoAffected
		}
		return
	}
	return orm.ErrNoRows
}

// Read execute select based on data struct that already
// assigned.
func (m *Link) Read(fields ...string) error {
	o := orm.NewOrm()
	return o.Read(m, fields...)
}

// Set shorturl from host and hash
func (m *Link) SetShortUrl(hash string) {
	m.ShortUrl = m.Host + "/" + hash
}
