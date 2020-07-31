package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/celsoRodrigues/cloop/layoutfunc"
	"github.com/celsoRodrigues/cloop/model"
	"github.com/celsoRodrigues/cloop/section"
	"github.com/nguyenthenguyen/docx"
)

var fn = template.FuncMap{
	"AddDash":       layoutfunc.AddDash,
	"Xplode":        layoutfunc.Xplode,
	"XplodeLast":    layoutfunc.XplodeLast,
	"XplodeCta":     layoutfunc.XplodeCta,
	"RemoveSpecial": layoutfunc.RemoveSpecial,
}

func main() {

	promoNurserySec := model.PromoStruct{}
	promoExistingSec := model.PromoStruct{}
	mainSec := model.Campaign{}
	secSec := model.Campaign{}
	featSec := model.Campaign{}
	prodCatSec := model.Campaign{}
	brandSec := model.Campaign{}
	myDoc := model.Document{}

	var mkt model.Marketing
	docxFileName := "hp.docx"

	fd, err := os.OpenFile(filepath.Join("./docx", docxFileName), os.O_EXCL|os.O_RDWR, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fd.Close()
		r, err := docx.ReadDocxFile(filepath.Join("./docx", docxFileName))
		if err != nil {
			panic(err)
		}
		defer r.Close()
		err = xml.Unmarshal([]byte(r.Content), &myDoc)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}

	section.GetPromo(myDoc, "nursery", &promoNurserySec)
	section.GetPromo(myDoc, "existing", &promoExistingSec)

	section.GetSection(myDoc, "main", &mainSec)
	section.GetSection(myDoc, "sec", &secSec)
	section.GetSection(myDoc, "feat", &featSec)
	section.GetSection(myDoc, "prodcat", &prodCatSec)
	section.GetSection(myDoc, "brand", &brandSec)
	getWeek(myDoc, &mkt)

	getTopLinks(myDoc)

	mkt.Campaigns = []model.Campaign{mainSec, secSec, featSec, prodCatSec, brandSec}
	mkt.Promos = []model.PromoStruct{promoNurserySec, promoExistingSec}
	fd, err = os.OpenFile(filepath.Join("./bin", "page.html"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer fd.Close()

	tmpl := template.Must(template.New("layout.html").Funcs(fn).ParseFiles(filepath.Join("./layout", "layout.html")))
	err = tmpl.Execute(fd, mkt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func getTopLinks(myDoc model.Document) {

	for i := 0; i <= len(myDoc.Body.P)-1; i++ {

		//if myDoc.Body.P[i].Sdt.SdtPr.Tag.Val == "top_link" {
		//	fmt.Printf("%s: ", myDoc.Body.P[i].Sdt.SdtPr.Tag.Val)
		//	fmt.Printf("%v\n", myDoc.Body.P[i].Sdt.SdtContent.Hyperlink.R.T)
		//}
	}

}

func getWeek(myDoc model.Document, mkt *model.Marketing) {

	for i := 0; i <= len(myDoc.Body.P)-1; i++ {

		if myDoc.Body.P[i].Sdt.SdtPr.Tag.Val == "week" {
			mkt.Week = myDoc.Body.P[i].Sdt.SdtContent.R[0].T.Text
		}
	}
}
