<html>
<head>
<title>Edit User</title>
<link rel="stylesheet" href="/main.css" />
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-120199721-1"></script> <script>   window.dataLayer = window.dataLayer || [];   function gtag(){dataLayer.push(arguments);}   gtag('js', new Date());    gtag('config', 'UA-120199721-1'); </script></head>
<body>
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

<h2>Edit User</h2>
    <form action="/admin/user" method="post">
        <label for="name">Name:</label>
        <input type="text" name="name" value="{{.Name}}"><br>

        <label for="email">Email:</label>
        <input type="email" name="email" value="{{.Email}}"><br>
        
        <label for="password">Password:</label>
        <input type="password" name="password"><br>
        
        <label for="email2">Email (School):</label>
        <input type="email" name="email2" value="{{.Email2}}"><br>

        <label for="year">Year:</label>
        <input type="number" name="year" value="{{.Year}}"><br>

        <label for="role">Role: [{{.Role}}]</label>
        <select name="role">
            <option value="none" selected>select new role</option>
            <option value="advisor">Advisor</option>
            <option value="client">Client</option>
            <option value="admin">Admin</option>
            <option value="inactive">Inactive</option>
        </select><br>

        <label for="status">Status:</label>
        <textarea name="status" rows="3" cols="60">{{.Status}}</textarea><br>

        <label for="profile">Profile:</label>
        <textarea name="profile" rows="8" cols="60">{{.Profile}}</textarea><br>
<BR><BR>
        <input type="hidden" name="editemail" value="{{.Email}}">
        <input type="submit" value="UPDATE">
    </form>

<BR><BR><BR>
</div>
</body>
</html>