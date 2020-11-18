package wdservice

import (
	"log"

	"github.com/tebeka/selenium"
)

//WebDrivceService webdrvice服务
type WebDrivceService struct {
}

//GetEleStringByValue 通过 字符串的方式 制定 对应的 查询类型 以及查询的 元素条件
func (l *WebDrivceService) GetEleStringByValue(wd selenium.WebDriver, ByString, Value string) selenium.WebElement {
	element, err := wd.FindElement(ByString, Value)
	if err != nil {
		log.Print("获取"+Value+"的时候发生错误-->", err)
		return nil
	}
	return element
}

//GetEleByXPathID 通过XPATH进行搜索dom节点
func (l *WebDrivceService) GetEleByXPathID(wd selenium.WebDriver, Value string) selenium.WebElement {
	element := l.GetEleStringByValue(wd, selenium.ByXPATH, Value)

	return element
}
