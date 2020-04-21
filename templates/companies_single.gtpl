<!-- company_vanity_profile.html -->
<!DOCTYPE html>
<html lang="en-US">
  <head>
    <meta charset="UTF-8">
    <title>Company Profile | Syndica</title>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="dns-prefetch" href="//fonts.googleapis.com">

    <!-- Twitter Bootstrap -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">

    <!-- Flatly Bootstrap theme | https://bootswatch.com -->
    <link rel="stylesheet" href="/flatly.css">

    <!-- Icon set | https://useiconic.com/open/ -->
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/open-iconic/1.1.1/font/css/open-iconic-bootstrap.min.css">

  <script async src="https://www.googletagmanager.com/gtag/js?id=UA-120199721-1"></script> <script>   window.dataLayer = window.dataLayer || [];   function gtag(){dataLayer.push(arguments);}   gtag('js', new Date());    gtag('config', 'UA-120199721-1'); </script></head>
  <body class="bg-light" onload='LoadCompanyProfile("{{.C.Name}}");LoadCompanyQuestions("{{.C.Name}}");LoadCompanyComments("{{.C.Name}}")'>

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
<div class="companyrow1">

  <div class="jumbotron jumbotron-fluid bg-light mb-0" style="background-color: #95a5a6;padding-top:20px;padding-bottom:20px">
    <div class="container">
      <div class="companylogo">
        <img class="bg-light p-1 border" style="max-height:65px; height:auto;" src="{{.C.Logo}}">
      </div>
      <h1 class="display-3">{{.C.Name}}</h1>
      <p>{{.C.Short | html}}</p>
    </div>
  </div>

  <div class="bg-white">
    <div class="container">
      <div class="row py-4">

        <div class="col-md-8">
          <img class="img-fluid mb-5" src="{{.C.PrimaryPic}}">

          <div class="mt-1" style="word-wrap: break-word; overflow-wrap: break-word; white-space:pre-wrap;">{{.C.Description | html}}</div>

        </div><!-- .col-8 -->

        <div class="col-md-4 pt-4 pt-md-0">
          <u>Company Info</u>
          <ul>
            <li>Founders: {{.C.Founders}}</li>
            <li>Industry: {{.C.Category}}</li>
            <li>Stage: {{.C.Stage}}</li>
          </ul>
          <u>More Info</u>
          <ul>
            {{if .C.Website}}<li>Website: <a href="{{.C.Website}}" target="_blank">{{.C.Website}}</a></li>{{end}}
            {{if .C.Twitter}}<li>Twitter: <a href="https://twitter.com/{{.C.Twitter}}" target="_blank">{{.C.Twitter}}</a></li>{{end}}
            {{if .C.Facebook}}<li>Facebook: <a href="https://www.facebook.com/{{.C.Facebook}}" target="_blank">{{.C.Facebook}}</a></li>{{end}}
            {{if .C.AngelList}}<li>AngelList: <a href="https://angel.co/{{.C.AngelList}}" target="_blank">{{.C.AngelList}}</a></li>{{end}}
            {{if .C.LinkedIn}}<li>LinkedIn: <a href="https://www.linkedin.com/company/{{.C.LinkedIn}}" target="_blank">{{.C.LinkedIn}}</a></li>{{end}}
            {{if .C.Instagram}}<li>Instagram: <a href="https://www.instagram.com/{{.C.Instagram}}" target="_blank">{{.C.Instagram}}</a></li>{{end}}
            {{if .C.Youtube}}<li>Youtube: <a href="https://www.youtube.com/{{.C.Youtube}}" target="_blank">{{.C.Youtube}}</a></li>{{end}}
            {{if .C.GooglePlus}}<li>Google+: <a href="https://plus.google.com/+{{.C.GooglePlus}}" target="_blank">{{.C.GooglePlus}}</a></li>{{end}}
          </ul>
          <!--<u>Advisory Team</u>
          <br>-->
          <div id="myteam" style="visibility:hidden;"></div>
          <form method="post" action="/jointeam">
            {{if .T}}<a class="btn btn-block btn-outline-primary" href="/company?name={{.C.Name}}&team=1">View Team Discussions</a>
            {{else}}<input class="btn btn-block btn-outline-primary" type="submit" value="Add to Portfolio"><input type="hidden" name="company" value="{{.C.Name}}">
            {{end}}
          </form>
        </div><!-- .col-4 -->

      </div><!-- .row -->


    </div><!-- .container -->

  </div><!-- .bg-white -->

</div><!-- .companyrow1 -->

<!-- /MAIN CONTENT -->

      </div><!-- .content -->

      <div class="container bg-light">

        <div class="company-open-questions mt-5">
          <div class="card bg-white">
            <div class="card-header bg-light">
              <div class="h5 mb-0">Open questions by {{.C.Name}}:</div>
            </div>
            <div class="card-body p-0">
              <div id="questions"></div><!-- #questions -->
            </div><!-- .card-body -->
          </div><!-- .card -->
        </div><!-- company-open-questions -->

        <div class="company-open-comments mt-5 mb-3">
          <div class="card bg-white">

            <div class="card-header bg-primary">
              <div class="h5 mb-0 text-white text-center">What are your thoughts on {{.C.Name}}?</div>
            </div><!-- .card-header -->

            <div class="your-response mt-3" style="margin-top:0px !important;">
              <form class="bg-light rounded p-3 clearfix" action="/comment" method="post">
                <div class="small text-muted">
                  <p>
                    What are your thoughts on <strong>{{.C.Name}}</strong>? <br>Provide your comments and feedback in this section - all responses are welcome and appreciated.
                  </p>
                </div>
                <div class="form-group mb-2">
                  <textarea class="form-control" rows="3" required name="comment"></textarea>
                </div>
                <input id="companyid" type="hidden" name="company" value="{{.C.Name}}">
                <input type="hidden" name="reply_id" value="0">
                <input class="submitresponse btn btn-secondary btn-block" type="submit" name="submitresponse" value="SUBMIT">
              </form>
            </div><!-- .your-response -->

            <div class="card-body p-4">
              <div id="comments" style="min-height:500px;"></div><!-- #comments -->
            </div><!-- .card-body -->
          </div><!-- .card -->
        </div><!-- company-open-comments -->

      </div><!-- .container -->

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
<!-- /company_vanity_profile.html -->
