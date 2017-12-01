package graph_verification

import (

	"html/template"
	"log"
	"net/http"
	"fmt"
)

// 图片框长宽大小
const (
	dx = 300
	dy = 100
)


// 生成图片验证码,前后端不分离
//func main() {
//
//	//所有的相对路径都是相对于该项目开始
//	err := ReadFonts("src/graph.verification/example/fonts", ".ttf")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	http.HandleFunc("/", Index)
//	http.HandleFunc("/get/", Get)
//	fmt.Println("服务已启动...")
//	err = http.ListenAndServe(":8800", nil)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("src/graph.verification/example/tpl/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

func Get(w http.ResponseWriter, r *http.Request) {
	captchaImage, err := NewCaptchaImage(dx, dy, RandLightColor())

	captchaImage.DrawNoise(CaptchaComplexLower)

	captchaImage.DrawTextNoise(CaptchaComplexLower)

	//生成验证码的位数
	code := RandText(4)
	fmt.Printf("code:%s",code)
	fmt.Println()
	captchaImage.DrawText(code)

	captchaImage.DrawBorder(ColorToRGB(0x17A7A7A))

	captchaImage.DrawSineLine()

	if err != nil {
		fmt.Println(err)
	}

	captchaImage.SaveImage(w, ImageFormatJpeg)
}
