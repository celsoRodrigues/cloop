package section

import "github.com/celsoRodrigues/cloop/model"

//GetSection to be exported
func GetSection(myDoc model.Document, s string, sec *model.Campaign) {

	for i := 0; i <= len(myDoc.Body.P)-1; i++ {

		if myDoc.Body.P[i].Sdt.SdtPr.Tag.Val == "image_"+s {
			sec.Image = myDoc.Body.P[i].Sdt.SdtContent.R[0].T.Text
		}

		if myDoc.Body.P[i].Sdt.SdtPr.Tag.Val == "link_"+s {
			sec.Link = myDoc.Body.P[i].Sdt.SdtContent.Hyperlink.R.T
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
