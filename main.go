package main

import (
	"crypto/tls"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	jsoniter "github.com/json-iterator/go"
	"io"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var version = "0.0.7"
var (
	Started    bool
	start      *widget.Button
	name       *widget.Entry
	signature  *widget.Entry
	seessionId *widget.Entry
	phone      *widget.Entry
	chepai     *widget.Entry
	textArea   *widget.Entry
)

func getDate() []byte {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", "http://newretail.pingan.com.cn/ydt/reserve/store/bookingTime?storefrontseq=39807&businessType=14", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", "")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 Edg/107.0.1418.23")
	req.Header.Set("sec-ch-ua", `"Edge";v="107", "Chromium";v="107", "Not=A?Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return bodyText
}
func yueyue(applicantIdCard string, contactName string, contactTelephone string, vehicleNo string, bookingTime string, bookingDate string, idBookingSurvey string, signature string, sessionId string) []byte {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}

	var data = strings.NewReader("{\"businessName\":\"承保验车\",\"storefrontName\":\"摩托车投保预约\",\"detailaddress\":\"北京市朝阳区世纪财富中心2号楼2层平安门店\",\"bookingType\":\"1\",\"bookingDate\":\"" + bookingDate + "\",\"bookingTime\":\"" + bookingTime + "\",\"storefrontseq\":\"39807\",\"storefrontTelephone\":\"95511\",\"businessType\":\"14\",\"bookContent\":\"\",\"idBookingSurvey\":\"" + idBookingSurvey + "\",\"deptCode\":\"39807\",\"contactName\":\"" + contactName + "\",\"contactTelephone\":\"" + contactTelephone + "\",\"applicantName\":\"\",\"applicantIdCard\":\"" + applicantIdCard + "\",\"bookingSource\":\"miniApps\",\"businessKey\":null,\"agentFlag\":\"0\",\"newCarFlag\":\"0\",\"noPolicyFlag\":\"0\",\"vehicleNo\":\"" + vehicleNo + "\",\"inputPolicyNo\":\"\",\"latitude\":\"\",\"longitude\":\"\",\"offlineItemList\":[]}")
	req, err := http.NewRequest("POST", "http://newretail.pingan.com.cn/ydt/reserve/reserveOffline?time=1687830149279", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) MicroMessenger/7.0.0(0x17000000) MacWechat/3.6.2(0x13060211) MiniProgramEnv/Mac MiniProgram")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Host", "newretail.pingan.com.cn")
	req.Header.Set("Accept-Language", "zh-CN,zh-Hans;q=0.9")
	req.Header.Set("signature", signature)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("sessionId", sessionId)
	req.Header.Set("Origin", "https://newretail.pingan.com.cn")
	req.Header.Set("Referer", "https://newretail.pingan.com.cn/ydt/newretail/")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	s := fmt.Sprintf("%s", bodyText)
	textArea.Text = s + "\n" + textArea.Text
	textArea.Refresh()
	return bodyText
}
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
func yuyuepre(applicantIdCard string, contactName string, contactTelephone string, vehicleNo string, bookingTime string, bookingDate string, idBookingSurvey string, signature string, sessionId string) {
	for true {
		now_time := time.Now()
		if strings.Contains(now_time.String(), "17:4") {
			textArea.Text = "捡漏模式关闭\n" + textArea.Text
			textArea.Refresh()
			break
		}
		yueyue(applicantIdCard, contactName, contactTelephone, vehicleNo, bookingTime, bookingDate, idBookingSurvey, signature, sessionId)
		time.Sleep(time.Millisecond * 200)
		textArea.Text = "已经开启捡漏模式，请在17.40左右关闭\n" + textArea.Text
		textArea.Refresh()
	}
}

func startTask(keys []string, js map[string]string) {
	rand.Seed(time.Now().UnixNano())
	sig := keys[0]
	i := 0
	start_bh := 0
	for Started {
		result := getDate()
		r := make(map[string]interface{})
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		json.Unmarshal(result, &r)
		datas := r["data"].([]interface{})
		textArea.Text = "第" + strconv.Itoa(i) + "次刷新\n" + textArea.Text
		i += 1
		index_num := 10
		if len(datas) > 0 {
			dd := datas[len(datas)-1]
			for Started {
				if index_num < 0 {
					time.Sleep(time.Second * 2)
					textArea.Text = "预约结束请查看日志\n" + textArea.Text
					phone.Enable()
					Started = false
					name.Enable()
					start.Enable()
					start.Text = "开始"
					chepai.Enable()
					start.Refresh()
					textArea.Refresh()
					return
				}
				index_num -= 1
				m := dd.(map[string]interface{})
				bookingDate := m["bookingDate"]
				bookingDates := fmt.Sprintf("%s", bookingDate)
				bookingRules := m["bookingRules"].([]interface{})
				for i := 1; i < len(bookingRules); i++ {
					bookingRule := bookingRules[len(bookingRules)-i].(map[string]interface{})
					startTime := bookingRule["startTime"]
					endTime := bookingRule["endTime"]
					idBookingSurvey := bookingRule["idBookingSurvey"]
					bookingTime := fmt.Sprintf("%s-%s", startTime, endTime)
					idBookingSurveys := fmt.Sprintf("%s", idBookingSurvey)
					go yueyue("", name.Text, phone.Text, chepai.Text, bookingTime, bookingDates, idBookingSurveys, sig, js[sig])
					if start_bh < 6 {
						start_bh += 1
						go yuyuepre("", name.Text, phone.Text, chepai.Text, bookingTime, bookingDates, idBookingSurveys, sig, js[sig])
					}

				}
			}
		}
		time.Sleep(time.Microsecond * 0)
	}
}

func getKeys3(m map[string]string) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	a := app.New()
	var t = &MyTheme{}
	a.Settings().SetTheme(t)
	w := a.NewWindow("摩托三者预约")
	w.Resize(fyne.NewSize(400, 300))
	name = widget.NewEntry()
	name.SetPlaceHolder("输入姓名")

	phone = widget.NewEntry()
	phone.SetPlaceHolder("手机号")

	seessionId = widget.NewEntry()
	seessionId.SetPlaceHolder("seessionId")
	signature = widget.NewEntry()
	signature.SetPlaceHolder("signature")

	chepai = widget.NewEntry()
	chepai.SetPlaceHolder("车牌,全部大写！格式：京B-XXXX 或者 京A-XXXX")
	textArea = widget.NewMultiLineEntry()

	start = widget.NewButton("开始", func() {
		wx := "{\"" + signature.Text + "\": \"" + seessionId.Text + "\"}"
		r := make(map[string]string)
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		json.Unmarshal([]byte(wx), &r)
		keys := getKeys3(r)
		if Started {
			start.Text = "开始"
			Started = false
			textArea.Text = "停止预约\n" + textArea.Text
			chepai.Enable()
			name.Enable()
			phone.Enable()
		} else {
			if !VerifyMobileFormat(phone.Text) {
				dialog.NewInformation("错误", "手机号格式错误", w).Show()
				return
			}

			if !strings.Contains(chepai.Text, "京A-") && !strings.Contains(chepai.Text, "京B-") {
				dialog.NewInformation("错误", "车牌格式错误", w).Show()
				return
			}

			if name.Text == "" || phone.Text == "" {
				dialog.NewInformation("错误", "姓名，手机号均不能为空", w).Show()
				return
			}
			start.Text = "停止"
			Started = true
			textArea.Text = "开始预约\n" + textArea.Text
			chepai.Disable()
			name.Disable()
			phone.Disable()
			go startTask(keys, r)
		}
		textArea.Refresh()
		start.Refresh()
	})

	content := container.NewVBox(name, phone, chepai, seessionId, signature, start, textArea)

	w.SetContent(content)
	w.ShowAndRun()
}
