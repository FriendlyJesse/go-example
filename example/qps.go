package example

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var QpsCount = 0
var START_TIME = time.Now().Unix()
var TimeList = make(map[int64]int)

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
func logContent(str string, filename string) {
	filePath := filename
	file, openErr := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if openErr != nil {
		fmt.Println("文件打开失败", openErr)
		createFile, createErr := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if createErr != nil {
			fmt.Println("文件创建失败", openErr)
		} else {
			file = createFile
		}
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	//for i := 0; i < 5; i++ {
	write.WriteString(str + " \n")
	//}
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

func indexFunc(c *gin.Context) {
	// 首页
	QpsCount += 1
	var timeInterval int64 = 1
	currentTime := time.Now().Unix()
	if currentTime%timeInterval == 0 { // 是当前秒数
		if _, ok := TimeList[currentTime]; !ok {
			TimeList[currentTime] = QpsCount
			addCount := QpsCount - TimeList[currentTime-timeInterval]
			logContent(strconv.FormatInt(currentTime, 10)+
				"   运行时间:"+strconv.FormatInt(currentTime-START_TIME, 10)+
				"  请求总数"+strconv.Itoa(QpsCount)+
				"   请求增量"+strconv.Itoa(addCount), "count_log")
		}
	}
	c.String(http.StatusOK, "ok")
}

func ExecQPS() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("", indexFunc)
	router.Run("127.0.0.1:8023")
}
