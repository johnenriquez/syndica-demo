<html>
    <head>
    <title>Register</title>
    <link rel="stylesheet" href="/main.css" />
    <script async src="https://www.googletagmanager.com/gtag/js?id=UA-120199721-1"></script> <script>   window.dataLayer = window.dataLayer || [];   function gtag(){dataLayer.push(arguments);}   gtag('js', new Date());    gtag('config', 'UA-120199721-1'); </script></head>
    <body>
    <div class="maincontainer">
        <h2>Create New User Account</h2>
        <form action="/verify" method="post">
            <label for="name">Your Full Name:</label>
            <input type="text" name="name" required><br>

            <label for="email">School Email:</label>
            <input type="email" name="email" value="{{.Email}}" readonly><br>
            
            <label for="password">Password:</label>
            <input type="password" name="password" required><br>
            <input type="hidden" name="hash" value="{{.Hash}}"><br>

            <label for="email2">Personal Email:</label>
            <input type="email" name="email2"><br>

            <label for="year">Year Graduated:</label>
            <input type="number" name="year"><br>
            <br>
            <input type="hidden" name="app" value="syndica">
            <input type="submit" value="Create Account">
        </form>
    </div>
    <BR><BR><BR>       
    </body>
</html>