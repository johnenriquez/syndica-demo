<!-- advisorsview_portfolio.html -->
<html lang="en-US">
  <head>
    <meta charset="UTF-8">
    <title>Portfolio | SYNDICA</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="dns-prefetch" href="//fonts.googleapis.com">

    <!-- Twitter Bootstrap -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">

    <!-- Flatly Bootstrap theme | https://bootswatch.com -->
    <link rel="stylesheet" href="/flatly.css">

    <!-- Icon set | https://useiconic.com/open/ -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/open-iconic/1.1.1/font/css/open-iconic-bootstrap.min.css">

  <script async src="https://www.googletagmanager.com/gtag/js?id=UA-120199721-1"></script> <script>   window.dataLayer = window.dataLayer || [];   function gtag(){dataLayer.push(arguments);}   gtag('js', new Date());    gtag('config', 'UA-120199721-1'); </script></head>
  <body class="bg-light">

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
                <a class="nav-link" href="/company">Companies</a>
              </li>
              <li class="nav-item">
                <a class="nav-link active" href="/portfolio">Portfolio</a>
              </li>
              <li class="nav-item dropdown">
                <a class="nav-link dropdown-toggle" href="/company_profile_view" id="navbarDropdown" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                  My Account
                </a>
                <div class="dropdown-menu dropdown-menu-right" aria-labelledby="navbarDropdown">
                  <a class="dropdown-item" href="/activity">Notifications</a>
                  <div class="dropdown-divider"></div>
                  <a class="dropdown-item" href="/messages">Messages</a>
                  <a class="dropdown-item" href="/profile">Settings</a>
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

<div class="bg-white py-5">
  <div class="container">

    <div class="d-flex text-center justify-content-center align-items-center">

      <!--<div class="px-3">
        <img class="img-thumbnail" style="max-width: 10em" src="http://math.utsa.edu/wp-content/uploads/2017/11/temporary-profile-placeholder-1.jpg" alt="Profile photo" />
      </div>-->
      <div class="px-3 lead">
        <h3 class="h2">{{.Advisor.Name}}</h3>
        <ul class="list-unstyled text-muted">
          <li>Master of Business Administration</li>
          <li>UCLA Anderson School of Management</li>

          <li>{{.Advisor.Title}} · {{.Advisor.Job}}</li>
          <!-- <li>{{.Advisor.Experience}}</li> -->

          <li>{{.Advisor.Profile | html}}</li>
        </ul>
      </div>
    </div>

  </div>
</div>

<div class="container">
  <div class="text-center">
    <h3 class="py-5 m-0 h4">Portfolio Companies</h3>
  </div>
  <div class="row">

    {{range .Companies}}
    <div class="companygriditem col-md-4" style="max-width: 30em;">
      <div class="card mb-4">
        <a href="/company?name={{.Name}}&team=1">
          <div class="card-img-top bg-white" style=" background-image: url('{{.Logo}}'); padding-top: 75%; width: 100%; background-size: contain; background-repeat: no-repeat; background-position: 50% 50%;">            
            <img class="d-none"  src="{{.Logo}}">
          </div>
        </a>
        <div class="card-body bg-white">
          <h5 class="card-title"><a class="text-primary" href="/company?name={{.Name}}&team=1">{{.Name}}</a></h5>
          <p class="description card-text mb-4 text-muted">{{.Short |html}}</p>
          <a class="btn btn-block btn-secondary" href="/company?name={{.Name}}&team=1">View Team Discussions</a>
        </div>
      </div><!-- .card -->
    </div><!-- .companygriditem -->
    {{end}}

  </div><!-- .row -->
</div><!-- .container -->

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
<!-- /advisorsview_portfolio.html -->
