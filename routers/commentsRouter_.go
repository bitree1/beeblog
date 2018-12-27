package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beeblog/controllers:PageController"] = append(beego.GlobalControllerRouter["beeblog/controllers:PageController"],
        beego.ControllerComments{
            Method: "ServerDemined",
            Router: `/403`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeblog/controllers:PageController"] = append(beego.GlobalControllerRouter["beeblog/controllers:PageController"],
        beego.ControllerComments{
            Method: "PageNotFound",
            Router: `/404`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeblog/controllers:PageController"] = append(beego.GlobalControllerRouter["beeblog/controllers:PageController"],
        beego.ControllerComments{
            Method: "ServerError",
            Router: `/500`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeblog/controllers:PageController"] = append(beego.GlobalControllerRouter["beeblog/controllers:PageController"],
        beego.ControllerComments{
            Method: "Blog",
            Router: `/iframe/blog`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeblog/controllers:PageController"] = append(beego.GlobalControllerRouter["beeblog/controllers:PageController"],
        beego.ControllerComments{
            Method: "IframeNote",
            Router: `/iframe/note`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeblog/controllers:PageController"] = append(beego.GlobalControllerRouter["beeblog/controllers:PageController"],
        beego.ControllerComments{
            Method: "IframeUser",
            Router: `/iframe/user`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeblog/controllers:PageController"] = append(beego.GlobalControllerRouter["beeblog/controllers:PageController"],
        beego.ControllerComments{
            Method: "UsPage",
            Router: `/us`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
