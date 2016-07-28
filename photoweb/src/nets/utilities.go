package nets

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

//检查路径是否存在
func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

//模板渲染
func RenderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
	t := template.Must(template.ParseFiles(Templates[tmpl]))
	err := t.Execute(w, locals)
	CheckError(err)
}

func SafeHandler(fn http.HandlerFunc) http.HandlerFunc {
	fmt.Println("......")
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)

				// 或者输出自定义的 50x 错误页面
				// w.WriteHeader(http.StatusInternalServerError)
				// renderHtml(w, "error", e.Error())

				// logging
				log.Println("WARN: panic fired in %v.panic - %v", fn, e)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}

//静态资源处理
func StaticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & ListDir) == 0 {
			if exists := IsExists(file); !exists {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}
