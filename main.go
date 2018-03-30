package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gx/ipfs/QmXY77cVe7rVRQXZZQRioukUM7aRW3BTcAgJe12MCtb3Ji/go-multiaddr"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

func main() {
	cert := flag.String("cert", "", "/path/to/server.crt")
	key := flag.String("key", "", "/path/to/server.key")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}

		var addrs []string
		err = json.Unmarshal(b, &addrs)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, err.Error())
			return
		}

		for _, addr := range addrs {
			ma, err := multiaddr.NewMultiaddr(addr)
			if err != nil {
				continue
			}
			port, portErr := ma.ValueForProtocol(multiaddr.P_TCP)
			if portErr != nil {
				continue
			}
			network := "tcp4"
			ip, ipErr := ma.ValueForProtocol(multiaddr.P_IP4)
			if ipErr != nil {
				ip, ipErr = ma.ValueForProtocol(multiaddr.P_IP6)
				network = "tcp6"
				ip = "[" + ip + "]"
			}

			if ipErr == nil && portErr == nil {
				conn, err := net.Dial(network, ip+":"+port)
				if err == nil {
					conn.Close()
					fmt.Fprintf(w, `{"success": true}`)
					return
				}
			}
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, `{"success": false}`)
	})

	if *cert == "" && *key != "" || *key == "" && *cert != "" {
		log.Fatal("Must provide both cert and key if using SSL")
	}

	var err error
	if *cert == "" {
		err = http.ListenAndServe(":80", nil)
	} else {
		err = http.ListenAndServeTLS(":443", *cert, *key, nil)
	}
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
