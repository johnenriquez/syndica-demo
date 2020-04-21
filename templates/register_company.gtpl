<html>
    <head>
    <title>Register Startup</title>
    <link rel="stylesheet" href="/main.css" />
    <script async src="https://www.googletagmanager.com/gtag/js?id=UA-120199721-1"></script> <script>   window.dataLayer = window.dataLayer || [];   function gtag(){dataLayer.push(arguments);}   gtag('js', new Date());    gtag('config', 'UA-120199721-1'); </script></head>
    <body>
    <form method="post" action="/home">
        <button type="submit">Home</button>
    </form>
    <form method="post" action="/admn">
        <button type="submit">Admin</button>
    </form>
    <form method="post" action="/logout">
        <button type="submit">Logout</button>
    </form>
    <button onclick="javascript:window.history.back();">Go Back</button>

    <div class="maincontainer">

        <h2>Create New Startup</h2>
        <form action="/register/company" method="post">
            <label for="name">Company Name:</label>
            <input type="text" name="name" required><br>

            <label for="short">Short Description:</label>
            <input type="text" name="short" required><br>

            <label for="description">Long Description:</label>
            <textarea name="description" rows="6" cols="60"></textarea><br>

            <label for="logo">Logo Image URL:</label>
            <input type="text" name="logo" required><br>

            <label for="primarypic">Product Image URL:</label>
            <input type="text" name="primarypic" required><br>

            <label for="founders">Founders:</label>
            <input type="text" name="founders" required><br>

            <label for="category">Category:</label>
            <input type="text" name="category" required><br>

            <label for="stage">Stage:</label>
            <input type="text" name="stage" required><br>
           
            <label for="website">Website URL:</label>
            <input type="text" name="website" required><br>

            <label for="twitter">Twitter:</label>
            <input type="text" name="twitter"><br>

            <label for="facebook">Facebook:</label>
            <input type="text" name="facebook"><br>

            <label for="angellist">AngelList:</label>
            <input type="text" name="angellist"><br>

            <label for="linkedin">LinkedIn:</label>
            <input type="text" name="linkedin"><br>

            <label for="instagram">Instagram:</label>
            <input type="text" name="instagram"><br>

            <label for="youtube">Youtube:</label>
            <input type="text" name="youtube"><br>

            <label for="googleplus">GooglePlus:</label>
            <input type="text" name="googleplus"><br>

            <label for="contact">Client Contact Name:</label>
            <input type="text" name="contact" required><br>

            <label for="email">Client Email:</label>
            <input type="email" name="email" required><br>

            <BR><BR>
            <input type="hidden" name="app" value="syndica">
            <input type="hidden" name="token" value="{{.}}">            
            <input type="submit" value="Create Company">
        </form>

        <BR><BR><BR>
        </div>
    </body>
</html>