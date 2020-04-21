<html>
<head>
<title>Edit Users</title>
<link rel="stylesheet" href="/main.css" />
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-120199721-1"></script> <script>   window.dataLayer = window.dataLayer || [];   function gtag(){dataLayer.push(arguments);}   gtag('js', new Date());    gtag('config', 'UA-120199721-1'); </script></head>
<body>
<h3>All Users</h3>
<hr/>
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
<hr>
<div class="maincontainer">

<div id="users">
  <input class="search" name="user_search_bar" placeholder="Search here" />
  <table>
    <thead>
      <tr><th>Name</th><th>Email</th><th>Email2</th><th>Role</th><th>Status</th></tr>
    </thead>
    <tbody class="list">
        {{range .}}
        <tr>
            <td class="user_name"><a href="/admin/user?email={{.Email}}">{{.Name}}</a></td>
            <td class="user_email">{{.Email}}</td>
            <td class="user_email2">{{.Email2}}</td>
            <td class="user_role">{{.Role}}</td>
            <td class="user_status">{{.Status}}</td>
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
  valueNames: [ 'user_name', 'user_email', 'user_email2', 'user_year', 'user_role', 'user_status' ]
};
var userList = new List('users', options);
</script>

</body>
</html>