<html>
<head>
<title>Syndica Startups</title>
<link rel="stylesheet" href="/main.css" />
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-120199721-1"></script> <script>   window.dataLayer = window.dataLayer || [];   function gtag(){dataLayer.push(arguments);}   gtag('js', new Date());    gtag('config', 'UA-120199721-1'); </script></head>
<body>
<form method="post" action="/home">
    <button type="submit">Home</button>
</form>
<form method="post" action="/admin">
    <button type="submit">Admin Page</button>
</form>
<form method="post" action="/logout">
    <button type="submit">Logout</button>
</form>
<button onclick="javascript:window.history.back();">Go Back</button>

<div class="maincontainer">

<div id="companies">
  <input class="search" name="user_search_bar" placeholder="Search here" />
  <table>
    <thead>
      <tr><th>Name</th><th>Contact</th><th>Email</th><th>Website</th><th>Status</th></tr>
    </thead>
    <tbody class="list">
        {{range .}}
        <tr>
            <td class="company_name"><a href="/admin/company?name={{.Name}}">{{.Name}}</a></td>
            <td class="company_contact">{{.Contact}}</td>
            <td class="company_email">{{.Email}}</td>
            <td class="company_website">{{.Website}}</td>
            <td class="company_status">{{.Status}}</td>
        </tr>
        {{end}}
    </tbody>
  </table>
</div>

<BR><BR><BR>
</div>
<script src="//cdnjs.cloudflare.com/ajax/libs/list.js/1.5.0/list.min.js"></script>
<script>
var options = {
  valueNames: [ 'company_name', 'company_contact', 'company_email', 'company_website' ]
};
var userList = new List('companies', options);
</script>

</body>
</html>