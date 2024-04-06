package example

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	chromeDriver = "/home/jesse/project/go-study/chromedriver"
	port         = 8080
)

type job struct {
	Name      string `json:"name"`
	Area      string `json:"area"`
	Salary    string `json:"salary"`
	Exp       string `json:"exp"`
	Education string `json:"education"`
	Tags      string `json:"tags"`
	Desc      string `json:"desc"`
	Publis    string `json:"publis"`
	Cmp       string `json:"cmp"`
	Scale     string `json:"scale"`
}

func getBrowser() (browser selenium.WebDriver, service *selenium.Service) {
	// 开启 webdriver 服务
	service, _ = selenium.NewChromeDriverService(chromeDriver, port)
	var caps = selenium.Capabilities{}
	// 设置 chrome 特定功能
	var chromeCaps = chrome.Capabilities{
		// 使用开发者调试模式
		ExcludeSwitches: []string{"enable-automation"},
		Args: []string{
			"--no-sandbox",
			//设置请求头
			"--user-agent=Mozilla/5,0 (Windows NT 10.0; Win64;" +
				"x64) AppleWebKit/537.36 (KHTML, like Gecko) " +
				"Chrome/94.0.4606.61 Safari/537,36",
		},
	}
	// 将 chrome 特定模式添加到 caps
	caps.AddChrome(chromeCaps)

	// 根据浏览器功能连接 Selenium
	var urlPrefix = fmt.Sprintf("http://localhost:%d/wd/hub", port)
	browser, _ = selenium.NewRemote(caps, urlPrefix)

	return browser, service
}

func getJobs(browser selenium.WebDriver) (jobs []job) {

	browser.Wait(func(browser selenium.WebDriver) (bool, error) {
		_, err := browser.FindElement(selenium.ByClassName, "search-job-result")
		if err == nil {
			return true, nil
		} else {
			return false, nil
		}
	})
	var jobCards, err = browser.FindElements(selenium.ByClassName, "job-card-wrapper")
	if err != nil {
		fmt.Println("获取不到元素！")
	}
	for _, v := range jobCards {
		var job job

		// 职务
		var name, _ = v.FindElement(selenium.ByClassName, "job-name")
		job.Name, _ = name.Text()
		// 工作地点
		var area, _ = v.FindElement(selenium.ByClassName, "job-area")
		job.Area, _ = area.Text()
		// 薪资
		var salary, _ = v.FindElement(selenium.ByClassName, "salary")
		job.Salary, _ = salary.Text()
		// 经验
		var exp, _ = v.FindElement(selenium.ByCSSSelector, ".tag-list li:first-child")
		job.Exp, _ = exp.Text()
		// 学历
		var education, _ = v.FindElement(selenium.ByCSSSelector, ".job-card-body .tag-list li:last-child")
		job.Education, _ = education.Text()
		// 职位
		var tags, _ = v.FindElement(selenium.ByCSSSelector, ".job-card-footer .tag-list")
		job.Tags, _ = tags.Text()
		// 公司福利
		var desc, _ = v.FindElement(selenium.ByClassName, "info-desc")
		job.Desc, _ = desc.Text()
		// 人事信息
		var publis, _ = v.FindElement(selenium.ByClassName, "info-public")
		job.Publis, _ = publis.Text()
		// 公司名称
		var cmp, _ = v.FindElement(selenium.ByCSSSelector, ".company-name > a")
		job.Cmp, _ = cmp.Text()
		// 公司行业和规模
		var scale, _ = v.FindElement(selenium.ByCSSSelector, `.company-tag-list li:last-child`)
		job.Scale, _ = scale.Text()

		jobs = append(jobs, job)
	}
	return jobs
}

func saveData(jobs []job) {
	var fs, _ = os.OpenFile("jobs.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	var encoder = json.NewEncoder(fs)
	var err = encoder.Encode(jobs)
	if err != nil {
		fmt.Println("JSON写入失败：", err)
	} else {
		fmt.Println("JSON写入成功")

	}
}

func ExecAutoTest() {
	// 获取浏览器对象
	var browser, service = getBrowser()
	// 关闭服务
	defer service.Stop()
	// 关闭浏览器对象
	defer browser.Quit()

	// 访问网址
	browser.Get("https://www.zhipin.com/")
	// 最大化窗口
	browser.MaximizeWindow("")
	time.Sleep(5 * time.Second)

	// 输入查询职位
	var query, _ = browser.FindElement(selenium.ByName, "query")
	query.SendKeys("go 语言")
	time.Sleep(2 * time.Second)
	// 单击搜索按钮
	var search, _ = browser.FindElement(selenium.ByCSSSelector, `[class="btn btn-search"]`)
	search.Click()
	time.Sleep(2 * time.Second)

	// 获取第一页的职位信息
	var jobs = getJobs(browser)
	// 翻页执行
	// for {
	// 	var nextPage, err = browser.FindElement(selenium.ByCSSSelector, `.options-pages a:last-child`)
	// 	if err != nil {
	// 		break
	// 	} else {
	// 		nextPage.Click()
	// 		time.Sleep(2 * time.Second)
	// 		jobs = append(jobs, getJobs(browser)...)
	// 	}
	// }
	saveData(jobs)

}
