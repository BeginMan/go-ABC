package nets

import (
	"io/ioutil"
	"log"
	"path"
)

const (
	UPLOAD_DIR   = "./uploads"
	TEMPLATE_DIR = "./views"
	ListDir      = 0x0001
)

//声明并初始化一个全局变量 Templates,用于存放所有模板内容
//Templates 是一个map类型的复合结构,
//map的键(key)是字符串类型,即模板的名字,值(value) 是 *template.Template 类型
var Templates = make(map[string]string)

func Init() {
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
	CheckError(err)

	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:", templatePath)
		//		t := template.Must(template.ParseFiles(templatePath))
		Templates[templateName] = templatePath
	}
}
