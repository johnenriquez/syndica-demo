<ul class="list-group list-group-flush">
  {{range .}}
  <li class="list-group-item d-flex align-items-center">
    <h5 class="h5 my-2">
    <a class="text-primary" href="/company?name={{.Company}}&q={{.ID}}">{{.Question}}</a>
    </h5>
  </li>
  {{end}}
</ul>
