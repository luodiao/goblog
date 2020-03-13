package admin

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type AdminController struct {
	baseController
}

type clipSize struct {
	x0 int
	y0 int
	x1 int
	y1 int
}

func getClipSize() clipSize {
	x0 := rand.Intn(270)
	y0 := rand.Intn(120)
	return clipSize{
		x0: x0,
		y0: y0,
		x1: x0 + 30,
		y1: y0 + 30,
	}
}

// 登录
func (c *AdminController) Login() {
	if c.Ctx.Request.Method == "POST" {
		if c.Ctx.Request.PostFormValue("action") == "verify" {
			if a, _ := strconv.ParseFloat(c.Ctx.Request.PostFormValue("accuracy"), 10); math.Abs(a) > 3 {
				c.Data["json"] = getResponse(-1, "验证失败", "")
			} else {
				c.Data["json"] = getResponse(0, "success", "")
			}

			c.ServeJSON()
			return
		}
		uname := c.GetString("username")
		upass := c.GetString("password")

		beego.Info(uname, upass)
	}

	// fmt.Println(c.controller, c.action)

	// passwordOK := "admin"

	// hash, err := bcrypt.GenerateFromPassword([]byte(passwordOK), bcrypt.DefaultCost)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(hash)

	size := getClipSize()
	var srcArr = [3]string{"./static/image/t1.png", "./static/image/t2.png", "./static/image/t3.png"}
	src := srcArr[rand.Intn(2)+1]

	// 裁剪图片
	clip(src, "./static/image/watermark.png", size.x0, size.y0, size.x1, size.y1)

	// 拼图
	glue("./static/image/verify_bg.png", "./static/image/watermark.png", "./static/image/out.png", size.x0, size.y0, false)

	// 原图
	glue(src, "./static/image/watermark.png", "./static/image/out_bg.png", size.x0, size.y0, true)

	v := strconv.FormatInt(time.Now().Unix(), 10)
	c.Data["offLeft"] = size.x0 - (size.x0 * 2)
	c.Data["verifyImg"] = "/static/image/out.png?v=" + v
	c.Data["verifyImgBg"] = "/static/image/out_bg.png?v=" + v
	c.TplName = c.controller + "/login.tpl"
}

type user struct {
	a int
	b int
}

// 注册
func (this *AdminController) Register() {
	if this.Ctx.Request.Method == "GET" {
		this.TplName = this.controller + "/register.tpl"
		return
	}

	// add
	// if this.Ctx.Request.PostFormValue("action") == "add" {
	u := params{}
	u.Model = user{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
	}

	fmt.Printf("%T,%#v\n", u, u)
	this.Data["json"] = getResponse(0, "success", "")
	this.ServeJSON()
	return
	// }
}

/*
	剪裁图片
	inName 输入图片名称
	outName 输出图片名称
	x0 剪裁x起点
	y0 剪裁y起点
	x1 剪裁x终点
	y1 剪裁y终点
*/
func clip(inName string, outName string, x0 int, y0 int, x1 int, y1 int) error {
	fIn, _ := os.Open(inName)
	defer fIn.Close()

	fOut, _ := os.Create(outName)
	defer fOut.Close()

	origin, fm, err := image.Decode(fIn)
	if err != nil {
		return err
	}

	switch fm {
	case "jpeg":
		img := origin.(*image.YCbCr)
		subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.YCbCr)

		return jpeg.Encode(fOut, subImg, &jpeg.Options{100})
	case "png":
		switch origin.(type) {
		case *image.Paletted:
			img := origin.(*image.Paletted)
			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.Paletted)
			return png.Encode(fOut, subImg)
		case *image.NRGBA:
			img := origin.(*image.NRGBA)
			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.NRGBA)
			return png.Encode(fOut, subImg)
		case *image.RGBA:
			img := origin.(*image.RGBA)
			subImg := img.SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
			return png.Encode(fOut, subImg)
		}

	default:
		return errors.New("Error format")
	}

	return nil
}

/*
	合成图片
	inName 输入图片名称
	watermarkName 水印图片
	x0 剪裁x起点
	y0 剪裁y起点
	hd 颜色是否反转
*/
func glue(inName string, watermarkName string, outName string, x0 int, y0 int, hd bool) {
	fIn, _ := os.Open(inName)
	fInImg, _ := png.Decode(fIn)
	defer fIn.Close()

	fOut, _ := os.Open(watermarkName)
	fWatermarkName, _ := png.Decode(fOut)
	defer fOut.Close()

	if hd == true {
		fWatermarkName = hdImage(fWatermarkName)
	}

	offset := image.Pt(x0, y0)
	switch fInImg.(type) {
	case *image.Paletted:
		m := image.NewPaletted(fInImg.Bounds(), []color.Color{
			color.Gray{Y: 255},
			color.Gray{Y: 160},
			color.Gray{Y: 70},
			color.Gray{Y: 35},
			color.Gray{Y: 0},
		})
		draw.Draw(m, fInImg.Bounds(), fInImg, image.ZP, draw.Src)
		draw.Draw(m, fWatermarkName.Bounds().Add(offset), fWatermarkName, image.ZP, draw.Over)

		//生成新图片
		out, _ := os.Create(outName)
		png.Encode(out, m)

		defer out.Close()
	case *image.NRGBA:
		m := image.NewNRGBA(fInImg.Bounds())
		draw.Draw(m, fInImg.Bounds(), fInImg, image.ZP, draw.Src)
		draw.Draw(m, fWatermarkName.Bounds().Add(offset), fWatermarkName, image.ZP, draw.Over)

		//生成新图片
		out, _ := os.Create(outName)
		png.Encode(out, m)

		defer out.Close()
	case *image.RGBA:
		m := image.NewRGBA(fInImg.Bounds())
		draw.Draw(m, fInImg.Bounds(), fInImg, image.ZP, draw.Src)
		draw.Draw(m, fWatermarkName.Bounds().Add(offset), fWatermarkName, image.ZP, draw.Over)

		//生成新图片
		out, _ := os.Create(outName)
		png.Encode(out, m)

		defer out.Close()
	}
}

//图片灰色处理
func hdImage(m image.Image) *image.RGBA {
	bounds := m.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	newRgba := image.NewRGBA(bounds)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			// colorRgb := m.At(i, j)
			// _, g, _, a := colorRgb.RGBA()
			// g_uint8 := uint8(g >> 8)
			// a_uint8 := uint8(a >> 8)

			// newRgba.SetRGBA(i, j, color.RGBA{g_uint8, g_uint8, g_uint8, a_uint8})
			newRgba.SetRGBA(i, j, color.RGBA{255, 255, 255, 255})
		}
	}

	return newRgba
}
