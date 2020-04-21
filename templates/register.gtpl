<html>
    <head>
    <title>Register</title>
    <link rel="stylesheet" href="/main.css" />
    <script async src="https://www.googletagmanager.com/gtag/js?id=UA-120199721-1"></script> <script>   window.dataLayer = window.dataLayer || [];   function gtag(){dataLayer.push(arguments);}   gtag('js', new Date());    gtag('config', 'UA-120199721-1'); </script></head>
    <body>
    <hr/>
    <form method="post" action="/home">
        <button type="submit">Home</button>
    </form>
    <form method="post" action="/logout">
        <button type="submit">Logout</button>
    </form>
    <button onclick="javascript:window.history.back();">Go Back</button>
    <hr>
    <div class="maincontainer">
        <h2>Create New User</h2>
        <form action="/register" method="post">
            <label for="name">Name:</label>
            <input type="text" name="name" required><br>

            <label for="email">School Email:</label>
            <input type="email" name="email" required><br>
            
            <label for="password">Password:</label>
            <input type="password" name="password" required><br>
            
            <label for="email2">Personal Email:</label>
            <input type="email" name="email2" required><br>

            <label for="year">Year:</label>
            <input type="number" name="year" required><br>

            <label for="role">Role:</label>
            <select name="role" required>
                <option value="advisor" selected>Advisor</option>
                <option value="client">Client</option>
                <option value="admin">Admin</option>
                <option value="inactive">Inactive</option>
                <option value="pending">Pending Verification</option>
            </select><br>

            <label for="status">Status:</label>
            <textarea name="status" rows="3" cols="60"></textarea><br>

            <label for="profile">Profile:</label>
            <textarea name="profile" rows="8" cols="60"></textarea><br><BR><BR>

            <input type="hidden" name="app" value="syndica">
            <input type="hidden" name="token" value="{{.}}">            
            <input type="submit" value="Create User">
        </form>
    </div>
    <BR><BR><BR>        
    </body>
</html>