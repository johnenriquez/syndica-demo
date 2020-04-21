<html>
<head>
<title>Admin</title>
<link rel="stylesheet" href="/main.css" />
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-120199721-1"></script> <script>   window.dataLayer = window.dataLayer || [];   function gtag(){dataLayer.push(arguments);}   gtag('js', new Date());    gtag('config', 'UA-120199721-1'); </script></head>
<body>
<h2>Admin Section</h2>
<hr/>
<form method="get" action="/home">
    <button type="submit">Home</button>
</form>
<form method="get" action="/admin/messages">
    <button type="submit">Admin Messages</button>
</form>
<form method="get" action="/register">
    <button type="submit">Add New User</button>
</form>
<form method="get" action="/admin/user">
    <button type="submit">Edit Users</button>
</form>
<form method="get" action="/register">
    <button type="submit">Add New Advisor</button>
</form>
<form method="get" action="/admin/advisor">
    <button type="submit">Edit Advisors</button>
</form>
<form method="get" action="/register/company">
    <button type="submit">Add New Company</button>
</form>
<form method="get" action="/admin/company">
    <button type="submit">Edit Companies</button>
</form>
<form method="get" action="/logout">
    <button type="submit">Logout</button>
</form>
</body>
</html>