package restyHttp

import (
	"github.com/go-resty/resty/v2"
	"time"
)

func GetHttpClient(proxys ...string) *resty.Client {
	client := resty.New()
	if len(proxys) > 0 {
		proxy := proxys[0]
		client.SetProxy(proxy)
	}
	client.SetTimeout(time.Second * 5)
	return client
}

func GetMobileHttpRequest() *resty.Request {
	return GetHttpClient().SetHeader("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1").R()
}
