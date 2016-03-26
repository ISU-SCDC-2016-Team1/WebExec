package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var function string = r.URL.Query().Get("function")
		var project string = r.URL.Query().Get("project")
		var runner string = r.URL.Query().Get("runner")
		var user string = r.URL.Query().Get("user")
		var redirect string = r.URL.Query().Get("redirect")
		var method string = r.URL.Query().Get("method")

		if (function == "stdin") {
			var out bytes.Buffer
			cmd := exec.Command("crconsole", "-u", user, "-i", "/tmp/" + user + ".priv",  "-p", project, "-r", runner, "stdin", "-f", "/tmp/stdin_" + user + "_stdin.tmp")
			cmd.Stdout = &out
			cmd.Stderr = &out
			cmd.Run()
			fmt.Fprintln(w, out.String())
		} else if (function == "get") {
			var out bytes.Buffer
			cmd := exec.Command("crconsole", "-u", user, "-i", "/tmp/" + user + ".priv",  "-p", project, "-r", runner, "get", "-m", method)
			cmd.Stdout = &out
			cmd.Stderr = &out
			cmd.Run()
			fmt.Fprintln(w, out.String())
		} else if (function == "run") {
			var out bytes.Buffer
			cmd := exec.Command("crconsole", "-u", user, "-i", "/tmp/" + user + ".priv",  "-p", project, "-r", runner, "run", "-x", redirect)
			cmd.Stdout = &out
			cmd.Stderr = &out
			cmd.Run()
			fmt.Fprintln(w, out.String())
		} else {
			var out bytes.Buffer
			cmd := exec.Command("crconsole", "-u", user, "-i", "/tmp/" + user + ".priv",  "-p", project, "-r", runner, function)
			cmd.Stdout = &out
			cmd.Stderr = &out
			cmd.Run()
			fmt.Fprintln(w, out.String())
		}
	})

	http.ListenAndServe("127.0.0.1:5634", nil)
	fmt.Println("Running and serving on :5634")
}
