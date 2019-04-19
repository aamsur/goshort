// Copyright 2017 PT. Qasico Teknologi Indonesia. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package model

import (
	"fmt"

	"git.qasico.com/cuxs/common/faker"
)

// DummyLink make a dummy data for model Link
func DummyLink() *Link {
	var m Link
	faker.Fill(&m, "ID")

	if e := m.Save(); e != nil {
		fmt.Printf("error saving %s", e.Error())
	}
	return &m
}

// DummyLinkLog make a dummy data for model LinkLog
func DummyLinkLog() *LinkLog {
	var m LinkLog
	faker.Fill(&m, "ID")

	if e := m.Save(); e != nil {
		fmt.Printf("error saving %s", e.Error())
	}
	return &m
}

// DummyUser make a dummy data for model User
func DummyUser() *User {
	var m User
	faker.Fill(&m, "ID")

	if e := m.Save(); e != nil {
		fmt.Printf("error saving %s", e.Error())
	}
	return &m
}
