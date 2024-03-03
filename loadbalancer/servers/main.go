package servers

import (
	"fmt"
	"net/http"
	"strings"
)

const discoveryUrl = "http://localhost:4000/register"

func ServerGenerator(port string, count int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request from", strings.Split(r.RemoteAddr, ":")[0])
		fmt.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
		fmt.Println("Host:", r.Host)
		for key, value := range r.Header {
			fmt.Printf("%s: %s\n", key, value[0])
		}
		fmt.Printf("Response from server: %s 200 OK\n", r.Proto)
		fmt.Println("=====================================")

		w.Header().Add("status", "200")
		w.Write([]byte("Hello from the backend server\n"))
	})

	fmt.Printf("Server %d on port: %s", count, port)
	_, err := http.Get(discoveryUrl + "?url=http://localhost:" + port)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(port, nil)
}
