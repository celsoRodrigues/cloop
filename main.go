package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/nguyenthenguyen/docx"
)

//Campaign to be exported
type campaign struct {
	Title string
	Bodie string
	Cta   string
}

type campaigns struct {
	Mkt []campaign
}

func main() {

	r, err := docx.ReadDocxFile(filepath.Join("./docx", "hp.docx"))
	defer r.Close()

	if err != nil {
		panic(err)
	}

	mycapaign := campaign{}

	var cr = regexp.MustCompile(`</w:p>`)
	var re = regexp.MustCompile(`[<][^>]*[>]`)

	str := cr.ReplaceAll([]byte(r.Content), []byte("\n"))
	str = re.ReplaceAll([]byte(str), []byte(""))
	fmt.Print(string(str))
	scanner := bufio.NewScanner(strings.NewReader(string(str)))

	var titleLines []string
	var bodyLines []string
	var ctaLines []string

	for scanner.Scan() {

		if strings.Contains(scanner.Text(), "Title copy") {
			scanner.Scan()
			titleLines = append(titleLines, scanner.Text())
		}

		if strings.Contains(scanner.Text(), "Body copy") {
			scanner.Scan()
			bodyLines = append(bodyLines, scanner.Text())
		}

		if strings.Contains(scanner.Text(), "CTA copy") {
			scanner.Scan()
			ctaLines = append(ctaLines, scanner.Text())
		}

	}

	if len(titleLines) != len(bodyLines) || len(ctaLines) != len(titleLines) || len(bodyLines) != len(ctaLines) {
		fmt.Println("There is a field missing in the page\n You have:\n titles:", len(titleLines), "Bodies:", len(bodyLines), "Cta's:", len(ctaLines))
	}

	var mycampaigns []campaign
	for i := 0; i < len(titleLines); i++ {

		mycapaign.Title = titleLines[i]
		mycapaign.Bodie = bodyLines[i]
		mycapaign.Cta = ctaLines[i]

		mycampaigns = append(mycampaigns, mycapaign)

	}

	fd, err := os.OpenFile(filepath.Join("./bin", "page.html"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)

	tmpl := template.Must(template.ParseFiles(filepath.Join("./layout", "layout.html")))
	err = tmpl.Execute(fd, mycampaigns)
	if err != nil {
		log.Fatal(err)
	}

}
