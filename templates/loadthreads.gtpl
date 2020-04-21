<ul class="list-group list-group-flush">
  {{range .}}
  <li class="list-group-item">
  <div class="px-2 small">
    <h6 class="card-title mb-1" style="margin-bottom:0 !important">
    {{if .IsCompanyLink}}<a class="text-primary" href="{{.Url}}" target="_blank">{{.Title}}</a>
    {{else}}<a class="text-primary" href="/thread?tid={{.ID}}">{{.Title}}</a>
    {{end}}
    </h6>
    <p class="description card-text mb-1 text-muted">Posted by {{.Name}} &nbsp; {{.TimeAgo}}{{if .IsCompanyLink}}{{else}} &nbsp; {{.Count}} comments{{end}}</p>
  </div>
  </li>
  {{end}}
</ul>
