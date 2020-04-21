<ul class="list-group list-group-flush">
{{range .}}
<li class="questionitem list-group-item d-flex justify-content-between align-items-center">
    <h5 class="my-3 h5">
      <a class="text-primary" href="/company?name={{.Company}}&q={{.ID}}">{{.Question}}</a>
    </h5>
    <div class="ml-3 small">
      <a class="text-muted" href="/question_edit?name={{.Company}}&q={{.ID}}">EDIT</a> |
      <a class="text-muted" href="/question_delete?name={{.Company}}&q={{.ID}}" onclick="return confirm('Are you sure you want to delete this question?');">DELETE</a>
    </div>
</li>
{{end}}
</ul>
