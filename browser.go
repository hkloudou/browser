package browser

import (
	"io"
	"io/ioutil"
	"net/http"
)

//Browser 浏览器
type Browser struct {
	cookieJar *Jar
	userAgent string
}

//GetHTTPClient 获得GetHTTPClient对象
func (me *Browser) GetHTTPClient() *http.Client {
	return &http.Client{
		Jar: me.cookieJar,
	}
}

func (me *Browser) GetJar() *Jar {
	return me.cookieJar
}

func (me *Browser) SetJar(tmp *Jar) {
	me.cookieJar.DeepCopyFrom(tmp)
}

func (me *Browser) SetUserAgent(str string) string {
	return me.userAgent
}

//GetRequestObject 获得请求对象
func (me *Browser) GetRequestObject(method string, url string, body io.Reader) (*http.Request, error) {
	return me.GetRequestWithHeadObject(method, url, nil, body)
}

//GetRequestWithHeadObject 使用指定头去提交
func (me *Browser) GetRequestWithHeadObject(method string, url string, headers map[string]string, body io.Reader) (*http.Request, error) {
	httpReq, err := http.NewRequest(method, url, body)
	if err == nil {
		if me.userAgent != "" {
			httpReq.Header.Set("User-Agent", me.userAgent)
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
	if httpResp != nil {
		//https://blog.csdn.net/hello_ufo/article/details/92994573 提前决定是否要关闭，修复一个可能的内存泄漏
		defer httpResp.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	body, errReadAll := ioutil.ReadAll(httpResp.Body)
	if errReadAll != nil {
		return nil, errReadAll
	}
	return body, nil
}

// func (me *Browser) Release() {
// 	pools.Put(me)
// }

//NewBrowser 新建一个全新的浏览器
func NewBrowser() *Browser {
	jar, err := NewJar(nil)
	if err != nil {
		panic(err)
	}
	return &Browser{
		cookieJar: jar,
		userAgent: "",
	}
	// x := pools.Get().(*Browser)
	// x.UserAgent = ""
	// x.cookieJar.Clear()
	// return x
}

//NewBrowserWithJar 新建一个带Cookie的浏览器
func NewBrowserWithJar(jar *Jar) *Browser {
	// bro := NewBrowser()
	// bro.cookieJar = jar
	// return bro
	return &Browser{
		cookieJar: jar,
		userAgent: "",
	}
}
