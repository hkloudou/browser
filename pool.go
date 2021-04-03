package browser

import "sync"

var pools = sync.Pool{
	New: func() interface{} {
		return &Browser{
			CookieJar: nil,
			UserAgent: "",
		}
	},
}
