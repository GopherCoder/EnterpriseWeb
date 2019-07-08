package cmd

import (
	"EnterpriseWeb/EnterpriseWithMux/unicorn/pkg/database"
	"EnterpriseWeb/EnterpriseWithMux/unicorn/web/model"
	"io"
	"log"
	"net/http"
	"strings"
	"unicode"

	"github.com/tidwall/gjson"

	"golang.org/x/net/html"

	"github.com/spf13/cobra"
)

var dataCMD = &cobra.Command{
	Use: "data",
	Run: func(cmd *cobra.Command, args []string) {
		database.EngineInit()
		defer database.Engine.Close()
		database.Engine.LogMode(true)
		var unicorn UnicornCompanyRequest
		unicorn = UnicornCompanyRequest{
			Request:   "https://dujiaoshou.io/#",
			ParseFunc: Parse,
		}
		Run(unicorn)

	},
}

type UnicornCompanyRequest struct {
	Request   string `json:"request"`
	ParseFunc func(io.Reader) interface{}
}

func Run(vars UnicornCompanyRequest) {
	request, _ := http.NewRequest(http.MethodGet, vars.Request, nil)
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		log.Println("err : ", err.Error())
		return
	}
	vars.ParseFunc(response.Body)
}

func Parse(r io.Reader) interface{} {
	doc, _ := html.Parse(r)
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			if strings.Contains(n.Data, "dataManager") {
				data := gjson.Parse(n.Data)
				for _, i := range data.Get("props.pageProps.companies").Array() {
					//fmt.Println(i.Get("name"), i.Get("link"), i.Get("country"), i.Get("last_funding_on"), i.Get("category"), i.Get("post_money_val"))

					var country model.Country

					name := HandleCountry(i.Get("country").String())
					if dbError := database.Engine.Where("name = ?", name).First(&country).Error; country.ID == 0 && dbError != nil {
						country.Name = name
						database.Engine.Save(&country)
					}

					var category model.Category

					if dbError := database.Engine.Where("name = ?", i.Get("category").String()).First(&category).Error; category.ID == 0 && dbError != nil {
						category.Name = i.Get("category").String()
						database.Engine.Save(&category)
					}

					var company model.Company
					if dbError := database.Engine.Where("name = ?", i.Get("name").String()).First(&company).Error; dbError != nil && company.ID != 0 {
						company = model.Company{
							Name:          i.Get("name").String(),
							WebSite:       i.Get("link").String(),
							Valuation:     uint(i.Get("post_money_val").Uint()),
							ValuationDate: i.Get("last_funding_on").Time(),
							CountryID:     country.ID,
							CategoryID:    category.ID,
						}
						database.Engine.Save(&company)
					} else {
						company.CategoryID = category.ID
						database.Engine.Save(&company)
					}

				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return nil
}

func HandleCountry(v string) string {
	var buf strings.Builder
	for _, i := range v {
		if unicode.IsLetter(i) {
			buf.WriteRune(i)
		}
	}
	return buf.String()
}
