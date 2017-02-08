DynDNS with UnitedDomains Reselling
===

This small go script fetches your remote IP Address ( ipv4 and ipv6 ) and writes it into a domain from united domains reselling.

**Details:**


* once started runs every minute
* receives all configuration ( domain settings and access data ) from simple config file
* checks your external IP addresses ( ipv4 and ipv6 )
* if at least one address changes do the following steps else wait one further minute
* send one command to United Domains API with username and password and change your DNS Zone
  * *domain.tld*. 3600 IN CNAME *CNameMaster*.
  * \*.*domain.tld*. 3600 IN CNAME *CNameMaster*.
  * *subdomain.domain.tld*. 600 IN A *publicIpv4*
  * *subdomain.domain.tld*. 600 IN AAAA *publicIpv6*
  * *subdomain.domain.tld*. 600 IN TXT *current time in RFC3339 format*

Requirements
---

In order to run this script you just need to install go lang. For more information see https://golang.org/doc/install 

Additionally you need to have access to an united domains reselling account with a domain and an existing DNS Zone.

Installation
---

* In order to install this please clone this repository to a location you like.
* Install this package to get all requirements for your new go script: ```go get "github.com/BurntSushi/toml"```
* Build the go script: ```GOBIN=$(pwd) go install src/server.go```
* Copy the file ```properties.ini.default``` to ```properties.ini``` and adjust your settings.
  * ```cnamemaster``` what should be the CNAME target for the domain ( not the subdomain )
  * ```subdomain``` the subdomain + domain you use for your dyndns
  * ```domain``` the domain you use for your dyndns. It gets the CNameMaster Value as CNAME
  * ```user``` your United Domains Reselling user
  * ```pass``` your United Domains Reselling user password
* Run ```./server```, sit down and take a look on the log.

Known Todo
---

* Create a MakeScript for Debian / Ubuntu based systems
* Create a systemD unit file to start and monitor this script from system startup

Support and supporting
---

There is no free support for this script. If you need help open a ticket and wait for response, if you need paid support write us a mail.

In order to support this script please do pull request or buy paid support.
Supporting OpenSource is always nice to help receiving public scripts for everyone.

How to report bugs
---

If you want to report bugs please ensure to give the following information:

* git branch is master and latest version
* when did you compile the ```server``` binary?
* a copy of your config file WITHOUT PASSWORD would be nice
* What is the expected output?
* What is happening?
* Is the IP in your log the correct public ip?

CopyRight and Law Stuff
---

This script is free to use and modify for anyone. If you use this script you accept that you are using this script on your own risk. It would be nice if you'd fork publicly and 

Created in 2017 from Bastian Bringenberg <bastian@agentur-pottkinder.de>