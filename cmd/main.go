package main

import "sample_go_project/api"

func main() {
	api.InitServer()
}

// {{ range $data:=.datas}}
// <td>{{ .Name }}</td>
// {{end}}