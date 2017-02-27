package mvc

import (
	"regexp"
	"reflect"
	"strings"
)

var StaticDir map[string]string

type controllerInfo struct {
	regex *regexp.Regexp
	params map[int]string
	controllerType reflect.Type
}

type ControllerRegistor struct {
	routers []*controllerInfo
	Application *App
}

type App struct {
}

type ControllerInterface struct {

}

func (p *ControllerRegistor) Add(pattern string, c ControllerInterface)  {
	parts := strings.Split(pattern, "/")

	j := 0
	params := make(map[int]string)

	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			expr := "([^/]+)"

			//a user may choose to override the defult expression
			// similar to expressjs: ‘/user/:id([0-9]+)’

			if index := strings.Index(part, "("); index != -1 {
				expr = part[index:]
				part = part[:index]
			}
			params[j] = part
			parts[i] = expr
			j++
		}
	}

	pattern = strings.Join(parts, "/")
	regex, regexErr := regexp.Compile(pattern)

	if regexErr != nil {
		panic(regexErr)
		return
	}

	t := reflect.Indirect(reflect.ValueOf(c)).Type()
	route := &controllerInfo{}
	route.regex = regex
	route.params = params
	route.controllerType = t

	p.routers = append(p.routers, route)
}

// sample mvc.SetStaticPath("/img","/static/img")
func (app *App) SetStaticPath(url string, path string) *App {
	StaticDir[url] = path
	return app
}