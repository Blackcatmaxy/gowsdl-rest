package gowsdl

var serverTmpl = `
var soapClient *soap.Client

func AddEndPoints(server *gin.Engine, client *soap.Client) {
	{{range . -}}
		{{range .Operations -}}
				{{$requestType := findType .Input.Message | replaceReservedWords | makePublic -}} ` + `
  				server.POST("/{{$requestType | lower }}", {{$requestType}}Func)
		{{end}}
	{{end}}
	soapClient = client
}

{{range .}}
	{{range .Operations}}
		{{$responseType := findType .Output.Message | replaceReservedWords | makePublic}}
		{{$requestType := findType .Input.Message | replaceReservedWords | makePublic}}
		{{$requestTypeSource := findType .Input.Message | replaceReservedWords }}
func {{$requestType}}Func(context *gin.Context) {
	{{$privateType := makePrivate $requestType -}}
	var {{$privateType}} {{$requestType}}
	if err := context.ShouldBindJSON(&{{$privateType}}); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := NewServicePortType(soapClient)
	response, err := service.{{$requestType}}(&{{$privateType}})
	if (err != nil) {
		context.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"result": response})
}
	{{end}}
{{end}}

`
