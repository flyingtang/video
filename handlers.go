package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"log"
)

// 播放视频
func streamHandle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 路径
	vid := ps.ByName("vid-id")
	videoPath := video_dir + vid

	// 打开
	video, err := os.Open(videoPath)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "no find video")
		return
	}
	// 发送
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	// 关闭
	defer video.Close()
}

// 上传视频
func uploadHandle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	http.MaxBytesReader(w, r.Body, MaxSize)
	if err := r.ParseMultipartForm(MaxSize); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "file is too big")
		return
	}
	// 中间参数可用验证content
	file, h, err := r.FormFile("file")
	if err != nil {
		log.Printf("parse file err: %s", err.Error())
		sendErrorResponse(w, http.StatusInternalServerError, "internal server error ")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("read file err: %s", err.Error())
		sendErrorResponse(w, http.StatusInternalServerError, "internal server error ")
	}

	//fn := ps.ByName("id")
	err = ioutil.WriteFile(video_dir+h.Filename, data, 0666)
	if err != nil {
		log.Printf("write file err: %s", err.Error())
		sendErrorResponse(w, http.StatusInternalServerError, "internal server error ")
	}
	return
}

// 测试上传视频前端模板

func testHanle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

//	const tpl = `
//		<!DOCTYPE html>
//<html lang="en">
//<head>
//    <meta charset="UTF-8">
//    <title>upload</title>
//</head>
//<body>
//    <form enctype="multipart/form-data" action="http://127.0.0.1:4001/upload/123" method="post">
//        <input type="file" name="file">
//        <button type="submit" value="input file" >上传</button>
//    </form>
//</body>
//</html>
//	`
//	t := template.New("new template")
//	t , _ = t.ParseFiles("upload.html")
 t , _:=template.ParseFiles("upload.html")
	t.Execute(w, nil)
}