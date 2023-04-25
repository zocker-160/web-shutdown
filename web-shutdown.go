package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)


func handleShutdown(w http.ResponseWriter, r *http.Request) {
	fmt.Println("shutdown triggered")

	cmd := exec.Command("shutdown", "-k", "now")
	go cmd.Run()

	fmt.Fprint(w, "OK")
}

func handleReboot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reboot triggered")

	cmd := exec.Command("reboot", "now")
	go cmd.Run()

	fmt.Fprint(w, "OK")
}


const indexhtml = `
<!DOCTYPE html>
<html>
    <header>
        <title>web-shutdown</title>
    </header>
    <body>
        <a href="/cmd/shutdown">shutdown</a>
        <a href="/cmd/reboot">reboot</a>
    </body>
</html>
`

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, indexhtml)
}

func main() {
	http.HandleFunc("/cmd/shutdown", handleShutdown)
	http.HandleFunc("/cmd/reboot", handleReboot)
	http.HandleFunc("/", handleIndex)
	
	var port int

	if (len(os.Args) > 1) {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		port = i
	} else {
		port = 8080
	}

	fmt.Print("web-shutdown by zocker_160\n\n")
	fmt.Printf("Listening on port %d...\n", port)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
