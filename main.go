package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func makeHtml(data Resume, tmpl string) error {
	funcs := template.FuncMap{
		"nl2br": func(text string) template.HTML {
			return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
		},
	}

	t, err := template.New("resume").Funcs(funcs).ParseFiles("template/index.tmpl", "template/profile.tmpl", "template/experience.tmpl", "template/project.tmpl", "template/education.tmpl")
	if err != nil {
		log.Println(err)
		return err
	}

	if err := t.ExecuteTemplate(os.Stdout, "index", data); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func main() {
	var err error
	var file []byte
	var result *Resume
	tempFile := flag.String("template", "template.tmpl", "template file")
	userFile := flag.String("user", "user.yaml", "user resume informance file path")

	flag.Parse()

	file, err = ioutil.ReadFile(*userFile)
	if err != nil {
		log.Fatal(err)
	}
	result, err = parseYaml(file)
	if err != nil {
		log.Fatal(err)
	}

	err = makeHtml(*result, *tempFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	os.Exit(0)
}
