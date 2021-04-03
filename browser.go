package browser

import (
	"io"
	"io/ioutil"
	"net/http"
)

//Browser 浏览器
type Browser struct {
	CookieJar *Jar
	UserAgent string
}

//GetHTTPClient 获得GetHTTPClient对象
func (me *Browser) GetHTTPClient() *http.Client {
	return &http.Client{
		Jar: me.CookieJar,
	}
}

//GetRequestObject 获得请求对象
func (me *Browser) GetRequestObject(method string, url string, body io.Reader) (*http.Request, error) {
	return me.GetRequestWithHeadObject(method, url, nil, body)
}

//GetRequestWithHeadObject 使用指定头去提交
func (me *Browser) GetRequestWithHeadObject(method string, url string, headers map[string]string, body io.Reader) (*http.Request, error) {
	httpReq, err := http.NewRequest(method, url, body)
	if err == nil {
		if me.UserAgent != "" {
			httpReq.Header.Set("User-Agent", me.UserAgent)
		}
	}
	for k, v := range headers {
		httpReq.Header.Set(k, v)
	}
	return httpReq, err
}

//Do 打开网页
func (me *Browser) Do(httpClient *http.Client, httpReq *http.Request) ([]byte, error) {
	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	body, errReadAll := ioutil.ReadAll(httpResp.Body)
	if errReadAll != nil {
		return nil, errReadAll
	}
	return body, nil
}

func (me *Browser) Release() {
	pools.Put(me)
}

//NewBrowser 新建一个全新的浏览器
func NewBrowser() *Browser {
	x := pools.Get().(*Browser)
	x.CookieJar, _ = NewJar(nil)
	x.UserAgent = ""
	return x
}

// //NewBrowserWithJar 新建一个带Cookie的浏览器
// func NewBrowserWithJar(jar *Jar) *Browser {
// 	bro := &Browser{
// 		CookieJar: jar,
// 		UserAgent: "",
// 	}
// 	return bro
// }
