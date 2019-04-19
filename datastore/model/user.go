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
	orm.RegisterModel(new(User))
}

// User model for user table.
type User struct {
	ID           int64     `orm:"column(id);auto" json:"-"`
	FullName     string    `orm:"column(full_name);size(100)" json:"full_name"`
	Email        string    `orm:"column(email);size(100)" json:"email"`
	Address      string    `orm:"column(address)" json:"address"`
	Username     string    `orm:"column(username);size(100)" json:"username"`
	Password     string    `orm:"column(password);size(145)" json:"password"`
	LastLogin    time.Time `orm:"column(last_login);type(timestamp);null" json:"last_login"`
	CreatedAt    time.Time `orm:"column(created_at);type(timestamp)" json:"created_at"`
	UpdatedAt    time.Time `orm:"column(updated_at);type(timestamp);null" json:"updated_at"`
	LastLogoutAt time.Time `orm:"column(last_logout_at);type(timestamp);null" json:"last_logout_at"`
}

// MarshalJSON customized data struct when marshaling data
// into JSON format, all Primary key & Foreign key will be encrypted.
func (m *User) MarshalJSON() ([]byte, error) {
	type Alias User

	return json.Marshal(&struct {
		ID       string `json:"id"`
		Password string `json:"password"`
		*Alias
	}{
		ID:       common.Encrypt(m.ID),
		Password: "******",
		Alias:    (*Alias)(m),
	})
}

// Save inserting or updating User struct into user table.
// It will updating if this struct has valid Id
// if not, will inserting a new row to user.
// The field parameter is an field that will be saved, it is
// usefull for partial updating data.
func (m *User) Save(fields ...string) (err error) {
	o := orm.NewOrm()
	if m.ID > 0 {
		_, err = o.Update(m, fields...)
	} else {
		m.ID, err = o.Insert(m)
	}
	return
}

// Delete permanently deleting user data
// this also will truncated all data from all table
// that have relation with this user.
func (m *User) Delete() (err error) {
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
func (m *User) Read(fields ...string) error {
	o := orm.NewOrm()
	return o.Read(m, fields...)
}
