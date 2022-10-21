package main

import (
	"fmt"
	"github.com/Mmx233/BitSrunLoginGo/controllers"
	"github.com/Mmx233/BitSrunLoginGo/global"
	"github.com/Mmx233/BitSrunLoginGo/util"
	log "github.com/sirupsen/logrus"
)

func main() {
	if global.Flags.RunningDaemon {
		//后台挂起模式中
		controllers.Guardian()
	} else if global.Config.Settings.Guardian.Enable {
		//进入守护模式流程
		controllers.EnterGuardian()
	} else {
		//登录流程
		var err error
		if global.Config.Settings.Basic.Interfaces == "" { //单网卡
			if err = controllers.Login(nil); err != nil {
				log.Errorln("登录出错: ", err)
				if !global.Config.Settings.Log.DebugLevel {
					fmt.Printf("开启调试日志（debug_level）获取详细信息")
				}
				return
			}
		} else { //多网卡
			log.Debugln("多网卡模式")
			interfaces, _ := util.GetInterfaceAddr()
			for _, eth := range interfaces {
				log.Debugln("使用网卡: ", eth.Name)
				if err = controllers.Login(eth.Addr); err != nil {
					log.Errorln("登录出错: ", err)
				}
			}
		}
	}
}
