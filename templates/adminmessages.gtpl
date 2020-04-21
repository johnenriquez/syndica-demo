<!DOCTYPE html>
<html lang="en-US">
<meta charset="UTF-8" />
<title>Syndica</title>
<meta name="viewport" content="width=device-width, initial-scale=1" />
<link rel='dns-prefetch' href='//fonts.googleapis.com' />
<link rel="stylesheet" href="/theme.css" />
<link rel="stylesheet" href="/main.css" />
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-120199721-1"></script> <script>   window.dataLayer = window.dataLayer || [];   function gtag(){dataLayer.push(arguments);}   gtag('js', new Date());    gtag('config', 'UA-120199721-1'); </script></head>
<body>
<div class="maincontainer">
<hr><div class="headnav"><a href="/company">COMPANIES</a> | <a href="/discussions">DISCUSSIONS</a> | <a href="/activity">ACTIVITY</a> | <a href="/portfolio">PORTFOLIO</a> | <a href="/messages"><b>MESSAGES</b></a> | <a href="/profile">PROFILE</a> | <a href="/logout">LOGOUT</a></div>
<BR><BR>
SEND A MESSAGE:
<BR><BR>
<form style="600px;" action="/admin/send_message" method="post">
    TO: <input style="margin-left:20px;width:250px" type="text" name="to"><BR><BR>
    <textarea cols=60 rows=5 required name="message"></textarea><BR><BR>
    <input class="submitresponse" type="submit" name="submitresponse" value="Submit">
</form>

<BR><BR>
<div id="messages">
    <UL>
    {{range .}}
    <LI>{{.Time}} - <b>From: {{.From}}</b>  -  <b>To: {{.User}}</b>  -  {{.Message}}</LI>
    {{end}}
    </UL>
</div>
<BR><BR><BR><BR>


<script src="/main.js?v=1.2.0"></script>
</body>
</html>