package zip

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"log"
	"os"
	"strings"
	"testing"
)

func TestZipOneFile(t *testing.T) {
	err := ZipOneFile("../README.md")
	t.Log(err)
}

func TestZipMultiFile(t *testing.T) {
	err := ZipMultiFile("zip_test.zip", []string{"../go.mod", "zip.go"})
	t.Log(err)
}

func TestUnzipFile(t *testing.T) {
	err := UnzipFile("zip_test.zip",".")
	t.Log(err)
}

func TestHttp(t *testing.T) {
	//url := "https://www.instagram.com/skuukzky/"
	//req,_ := http.NewRequest(http.MethodGet, url, nil)
	//client := http.Client{}
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36")
	//req.Header.Set("Cookie", "mid=YR4dpAALAAHC7ZWenKwNjc3IDDXg; ig_did=026AB2CF-85B6-4664-9E66-5DBDBDB78951; ig_nrcb=1; shbid=\"276\\0542933823632\\0541660900212:01f7ddee73f4ae284a25d3d7df95427aafe4e3fa3e643ae316cde52304f633623bbe0984\"; shbts=\"1629364212\\0542933823632\\0541660900212:01f7ff7431405b3741021657c32335359547b45deb8af75ab627bb8f4e0914f3105cade5\"; csrftoken=611F7GkCph5apy7aWtbMOxPqxnuw4EP1; ds_user_id=2933823632; sessionid=2933823632%3Am6FQfrJap39R2r%3A24; rur=\"ATN\\0542933823632\\0541660904270:01f7239d6771b2ea592b6bdce0a35bcfbc821c45c2fa073289b456727e5c05776ead7eec")
	//resp, err := client.Do(req)
	//contents ,_ :=ioutil.ReadAll(resp.Body)
	//if err !=  nil {
	//	log.Fatal(err)
	//}
	//f, err := os.OpenFile("index.html", os.O_RDWR| os.O_CREATE | os.O_TRUNC, os.ModePerm)
	//defer resp.Body.Close()
	//if err !=  nil {
	//	log.Fatal(err)
	//}
	//r := bufio.NewWriter(f)
	//r.Write(contents)
	//r.Flush()
	//f.Close()
	f , err := os.Open("index.html")
	if err !=  nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}
	//if item.text().strip().startswith('window._sharedData'):
	//js_data = json.loads(item.text()[21:-1], encoding='utf-8')
	//edges = js_data["entry_data"]["ProfilePage"][0]["graphql"]["user"]["edge_owner_to_timeline_media"]["edges"]
	//for edge in edges:
	//url = edge['node']['display_url']
	//print(url)
	//urls.append(url)
	pictures := make([]string, 0)
	doc.Find("script[type=\"text/javascript\"]").Each(func(i int, selection *goquery.Selection) {
		//if len(strings.TrimSpace(selection.Text())) > 21 {
		//	fmt.Println(i, strings.TrimSpace(selection.Text())[:21])
		//}
		if strings.HasPrefix(strings.TrimSpace(selection.Text()), "window._sharedData") {
			jsData := selection.Text()[21:len(selection.Text())-1]
			result := gjson.Get(jsData,"entry_data.ProfilePage.0.graphql.user.edge_owner_to_timeline_media.edges")
			result.ForEach(func(key, value gjson.Result) bool {
				fmt.Println(value.Get("node.display_url"))
				pictures = append(pictures, value.Get("node.display_url").String())
				return true // keep iterating
			})
		}
	})
	fmt.Println(len(pictures))
	fmt.Println(pictures)
}