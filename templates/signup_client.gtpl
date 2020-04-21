<html>
    <head>
    <title>Register</title>
    <link rel="stylesheet" href="/main.css" />
    <script async src="https://www.googletagmanager.com/gtag/js?id=UA-120199721-1"></script> <script>   window.dataLayer = window.dataLayer || [];   function gtag(){dataLayer.push(arguments);}   gtag('js', new Date());    gtag('config', 'UA-120199721-1'); </script></head>
    <body>
    <div class="maincontainer">
        <h2>Submit Your Startup Information</h2>
        <form action="/verify" method="post">

            <label for="companyname">Company Name:</label>
            <input type="text" name="companyname" required><br>

            <label for="name">Your Name:</label>
            <input type="text" name="name" required><br>

            <label for="email">Your Business Email:</label>
            <input type="email" name="email" value="{{.Email}}" readonly><br>
            
            <label for="website">Website URL:</label>
            <input type="text" name="website"><br>

            <label for="short">Short Description:</label>
            <input type="text" name="short" required><br>

            <label for="description">Long Description:</label>
            <textarea name="description" rows="6" cols="60"></textarea><br>

            <BR><BR>
            <input type="hidden" name="hash" value="{{.Hash}}"><br>
            <input type="hidden" name="app" value="syndica">
            <input type="submit" value="Submit Request">
        </form>
    </div>
    <BR><BR><BR>       
    </body>
</html>