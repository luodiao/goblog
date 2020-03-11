package admin

import (
	"fmt"
	"io/ioutil"

	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

type AdminController struct {
	baseController
}

// 登录
func (c *AdminController) Login() {
	if c.Ctx.Request.Method == "POST" {
		uname := c.GetString("username")
		upass := c.GetString("password")

		beego.Info(uname, upass)
	}

	fmt.Println(c.controller, c.action)

	passwordOK := "admin"

	hash, err := bcrypt.GenerateFromPassword([]byte(passwordOK), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(hash)

	c.Data["offLeft"] = -100
	c.TplName = c.controller + "/login.tpl"
}

// 注册
func (c *AdminController) Register() {
	c.TplName = c.controller + "/register.tpl"
}

// 验证码图片
func (c *AdminController) VerifyImg() {
	src := "./static/image/block_2.jpg"

	c.Ctx.Output.Header("Content-Type", "image/jpg")
	c.Ctx.Output.Header("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", src))

	file, err := ioutil.ReadFile(src)
	if err != nil {
		beego.Info("文件不存在")
		return
	}

	c.Ctx.WriteString(string(file))

	// f1, err := os.Open(src)
	// if err != nil {
	// 	panic(err)
	// }
	// defer f1.Close()

	// buf := bytes.NewBuffer(nil)
	// if _, err := io.Copy(buf, f1); err != nil {
	// 	panic(err)
	// }

	// c.Ctx.WriteString(string(buf.Bytes()))
}
