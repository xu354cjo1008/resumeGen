package main

import (
	"errors"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

var config struct {
	templatePath string
	templateFile []string
}

func configure() error {

	viper.SetConfigName("app")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		return errors.New("Config file not found...")
	} else {
		config.templatePath = viper.GetString("global.templatePath")
		config.templateFile = viper.GetStringSlice("global.templateFile")
	}

	log.Printf("\n Read config: \n template = %s\n"+
		"template components = %v\n",
		config.templatePath,
		config.templateFile)

	return nil
}

func makeHtml(data Resume, tmpl string) error {
	funcs := template.FuncMap{
		"nl2br": func(text string) template.HTML {
			return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
		},
	}

	files := make([]string, len(config.templateFile))
	for i, f := range config.templateFile {
		files[i] = path.Join(tmpl, f)
	}

	t, err := template.New("resume").Funcs(funcs).ParseFiles(files...)
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
	tempFile := flag.String("template", "example", "template file")
	userFile := flag.String("user", "user.yaml", "user resume informance file path")

	flag.Parse()

	err = configure()
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	file, err = ioutil.ReadFile(*userFile)
	if err != nil {
		log.Fatal(err)
	}
	result, err = parseYaml(file)
	if err != nil {
		log.Fatal(err)
	}

	err = makeHtml(*result, path.Join(config.templatePath, *tempFile))
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

	os.Exit(0)
}
