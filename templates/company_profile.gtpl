<!-- dl/edit_company_profile.gtpl -->
<!DOCTYPE html>
<html lang="en-US">
  <head>
    <meta charset="UTF-8">
    <title>Edit Company Profile | Syndica</title>
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

<div class="container">
  <div class="card my-4">
    <div class="card-header bg-primary text-white">
      <div class="h4 mb-0">Edit Company Profile</div>
    </div>
    <div class="card-body">
      <form action="/company_profile" method="post">

        <div class="form-group">
          <label for="name"><b>Company Name:</b></label>
          <input class="form-control" type="text" name="name" value="{{.Name}}" readonly="readonly">
        </div>

        <div class="form-group">
          <label for="contact"><b>Client Contact Name:</b></label>
          <input class="form-control" type="text" name="contact" value="{{.Contact}}" required>
        </div>

        <div class="form-group">
          <label for="email"><b>Client Email:</b></label>
          <input class="form-control" type="email" name="email" value="{{.Email}}" required>
        </div>

        <div class="form-group">
          <label for="short"><b>Company Short Description:</b></label>
          <input class="form-control" type="text" name="short" required value="{{.Short}}">
        </div>

        <div class="form-group">
          <label for="description"><b>Company Long Description:</b></label>
          <textarea class="form-control" name="description" rows="6">{{.Description}}</textarea>
        </div>

        <div class="form-group">
          <label for="logo"><b>Logo Image URL:</b></label>
          <input class="form-control" type="text" name="logo" value="{{.Logo}}" required>
        </div>

        <div class="form-group">
          <label for="primarypic"><b>Product Image URL:</b></label>
          <input class="form-control" type="text" name="primarypic" value="{{.PrimaryPic}}" required>
        </div>

        <div class="form-group">
          <label for="founders"><b>Founders:</b></label>
          <input class="form-control" type="text" name="founders" value="{{.Founders}}" required>
        </div>

        <div class="form-group">
          <label for="category"><b>Category:</b></label>
          <input class="form-control" type="text" name="category" value="{{.Category}}" required>
        </div>

        <div class="form-group">
          <label for="stage"><b>Stage:</b></label>
          <input class="form-control" type="text" name="stage" value="{{.Stage}}" required>
        </div>

        <div class="form-group">
          <label for="website"><b>Website URL:</b></label>
          <input class="form-control" type="text" name="website" value="{{.Website}}" required>
        </div>

        <div class="form-group">
          <label for="twitter"><b>Twitter:</b></label>
          <input class="form-control" type="text" name="twitter" value="{{.Twitter}}">
        </div>

        <div class="form-group">
          <label for="facebook"><b>Facebook:</b></label>
          <input class="form-control" type="text" name="facebook" value="{{.Facebook}}">
        </div>

        <div class="form-group">
          <label for="angellist"><b>AngelList:</b></label>
          <input class="form-control" type="text" name="angellist" value="{{.AngelList}}">
        </div>

        <div class="form-group">
          <label for="linkedin"><b>LinkedIn:</b></label>
          <input class="form-control" type="text" name="linkedin" value="{{.LinkedIn}}">
        </div>

        <div class="form-group">
          <label for="instagram"><b>Instagram:</b></label>
          <input class="form-control" type="text" name="instagram" value="{{.Instagram}}">
        </div>

        <div class="form-group">
          <label for="youtube"><b>Youtube:</b></label>
          <input class="form-control" type="text" name="youtube" value="{{.Youtube}}">
        </div>

        <div class="form-group">
          <label for="googleplus"><b>GooglePlus:</b></label>
          <input class="form-control" type="text" name="googleplus" value="{{.GooglePlus}}">
        </div>

        <input type="hidden" name="app" value="syndica">
        <input class="btn btn-block btn-primary" type="submit" value="UPDATE">
      </form>

    </div>
  </div>
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
<!-- dl/edit_company_profile.gtpl -->
