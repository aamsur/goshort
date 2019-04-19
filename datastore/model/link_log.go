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
	orm.RegisterModel(new(LinkLog))
}

// LinkLog model for link_log table.
type LinkLog struct {
	ID        int64     `orm:"column(id);auto" json:"-"`
	ShortUrl  string    `orm:"column(short_url);size(100);null" json:"short_url"`
	ClickedAt time.Time `orm:"column(clicked_at);type(timestamp);null" json:"clicked_at"`
}

// MarshalJSON customized data struct when marshaling data
// into JSON format, all Primary key & Foreign key will be encrypted.
func (m *LinkLog) MarshalJSON() ([]byte, error) {
	type Alias LinkLog

	return json.Marshal(&struct {
		ID string `json:"id"`
		*Alias
	}{
		ID:    common.Encrypt(m.ID),
		Alias: (*Alias)(m),
	})
}

// Save inserting or updating LinkLog struct into link_log table.
// It will updating if this struct has valid Id
// if not, will inserting a new row to link_log.
// The field parameter is an field that will be saved, it is
// usefull for partial updating data.
func (m *LinkLog) Save(fields ...string) (err error) {
	o := orm.NewOrm()
	if m.ID > 0 {
		_, err = o.Update(m, fields...)
	} else {
		m.ID, err = o.Insert(m)
	}
	return
}

// Delete permanently deleting link_log data
// this also will truncated all data from all table
// that have relation with this link_log.
func (m *LinkLog) Delete() (err error) {
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
func (m *LinkLog) Read(fields ...string) error {
	o := orm.NewOrm()
	return o.Read(m, fields...)
}
