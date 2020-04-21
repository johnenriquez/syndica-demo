{{range .}}
    <div class="companygriditem">
        <img src="{{.Logo}}">
        {{.Name}}
        {{.Short}}
    </div>
{{end}}
