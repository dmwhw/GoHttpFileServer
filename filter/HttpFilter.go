package filter

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type DoFilter func(w http.ResponseWriter, req *http.Request) (err error,
	code int,
	msg string,
	isHanlded bool)

type DoAction func(w http.ResponseWriter, req *http.Request) //实际会强转

type FilterHandler struct {
	Filters []DoFilter //前置拦截器
	Action  DoAction   //任务
	//错误处理？
	//后置？
}

func (filterHandler *FilterHandler) AddFilter(doFilter DoFilter) (result bool) {
	if filterHandler.Filters == nil {
		filterHandler.Filters = make([]DoFilter, 0, 2)
	}
	filterHandler.Filters = append(filterHandler.Filters, doFilter)
	return true
}
func (filterHandler *FilterHandler) SetAction(doAction DoAction) (result bool) {
	filterHandler.Action = doAction
	return false
}

/////////////////////fiterRouting
type ErrorHandler func(w http.ResponseWriter, req *http.Request, err error,
	code int,
	msg string,
	isHanlded bool) //实际会强转

type FilterRouter struct {
	routerMapping      map[string]*FilterHandler
	GlobalErrorHandler ErrorHandler //全局处理
}

func NewFilterRouter() (filterRouter *FilterRouter) {
	filterRouter = new(FilterRouter)
	filterRouter.routerMapping = make(map[string]*FilterHandler)
	return
}

func (filterRouter *FilterRouter) hanldeError(w http.ResponseWriter, req *http.Request, err error,
	code int,
	msg string,
	isHanlded bool) {
	if filterRouter.GlobalErrorHandler == nil {
		w.Header().Set("Content-Type", "application/xml")
		w.Write([]byte("server error "))
		return
	}
	filterRouter.GlobalErrorHandler(w, req, err, code, msg, isHanlded) //全局处理
}

func (filterRouter *FilterRouter) AddHanlder(uri string, filterHandler *FilterHandler) {
	filterRouter.routerMapping[uri] = filterHandler
	return

}
func (filterRouter *FilterRouter) execute(filterHandler *FilterHandler, w http.ResponseWriter, req *http.Request) {
	if filterHandler.Filters == nil {
		filterHandler.Action(w, req)
		return
	}
	for _, myfilter := range filterHandler.Filters {
		err, code, msg, isHanlded := myfilter(w, req)
		if err != nil {
			filterRouter.hanldeError(w, req, err, code, msg, isHanlded)
			return
		}
		if isHanlded {
			break
		}
	}
	filterHandler.Action(w, req)
	return
}
func (filterRouter *FilterRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	filterRouter.Routing(w, req)
}

func (filterRouter *FilterRouter) Routing(w http.ResponseWriter, req *http.Request) {
	////正侧表达式
	//匹配路径，调用对应的
	//优先匹配精确路径
	//通配符路径
	var uri string = req.RequestURI
	fmt.Println("request-->", uri)
	uri2 := uri
	if !strings.HasSuffix(uri, "/") {
		uri2 += "/"
	}
	filterHandler, hasKey := filterRouter.routerMapping[uri]
	if hasKey {
		filterRouter.execute(filterHandler, w, req)
		return
	}
	filterHandler2, hasKey2 := filterRouter.routerMapping[uri2]
	if hasKey2 {
		filterRouter.execute(filterHandler2, w, req)
		return
	}
	//regex
	for key, filterHandler := range filterRouter.routerMapping {
		matched, _ := regexp.MatchString(key, uri)
		if matched {
			filterRouter.execute(filterHandler, w, req)
			return
		}
	}
	//404
	//writeHTML(w, "filter:404NotFound", 404)
	http.NotFoundHandler().ServeHTTP(w, req)
}

func writeHTML(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(msg))
}
