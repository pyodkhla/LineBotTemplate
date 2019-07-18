// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				log.Println("===== " + message.Text + " =====")
				mukSeaw(event.ReplyToken, message.Text)
				/*if message.Text == "Hi" {
					if _, err = bot.ReplyMessage(event.ReplyToken,
						linebot.NewTextMessage("Hi, how are you today?")).Do(); err != nil {
						log.Print(err)
					}
				} else {
					if _, err = bot.ReplyMessage(event.ReplyToken,
						linebot.NewTextMessage(message.ID+":" + message.Text+" OK!")).Do(); err != nil {
						log.Print(err)
					}
				}*/
			default:
				if _, err = bot.ReplyMessage(event.ReplyToken,
					linebot.NewTextMessage("Sorry, this command is not support.")).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}

func sendReplyMessage(replyToken string, message string) error {
	if _, err := bot.ReplyMessage(replyToken,
		linebot.NewTextMessage(message)).Do(); err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func mukSeaw(replyToken string, message string) bool {
	if message == "" {
		return false
	}

	if m, _ := regexp.MatchString("mbot ด่า.*?ให้หน่อย", message); m {
		name := strings.TrimSpace(message[len("mbot ด่า") : len(message)-len("ให้หน่อย")])

		dar := []string{"อีข้อศอกหมี", "อีตาปลาถูกตัดที่ร้านทำเล็บ", "อีกิ้งกือตัดต่อพันธุกรรม", "อีเล็บขบของไส้เดือน", "ไอ้แตงกวาดอง", "ไอ้กะหล่ำปลี", "อีเห็ดสด",
			"อีแมวน้ำ", "ไอ้ปูปู้", "อิอมีบาวิ่งผ่านพารามีเซียม", "อีปลวกมีปีก", "อีแบรนด์ซุปไก่สกัด", "อิโดเรม่อนไม่มีกระเป๋าวิเศษ", "อิกระดาษโดนน้ำ", "อีสายพานจักรยาน",
			"อีmouseไม่มีwheel", "อีCPU single core", "อีpower bank แบตหมด", "อีสาย usb หักใน", "อิหอยกาบสแกนดิเนเวีย", "อิต่อต้านอนุมูลอิสระ",
			"อีส้มตำไม่ใส่มะละกอ", "อี Ferrari ยกสูง", "อิน้ำยาปรับผ้านุ่ม", "อิดาบเจ็ดสี มณีเจ็ดแสง", "อีCPUริมๆWafer", "อีPower supply 200W",
			"อี Protoss ไม่มี carrier", "อีไข่เจียวไม่ใส่หมูสับ", "อี DNA เส้นตรง", "ไอ้ตุ๊กตาปูปู้", "ไอ้ผัดไทยห่อไข่ดาว", "ไอ้กระทู้พันทิป", "ไอ้แว่นตาเลนส์เว้า",
			"ไอ้หลอดไฟสี daylight", "ไอ้เสื้อยืดคอเต่า", "ไอ้เสื้อลายสก๊อต", "ไอ้หนังสือพิมพ์เปื้อนนิ้ว", "ไอ้นาฬิกา Kinetic", "ไอ้ Siri text mode",
			"ไอ้ดอกกุหลาบหนามแหลม", "อี Twitter limit 140 ตัวอักษร", "อีเบียร์ใส่น้ำแข็ง", "อีไวน์หมัก10ปี"}

		sendReplyMessage(replyToken, name+" "+dar[rand.Intn(len(dar))])
		return true
	}

	if message == "mbot help" {
		sendReplyMessage(replyToken,"คิดเองเดะ")
		return true
	}
	if message == "mbot /?" {
		sendReplyMessage(replyToken,"ไม่ช่วย ไม่ตอบ")
		return true
	}
	if message == "mbot แสด" {
		sendReplyMessage(replyToken,"ด่าตัวเองหรอ?")
		return true
	}
	if strings.HasPrefix(message, "mbot thank") {
		sendReplyMessage(replyToken,"เก็บคำนั้นไว้กับนายเถอะ")
		return true
	}
	if message == "mbot resurrect" {
		sendReplyMessage(replyToken,"ชั้นจะกลับมาเสมอ แม้นายจะเตะชั้นอีกกี่ครั้ง")
		return true
	}
	if message == "วันก่อนครับ" {
		sendReplyMessage(replyToken,"ทำไมหรอครับ?")
		return true
	}
	if message == "มีคุณยายขึ้นรถเมล์ แม่งไม่มีคนลุกเลยครับ" {
		sendReplyMessage(replyToken,"ไม่มีน้ำใจกันเลยนะครับ")
		return true
	}
	if message == "ซักพักมีผู้ชายคนนึงทนไม่ไหว ลุกให้ยายนั่ง คนร้องกันทั้งรถเลยครับ" {
		sendReplyMessage(replyToken,"เพราะชื่นชมที่เค้าเป็นสุภาพบุรุษ?")
		return true
	}
	if message == "เปล่า คนที่ลุกให้ยายนั่งอะ คนขับ" {
		sendReplyMessage(replyToken,"...")
		return true
	}
	if message == "ไปสวนสาธารณะเปิดใหม่มา" {
		sendReplyMessage(replyToken,"ไปเดินเล่นหรอครับ?")
		return true
	}
	if message == "ไปถึงนี่ ไม่มีที่ให้นั่งเลยครับ" {
		sendReplyMessage(replyToken,"คนเยอะมาก ใครๆก็ไป จนไม่มีที่นั่ง?")
		return true
	}
	if message == "เปล่า มีแต่ม้านั่งครับ..." {
		sendReplyMessage(replyToken,"แสดด")
		return true
	}
	if message == "mbot เก่งมาก" {
		sendReplyMessage(replyToken,"ไม่ต้องมาแกล้งชมชั้นหรอก")
		return true
	}
	if message == "เฮ้ย ชมจริงๆ" {
		sendReplyMessage(replyToken,"อ่ะๆ กองไว้ตรงนั้นแหละ")
		return true
	}
	if message == "mbot ขอบใจนะ" {
		sendReplyMessage(replyToken,"เก็บคำนั้นไว้เถอะ")
		return true
	}
	if message == "mbot เขียนโปรแกรมให้หน่อยได้มะ" {
		sendReplyMessage(replyToken,"วันก่อนครับ")
		return true
	}
	if message == "ทำไมหรอครับ??" {
		sendReplyMessage(replyToken,"มีฝรั่งดูโค้ดผม บอกว่าโค้ดผมสะอาดมากเลยครับ")
		return true
	}
	if message == "เค้าพูดว่าไรหรอครับ??" {
		sendReplyMessage(replyToken,"ยัวร์ โค้ด ซัก")
		return true
	}

	return false
}