package click

import (
	"github.com/aamsur/goshort/datastore/model"
	"git.qasico.com/cuxs/orm"
)

func GetByShortUrl(su string) (sl *model.Link, e error) {
	o := orm.NewOrm()
	if e = o.Raw("SELECT * FROM link WHERE short_url LIKE '%" + su + "' limit 1;").QueryRow(&sl); e == nil {
		return sl, nil
	}

	return nil, e
}
