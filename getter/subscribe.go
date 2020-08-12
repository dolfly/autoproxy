package getter

import (
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/zu1k/proxypool/proxy"
	"github.com/zu1k/proxypool/tool"
)

func init() {
	Register("subscribe", NewSubscribe)
}

type Subscribe struct {
	Url string
}

func (s *Subscribe) Get() []proxy.Proxy {
	resp, err := http.Get(s.Url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	nodesString, err := tool.Base64DecodeString(string(body))
	if err != nil {
		return nil
	}
	nodesString = strings.ReplaceAll(nodesString, "\t", "")

	nodes := strings.Split(nodesString, "\n")
	return StringArray2ProxyArray(nodes)
}

func (s *Subscribe) Get2Chan(pc chan proxy.Proxy, wg *sync.WaitGroup) {
	wg.Add(1)
	nodes := s.Get()
	for _, node := range nodes {
		pc <- node
	}
	wg.Done()
}

func NewSubscribe(options tool.Options) Getter {
	url, found := options["url"]
	if found {
		return &Subscribe{
			Url: url.(string),
		}
	}
	return nil
}
