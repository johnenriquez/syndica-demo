<html>
<head>
<title>Home</title>
<link rel="stylesheet" href="/main.css" />
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-120199721-1"></script> <script>   window.dataLayer = window.dataLayer || [];   function gtag(){dataLayer.push(arguments);}   gtag('js', new Date());    gtag('config', 'UA-120199721-1'); </script></head>
<body>
<h2>Hello {{.}}</h2>
<hr/>
<form method="get" action="/admin">
    <button type="submit">Admin Page</button>
</form>
<form method="get" action="/profile">
    <button type="submit">My Profile</button>
</form>
<form method="get" action="/advisor">
    <button type="submit">View Advisors</button>
</form>
<form method="get" action="/company">
    <button type="submit">View Companies</button>
</form>
<form method="post" action="/logout">
    <button type="submit">Logout</button>
</form>
</body>
</html>