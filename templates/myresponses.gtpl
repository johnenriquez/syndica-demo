{{range .Responses}}
<div class="q_response mt-3" id="response_{{.ID}}" data-responseid="{{.ID}}" data-replyid="{{.ReplyID}}">
  <div class="q_response_inside">
    <div class="text-dark small"><b>{{.Name}}</b>{{if .Title}} &mdash; <span class="text-muted">{{.Title}} &middot; {{.Job}}</span>{{end}}</div>
    <div class="py-2" style="word-wrap: break-word; overflow-wrap: break-word; white-space:pre-wrap;">{{.Answer}}</div>
    {{if .IsSelf}}
      <div class="small">
        <a class="text-muted likes" id="like_{{.ID}}">Acknowledge ({{.Likes}})</a> &middot;
        <a class="text-muted reply" href="#reply_10028" data-toggle="collapse" onclick="doReply({{.ID}});">Reply</a> &middot;
        <a class="text-muted reply" style="cursor:pointer" onclick="doEditResponse({{.ID}});">Edit</a> &middot;
        <a class="text-muted reply" style="cursor:pointer" onclick="doDeleteResponse({{.ID}});">Delete</a>
      </div>
    {{else}}
      <div class="small">
        <a class="text-muted likes" id="like_{{.ID}}" style="cursor:pointer" onclick="doLike({{.ID}});">Acknowledge ({{.Likes}})</a> &middot;
        <a class="text-muted reply" href="#reply_10028" data-toggle="collapse" onclick="doReply({{.ID}});">Reply</a> &middot;
      </div>
    {{end}}
  </div><!-- .q_response_inside -->
  <div id="reply_{{.ID}}" class="replybox mt-2 collapse">
    <form class="bg-light rounded p-3 clearfix" action="/response" method="post">
      <div class="form-group mb-2">
        <textarea class="form-control" rows="3" id="reply_{{.ID}}_textarea" required="" name="answer"></textarea>
      </div>
      <input type="hidden" name="question_id" value="{{.QuestionID}}">
      <input type="hidden" name="reply_id" value="{{.ID}}">
      <input id="reply_{{.ID}}_company" type="hidden" name="company" value="">
      <input class="submitresponse btn btn-secondary btn-block" type="submit" name="submitresponse" value="Submit">
    </form>
  </div><!-- #reply_ID -->

</div><!-- #response_ID -->
{{end}}

<div class="your-response mt-3">
  <form class="bg-light rounded p-3 clearfix" action="/response" method="post">
    <div class="form-group mb-2">
      <label for="answer">Your Response</label>
      <textarea class="form-control" rows="3" required name="answer"></textarea>
    </div>
    <input id="questionid" type="hidden" name="question_id" value="{{.QuestionID}}">
    <input id="replyid" type="hidden" name="reply_id" value="{{.QuestionID}}">
    <input id="companyid" type="hidden" name="company" value="{{.Company}}">
    <input class="submitresponse btn btn-secondary btn-block" type="submit" name="submitresponse" value="SUBMIT">
  </form>
</div><!-- .your-response -->
