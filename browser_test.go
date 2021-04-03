package browser

import (
	"net/url"
	"testing"
)

func Test_pool(t *testing.T) {
	x := NewBrowser()
	x.UserAgent = "测试"
	x.CookieJar.SetCookies(&url.URL{}, nil)
	t.Logf("%p\n", x)
	t.Log(x, x.CookieJar.Entries, x.CookieJar.NextSeqNum)
	x.Release()
	// runtime.GC()
	y := NewBrowser()
	t.Logf("%p\n", y)
	t.Log(y, y.UserAgent, y.CookieJar.Entries, y.CookieJar.NextSeqNum)
}
