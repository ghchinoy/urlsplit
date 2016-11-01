package main

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"text/template"
)

const fmtmpl = `<#assign 
  url = message.getProperty("http.request.line")?split(" ")
  path = url[1]?split("/")
  id = path[{{.Value}}] <#-- {{.Name}} -->
>`

func main() {

	// Get the args
	args := os.Args

	// reject if there isn't a URL (2nd, 1st is the binary)
	if len(args) != 2 {
		fmt.Println(`Usage:
	urlsplit URL`)
		os.Exit(1)
	}

	// Parse the URL
	u, err := url.Parse(args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Print out the URL
	log.Println(u.Path)
	n := len(strings.Split(u.Path, "/"))
	for k, v := range strings.Split(u.Path, "/") {
		log.Println(k, v)
	}

	// Output FreeMarker snippet
	tmpl, err := template.New("freemarker").Parse(fmtmpl)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	name := strings.Split(u.Path, "/")[n-1]
	data := struct {
		Value int
		Name  string
	}{
		n - 1,
		name,
	}
	var outputstr bytes.Buffer
	err = tmpl.Execute(&outputstr, data)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(outputstr.String())
}
