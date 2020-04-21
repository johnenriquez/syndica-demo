<!DOCTYPE html>
<html lang="en-US">
<meta charset="UTF-8" />
<title>Syndica Advisors</title>
<meta name="viewport" content="width=device-width, initial-scale=1" />
<link rel='dns-prefetch' href='//fonts.googleapis.com' />
<link rel="stylesheet" href="/theme.css" />
<link rel="stylesheet" href="/main.css" />
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-120199721-1"></script> <script>   window.dataLayer = window.dataLayer || [];   function gtag(){dataLayer.push(arguments);}   gtag('js', new Date());    gtag('config', 'UA-120199721-1'); </script></head>
<body>
<div class="maincontainer">
<h3>Syndica Advisors</h3>
<hr/>
<div class="headnav"><a href="/home">Home</a> | <a href="/logout">Logout</a></div>
<BR><BR>
<div id="advisors">
  <input class="search" name="user_search_bar" placeholder="Search here" />
  <table>
    <thead>
      <tr><th>Name</th><th>Email</th><th>Year</th></tr>
    </thead>
    <tbody class="list">
        {{range .}}
        <tr>
            <td class="user_name"><a href="/advisor?email={{.Email}}">{{.Name}}</a></td>
            <td class="user_email">{{.Email}}</td>
            <td class="user_year">{{.Year}}</td>
        </tr>
        {{end}}
    </tbody>
  </table>
</div>

<BR><BR><BR>

<script src="//cdnjs.cloudflare.com/ajax/libs/list.js/1.5.0/list.min.js"></script>
<script>
var options = {
  valueNames: [ 'user_name', 'user_email', 'user_email2', 'user_year' ]
};
var userList = new List('advisors', options);
</script>

</body>
</html>