<!-- signup.gtpl -->
<!DOCTYPE html>
<html lang="en-US">
  <head>
    <meta charset="UTF-8">
    <title>Sign Up | SYNDICA</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="dns-prefetch" href="//fonts.googleapis.com">

    <!-- Twitter Bootstrap -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">

    <!-- Flatly Bootstrap theme | https://bootswatch.com -->
    <link rel="stylesheet" href="/flatly.css">

    <!-- Icon set | https://useiconic.com/open/ -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/open-iconic/1.1.1/font/css/open-iconic-bootstrap.min.css">

    <!-- CAPTCHA -->
    <script src='https://www.google.com/recaptcha/api.js'></script>

  <script async src="https://www.googletagmanager.com/gtag/js?id=UA-120199721-1"></script> <script>   window.dataLayer = window.dataLayer || [];   function gtag(){dataLayer.push(arguments);}   gtag('js', new Date());    gtag('config', 'UA-120199721-1'); </script></head>
  <body id="homepage" class="text-light" style="background-color: #000">

    <div class="maincontainer">

      <div class="headnav">
        <nav class="navbar navbar-expand-md navbar-dark" style="background-color: #000;">
          <a class="navbar-brand" href="/home">SYNDICA</a>
          <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
          </button>
          <div class="collapse navbar-collapse" id="navbarNav">
            <ul class="navbar-nav ml-auto">
              <li class="nav-item">
                <a class="nav-link" href="/login">LOGIN</a>
              </li>
            </ul>
          </div>
        </nav>
      </div><!-- .headnav -->

      <div class="content">

<!-- MAIN CONTENT -->

<div class="mx-auto" style="max-width: 40em;">
  <div class="h3 py-4 text-center">
    Create a User Account
  </div>
  <div>
    <div class="h4">UCLA Anderson MBAs</div>
    <p class="text-secondary">If you are an Anderson MBA, please provide your Anderson email address. (example@anderson.ucla.edu)<br>
    We will send you a confirmation email in order to authenticate your account.</p>
    <br>
    <div class="h4">Tech Startups</div>
    <p class="text-secondary">If you represent a startup company, please provide your company email address below.
We will immediately send you a follow up email to proceed with next steps.</p>
  </div>

  <form class="mt-4 bg-primary rounded p-3" action="/signup" method="post">

      <div class="form-group mb-4">
        <label for="email"><b>Email:</b></label>
        <input class="form-control bg-primary rounded-0" style="border-width: 0 0 1px; border-color: #999;" type="email" name="email" placeholder="example@domain.com" required>
      </div>

      <input type="hidden" name="app" value="syndica">

      <div class="g-recaptcha pb-4 mx-auto" data-sitekey="6Lc2JFEUAAAAAPCFS9ilr14QtbRi7EWVTn58VmD-"></div>

      <input class="btn btn-block btn-secondary" type="submit" value="SUBMIT">
  </form>

  <p class="text-secondary small my-3">
    *By using this site, you agree to our <a class="text-white" href="/terms.html" target="_blank">Terms</a> and
    confirm that you have read our <a class="text-white" href="/privacy.html" target="_blank">privacy policy</a>.
  </p>
</div>

<!-- /MAIN CONTENT -->

      </div><!-- .content -->

    </div><!-- .maincontainer -->

    <script src="/main.js?v=1.2.0"></script>

    <!-- Twitter Bootstrap block -->
    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js" integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
    <!-- /Twitter Bootstrap Block -->
  </body>
</html>
<!-- signup.gtpl -->
