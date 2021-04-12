package browser

import "net/http"

func (me *Browser) SetAcceptAll(req *http.Request) {
	req.Header.Set("Accept", "*/*")
}

func (me *Browser) SetAcceptHTML(req *http.Request) {
	req.Header.Set("Accept", "text/html, application/xhtml+xml, image/jxr, */*")
}

func (me *Browser) SetAcceptJSON(req *http.Request) {
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
}

func (me *Browser) SetNoCache(req *http.Request) {
	req.Header.Set("Cache-Control", "no-cache")
}

func (me *Browser) SetAcceptLang(req *http.Request) {
	req.Header.Set("Accept-Language", "zh-CN")
}

// X-Requested-With: XMLHttpRequest
func (me *Browser) SetAjax(req *http.Request) {
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
}

func (me *Browser) SetReferer(req *http.Request, str string) {
	req.Header.Set("Referer", str)
}
func (me *Browser) SetHeader(req *http.Request, key, value string) {
	// req.Header.Set("Referer", str)
	// req.Header.Set("c-token", me.CSRFToken)
	req.Header.Set(key, value)
}

func (me *Browser) SetContentJSON(req *http.Request) {
	// req.Header.Set("Referer", str)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
}

func (me *Browser) SetContentFormPost(req *http.Request) {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}
