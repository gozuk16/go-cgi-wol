package main

import (
	"fmt"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/linde12/gowol"
	"net/http"
	"net/http/cgi"
)

func wol(mac string) (err error) {

	if packet, err := gowol.NewMagicPacket(mac); err == nil {
		packet.Send("255.255.255.255")          // send to broadcast
		packet.SendPort("255.255.255.255", "7") // specify receiving port
		return nil
	} else {
		return err
	}

}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "Welcome to Wake on Lan Server.\n")

	mac := r.URL.RawQuery
	err := validation.Validate(mac,
		validation.Required,
		is.MAC,
	)
	if err != nil {
		fmt.Fprintf(w, "parameter error: %s.\n%v", mac, err)
		return
	}

	if err := wol(r.URL.RawQuery); err == nil {
		fmt.Fprintf(w, "%s will soon wake up!\n", mac)
	} else {
		fmt.Fprintf(w, "error. %s is still sleeping. Zzz\n%v", mac, err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	cgi.Serve(nil)
}
