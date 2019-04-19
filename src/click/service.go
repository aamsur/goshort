package click

import (
	"github.com/aamsur/goshort/datastore/model"
	"git.qasico.com/cuxs/orm"
	"regexp"
	"log"
	"time"
)

func GetByShortUrl(hash string) (l *model.Link, e error) { // Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	hash = reg.ReplaceAllString(hash, "")

	o := orm.NewOrm()
	if e = o.Raw("SELECT * FROM link WHERE `short_url` LIKE '%" + hash + "' limit 1;").QueryRow(&l); e == nil {
		return l, nil
	}

	return nil, e
}

func CreateLinkLog(l *model.Link) {
	l.Clicked = l.Clicked + 1
	l.LastClickedAt = time.Now()
	l.UpdatedAt = l.LastClickedAt
	l.Save("Clicked", "LastClickedAt", "UpdatedAt")

	var ll = &model.LinkLog{ShortUrl: l.ShortUrl, ClickedAt: l.LastClickedAt}
	ll.Save()
}

func CreateJustRedirectLog(shortUrl string) {
	var ll = &model.LinkLog{ShortUrl: shortUrl, ClickedAt: time.Now()}
	ll.Save()
}
