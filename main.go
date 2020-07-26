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
	Week  string
	Title string
	Body  string
	Cta   string
}

type marketing struct {
	Campaigns []campaign
	Links     []string
}

//AddDash to be exported
func AddDash(s string) (sss string) {
	scanner := bufio.NewScanner(strings.NewReader(strings.Trim(s, " ")))
	scanner.Split(bufio.ScanWords)
	var ss string
	for scanner.Scan() {
		ss += scanner.Text() + "-"
	}
	return ss[0 : len(ss)-1]
}

var fn = template.FuncMap{
	"AddDash": AddDash,
}

func main() {

	r, err := docx.ReadDocxFile(filepath.Join("./docx", "hp2.docx"))
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

	var week string
	var links []string
	var titleLines []string
	var bodyLines []string
	var ctaLines []string

	for scanner.Scan() {

		if strings.Contains(scanner.Text(), "Week:") {
			index := strings.Index(scanner.Text(), ":")
			week = strings.Trim(scanner.Text()[index+1:], " ")
		}

		if strings.Contains(scanner.Text(), "Link:") {
			index := strings.Index(scanner.Text(), ":")
			link := strings.Trim(scanner.Text()[index+1:], " ")
			links = append(links, link)
		}

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
	mycapaign.Week = week
	for i := 0; i < len(titleLines); i++ {

		mycapaign.Title = titleLines[i]
		mycapaign.Body = bodyLines[i]
		mycapaign.Cta = ctaLines[i]

		mycampaigns = append(mycampaigns, mycapaign)

	}

	mkt := marketing{
		Campaigns: mycampaigns,
		Links:     links,
	}

	fd, err := os.OpenFile(filepath.Join("./bin", "page.html"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	defer fd.Close()

	tmpl := template.Must(template.New("layout.html").Funcs(fn).ParseFiles(filepath.Join("./layout", "layout.html")))
	err = tmpl.Execute(fd, mkt)
	if err != nil {
		log.Fatal(err)
	}

}
