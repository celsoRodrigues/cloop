package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/celsoRodrigues/cloop/model"
	"github.com/nguyenthenguyen/docx"
)

//Campaign to be exported
type campaign struct {
	Sticker string
	image   string
	link    string
	Title   string
	Body    string
	Cta     string
}

type linkStruct struct {
	URL  string
	Copy string
}

type marketing struct {
	Week      string
	Campaigns []campaign
	Links     []linkStruct
}

// SplitAt returns a bufio.SplitFunc closure, splitting at a substring
// scanner.Split(SplitAt("\n# "))
func SplitAt(substring string) func(data []byte, atEOF bool) (advance int, token []byte, err error) {

	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {

		// Return nothing if at end of file and no data passed
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		// Find the index of the input of the separator substring
		if i := strings.Index(string(data), substring); i >= 0 {
			return i + len(substring), data[0:i], nil
		}

		// If at end of file with data return the data
		if atEOF {
			return len(data), data, nil
		}

		return
	}
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

//RemoveSpecial to be exported
func RemoveSpecial(s string) (sss string) {

	scanner := bufio.NewScanner(strings.NewReader(strings.Trim(s, " ")))
	scanner.Split(bufio.ScanRunes)
	var ss string
	for scanner.Scan() {
		if scanner.Text() != "!" || scanner.Text() != "?" || scanner.Text() != "'" || scanner.Text() != "." || scanner.Text() != ";" || scanner.Text() != "," || scanner.Text() != "£" {
			ss += scanner.Text()
		}

	}

	return ss
}

//Xplode to be exported
func Xplode(s string) (sss []string) {
	scanner := bufio.NewScanner(strings.NewReader(strings.Trim(s, " ")))
	scanner.Split(bufio.ScanWords)
	var ss []string

	for scanner.Scan() {
		ss = append(ss, scanner.Text())
	}
	return ss
}

//XplodeLast to be exported
func XplodeLast(s string) (sss []string) {
	scanner := bufio.NewScanner(strings.NewReader(strings.Trim(s, " ")))
	scanner.Split(bufio.ScanWords)
	var ss []string

	for scanner.Scan() {
		ss = append(ss, scanner.Text())
	}
	s1 := strings.Join(ss[:len(ss)-1], " ")
	s2 := strings.Join(ss[len(ss)-1:], " ")

	s3 := []string{s1, s2}
	return s3
}

//XplodeCta to be exported
func XplodeCta(s string) (sss []string) {
	scanner := bufio.NewScanner(strings.NewReader(strings.Trim(s, " ")))
	scanner.Split(SplitAt(","))
	var ss []string

	for scanner.Scan() {
		ss = append(ss, scanner.Text())
	}
	return ss
}

var fn = template.FuncMap{
	"AddDash":       AddDash,
	"Xplode":        Xplode,
	"XplodeLast":    XplodeLast,
	"XplodeCta":     XplodeCta,
	"RemoveSpecial": RemoveSpecial,
}

func main() {

	mainSec := campaign{}
	secSec := campaign{}
	featSec := campaign{}
	prodCatSec := campaign{}
	brandSec := campaign{}

	var mkt marketing

	r, err := docx.ReadDocxFile(filepath.Join("./docx", "hp2.docx"))
	defer r.Close()

	if err != nil {
		panic(err)
	}
	myDoc := model.Document{}

	err = xml.Unmarshal([]byte(r.Content), &myDoc)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	getSections(myDoc, "main", &mainSec)
	getSections(myDoc, "sec", &secSec)
	getSections(myDoc, "feat", &featSec)
	getSections(myDoc, "prodcat", &prodCatSec)
	getSections(myDoc, "brand", &brandSec)
	getWeek(myDoc, &mkt)

	getTopLinks(myDoc)

	mkt.Campaigns = []campaign{mainSec, secSec, featSec, prodCatSec, brandSec}
	fd, err := os.OpenFile(filepath.Join("./bin", "page.html"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	defer fd.Close()

	tmpl := template.Must(template.New("layout.html").Funcs(fn).ParseFiles(filepath.Join("./layout", "layout.html")))
	err = tmpl.Execute(fd, mkt)
	if err != nil {
		log.Fatal(err)
	}

}

func getSections(myDoc model.Document, s string, sec *campaign) {

	for i := 0; i <= len(myDoc.Body.P)-1; i++ {

		if myDoc.Body.P[i].Sdt.SdtPr.Tag.Val == "image_"+s {
			sec.image = myDoc.Body.P[i].Sdt.SdtContent.R[0].T.Text
		}

		if myDoc.Body.P[i].Sdt.SdtPr.Tag.Val == "link_"+s {
			sec.link = myDoc.Body.P[i].Sdt.SdtContent.Hyperlink.R.T
		}

		if myDoc.Body.P[i].Sdt.SdtPr.Tag.Val == "sticker_"+s {
			sec.Sticker = myDoc.Body.P[i].Sdt.SdtContent.R[0].T.Text
		}

		if myDoc.Body.P[i].Sdt.SdtPr.Tag.Val == "title_"+s {
			sec.Title = myDoc.Body.P[i].Sdt.SdtContent.R[0].T.Text
		}

		if myDoc.Body.P[i].Sdt.SdtPr.Tag.Val == "body_"+s {
			sec.Body = myDoc.Body.P[i].Sdt.SdtContent.R[0].T.Text
		}

		if myDoc.Body.P[i].Sdt.SdtPr.Tag.Val == "cta_"+s {
			sec.Cta = myDoc.Body.P[i].Sdt.SdtContent.R[0].T.Text
		}
	}
}

func getTopLinks(myDoc model.Document) {

	for i := 0; i <= len(myDoc.Body.P)-1; i++ {

		if myDoc.Body.P[i].Sdt.SdtPr.Tag.Val == "top_link" {
			fmt.Printf("%s: ", myDoc.Body.P[i].Sdt.SdtPr.Tag.Val)
			fmt.Printf("%v\n", myDoc.Body.P[i].Sdt.SdtContent.Hyperlink.R.T)
		}

	}

}

func getWeek(myDoc model.Document, mkt *marketing) {

	for i := 0; i <= len(myDoc.Body.P)-1; i++ {

		if myDoc.Body.P[i].Sdt.SdtPr.Tag.Val == "week" {
			mkt.Week = myDoc.Body.P[i].Sdt.SdtContent.R[0].T.Text
		}

	}

}
