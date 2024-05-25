package main

import (
	"flick_tickets/common/log"
	"flick_tickets/common/utils"
	"flick_tickets/configs"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

type MovieDetail struct {
	MovieTime int64
}

func main() {
	//r := gin.Default()
	// r.GET("/api/sse/:id", SSEHandler)

	//	fmt.Println(result)

	// r.Use(cors.AllowAll())
	// r.GET("/load", LoadFileHtml)
	// r.Run(":8080")
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(utils.GeneratePassword())

	// }
	configs.LoadConfig("./configs/configs.json")
	log.LoadLogger()
	// Example movie details
	//addTimeUseTest()
	//sortSearch()
	fmt.Println(utils.ConvertTimestampToDateTime(1716670567))
	// exportExcel()

}
func exportExcel() {
	// Tạo một slice chứa các instance của struct Users
	users := []Users{
		{UserName: "John Doe", ShiftNames: "Morning Shift", CinemaName: "Cinema 1", Age: 30, Address: "123 Main St"},
		{UserName: "Jane Smith", ShiftNames: "Evening Shift", CinemaName: "Cinema 2", Age: 25, Address: "456 Elm St"},
	}

	// Tạo một file Excel mới
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Users")
	if err != nil {
		fmt.Println("Error creating sheet:", err)
		return
	}

	// Tạo header cho sheet
	row := sheet.AddRow()
	row.AddCell().SetValue("User Name")
	row.AddCell().SetValue("Shift Names")
	row.AddCell().SetValue("Cinema Name")
	row.AddCell().SetValue("Age")
	row.AddCell().SetValue("Address")

	// Thêm dữ liệu từ slice users vào sheet
	for _, user := range users {
		row := sheet.AddRow()
		row.AddCell().SetValue(user.UserName)
		row.AddCell().SetValue(user.ShiftNames)
		row.AddCell().SetValue(user.CinemaName)
		row.AddCell().SetInt(user.Age)
		row.AddCell().SetValue(user.Address)
	}

	// Lưu file Excel
	err = file.Save("users.xlsx")
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	fmt.Println("Excel file saved successfully")
}

type Users struct {
	UserName   string `json:"user_name"`
	ShiftNames string `json:"shift_names"`
	CinemaName string `json:"cinema_name"`
	Age        int    `json:"age"`
	Address    string `json:"address"`
}

func addTimeUseTest() {
	// Get the current time
	currentTime := time.Now()

	// Add 15 minutes to the current time
	timeInFuture := currentTime.Add(19 * time.Minute)

	// Get the Unix timestamp of the future time
	timeInFutureTimestamp := timeInFuture.Unix()

	// Print the current time and the future time
	fmt.Println("Current Unix timestamp:", currentTime.Unix())
	fmt.Println("Future Unix timestamp (15 minutes later):", timeInFutureTimestamp)

}
func sortSearch() {
	listRespDetail := []MovieDetail{
		{MovieTime: 1715879227}, // 2021-06-01 10:00:00 UTC
		{MovieTime: 1622635200}, // 2021-06-02 10:00:00 UTC
		{MovieTime: 1622721600}, // 2021-06-03 10:00:00 UTC
		{MovieTime: 1715879227},
	}

	// Sort the list by MovieTime (assuming it is not already sorted)
	sort.Slice(listRespDetail, func(i, j int) bool {
		return listRespDetail[i].MovieTime < listRespDetail[j].MovieTime
	})

	// Get the current time as a Unix timestamp
	timeNowTypetimestamp := time.Now().Unix()

	// Find the first movie detail where MovieTime is greater than or equal to the current time
	index := sort.Search(len(listRespDetail), func(i int) bool {
		return listRespDetail[i].MovieTime >= timeNowTypetimestamp
	})

	// Check if such a movie detail was found
	if index < len(listRespDetail) {
		fmt.Printf("Next available movie detail: %+v\n", listRespDetail[index])
	} else {
		fmt.Println("No available movie details found")
	}
}

func LoadFileHtml(c *gin.Context) {
	path := "cmd/test/index.html"
	htmlBytes, err := os.ReadFile(path)
	if err != nil {
		// Xử lý lỗi nếu có
		c.String(http.StatusInternalServerError, "Lỗi khi đọc tệp HTML")
		return
	}

	// Trả về trang HTML
	c.Data(http.StatusOK, "text/html; charset=utf-8", htmlBytes)
}

func SSEHandler(c *gin.Context) {
	id := c.Param("id")

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")

	// Tạo kênh SSE
	ch := make(chan sse.Event)
	defer close(ch)

	// Goroutine để lấy dữ liệu từ use case khác và gửi SSE
	go func() {
		for {
			// Lấy dữ liệu từ use case khác
			data := GetDataFromUseCase(id)
			event := sse.Event{
				Event: "message",
				Data:  data,
			}
			ch <- event
			time.Sleep(1 * time.Second)
		}
	}()

	// Gửi dữ liệu SSE đến client
	c.Stream(func(w io.Writer) bool {
		if _, ok := <-ch; ok {
			sse.Encode(w, sse.Event{
				Data: GetDataFromUseCase(id),
			})
			w.Write([]byte("\n"))
			return true
		}
		return false
	})

	// Xử lý sự kiện đóng kết nối từ client
	<-c.Writer.CloseNotify()
	// Thực hiện các tác vụ khi kết nối bị đóng
	println("Connection closed by client")
}

type SinhVien struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var data = []SinhVien{
	SinhVien{
		Id:   "1",
		Name: "A 11",
	},
	SinhVien{
		Id:   "2",
		Name: "B 22",
	},
	SinhVien{
		Id:   "3",
		Name: "B 3",
	},
}

func GetDataFromUseCase(id string) []SinhVien {
	var list = make([]SinhVien, 0)

	for i := 0; i < len(data); i++ {
		if data[i].Id == id {
			list = append(list, data[i])
		}
	}
	return list
}
