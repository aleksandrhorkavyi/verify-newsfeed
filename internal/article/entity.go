package article

import "time"

type Article struct {
	id      string
	short   string
	full    string
	created time.Time
}
