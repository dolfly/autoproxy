package proxy

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/dolfly/autoproxy/assets"
	"github.com/oschwald/geoip2-golang"
)

var geoIp GeoIP

func InitGeoIpDB() {
	geoIp = NewGeoIP2()
}

// GeoIP2
type GeoIP struct {
	db       *geoip2.Reader
	emojiMap map[string]string
}

type CountryEmoji struct {
	Code  string `json:"code"`
	Emoji string `json:"emoji"`
}

// new geoip from db file
func NewGeoIP2() (geoip GeoIP) {
	data, err := assets.FS.ReadFile("GeoLite2-City.mmdb")
	if err != nil {
		log.Println("Geoip2 City库打开失败")
		os.Exit(1)
	} else {
		db, err := geoip2.FromBytes(data)
		if err != nil {
			log.Fatal(err)
		}
		geoip.db = db
	}
	if data, err := assets.FS.ReadFile("flags.json"); err != nil {
		log.Println("flags.json 读取失败")
		os.Exit(1)
	} else {
		var countryEmojiList = make([]CountryEmoji, 0)
		err = json.Unmarshal(data, &countryEmojiList)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}

		emojiMap := make(map[string]string)
		for _, i := range countryEmojiList {
			emojiMap[i.Code] = i.Emoji
		}
		geoip.emojiMap = emojiMap
	}
	return
}

// find ip info
func (g GeoIP) Find(ipORdomain string) (ip, country string, err error) {
	ips, err := net.LookupIP(ipORdomain)
	if err != nil {
		return "", "", err
	}
	ip = ips[0].String()

	var record *geoip2.City
	record, err = g.db.City(ips[0])
	if err != nil {
		return
	}
	countryIsoCode := record.Country.IsoCode
	if countryIsoCode == "" {
		country = fmt.Sprintf("🏁 ZZ")
	}
	emoji, found := g.emojiMap[countryIsoCode]
	if found {
		country = fmt.Sprintf("%v %v", emoji, countryIsoCode)
	} else {
		country = fmt.Sprintf("🏁 ZZ")
	}
	return
}
