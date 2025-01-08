package main

import "net/http"

func (app *application) webSecHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Burpsuite: https://portswigger.net/burp \n AltDNS: https://github.com/infosec-au/altdns \n Amass: https://owasp.org/www-project-amass/ \n SSL Labs Server Test: https://www.ssllabs.com/ssltest/"))
}
