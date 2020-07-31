package model

import "encoding/xml"

//Campaign to be exported
type Campaign struct {
	Sticker string
	Image   string
	Link    string
	Title   string
	Body    string
	Cta     string
}

//LinkStruct to be exported
type LinkStruct struct {
	URL  string
	Copy string
}

//PromoStruct to be exported
type PromoStruct struct {
	Link string
	Copy string
}

//Marketing to be exported
type Marketing struct {
	Week      string
	Campaigns []Campaign
	Links     []LinkStruct
	Promos    []PromoStruct
}

//Document to be exported
type Document struct {
	XMLName   xml.Name `xml:"document"`
	Text      string   `xml:",chardata"`
	Wpc       string   `xml:"wpc,attr"`
	Cx        string   `xml:"cx,attr"`
	Cx1       string   `xml:"cx1,attr"`
	Cx2       string   `xml:"cx2,attr"`
	Cx3       string   `xml:"cx3,attr"`
	Cx4       string   `xml:"cx4,attr"`
	Cx5       string   `xml:"cx5,attr"`
	Cx6       string   `xml:"cx6,attr"`
	Cx7       string   `xml:"cx7,attr"`
	Cx8       string   `xml:"cx8,attr"`
	Mc        string   `xml:"mc,attr"`
	Aink      string   `xml:"aink,attr"`
	Am3d      string   `xml:"am3d,attr"`
	O         string   `xml:"o,attr"`
	R         string   `xml:"r,attr"`
	M         string   `xml:"m,attr"`
	V         string   `xml:"v,attr"`
	Wp14      string   `xml:"wp14,attr"`
	Wp        string   `xml:"wp,attr"`
	W10       string   `xml:"w10,attr"`
	W         string   `xml:"w,attr"`
	W14       string   `xml:"w14,attr"`
	W15       string   `xml:"w15,attr"`
	W16cex    string   `xml:"w16cex,attr"`
	W16cid    string   `xml:"w16cid,attr"`
	W16       string   `xml:"w16,attr"`
	W16se     string   `xml:"w16se,attr"`
	Wpg       string   `xml:"wpg,attr"`
	Wpi       string   `xml:"wpi,attr"`
	Wne       string   `xml:"wne,attr"`
	Wps       string   `xml:"wps,attr"`
	Ignorable string   `xml:"Ignorable,attr"`
	Body      struct {
		Text string `xml:",chardata"`
		P    []struct {
			Text         string `xml:",chardata"`
			ParaID       string `xml:"paraId,attr"`
			TextID       string `xml:"textId,attr"`
			RsidR        string `xml:"rsidR,attr"`
			RsidRPr      string `xml:"rsidRPr,attr"`
			RsidRDefault string `xml:"rsidRDefault,attr"`
			RsidP        string `xml:"rsidP,attr"`
			PPr          struct {
				Text string `xml:",chardata"`
				Jc   struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"jc"`
				RPr struct {
					Text   string `xml:",chardata"`
					RFonts struct {
						Text       string `xml:",chardata"`
						ASCIITheme string `xml:"asciiTheme,attr"`
						HAnsiTheme string `xml:"hAnsiTheme,attr"`
					} `xml:"rFonts"`
					Sz struct {
						Text string `xml:",chardata"`
						Val  string `xml:"val,attr"`
					} `xml:"sz"`
					SzCs struct {
						Text string `xml:",chardata"`
						Val  string `xml:"val,attr"`
					} `xml:"szCs"`
					U struct {
						Text string `xml:",chardata"`
						Val  string `xml:"val,attr"`
					} `xml:"u"`
					B     string `xml:"b"`
					Color struct {
						Text       string `xml:",chardata"`
						Val        string `xml:"val,attr"`
						ThemeColor string `xml:"themeColor,attr"`
						ThemeShade string `xml:"themeShade,attr"`
					} `xml:"color"`
				} `xml:"rPr"`
				PStyle struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"pStyle"`
				NumPr struct {
					Text string `xml:",chardata"`
					Ilvl struct {
						Text string `xml:",chardata"`
						Val  string `xml:"val,attr"`
					} `xml:"ilvl"`
					NumID struct {
						Text string `xml:",chardata"`
						Val  string `xml:"val,attr"`
					} `xml:"numId"`
				} `xml:"numPr"`
			} `xml:"pPr"`
			R []struct {
				Text    string `xml:",chardata"`
				RsidR   string `xml:"rsidR,attr"`
				RsidRPr string `xml:"rsidRPr,attr"`
				RPr     struct {
					Text   string `xml:",chardata"`
					RFonts struct {
						Text       string `xml:",chardata"`
						ASCIITheme string `xml:"asciiTheme,attr"`
						HAnsiTheme string `xml:"hAnsiTheme,attr"`
					} `xml:"rFonts"`
					Sz struct {
						Text string `xml:",chardata"`
						Val  string `xml:"val,attr"`
					} `xml:"sz"`
					SzCs struct {
						Text string `xml:",chardata"`
						Val  string `xml:"val,attr"`
					} `xml:"szCs"`
					U struct {
						Text string `xml:",chardata"`
						Val  string `xml:"val,attr"`
					} `xml:"u"`
					B     string `xml:"b"`
					Color struct {
						Text       string `xml:",chardata"`
						Val        string `xml:"val,attr"`
						ThemeColor string `xml:"themeColor,attr"`
						ThemeShade string `xml:"themeShade,attr"`
					} `xml:"color"`
				} `xml:"rPr"`
				T struct {
					Text  string `xml:",chardata"`
					Space string `xml:"space,attr"`
				} `xml:"t"`
				LastRenderedPageBreak string `xml:"lastRenderedPageBreak"`
			} `xml:"r"`
			Sdt struct {
				Text  string `xml:",chardata"`
				SdtPr struct {
					Text string `xml:",chardata"`
					RPr  struct {
						Text   string `xml:",chardata"`
						RFonts struct {
							Text       string `xml:",chardata"`
							ASCIITheme string `xml:"asciiTheme,attr"`
							HAnsiTheme string `xml:"hAnsiTheme,attr"`
						} `xml:"rFonts"`
						Sz struct {
							Text string `xml:",chardata"`
							Val  string `xml:"val,attr"`
						} `xml:"sz"`
						SzCs struct {
							Text string `xml:",chardata"`
							Val  string `xml:"val,attr"`
						} `xml:"szCs"`
						U struct {
							Text string `xml:",chardata"`
							Val  string `xml:"val,attr"`
						} `xml:"u"`
						B     string `xml:"b"`
						Color struct {
							Text       string `xml:",chardata"`
							Val        string `xml:"val,attr"`
							ThemeColor string `xml:"themeColor,attr"`
							ThemeShade string `xml:"themeShade,attr"`
						} `xml:"color"`
					} `xml:"rPr"`
					Alias struct {
						Text string `xml:",chardata"`
						Val  string `xml:"val,attr"`
					} `xml:"alias"`
					Tag struct {
						Text string `xml:",chardata"`
						Val  string `xml:"val,attr"`
					} `xml:"tag"`
					ID struct {
						Text string `xml:",chardata"`
						Val  string `xml:"val,attr"`
					} `xml:"id"`
					Placeholder struct {
						Text    string `xml:",chardata"`
						DocPart struct {
							Text string `xml:",chardata"`
							Val  string `xml:"val,attr"`
						} `xml:"docPart"`
					} `xml:"placeholder"`
					ShowingPlcHdr string `xml:"showingPlcHdr"`
				} `xml:"sdtPr"`
				SdtEndPr struct {
					Text string `xml:",chardata"`
					RPr  struct {
						Text  string `xml:",chardata"`
						Color struct {
							Text string `xml:",chardata"`
							Val  string `xml:"val,attr"`
						} `xml:"color"`
						RFonts struct {
							Text       string `xml:",chardata"`
							ASCIITheme string `xml:"asciiTheme,attr"`
							HAnsiTheme string `xml:"hAnsiTheme,attr"`
						} `xml:"rFonts"`
						B  string `xml:"b"`
						Sz struct {
							Text string `xml:",chardata"`
							Val  string `xml:"val,attr"`
						} `xml:"sz"`
						SzCs struct {
							Text string `xml:",chardata"`
							Val  string `xml:"val,attr"`
						} `xml:"szCs"`
					} `xml:"rPr"`
				} `xml:"sdtEndPr"`
				SdtContent struct {
					Text string `xml:",chardata"`
					R    []struct {
						Text    string `xml:",chardata"`
						RsidR   string `xml:"rsidR,attr"`
						RsidRPr string `xml:"rsidRPr,attr"`
						RPr     struct {
							Text   string `xml:",chardata"`
							RFonts struct {
								Text       string `xml:",chardata"`
								ASCIITheme string `xml:"asciiTheme,attr"`
								HAnsiTheme string `xml:"hAnsiTheme,attr"`
							} `xml:"rFonts"`
							Color struct {
								Text       string `xml:",chardata"`
								Val        string `xml:"val,attr"`
								ThemeColor string `xml:"themeColor,attr"`
								ThemeShade string `xml:"themeShade,attr"`
							} `xml:"color"`
							Sz struct {
								Text string `xml:",chardata"`
								Val  string `xml:"val,attr"`
							} `xml:"sz"`
							SzCs struct {
								Text string `xml:",chardata"`
								Val  string `xml:"val,attr"`
							} `xml:"szCs"`
							U struct {
								Text string `xml:",chardata"`
								Val  string `xml:"val,attr"`
							} `xml:"u"`
							RStyle struct {
								Text string `xml:",chardata"`
								Val  string `xml:"val,attr"`
							} `xml:"rStyle"`
							B string `xml:"b"`
						} `xml:"rPr"`
						T struct {
							Text  string `xml:",chardata"`
							Space string `xml:"space,attr"`
						} `xml:"t"`
					} `xml:"r"`
					Hyperlink struct {
						Text    string `xml:",chardata"`
						ID      string `xml:"id,attr"`
						History string `xml:"history,attr"`
						R       struct {
							Text    string `xml:",chardata"`
							RsidR   string `xml:"rsidR,attr"`
							RsidRPr string `xml:"rsidRPr,attr"`
							RPr     struct {
								Text   string `xml:",chardata"`
								RStyle struct {
									Text string `xml:",chardata"`
									Val  string `xml:"val,attr"`
								} `xml:"rStyle"`
								RFonts struct {
									Text       string `xml:",chardata"`
									ASCIITheme string `xml:"asciiTheme,attr"`
									HAnsiTheme string `xml:"hAnsiTheme,attr"`
								} `xml:"rFonts"`
								B  string `xml:"b"`
								Sz struct {
									Text string `xml:",chardata"`
									Val  string `xml:"val,attr"`
								} `xml:"sz"`
								SzCs struct {
									Text string `xml:",chardata"`
									Val  string `xml:"val,attr"`
								} `xml:"szCs"`
							} `xml:"rPr"`
							T string `xml:"t"`
						} `xml:"r"`
					} `xml:"hyperlink"`
					ProofErr struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"proofErr"`
				} `xml:"sdtContent"`
			} `xml:"sdt"`
			ProofErr []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
			} `xml:"proofErr"`
		} `xml:"p"`
		Sdt []struct {
			Text  string `xml:",chardata"`
			SdtPr struct {
				Text string `xml:",chardata"`
				RPr  struct {
					Text   string `xml:",chardata"`
					RFonts struct {
						Text       string `xml:",chardata"`
						ASCIITheme string `xml:"asciiTheme,attr"`
						HAnsiTheme string `xml:"hAnsiTheme,attr"`
					} `xml:"rFonts"`
					Color struct {
						Text       string `xml:",chardata"`
						Val        string `xml:"val,attr"`
						ThemeColor string `xml:"themeColor,attr"`
						ThemeShade string `xml:"themeShade,attr"`
					} `xml:"color"`
					Sz struct {
						Text string `xml:",chardata"`
						Val  string `xml:"val,attr"`
					} `xml:"sz"`
					SzCs struct {
						Text string `xml:",chardata"`
						Val  string `xml:"val,attr"`
					} `xml:"szCs"`
					B string `xml:"b"`
				} `xml:"rPr"`
				Alias struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"alias"`
				Tag struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"tag"`
				ID struct {
					Text string `xml:",chardata"`
					Val  string `xml:"val,attr"`
				} `xml:"id"`
				Placeholder struct {
					Text    string `xml:",chardata"`
					DocPart struct {
						Text string `xml:",chardata"`
						Val  string `xml:"val,attr"`
					} `xml:"docPart"`
				} `xml:"placeholder"`
			} `xml:"sdtPr"`
			SdtContent struct {
				Text string `xml:",chardata"`
				P    struct {
					Text         string `xml:",chardata"`
					ParaID       string `xml:"paraId,attr"`
					TextID       string `xml:"textId,attr"`
					RsidR        string `xml:"rsidR,attr"`
					RsidRPr      string `xml:"rsidRPr,attr"`
					RsidRDefault string `xml:"rsidRDefault,attr"`
					RsidP        string `xml:"rsidP,attr"`
					PPr          struct {
						Text string `xml:",chardata"`
						RPr  struct {
							Text   string `xml:",chardata"`
							RFonts struct {
								Text       string `xml:",chardata"`
								ASCIITheme string `xml:"asciiTheme,attr"`
								HAnsiTheme string `xml:"hAnsiTheme,attr"`
							} `xml:"rFonts"`
							Color struct {
								Text       string `xml:",chardata"`
								Val        string `xml:"val,attr"`
								ThemeColor string `xml:"themeColor,attr"`
								ThemeShade string `xml:"themeShade,attr"`
							} `xml:"color"`
							Sz struct {
								Text string `xml:",chardata"`
								Val  string `xml:"val,attr"`
							} `xml:"sz"`
							SzCs struct {
								Text string `xml:",chardata"`
								Val  string `xml:"val,attr"`
							} `xml:"szCs"`
						} `xml:"rPr"`
					} `xml:"pPr"`
					R []struct {
						Text string `xml:",chardata"`
						RPr  struct {
							Text   string `xml:",chardata"`
							RFonts struct {
								Text       string `xml:",chardata"`
								ASCIITheme string `xml:"asciiTheme,attr"`
								HAnsiTheme string `xml:"hAnsiTheme,attr"`
							} `xml:"rFonts"`
							Color struct {
								Text       string `xml:",chardata"`
								Val        string `xml:"val,attr"`
								ThemeColor string `xml:"themeColor,attr"`
								ThemeShade string `xml:"themeShade,attr"`
							} `xml:"color"`
							Sz struct {
								Text string `xml:",chardata"`
								Val  string `xml:"val,attr"`
							} `xml:"sz"`
							SzCs struct {
								Text string `xml:",chardata"`
								Val  string `xml:"val,attr"`
							} `xml:"szCs"`
						} `xml:"rPr"`
						T struct {
							Text  string `xml:",chardata"`
							Space string `xml:"space,attr"`
						} `xml:"t"`
					} `xml:"r"`
					ProofErr []struct {
						Text string `xml:",chardata"`
						Type string `xml:"type,attr"`
					} `xml:"proofErr"`
				} `xml:"p"`
			} `xml:"sdtContent"`
		} `xml:"sdt"`
		SectPr struct {
			Text            string `xml:",chardata"`
			RsidR           string `xml:"rsidR,attr"`
			RsidRPr         string `xml:"rsidRPr,attr"`
			RsidSect        string `xml:"rsidSect,attr"`
			HeaderReference struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				ID   string `xml:"id,attr"`
			} `xml:"headerReference"`
			PgSz struct {
				Text string `xml:",chardata"`
				W    string `xml:"w,attr"`
				H    string `xml:"h,attr"`
			} `xml:"pgSz"`
			PgMar struct {
				Text   string `xml:",chardata"`
				Top    string `xml:"top,attr"`
				Right  string `xml:"right,attr"`
				Bottom string `xml:"bottom,attr"`
				Left   string `xml:"left,attr"`
				Header string `xml:"header,attr"`
				Footer string `xml:"footer,attr"`
				Gutter string `xml:"gutter,attr"`
			} `xml:"pgMar"`
			Cols struct {
				Text  string `xml:",chardata"`
				Space string `xml:"space,attr"`
			} `xml:"cols"`
			DocGrid struct {
				Text      string `xml:",chardata"`
				LinePitch string `xml:"linePitch,attr"`
			} `xml:"docGrid"`
		} `xml:"sectPr"`
	} `xml:"body"`
}
