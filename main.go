package main

//https://blog.csdn.net/idwtwt/article/details/81180460

import (
	"fmt"
	fileAction "httpFileServer/action/file"
	fileServerAction "httpFileServer/action/fileServer"

	"httpFileServer/filter"
	"httpFileServer/handler/errorHandler"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println(os.Args)
	////////////////////////////////embedding File Server start/////////////////////////////////////

	//http.Handle("/", http.FileServer(&fs)) //to FileServerfolder file without auth
	fileServerFilter := filter.NewFilterRouter()
	fileServerHandler := new(filter.FilterHandler)
	fileServerHandler.SetAction(fileServerAction.NewFromAssetfs().ServerFile)
	fileServerFilter.AddHanlder("^/(.*)$", fileServerHandler)
	////////////////////////////////embedding File Server end/////////////////////////////////////

	////////////////////////////////outer File Server start/////////////////////////////////////
	userFileFilter := filter.NewFilterRouter()
	userFileHandler := new(filter.FilterHandler)
	userFileHandler.SetAction(fileServerAction.NewFromFilePath("d:/").ServerFile)
	userFileHandler.AddFilter(fileServerAction.UserFileFilter)

	userFileFilter.AddHanlder("^/(.*)$", userFileHandler)

	////////////////////////////////outer File Server end/////////////////////////////////////

	////////////////////////////////API start/////////////////////////////////////
	downloadFileHandler := new(filter.FilterHandler)
	downloadFileHandler.SetAction(fileAction.DownloadFile)

	uploadFileHandler := new(filter.FilterHandler)
	uploadFileHandler.SetAction(fileAction.UploadFile)

	listFileHandler := new(filter.FilterHandler)
	listFileHandler.SetAction(fileAction.ListFile)

	httpAPIFilter := filter.NewFilterRouter()
	httpAPIFilter.GlobalErrorHandler = errorHandler.Handle

	httpAPIFilter.AddHanlder("/api/file/down", downloadFileHandler)
	httpAPIFilter.AddHanlder("/api/file/upload", uploadFileHandler)
	httpAPIFilter.AddHanlder("/api/file/list", listFileHandler)
	////////////////////////////////API end/////////////////////////////////////

	///AddHttp routing
	http.Handle("/", fileServerFilter) // provide download validate

	http.HandleFunc("/api/", httpAPIFilter.Routing) // provide download validate

	http.Handle("/file/", http.StripPrefix("/file/", userFileFilter)) // provide download validate

	log.Println(http.ListenAndServe(":8080", nil))
}
