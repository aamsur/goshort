// 
// 
// 

package generate

import (
	"github.com/aamsur/goshort/datastore/model"
	"strings"
	"git.qasico.com/cuxs/orm"
)

func Save(l *model.Link) (e error) {
	if l.CustomHash == "" {
		// generate hash
		l.SetShortUrl(genereateHash())
	} else {
		l.SetShortUrl(l.CustomHash)
	}

	if e = l.Save(); e != nil {
		isDuplicate := strings.Contains(e.Error(), "Duplicate entry")

		if l.CustomHash != "" && isDuplicate {
			return e
		} else if isDuplicate {
			Save(l)
		}
	}

	return e
}

func genereateHash() (hash string) {
	o := orm.NewOrm()
	if e := o.Raw("SELECT LEFT(UUID(), 5)").QueryRow(&hash); e == nil {
		return hash
	}

	return ""
}
