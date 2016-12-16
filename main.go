package main

import (
	"flag"
	"html/template"
	"log"
	"os"
	"strings"
)

type Profile struct {
	Name    string
	Label   string
	Image   string
	Address string
	Phone   string
	Email   string
	Github  string
	Summary string
}

type Experience struct {
	Company   string
	Role      string
	StartDate string
	EndDate   string
	Summary   string
}

type Project struct {
	Name      string
	Role      string
	StartDate string
	EndDate   string
	Summary   string
}
type Education struct {
	Location  string
	Major     string
	StartDate string
	EndDate   string
	Summary   string
}

type Resume struct {
	Profile     Profile
	Experiences []Experience
	Projects    []Project
	Educations  []Education
}

const tmpl = `<div class="col-md-6 col-sm-12">
				<h1>{{ .Profile.Name }}</h1>
				<h3>{{ .Profile.Label }}</h3>
			  </div>`

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
	tempFile := flag.String("template", "template.tmpl", "template file")

	flag.Parse()

	resume := Resume{
		Profile: Profile{
			Name:    "ChingChang Li",
			Label:   "Software Engineer",
			Email:   "xu354cjo1008@gmail.com",
			Github:  "https://github.com/xu354cjo1008",
			Summary: "A software engineer from Taiwan, who love technology, programing and mathmatics",
		},
		Experiences: []Experience{
			{
				Company:   "QNAP System",
				Role:      "Software Developer",
				StartDate: "2015.04",
				EndDate:   "2016.10",
				Summary: `1. Linux kernel storage system study and development
						2. SAS expander firmware development
						3. Linux hardware crypto system integration
						4. System performance analysis and tuning`,
			},
		},
		Projects: []Project{
			{
				Name:      "Linux storage development and system analysis",
				Role:      "Function developer and performance engineer",
				StartDate: "2015.06",
				EndDate:   "2016.10",
				Summary: `1. SAS expander firmware development\n
						2. SAS protocal development and analysis\n
						3. Linux storage stack study and function development\n
						4. Storage performance analysys and tuning\n`,
			},
		},
		Educations: []Education{
			{
				Location:  "National Taiwan University",
				Major:     "Electric Engineering graduate school",
				StartDate: "2012.09",
				EndDate:   "2014.07",
			},
		},
	}

	err := makeHtml(resume, *tempFile)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	os.Exit(0)
}
