package example

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func convertToString(src, srcdoe, tagCode string) (result string) {
	srcCoder := mahonia.NewDecoder(srcdoe)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result = string(cdata)
	return result
}

type Movie struct {
	gorm.Model
	Name        string `gorm:"type:varchar(50)"`
	Star        string `gorm:"type:varchar(50)"`
	Releasetime string `gorm:"type:varchar(50)"`
	Score       string `gorm:"type:varchar(50)"`
}

func getData(offset string) (result string) {
	var urls = "https://maoyan.com/board/4?offset=" + offset
	fmt.Println("url：", urls)

	// 定义请求对象
	var req, _ = http.NewRequest("GET", urls, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Winx64) AppleWebKit/537.36 (KHTML, like Gecko)Chrome/94.0.4606.81 Safari/537.36")
	client := &http.Client{}
	var res, _ = client.Do(req)
	// 获取网站响应内容
	body, _ := io.ReadAll(res.Body)
	result = string(body)
	time.Sleep(5 * time.Second)

	return result
}

func cleanData(data string) (result []map[string]string) {
	var dom, _ = goquery.NewDocumentFromReader(strings.NewReader(data))

	var info map[string]string

	var selection = dom.Find(".board-item-content")
	selection.Each(func(i int, s *goquery.Selection) {
		info = map[string]string{}
		name := selection.Find(".name").Text()
		star := selection.Find(".star").Text()
		releasetime := selection.Find(".releasetime").Text()
		score := selection.Find(".score").Text()

		info["name"] = strings.TrimSpace(name)
		info["star"] = strings.TrimSpace(star)
		info["releasetime"] = strings.TrimSpace(releasetime)
		info["score"] = strings.TrimSpace(score)

		result = append(result, info)
	})

	return result
}

func saveDataOnReq(data []map[string]string) {
	var dsn = `root:123456@tcp(127.0.0.1:3306)/test`
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	var DB, _ = db.DB()
	defer DB.Close()
	db.AutoMigrate(&Movie{})

	for _, v := range data {
		fmt.Println("当前数据：", v)
		var movies []Movie
		db.Where("name = ?", v["name"]).First(&movies)
		if len(movies) == 0 {
			var m1 = Movie{
				Name:        v["name"],
				Star:        v["star"],
				Releasetime: v["releasetime"],
				Score:       v["score"],
			}
			db.Create(&m1)
		} else {
			db.Where("name = ?", v["name"]).Find(&movies).Update("socre", v["socre"])
		}
	}

}

func ExecRequest() {

	for i := 0; i < 10; i++ {
		webData := getData(strconv.Itoa(i * 10))
		fmt.Println(webData)
		cleanData := cleanData(webData)
		saveDataOnReq(cleanData)
	}

}
