package getter

import (
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/zu1k/proxypool/tool"

	"github.com/zu1k/proxypool/proxy"
)

func init() {
	Register("webfuzz", NewWebFuzzGetter)
}

type WebFuzz struct {
	Url string
}

func (w *WebFuzz) Get() []proxy.Proxy {
	resp, err := http.Get(w.Url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	return FuzzParseProxyFromString(string(body))
}

func (w *WebFuzz) Get2Chan(pc chan proxy.Proxy, wg *sync.WaitGroup) {
	wg.Add(1)
	nodes := w.Get()
	for _, node := range nodes {
		pc <- node
	}
	wg.Done()
}

func NewWebFuzzGetter(options tool.Options) Getter {
	url, found := options["url"]
	if found {
		return &WebFuzz{Url: url.(string)}
	}
	return nil
}
