package main

/*
URL: http://myexternalip.com/#golang
*/

import (
	"net/http"
	"log"
	"os"
	"net/url"
	"time"
	"io/ioutil"
	"github.com/BurntSushi/toml"
)

type Config struct {
    Cnamemaster string
    Subdomain   string
    Domain      string
    User        string
    Pass        string
}


func getIp(requestType string) (string) {
	res, _ := http.Get("http://" + requestType  + ".myexternalip.com/raw")
	defer res.Body.Close()

        ip, _ := ioutil.ReadAll(res.Body)
	return string(ip[:])
}

func readConfigfile() Config {
    var configfile = "properties.ini"
    _, err := os.Stat(configfile)
    if err != nil {
        log.Fatal("Config file is missing: ", configfile)
    }

    var config Config
    if _, err := toml.DecodeFile(configfile, &config); err != nil {
        log.Fatal(err)
    }
    return config
}

func dyndns(ipv4 string, ipv6 string, config Config) {
	t := time.Now()
	formattedTime := t.Format(time.RFC3339)
	v := url.Values{}
	v.Add("s_login", config.User)
	v.Add("s_pw", config.Pass)
	v.Add("command", "UpdateDNSZone")
	v.Add("dnszone", config.Domain + ".")
	v.Add("rr0", "@ IN NS ns1a.dodns.net")
	v.Add("rr1", "@ IN NS ns2a.dodns.net")
	v.Add("rr2", config.Domain + ". 3600 IN CNAME " + config.Cnamemaster + ".")
	v.Add("rr3", "*." + config.Domain + ". 3600 IN CNAME " + config.Cnamemaster + ".")
	v.Add("rr4", config.Domain + ". 600 IN A " + ipv4)
	v.Add("rr5", config.Domain + ". 600 IN AAAA " + ipv6)
	v.Add("rr6", config.Domain + ". 600 IN TXT " + formattedTime)
	log.Print(v.Encode())
	resp, err := http.PostForm("https://api.domainreselling.de/api/call.cgi", v)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(resp.Status)
	defer resp.Body.Close()
}

func main() {
	config := readConfigfile()
	lastIpv4 := ""
	lastIpv6 := ""
	for {
		ipv4 := getIp("ipv4")
		ipv6 := getIp("ipv6")
		if(lastIpv4 != ipv4 || lastIpv6 != ipv6) {
			log.Print("Ipv4 or Ipv6 altered")
			log.Print("IPv4: " + ipv4)
			log.Print("IPv6: " + ipv6)
			lastIpv4 = ipv4
			lastIpv6 = ipv6
			dyndns(ipv4, ipv6, config)
		}
		time.Sleep(1 * time.Minute)
	}
}
