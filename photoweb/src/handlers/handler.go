package handlers

import (
	"io/ioutil"
	"net/http"
	"nets"
	"os"
)

// / 首页处理函数， 列出所有图片
func ListHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoAddr, err := ioutil.ReadDir(nets.UPLOAD_DIR) // fileInfoAddr: 文件对象数组
	nets.CheckError(err)

	locals := make(map[string]interface{})
	images := []string{}

	for _, fileInfo := range fileInfoAddr {
		imgid := fileInfo.Name()
		images = append(images, imgid)
	}

	locals["images"] = images
	nets.RenderHtml(w, "list.html", locals)

}

// 查看图片
func ViewHanlder(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id") // 获取请求参数
	imagePath := nets.UPLOAD_DIR + "/" + imageId
	if exists := nets.IsExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "image")
	//ServeFile()方法将该路径下的文件从磁盘中 读取并作为服务端的返回信息输出给客户端
	http.ServeFile(w, r, imagePath)
}

//上传图片
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		nets.RenderHtml(w, "upload.html", nil)

	} else if r.Method == "POST" {
		f, h, err := r.FormFile("image") // 从表单提交的字段中寻找image文件域并对其接值
		nets.CheckError(err)

		filename := h.Filename
		defer f.Close() // 关闭图片上传到服务器文件流的句柄
		t, err := os.Create(nets.UPLOAD_DIR + "/" + filename)
		nets.CheckError(err)

		defer t.Close() //关闭临时文件句柄
		nets.CheckError(err)

		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)

	}

}
