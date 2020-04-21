<!DOCTYPE html>
<html lang="en-US">
  <head>
    <meta charset="UTF-8">
    <title>HOME | Syndica</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="dns-prefetch" href="//fonts.googleapis.com">

    <!-- Twitter Bootstrap -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">

    <!-- Flatly Bootstrap theme | https://bootswatch.com -->
    <link rel="stylesheet" href="/flatly.css">

    <!-- Icon set | https://useiconic.com/open/ -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/open-iconic/1.1.1/font/css/open-iconic-bootstrap.min.css">

  <script async src="https://www.googletagmanager.com/gtag/js?id=UA-120199721-1"></script> <script>   window.dataLayer = window.dataLayer || [];   function gtag(){dataLayer.push(arguments);}   gtag('js', new Date());    gtag('config', 'UA-120199721-1'); </script></head>
  <body class="bg-light" onload="LoadThreads();">

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
                <a class="nav-link active" href="/home">Home</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/company">Companies</a>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="/portfolio">Portfolio</a>
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

        <div class="container pt-4">

          <div class="row">

            <div class="primary mb-0 companygriditem col-lg-8">
              <div class="d-flex py-2 justify-content-between align-items-center">
              <h5 class="m-0">Forum</h5>

              <form class="m-0" action="/newthread" method="get">
                <input class="btn btn-secondary" type="submit" name="newthread" value="+ New Thread">
              </form>
            </div>

            <div class="mb-3" id="threadsdiv"></div>
          </div><!-- .primary -->

            <div class="secondary pt-2 col-lg-4">

              <h5>Advisory Startups</h5>
              <div class="card mb-4">
                <div class="card-img-top bg-white">
                  <div class="media">
                    <a href="/company?name={{.Today.Name}}">
                      <img src="{{.Today.Logo}}" class="img-thumbnail m-2" style="height: auto; max-width: 7em">
                    </a>
                    <div class="media-body align-self-center px-2">
                      <h6 class="card-title mb-1"><a class="text-primary" href="/company?name={{.Today.Name}}">{{.Today.Name}}</a></h6>
                      <p class="description card-text mb-1 text-muted">{{.Today.Short |html}}</p>
                    </div><!-- .media-body -->
                  </div><!-- .media -->
                </div><!-- .card-img-top -->
                <div class="card-body bg-light">
                  <a class="btn btn-block btn-secondary" href="/company?name={{.Today.Name}}">Provide Feedback</a>
                </div><!-- .card-body -->
              </div><!-- .card -->

              <BR>

              <h5>Previous Startups</h5>
                {{range .Previous}}
                  <div class="card mb-4">
                    <div class="card-img-top bg-white">
                      <div class="media">
                        <a href="/company?name={{.Name}}">
                          <img src="{{.Logo}}" class="img-thumbnail m-2" style="height: auto; max-width: 4em">
                        </a>
                        <div class="media-body align-self-center px-2 small">
                          <h6 class="card-title mb-1"><a class="text-primary" href="/company?name={{.Name}}">{{.Name}}</a></h6>
                          <p class="description card-text mb-1 text-muted">{{.Short |html}}</p>
                        </div>
                      </div><!-- .media -->
                    </div><!-- .card-img-top -->
                  </div><!-- .card -->
                {{end}}

            </div> <!-- .secondary -->

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
