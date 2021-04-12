package browser

import "net/http"

func (me *Browser) setAcceptAll(req *http.Request) {
	req.Header.Set("Accept", "*/*")
}

func (me *Browser) setAcceptHTML(req *http.Request) {
	req.Header.Set("Accept", "text/html, application/xhtml+xml, image/jxr, */*")
}

func (me *Browser) setAcceptJSON(req *http.Request) {
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
}

func (me *Browser) setNoCache(req *http.Request) {
	req.Header.Set("Cache-Control", "no-cache")
}

func (me *Browser) setAcceptLang(req *http.Request) {
	req.Header.Set("Accept-Language", "zh-CN")
}

// X-Requested-With: XMLHttpRequest
func (me *Browser) setAjax(req *http.Request) {
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
}

func (me *Browser) setReferer(req *http.Request, str string) {
	req.Header.Set("Referer", str)
}
func (me *Browser) SetHeader(req *http.Request, key, value string) {
	// req.Header.Set("Referer", str)
	// req.Header.Set("c-token", me.CSRFToken)
	req.Header.Set(key, value)
}

func (me *Browser) setContentJSON(req *http.Request) {
	// req.Header.Set("Referer", str)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
}

func (me *Browser) setContentFormPost(req *http.Request) {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}
