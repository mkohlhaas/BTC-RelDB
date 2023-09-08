package protocol

import (
	"encoding/json"
	"io"
	"net"
	"net/http"
)

// IP is the standard IP address in IPv4 or IPv6 form.
type IP [16]byte // IP is big-endian

func (i IP) String() string {
	return net.IP(i[:]).String()
}

// IPAddr creates new ip from string.
func IPAddr(ipAddr string) IP {
	return *(*IP)(net.ParseIP(ipAddr).To16()) // *(*IP)(...) = IP(...)  but gopls doesn't like it
}

type ipQuery struct {
	Query string
}

func myPublicIPAddress() IP {
	var req *http.Response
	var err error
	if req, err = http.Get("http://ip-api.com/json/"); err != nil {
		return IP{}
	}
	defer req.Body.Close()

	var body []byte
	if body, err = io.ReadAll(req.Body); err != nil {
		return IP{}
	}

	var ipq ipQuery
	if err = json.Unmarshal(body, &ipq); err != nil {
		return IP{}
	}
	return IPAddr(ipq.Query)
}

type heightQuery struct {
	Height int32
}

func getLatestHeight() int32 {
	var hq heightQuery
	req, _ := http.NewRequest(http.MethodGet, "https://api.blockcypher.com/v1/btc/main", nil)
	res, _ := http.DefaultClient.Do(req)
	resBody, _ := io.ReadAll(res.Body)

	err := json.Unmarshal(resBody, &hq)
	if err != nil {
		return 0
	}
	return hq.Height
}
