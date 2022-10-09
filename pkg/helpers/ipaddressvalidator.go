package helpers

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

//IPAPI ...
var IPAPI string = "https://ipapi.co/"

//IPAPIResponse ....
type IPAPIResponse struct {
	IP                 string      `json:"ip"`
	Network            string      `json:"network"`
	Version            string      `json:"version"`
	City               string      `json:"city"`
	Region             string      `json:"region"`
	RegionCode         string      `json:"region_code"`
	Country            string      `json:"country"`
	CountryName        string      `json:"country_name"`
	CountryCode        string      `json:"country_code"`
	CountryCodeIso3    string      `json:"country_code_iso3"`
	CountryCapital     string      `json:"country_capital"`
	CountryTld         string      `json:"country_tld"`
	ContinentCode      string      `json:"continent_code"`
	InEu               bool        `json:"in_eu"`
	Postal             interface{} `json:"postal"`
	Latitude           float64     `json:"latitude"`
	Longitude          float64     `json:"longitude"`
	Timezone           string      `json:"timezone"`
	UtcOffset          string      `json:"utc_offset"`
	CountryCallingCode string      `json:"country_calling_code"`
	Currency           string      `json:"currency"`
	CurrencyName       string      `json:"currency_name"`
	Languages          string      `json:"languages"`
	CountryArea        float64     `json:"country_area"`
	CountryPopulation  int         `json:"country_population"`
	Asn                string      `json:"asn"`
	Org                string      `json:"org"`
}

//IsValidRequest validates the request origin client IP
func IsValidRequest(ipAddress string) bool {
	ipResponseData := IPAPIResponse{}
	url := IPAPI + ipAddress + "/json/" ///102.38.233.0/json/
	httpResponse, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	buf, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(buf, &ipResponseData)
	if err != nil {
		panic(err)
	}

	if ipResponseData.Country != "CY" {
		// return true
	}

	return true
}

//GetClientIP ...
func GetClientIP(r *http.Request) string {
	clientIP := ""
	// the default is the originating ip. but we try to find better options because this is almost
	// never the right IP
	if parts := strings.Split(r.RemoteAddr, ":"); len(parts) == 2 {
		clientIP = parts[0]
	}
	// If we have a forwarded-for header, take the address from there
	if xff := strings.Trim(r.Header.Get("X-Forwarded-For"), ","); len(xff) > 0 {
		addrs := strings.Split(xff, ",")
		lastFwd := addrs[len(addrs)-1]
		if ip := net.ParseIP(lastFwd); ip != nil {
			clientIP = ip.String()
		}
		// parse X-Real-Ip header
	} else if xri := r.Header.Get("X-Real-Ip"); len(xri) > 0 {
		if ip := net.ParseIP(xri); ip != nil {
			clientIP = ip.String()
		}
	}

	return clientIP
}
