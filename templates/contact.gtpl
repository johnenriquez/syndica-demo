<!DOCTYPE html>
<html lang="en-US">
<meta charset="UTF-8" />
<title>Syndica</title>
<meta name="viewport" content="width=device-width, initial-scale=1" />
<link rel='dns-prefetch' href='//fonts.googleapis.com' />
<link rel="stylesheet" href="/theme.css" />
<link rel="stylesheet" href="/main.css" />
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-120199721-1"></script> <script>   window.dataLayer = window.dataLayer || [];   function gtag(){dataLayer.push(arguments);}   gtag('js', new Date());    gtag('config', 'UA-120199721-1'); </script></head>
<body>
<div class="maincontainer">
<h3>Contact Us</h3>
<BR><BR>
<form method="post" action="/contact">
    <label for="name">Name:</label>
    <input type="text" name="name" required><br><BR>

    <label for="email">Email:</label>
    <input type="email" name="email" required><br><BR>
    
    <label for="comment">Message:</label>
    <textarea col=60 rows=6 required name="comment"></textarea>
    
    <input type="hidden" name="app" value="syndica">
    <BR>          <BR>
    <input type="submit" value="SEND">
</form>
</body>
</html>