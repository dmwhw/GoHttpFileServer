package fileAction

import (
	"fmt"
	"net/http"
)

func UploadFile(w http.ResponseWriter, req *http.Request) {
	fmt.Println("upload...")
	//
	//multipart.File, *multipart.FileHeader, error
	//file ,fileHeader ,error=req.FormFile("")
}

func DownloadFile(w http.ResponseWriter, req *http.Request) {
	// req.ParseForm()
	// form := req.Form
	// fmt.Println("down...")
	// for k, v := range req.Form {
	// 	fmt.Println("key:", k)
	// 	fmt.Println("val:", strings.Join(v, ""))
	// }
	// //namespace
	// //filePath
	// fp := path.Join("images", "foo.png")
	// log.Print(fp)
	// http.ServeFile(w, req, fp)
}

func ListFile(w http.ResponseWriter, req *http.Request) {
	fmt.Println("list...")
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("aaaaalistfile"))
}
