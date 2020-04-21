<!-- dl/question_single.html -->
<html lang="en-US">
  <head>
    <meta charset="UTF-8">
    <title>Question | Syndica</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="dns-prefetch" href="//fonts.googleapis.com">

    <!-- Twitter Bootstrap -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">

    <!-- Flatly Bootstrap theme | https://bootswatch.com -->
    <link rel="stylesheet" href="/flatly.css">

    <!-- Icon set | https://useiconic.com/open/ -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/open-iconic/1.1.1/font/css/open-iconic-bootstrap.min.css">

  <script async src="https://www.googletagmanager.com/gtag/js?id=UA-120199721-1"></script> <script>   window.dataLayer = window.dataLayer || [];   function gtag(){dataLayer.push(arguments);}   gtag('js', new Date());    gtag('config', 'UA-120199721-1'); </script></head>
  <body class="bg-light" onload='LoadCompanyDetails("{{.C.Name}}",{{.Q}});'>

    <div class="maincontainer">

      <div class="headnav">
        <nav class="navbar navbar-expand-md navbar-dark bg-primary">
          <a class="navbar-brand" href="/home">SYNDICA</a>
          <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
          </button>
          <div class="collapse navbar-collapse" id="navbarNav">
            <ul class="navbar-nav ml-auto">
              <li class="nav-item">
                <a class="nav-link" href="/home">Home</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/team">Advisory Team</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/company_profile_view">Company Profile</a><!-- COMPANY PROFILE -->
              </li>
              <li class="nav-item dropdown">
                <a class="nav-link dropdown-toggle" href="/company_profile_view" id="navbarDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                  My Account
                </a>
                <div class="dropdown-menu dropdown-menu-right" aria-labelledby="navbarDropdown">
                  <a class="dropdown-item" href="/client_profile">My Profile</a>
                  <a class="dropdown-item" href="/activity">Notifications</a>
                  <div class="dropdown-divider"></div>
                  <a class="dropdown-item" href="/messages">Messages</a>
                  <div class="dropdown-divider"></div>
                  <a class="dropdown-item" href="/logout">Logout</a>
                </div>
              </li>
            </ul>
          </div>
        </nav>
      </div><!-- .headnav -->

      <div class="content">

<!-- MAIN CONTENT -->

<style>
/* TODO: Please move this to a better place! */
  #myteam ul {
    display: inline;
    list-style: none;
    padding: 0;
  }
  #myteam ul li {
    display: inline;
  }
  #myteam ul li + li:before {
    content: ", ";
  }
</style>

<div class="bg-white">
  <div class="container py-4">

    <div class="text-muted small my-3">Question asked by <a class="text-primary" href="/company?name={{.C.Name}}"><b>{{.C.Name}}</b></a></div>

    <div id="myquestion"></div><!-- #myquestion -->

    <div class="small text-muted py-2">
      <b>Reward:</b>
      <div class="d-inline" id="myreward"></div><!-- #myreward -->
    </div>

    <div class="p-3 mt-3 bg-light small">
      <div class="mb-2">
        <div class="h4 mb-1"><a class="text-primary" href="/company?name={{.C.Name}}">{{.C.Name}}</a></div>
        <div>{{.C.Short | html}}</div>
      </div>
      <div class="text-muted">
        <b>Founders:</b> {{.C.Founders}}<br>
        <b>Industry:</b> {{.C.Category}}<br>
        <b>Stage:</b> {{.C.Stage}}<br>
        <b>Advisors:</b>
        <div id="myteam"></div><!-- #myteam -->
      </div>
    </div><!-- .p-3 .bg-light -->

  </div><!-- .container -->
</div><!-- .bg-white -->

<div class="container py-4">
  <div class="card">
    <div class="card-header bg-secondary text-white text-center">
      Responses
    </div>
    <div class="card-body">

      <div id="responses"></div><!-- #responses -->
      
    </div><!-- .card-body -->
  </div><!-- .card -->
</div><!-- .container -->

<!-- /MAIN CONTENT -->

      </div><!-- .content -->

    </div><!-- .maincontainer -->

    <script src="/cookie.js"></script>
    <script src="/main.js?v=1.2.0"></script>

    <!-- Twitter Bootstrap block -->
    <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js" integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" crossorigin="anonymous"></script>
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
    <!-- /Twitter Bootstrap Block -->
  </body>
</html>
<!-- dl/question_single.html -->
