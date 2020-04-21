<html>
<head>
<title>Edit Startup</title>
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
<h2>Edit Company</h2>
    <form action="/admin/company" method="post">
        <label for="contact">Company Contact:</label>
        <input type="text" name="contact" value="{{.Contact}}"><br>

        <label for="email">Contact Email:</label>
        <input type="email" name="email" value="{{.Email}}"><br>
        
        <label for="name">Company Name:</label>
        <input type="text" name="name" value="{{.Name}}"><br>

        <label for="short">Short Description:</label>
        <input type="text" name="short" value="{{.Short}}"><br>

        <label for="description">Long Description:</label>
        <textarea name="description" rows="6" cols="60">{{.Description}}</textarea><br>

        <label for="logo">Logo:</label>
        <input type="text" name="logo" value="{{.Logo}}"><br>

        <label for="primarypic">Product Image:</label>
        <input type="text" name="primarypic" value="{{.PrimaryPic}}"><br>

        <label for="founders">Founders:</label>
        <input type="text" name="founders" value="{{.Founders}}"><br>

        <label for="category">Category:</label>
        <input type="text" name="category" value="{{.Category}}"><br>
        
        <label for="stage">Stage:</label>
        <input type="text" name="stage" value="{{.Stage}}"><br>

        <label for="website">Website URL:</label>
        <input type="text" name="website" value="{{.Website}}"><br>

        <label for="twitter">Twitter name:</label>
        <input type="text" name="twitter" value="{{.Twitter}}"><br>

        <label for="facebook">Facebook name:</label>
        <input type="text" name="facebook" value="{{.Facebook}}"><br>

        <label for="angellist">AngelList name:</label>
        <input type="text" name="angellist" value="{{.AngelList}}"><br>

        <label for="linkedin">LinkedIn name:</label>
        <input type="text" name="linkedin" value="{{.LinkedIn}}"><br>

        <label for="instagram">Instagram name:</label>
        <input type="text" name="instagram" value="{{.Instagram}}"><br>

        <label for="youtube">Youtube name:</label>
        <input type="text" name="youtube" value="{{.Youtube}}"><br>

        <label for="googleplus">GooglePlus name:</label>
        <input type="text" name="googleplus" value="{{.GooglePlus}}"><br>

        <label for="status">Status:</label>
        <select name="status">
            <option value="active" selected>Active</option>
            <option value="pending">Pending</option>
        </select>

        <BR><BR><BR>
        <input type="hidden" name="editname" value="{{.Name}}">
        <input type="submit" value="SAVE COMPANY">
    </form>
    
    <BR><BR><BR><BR>

    <form action="/admin/delete_company" method="post">
        <input type="hidden" name="deletename" value="{{.Name}}">
        <input style="background-color:red;" type="submit" value="DELETE COMPANY">
    </form>

<BR><BR><BR>
</div>
</body>
</html>