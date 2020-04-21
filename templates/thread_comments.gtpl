<ul class="list-group list-group-flush">
{{range .}}
    <div class="q_response mt-3" id="response_{{.ID}}" data-responseid="{{.ID}}" data-replyid="{{.ReplyID}}">
    <div class="q_response_inside">
        <div class="text-dark small"><b>{{.Name}}</b> &nbsp; {{.TimeAgo}}</div>
        <div class="py-2" style="word-wrap: break-word; overflow-wrap: break-word; white-space:pre-wrap;line-height:1.0">{{.Comment}}</div>
        {{if .IsSelf}}
        <div class="small">
            <a class="text-muted reply" href="#reply_10028" data-toggle="collapse" onclick="doReplyThreadComment({{.ID}});">Reply</a> &middot;
            <a class="text-muted reply" style="cursor:pointer" onclick="doEditThreadComment({{.ID}},{{.DiscussionID}});">Edit</a> &middot;
            <a class="text-muted reply" style="cursor:pointer" onclick="doDeleteThreadComment({{.ID}},{{.DiscussionID}});">Delete</a>
        </div>
        {{else}}
        <div class="small">
            <a class="text-muted reply" href="#reply_10028" data-toggle="collapse" onclick="doReplyThreadComment({{.ID}});">Reply</a> &middot;
        </div>
        {{end}}
    </div><!-- .q_response_inside -->
    <div id="reply_{{.ID}}" class="replybox mt-2 collapse">
        <form class="bg-light rounded p-3 clearfix" action="/threadcommentreply" method="post">
        <div class="form-group mb-2">
            <textarea class="form-control" rows="3" cols=50 id="reply_{{.ID}}_textarea" required name="comment"></textarea>
        </div>
        <input type="hidden" name="reply_id" value="{{.ID}}">
        <input id="reply_{{.ID}}_thread" type="hidden" name="discussion_id" value="{{.DiscussionID}}">
        <input class="submitresponse btn btn-secondary btn-block" type="submit" name="submitresponse" value="Submit">
        </form>
    </div><!-- #reply_ID -->
    </div><!-- #response_ID -->
{{end}}
</ul>