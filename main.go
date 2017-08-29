package main

import (
	"os/exec"
	"fmt"
	"strings"
	"tools/mail"
	"time"
)

//func update(screen *ebiten.Image) error {
//	ebitenutil.DebugPrint(screen, "Hello worlddd!")
//	return nil
//}
//
//func main() {
//	ebiten.Run(update, 320, 240, 2, "Hello world!")
//}

var mailSend = false

func main() {
	for {
		findText()
		time.Sleep(2 * time.Second)
	}
}

func findText() {
	cmd := exec.Command("ls")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	outStr := string(out)
	outArr := strings.Split(outStr, "\n")
	finded := false
	for _, v := range outArr {
		if v == "1.txt" {
			finded = true
			break
		}
	}
	if !finded {
		if !mailSend {
			go mail.SendMailWithQQMail("找到一个bug", "服务器挂了！！！" + time.Now().Format("2006-01-02 15:04:05"), "332114994@qq.com")
			mailSend = true
		}
	} else {
		mailSend = false
	}
}
