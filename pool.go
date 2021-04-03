package browser

import "sync"

var pools = sync.Pool{
	New: func() interface{} {
		jar, err := NewJar(nil)
		if err != nil {
			panic(err)
		}
		return &Browser{
			CookieJar: jar,
			UserAgent: "",
		}
	},
}
