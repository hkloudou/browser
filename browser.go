package browser

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

//Browser 浏览器
type Browser struct {
	cookieJar *Jar
	userAgent string
	client    *http.Client
}

//GetHTTPClient 获得GetHTTPClient对象
func (me *Browser) GetHTTPClient() *http.Client {
	if me.client == nil {
		me.client = &http.Client{
			Jar: me.cookieJar,
		}
	}
	return me.client
}

func (me *Browser) WiseProxy(proxyURL *url.URL) {
	client := me.GetHTTPClient()
	if proxyURL != nil {
		client.Transport = &http.Transport{
			// DisableKeepAlives: true,
			Proxy: func(req *http.Request) (*url.URL, error) {
				return proxyURL, nil
			},
			IdleConnTimeout: 2 * time.Minute,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	} else {
		client.Transport = &http.Transport{
			// DisableKeepAlives: true,
			Proxy:           http.ProxyFromEnvironment,
			IdleConnTimeout: 2 * time.Minute,
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
}

func (me *Browser) SetTimeOut(timeout time.Duration) {
	client := me.GetHTTPClient()
	client.Timeout = timeout
}
func (me *Browser) Close() {
	client := me.GetHTTPClient()
	client.CloseIdleConnections()
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
func (me *Browser) Do(httpReq *http.Request) ([]byte, error) {
	client := me.GetHTTPClient()
	httpResp, err := client.Do(httpReq)
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
