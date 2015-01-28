//timeserver serves a web page displaying the current time
//of day. The default port number for the webserver is 8080.
//Timeserver only displays time for the time request.
//Using command-line argument -v can show the version
//number.
//
//Copyright 2015 Cici, Chunchao Zhang
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"log"
	"time"
	"html"
)

//handleTime: set up webpage format and display the current time
func handleTime(w http.ResponseWriter, r *http.Request) {
	const layout = "3:04:05PM"
	t := time.Now()
	content := fmt.Sprintf(`
<html>
<head>
<style>
p {font-size: xx-large}
span.time {color: red}
</style>
</head>
<body>
<p>The time is now <span class="time">%s</span>.</p>
</body>
</html>`, t.Format(layout))
	
	fmt.Fprintf(w, content)
}


//handleNoCookie: when there is no cookie,display login form
func handleNoCookie(w http.ResponseWriter, r *http.Request) {
	content := 
		`
<html>
<body>
<form action="login">
  What is your name, Earthling?
<input type="text" name="name" size="50">
<input type="submit">
</form>
</p>
</body>
</html>
`

	fmt.Fprintf(w, content)
}

//handleNotFound: customarized 404 page for non-time request
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	content :=
		`
<html>
<body>
<p>These are not the URLs you're looking for.</p>
</body>
</html>`

	fmt.Fprintf(w, content)

}

//handleQuery
func handleQuery(w http.ResponseWriter, r *http.Request) {
	name := html.EscapeString(r.FormValue("name"))

	if name != "" {
		uuid, err := exec.Command("uuidgen").Output()
		if err != nil {
			log.Fatal(err)
		}
			
	
		content := fmt.Sprintf(
		`
<html>
<body>
<p> Greetings, %s. uuid is %s.
</p>
</body>
</html>`, name, uuid)
		
		fmt.Fprint(w, content)
	}
}

func main() {
	portPtr := flag.Int("port", 8080, "http server port number")
	versionPtr := flag.Bool("v", false, "Display version number")
	flag.Parse()

	if *versionPtr {
		fmt.Println("1.0.0")
		return
	}

	http.HandleFunc("/time", handleTime)
	http.HandleFunc("/", handleNoCookie)
	http.HandleFunc("/login", handleQuery)

	error := http.ListenAndServe(fmt.Sprintf(":%v", *portPtr), nil)
	if error != nil {
		fmt.Printf("Start server with port %d failed: %v\n", *portPtr, error)
		os.Exit(1)
	}
}
