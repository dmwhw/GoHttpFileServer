package fileServerAction

import (
	"fmt"
	"httpFileServer/assets/web"
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
)

//httpFileServer.ServerFile

type HttpFileServerHandler struct {
	HttpHandler http.Handler
}

func (httpFileServerHandler *HttpFileServerHandler) ServerFile(w http.ResponseWriter, req *http.Request) {
	httpFileServerHandler.HttpHandler.ServeHTTP(w, req)
}

func UserFileFilter(w http.ResponseWriter, req *http.Request) (err error,
	code int,
	msg string,
	isHanlded bool) {
	fmt.Println("userFileFilter.....")
	return nil, 1, "", false
}

func NewFromAssetfs() *HttpFileServerHandler {
	fs := assetfs.AssetFS{
		Asset:     web.Asset,
		AssetDir:  web.AssetDir,
		AssetInfo: web.AssetInfo,
		Fallback:  "index.html",
		Prefix:    "static"}
	httpFileServer := new(HttpFileServerHandler)
	httpFileServer.HttpHandler = http.FileServer(&fs)
	return httpFileServer
}

func NewFromFilePath(path string) *HttpFileServerHandler {
	httpFileServer := new(HttpFileServerHandler)
	httpFileServer.HttpHandler = http.FileServer(http.Dir(path))
	return httpFileServer
}
