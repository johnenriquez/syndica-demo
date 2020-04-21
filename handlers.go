package main

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nu7hatch/gouuid"
)

const (
	HTTPPORT  = ":80"
	HTTPSPORT = ":443"
)

func handleRequests() {
	go http.ListenAndServe(HTTPPORT, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
	}))

	r := mux.NewRouter()
	// r.HandleFunc("/changepassword", changePasswordHandler)
	// r.HandleFunc("/resetpassword", resetPasswordHandler)
	r.HandleFunc("/activity", activityHandler)
	r.HandleFunc("/admin", adminHandler)
	r.HandleFunc("/admin/lock", adminLockHandler)
	r.HandleFunc("/admin/unlock", adminUnlockHandler)
	r.HandleFunc("/admin/user", adminUserHandler)
	r.HandleFunc("/admin/advisor", adminAdvisorHandler)
	r.HandleFunc("/admin/company", adminCompanyHandler)
	r.HandleFunc("/admin/delete_company", adminDeleteCompanyHandler)
	r.HandleFunc("/admin/send_message", adminSendMessageHandler)
	r.HandleFunc("/admin/messages", adminMessagesHandler)
	r.HandleFunc("/advisor", advisorHandler)
	r.HandleFunc("/client_profile", clientProfileHandler)
	r.HandleFunc("/company", companyHandler)
	// r.HandleFunc("/company_list", companyListHandler)
	r.HandleFunc("/company_profile", companyProfileHandler)
	r.HandleFunc("/company_profile_view", companyProfileViewHandler)
	r.HandleFunc("/companyquestions", companyQuestionsHandler)
	r.HandleFunc("/companyteamquestions", companyTeamQuestionsHandler)
	r.HandleFunc("/comment", commentHandler)
	r.HandleFunc("/comment_edit", commentEditHandler)
	r.HandleFunc("/comment_delete", commentDeleteHandler)
	r.HandleFunc("/commentlike", commentLikeHandler)
	r.HandleFunc("/contact", contactHandler)
	r.HandleFunc("/discussions", discussionsHandler)
	r.HandleFunc("/home", homeHandler)
	r.HandleFunc("/jointeam", joinTeamHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", logoutHandler)
	r.HandleFunc("/like", likeHandler)
	r.HandleFunc("/loadcompanycomments", loadCompanyCommentsHandler)
	r.HandleFunc("/loadcompanyquestions", loadCompanyQuestionsHandler)
	r.HandleFunc("/loadthreads", loadThreadsHandler)
	r.HandleFunc("/loadthreadcomments", loadThreadCommentsHandler)
	r.HandleFunc("/myresponses", myResponsesHandler)
	r.HandleFunc("/messages", myMessagesHandler)
	r.HandleFunc("/myreward", myRewardHandler)
	r.HandleFunc("/myteam", myTeamHandler)
	r.HandleFunc("/myquestion", myQuestionHandler)
	r.HandleFunc("/newquestion", newQuestionHandler)
	r.HandleFunc("/newthread", newThreadHandler)
	r.HandleFunc("/newmessage", newMessageHandler)
	r.HandleFunc("/profile", profileHandler)
	r.HandleFunc("/portfolio", portfolioHandler)
	r.HandleFunc("/question_edit", questionEditHandler)
	r.HandleFunc("/question_delete", questionDeleteHandler)
	r.HandleFunc("/register/company", registerCompanyHandler)
	r.HandleFunc("/register", registerHandler)
	r.HandleFunc("/response", responseHandler)
	r.HandleFunc("/response_edit", responseEditHandler)
	r.HandleFunc("/response_delete", responseDeleteHandler)
	r.HandleFunc("/send_message", sendMessageHandler)
	r.HandleFunc("/signup", signupHandler)
	r.HandleFunc("/team", teamHandler)
	r.HandleFunc("/thread", threadHandler)
	r.HandleFunc("/threadreply", threadReplyHandler)
	r.HandleFunc("/threadreply_edit", threadReplyEditHandler)
	r.HandleFunc("/threadreply_delete", threadReplyDeleteHandler)
	r.HandleFunc("/threadcommentreply", threadCommentReplyHandler)
	r.HandleFunc("/verify", verifyUserHandler)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("app")))
	log.Fatal(http.ListenAndServeTLS(HTTPSPORT, "cert.pem", "key.pem", r))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		errorString := ""
		e := r.URL.Query().Get("error")
		if e != "" {
			errorLevel, err := strconv.Atoi(e)
			if err != nil {
				errorLevel = 3
			}
			switch errorLevel {
			case 1:
				errorString = "Not authorized"
			case 2:
				errorString = "Incorrect login credentials or locked user"
			case 3:
				errorString = "Unknown error"
			default:
				errorString = "Unexpected error"
			}
		}
		token := GetToken()
		redirectURL := r.URL.Query().Get("r")
		type loginTpl struct {
			E string
			T string
			R string
		}
		t, _ := template.ParseFiles("templates/login.gtpl")
		t.Execute(w, loginTpl{E: errorString, T: token, R: redirectURL})
		return
	}
	// POST
	r.ParseForm()
	if r.Form.Get("app") != "syndica" {
		http.Redirect(w, r, "/login?error=1", 302)
		return
	}
	token := r.Form.Get("token")
	if token == "" {
		http.Redirect(w, r, "/login?error=1", 302)
		return
	}
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	if email == "" || password == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	ip := r.RemoteAddr
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		http.Redirect(w, r, "/login?error=9", 302)
		return
	}
	redirectURL := r.URL.Query().Get("r")
	if !CanLogin(email, password, ip) {
		log.Println("Failed login: ", email, password, ip)
		// fmt.Fprintln(w, "Unable to log in.")
		if redirectURL != "" {
			http.Redirect(w, r, "/login?error=2&r="+redirectURL, 302)
			return
		}
		http.Redirect(w, r, "/login?error=2", 302)
		return
	}
	nonce := GetRandomStr(32)
	SetSession(w, email, nonce)
	SetUserNonce(email, nonce)
	if redirectURL != "" {
		d, err := base64.URLEncoding.DecodeString(redirectURL)
		if err == nil {
			http.Redirect(w, r, string(d), 302)
			return
		}
	}
	http.Redirect(w, r, "/home", 302)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	ClearSession(w)
	http.Redirect(w, r, "/", 302)
}

func adminLockHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	log.Println("admin lock attempt by: ", email)
	if !IsAdmin(email) {
		http.Redirect(w, r, "/login", 302)
		return
	}
	e := r.URL.Query().Get("email")
	LockUserAdmin(e)
	fmt.Fprintf(w, "user %s is now locked.", e)
}

func adminUnlockHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	log.Println("admin unlock attempt by: ", email)
	if !IsAdmin(email) {
		http.Redirect(w, r, "/login", 302)
		return
	}
	e := r.URL.Query().Get("email")
	UnlockUserAdmin(e)
	fmt.Fprintf(w, "user %s is now unlocked.", e)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	user := GetUser(email)
	var t *template.Template
	switch user.Role {
	case "advisor":
		t, _ = template.ParseFiles("templates/home.gtpl")
		sortedCompanies := GetCompaniesSorted()
		t.Execute(w, struct {
			Today    Company
			Previous []Company
		}{
			Today:    sortedCompanies[0],
			Previous: sortedCompanies[1:],
		})
	case "client":
		// http.Redirect(w, r, "/company_profile_view", 302)
		// return
		company := GetMyCompany(email)
		t, _ = template.ParseFiles("templates/home_client.gtpl")
		t.Execute(w, company)
	case "admin":
		t, _ = template.ParseFiles("templates/home_admin.gtpl")
		t.Execute(w, user.Name)
	default:
		t, _ = template.ParseFiles("templates/home_inactive.gtpl")
		t.Execute(w, user.Name)
	}
}

func activityHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	var t *template.Template
	if IsClient(email) {
		t, _ = template.ParseFiles("templates/activity_client.gtpl")
	} else {
		t, _ = template.ParseFiles("templates/activity.gtpl")
	}
	t.Execute(w, GetActivitiesForUser(email))
}

func discussionsHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if !IsAdvisor(email) {
		http.Redirect(w, r, "/home", 302)
		return
	}
	t, _ := template.ParseFiles("templates/discussion.gtpl")
	t.Execute(w, GetDiscussions())
}

func loadThreadsHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if !IsAdvisor(email) {
		http.Redirect(w, r, "/home", 302)
		return
	}
	t, _ := template.ParseFiles("templates/loadthreads.gtpl")
	t.Execute(w, GetTopThreads(100))
}

func myMessagesHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	var t *template.Template
	if IsClient(email) {
		t, _ = template.ParseFiles("templates/mymessages_client.gtpl")
	} else {
		t, _ = template.ParseFiles("templates/mymessages.gtpl")
	}
	t.Execute(w, GetMessages(email))
}

func adminMessagesHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if !IsAdmin(email) {
		http.Redirect(w, r, "/login", 302)
		return
	}
	t, err := template.ParseFiles("templates/adminmessages.gtpl")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, GetMessages(email))
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if r.Method != "POST" {
		user := GetUser(email)
		t, _ := template.ParseFiles("templates/profile.gtpl")
		t.Execute(w, user)
		return
	}
	r.ParseForm()
	year, _ := strconv.Atoi(r.Form.Get("year"))
	u := User{
		Name:       r.Form.Get("name"),
		Email:      email,
		Password:   r.Form.Get("password"),
		Year:       year,
		Profile:    r.Form.Get("profile"),
		Job:        r.Form.Get("job"),
		Title:      r.Form.Get("title"),
		Experience: r.Form.Get("experience"),
	}
	err := EditProfile(u)
	if err != nil {
		fmt.Fprintln(w, "Problem updating profile: ", err.Error())
		log.Println("Problem updating profile: ", err.Error())
		return
	}
	log.Println("successful edit profile for: ", email)
	http.Redirect(w, r, "/portfolio", 302)
}

func clientProfileHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if r.Method != "POST" {
		user := GetUser(email)
		t, _ := template.ParseFiles("templates/client_profile.gtpl")
		t.Execute(w, user)
		return
	}
	r.ParseForm()
	u := User{
		Name:     r.Form.Get("name"),
		Email:    email,
		Password: r.Form.Get("password"),
	}
	err := EditClientProfile(u)
	if err != nil {
		fmt.Fprintln(w, "Problem updating client profile: ", err.Error())
		log.Println("Problem updating client profile: ", err.Error())
		return
	}
	log.Println("successful edit client profile for: ", email)
	http.Redirect(w, r, "/home", 302)
}

func companyProfileHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if r.Method != "POST" {
		company := GetMyCompany(email)
		t, _ := template.ParseFiles("templates/company_profile.gtpl")
		t.Execute(w, company)
		return
	}
	r.ParseForm()
	if r.Form.Get("app") != "syndica" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	name := GetCompanyName(email)
	if name == "" {
		fmt.Fprintln(w, "Company Name is missing.")
		return
	}
	newCompany := Company{
		Name:        name,
		Contact:     r.Form.Get("contact"),
		Email:       r.Form.Get("email"),
		Website:     r.Form.Get("website"),
		Logo:        r.Form.Get("logo"),
		PrimaryPic:  r.Form.Get("primarypic"),
		Short:       r.Form.Get("short"),
		Description: r.Form.Get("description"),
		Founders:    r.Form.Get("founders"),
		Category:    r.Form.Get("category"),
		Stage:       r.Form.Get("stage"),
		Twitter:     r.Form.Get("twitter"),
		Facebook:    r.Form.Get("facebook"),
		AngelList:   r.Form.Get("angellist"),
		LinkedIn:    r.Form.Get("linkedin"),
		Instagram:   r.Form.Get("instagram"),
		Youtube:     r.Form.Get("youtube"),
		GooglePlus:  r.Form.Get("googleplus"),
	}
	err := EditCompany(name, newCompany)
	if err != nil {
		fmt.Fprintln(w, "Problem creating new company: ", err.Error())
		log.Println("Problem creating new company: ", err.Error())
		return
	}
	log.Printf("successful updated company %s => %s\n", email, name)
	http.Redirect(w, r, "/company_profile_view", 302)
}

func companyProfileViewHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if !IsClient(email) {
		http.Redirect(w, r, "/home", 302)
		return
	}
	thisCompany := GetCompany(GetCompanyName(email))
	if thisCompany.Name == "" {
		http.Redirect(w, r, "/home", 302)
		return
	}
	t, _ := template.ParseFiles("templates/company_single.gtpl")
	// thisCompany.Description = strings.Replace(template.HTMLEscapeString(thisCompany.Description), "\n", "<br>", -1)
	t.Execute(w, thisCompany)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	// log.Println("Admin page attempted by: ", email)
	if !IsAdmin(email) {
		http.Redirect(w, r, "/login", 302)
		return
	}
	// log.Println("Admin page accessed by: ", email)
	t, _ := template.ParseFiles("templates/admin.gtpl")
	t.Execute(w, nil)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		t, _ := template.ParseFiles("templates/signup.gtpl")
		t.Execute(w, nil)
		return
	}
	r.ParseForm()
	if r.Form.Get("app") != "syndica" {
		return
	}
	email := r.Form.Get("email")
	if email == "" {
		return
	}
	ip := r.RemoteAddr
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		return
	}
	captcha := r.Form.Get("g-recaptcha-response")
	if !IsHuman(captcha, ip) {
		fmt.Fprintln(w, "Please go back and click on the box to verify that you are not a robot.")
		return
	}
	if IsEmailExists(email) {
		fmt.Fprintf(w, "The email %s is already being used. Please go back and try again.", email)
		return
	}
	newUUID, _ := uuid.NewV4()
	hash := newUUID.String()
	newUser := User{
		Email: email,
		Hash:  hash,
	}
	err := CreateUserInitial(newUser)
	if err != nil {
		fmt.Fprintln(w, "Problem signing up new user: ", err.Error())
		log.Println("Problem signing up new user: ", err.Error())
		return
	}
	if IsMBA(email) {
		go SendWelcomeEmailMBA(email, hash)
	} else {
		go SendWelcomeEmailClient(email, hash)
	}
	t, _ := template.ParseFiles("templates/signup2.gtpl")
	t.Execute(w, email)
	return
}

func verifyUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		email := r.URL.Query().Get("e")
		if email == "" {
			return
		}
		hash := r.URL.Query().Get("q")
		if hash == "" {
			return
		}
		if !ValidInitialUser(email, hash) {
			return
		}
		var t *template.Template
		if IsMBA(email) {
			t, _ = template.ParseFiles("templates/signup_mba.gtpl")
		} else {
			t, _ = template.ParseFiles("templates/signup_client.gtpl")
		}
		t.Execute(w, struct {
			Email string
			Hash  string
		}{
			Email: email,
			Hash:  hash,
		})
		return
	}
	// create initial user
	r.ParseForm()
	if r.Form.Get("app") != "syndica" {
		return
	}
	email := r.Form.Get("email")
	if email == "" {
		return
	}
	ip := r.RemoteAddr
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		return
	}
	hash := r.Form.Get("hash")
	if hash == "" {
		return
	}
	if !ValidInitialUser(email, hash) {
		return
	}
	name := r.Form.Get("name")
	if name == "" {
		return
	}

	if IsMBA(email) {
		year, _ := strconv.Atoi(r.Form.Get("year"))
		UpdateUser(User{
			Name:     name,
			Email:    email,
			Email2:   r.Form.Get("email2"),
			Password: r.Form.Get("password"),
			Year:     year,
			Status:   "active",
		})
	}
	SendPeterEmail("New User Created", fmt.Sprintf("Name: %s <BR>Email: %s", name, email))

	if IsClient(email) {
		companyname := r.Form.Get("companyname")
		website := r.Form.Get("website")
		CreateCompany(Company{
			Name:        r.Form.Get("companyname"),
			Contact:     r.Form.Get("name"),
			Email:       email,
			Website:     website,
			Short:       r.Form.Get("short"),
			Description: r.Form.Get("description"),
			// Logo:        r.Form.Get("logo"),
			// PrimaryPic:  r.Form.Get("primarypic"),
			// Founders:    r.Form.Get("founders"),
			// Category:    r.Form.Get("category"),
			// Stage:       r.Form.Get("stage"),
			// Twitter:     r.Form.Get("twitter"),
			// Facebook:    r.Form.Get("facebook"),
			// AngelList:   r.Form.Get("angellist"),
			// LinkedIn:    r.Form.Get("linkedin"),
			// Instagram:   r.Form.Get("instagram"),
			// Youtube:     r.Form.Get("youtube"),
			// GooglePlus:  r.Form.Get("googleplus"),
			Status: "pending",
		})
		SendPeterEmail("New Company Created", fmt.Sprintf("Name: %s <BR>Website: %s", companyname, website))
	}
	http.Redirect(w, r, "/login", 302)
	return
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	log.Println("Register page attempted by: ", email)
	if !IsAdmin(email) {
		http.Redirect(w, r, "/login", 302)
		return
	}
	log.Println("Register page accessed by: ", email)
	if r.Method != "POST" {
		t, _ := template.ParseFiles("templates/register.gtpl")
		t.Execute(w, GetToken())
		return
	}
	// create new user
	r.ParseForm()
	if r.Form.Get("app") != "syndica" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	token := r.Form.Get("token")
	if token == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	newEmail := r.Form.Get("email")
	if newEmail == "" {
		fmt.Fprintln(w, "Email is missing.")
		return
	}
	if IsEmailExists(newEmail) {
		fmt.Fprintln(w, "Email already exists.")
		return
	}
	year, _ := strconv.Atoi(r.Form.Get("year"))
	newUser := User{
		Name:     r.Form.Get("name"),
		Email:    newEmail,
		Email2:   r.Form.Get("email2"),
		Password: r.Form.Get("password"),
		Year:     year,
		Role:     r.Form.Get("role"),
		Status:   r.Form.Get("status"),
		Profile:  r.Form.Get("profile"),
	}
	err := CreateUser(newUser)
	if err != nil {
		fmt.Fprintln(w, "Problem creating new user: ", err.Error())
		log.Println("Problem creating new user: ", err.Error())
		return
	}
	log.Printf("successful create new user %s => %s\n", email, newEmail)
	http.Redirect(w, r, "/admin", 302)
}

func registerCompanyHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	//log.Println("Register company page attempted by: ", email)
	if !IsAdmin(email) {
		http.Redirect(w, r, "/login", 302)
		return
	}
	//log.Println("Register company page accessed by: ", email)
	if r.Method != "POST" {
		t, _ := template.ParseFiles("templates/register_company.gtpl")
		t.Execute(w, GetToken())
		return
	}
	// create new company
	r.ParseForm()
	if r.Form.Get("app") != "syndica" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	token := r.Form.Get("token")
	if token == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	newName := r.Form.Get("name")
	if newName == "" {
		fmt.Fprintln(w, "Company Name is missing.")
		return
	}
	if IsCompanyExists(newName) {
		fmt.Fprintln(w, "Company already exists.")
		return
	}
	newCompany := Company{
		Name:        r.Form.Get("name"),
		Contact:     r.Form.Get("contact"),
		Email:       r.Form.Get("email"),
		Website:     r.Form.Get("website"),
		Logo:        r.Form.Get("logo"),
		PrimaryPic:  r.Form.Get("primarypic"),
		Short:       r.Form.Get("short"),
		Description: r.Form.Get("description"),
		Founders:    r.Form.Get("founders"),
		Category:    r.Form.Get("category"),
		Stage:       r.Form.Get("stage"),
		Twitter:     r.Form.Get("twitter"),
		Facebook:    r.Form.Get("facebook"),
		AngelList:   r.Form.Get("angellist"),
		LinkedIn:    r.Form.Get("linkedin"),
		Instagram:   r.Form.Get("instagram"),
		Youtube:     r.Form.Get("youtube"),
		GooglePlus:  r.Form.Get("googleplus"),
		Status:      "pending",
	}
	err := CreateCompany(newCompany)
	if err != nil {
		fmt.Fprintln(w, "Problem creating new company: ", err.Error())
		log.Println("Problem creating new company: ", err.Error())
		return
	}
	newUser := User{
		Name:       r.Form.Get("contact"),
		Email:      r.Form.Get("email"),
		Email2:     r.Form.Get("email"),
		Password:   "#syndica2018",
		Role:       "client",
		Status:     "active",
		Profile:    r.Form.Get("short"),
		Job:        r.Form.Get("name"),
		Title:      "Founder",
		Experience: r.Form.Get("category"),
	}
	err = CreateUser(newUser)
	if err != nil {
		fmt.Fprintln(w, "Problem creating new user: ", err.Error())
		log.Println("Problem creating new user: ", err.Error())
		return
	}
	log.Printf("successful create new company %s => %s\n", email, newName)
	http.Redirect(w, r, "/admin", 302)
}

func adminUserHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if !IsAdmin(email) {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if r.Method != "POST" {
		// determine viewing 1 user or listing all users
		e := r.URL.Query().Get("email")
		if e != "" {
			log.Println("editing user:", e)
			t, _ := template.ParseFiles("templates/users_single.gtpl")
			t.Execute(w, GetUser(e))
		} else {
			t, _ := template.ParseFiles("templates/users.gtpl")
			t.Execute(w, GetUsers())
		}
		return
	}
	// update user
	r.ParseForm()
	editEmail := r.Form.Get("editemail")
	if editEmail == "" {
		log.Println("invalid form post: missing email")
		return
	}
	status := r.Form.Get("status")
	if status == "delete" {
		err := DeleteUser(editEmail)
		if err != nil {
			fmt.Fprintln(w, "Problem deleting user: ", err.Error())
			log.Println("Problem deleting user: ", err.Error())
			return
		}
		log.Println("successful delete user: ", editEmail)
		http.Redirect(w, r, "/admin/user", 302)
	}
	year, _ := strconv.Atoi(r.Form.Get("year"))
	u := User{
		Name:     r.Form.Get("name"),
		Email:    r.Form.Get("email"),
		Email2:   r.Form.Get("email2"),
		Password: r.Form.Get("password"),
		Year:     year,
		Role:     r.Form.Get("role"),
		Status:   status,
		Profile:  r.Form.Get("profile"),
	}
	err := EditUser(editEmail, u)
	if err != nil {
		fmt.Fprintln(w, "Problem updating user: ", err.Error())
		log.Println("Problem updating user: ", err.Error())
		return
	}
	log.Println("successful edit user: ", editEmail)
	http.Redirect(w, r, "/admin/user", 302)
}

func advisorHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	// determine viewing 1 advisor or listing all advisors
	e := r.URL.Query().Get("email")
	if e != "" {
		t, _ := template.ParseFiles("templates/advisors_single.gtpl")
		t.Execute(w, GetAdvisor(e))
	} else {
		t, _ := template.ParseFiles("templates/advisors.gtpl")
		t.Execute(w, GetAdvisors())
	}
}

func portfolioHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	user := GetUser(email)
	if user.Role == "client" {
		http.Redirect(w, r, "/home", 302)
		return
	}
	t, _ := template.ParseFiles("templates/portfolio.gtpl")
	t.Execute(w, struct {
		Advisor   User
		Companies []Company
		Questions []Question
	}{
		Advisor:   GetAdvisor(email),
		Companies: GetMyCompanies(email),
		Questions: GetAdvisorTeamsQuestions(email),
	})
}

func companyHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		path := base64.URLEncoding.EncodeToString([]byte(r.URL.RequestURI()))
		http.Redirect(w, r, "/login?r="+path, 302)
		return
	}
	user := GetUser(email)
	myCompany := ""
	if user.Role == "client" {
		myCompany = GetCompanyName(email)
	}
	// determine viewing 1 company or listing all companies
	c := r.URL.Query().Get("name")
	if user.Role == "client" && c != myCompany {
		http.Redirect(w, r, "/home", 302)
		return
	}

	var t *template.Template
	if c == "" {
		t, _ = template.ParseFiles("templates/companies.gtpl")
		t.Execute(w, GetCompanies())
		return
	}

	thisCompany := GetCompany(c)
	if thisCompany.Name == "" {
		http.Redirect(w, r, "/home", 302)
		return
	}
	q, err := strconv.Atoi(r.URL.Query().Get("q"))
	if err != nil || q == 0 {
		isOnTeam := IsOnTeam(email, c)

		team, err := strconv.Atoi(r.URL.Query().Get("team"))
		if err != nil || team == 0 {
			t, _ = template.ParseFiles("templates/companies_single.gtpl")
			t.Execute(w, struct {
				T bool
				C Company
			}{
				T: isOnTeam,
				C: thisCompany,
			})
			return
		}

		// if asking for team questions
		t, _ = template.ParseFiles("templates/team_advisor.gtpl")
		members := GetTeamMembersFull(c)
		questions := GetCompanyTeamQuestions(c)
		t.Execute(w, struct {
			C string
			Q []Question
			M []User
		}{
			C: c,
			Q: questions,
			M: members,
		})
		return
	}

	if user.Role == "client" {
		t, _ = template.ParseFiles("templates/companies_single_discussion.gtpl")
	} else {
		t, _ = template.ParseFiles("templates/companies_single_discussion_advisor.gtpl")
	}
	t.Execute(w, struct {
		Q int
		C Company
	}{
		Q: q,
		C: thisCompany,
	})
}

func questionDeleteHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	user := GetUser(email)
	myCompany := ""
	if user.Role != "client" {
		http.Redirect(w, r, "/home", 302)
		return
	}
	myCompany = GetCompanyName(email)
	c := r.URL.Query().Get("name")
	if c != myCompany {
		http.Redirect(w, r, "/home", 302)
		return
	}
	q, err := strconv.Atoi(r.URL.Query().Get("q"))
	if err != nil || q == 0 {
		http.Redirect(w, r, "/home", 302)
		return
	}
	DeleteQuestion(c, q)
	http.Redirect(w, r, "/home", 302)
	return
}

func adminAdvisorHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if !IsAdmin(email) {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if r.Method != "POST" {
		// determine viewing 1 advisor or listing all advisors
		e := r.URL.Query().Get("email")
		if e != "" {
			t, _ := template.ParseFiles("templates/advisors_single_admin.gtpl")
			t.Execute(w, GetAdvisor(e))
		} else {
			t, _ := template.ParseFiles("templates/advisors_admin.gtpl")
			t.Execute(w, GetAdvisors())
		}
		return
	}
	// update advisor
	r.ParseForm()
	editEmail := r.Form.Get("editemail")
	if editEmail == "" {
		log.Println("invalid form post: missing email")
		return
	}
	year, _ := strconv.Atoi(r.Form.Get("year"))
	u := User{
		Name:    r.Form.Get("name"),
		Email:   r.Form.Get("email"),
		Email2:  r.Form.Get("email2"),
		Year:    year,
		Status:  r.Form.Get("status"),
		Profile: r.Form.Get("profile"),
	}
	err := EditUser(editEmail, u)
	if err != nil {
		fmt.Fprintln(w, "Problem updating advisor: ", err.Error())
		log.Println("Problem updating advisor: ", err.Error())
		return
	}
	log.Println("successful edit advisor: ", editEmail)
	http.Redirect(w, r, "/admin/advisor", 302)
}

func adminCompanyHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if !IsAdmin(email) {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if r.Method != "POST" {
		// determine viewing 1 company or listing all companies
		c := r.URL.Query().Get("name")
		if c != "" {
			t, _ := template.ParseFiles("templates/companies_single_admin.gtpl")
			t.Execute(w, GetCompany(c))
		} else {
			t, _ := template.ParseFiles("templates/companies_admin.gtpl")
			t.Execute(w, GetCompaniesAdmin())
		}
		return
	}
	// update company
	r.ParseForm()
	editName := r.Form.Get("editname")
	if editName == "" {
		log.Println("invalid form post: missing name")
		return
	}
	c := Company{
		Name:        r.Form.Get("name"),
		Contact:     r.Form.Get("contact"),
		Email:       r.Form.Get("email"),
		Website:     r.Form.Get("website"),
		Short:       r.Form.Get("short"),
		Description: r.Form.Get("description"),
		Logo:        r.Form.Get("logo"),
		PrimaryPic:  r.Form.Get("primarypic"),
		Founders:    r.Form.Get("founders"),
		Category:    r.Form.Get("category"),
		Stage:       r.Form.Get("stage"),
		Twitter:     r.Form.Get("twitter"),
		Facebook:    r.Form.Get("facebook"),
		AngelList:   r.Form.Get("angellist"),
		LinkedIn:    r.Form.Get("linkedin"),
		Instagram:   r.Form.Get("instagram"),
		Youtube:     r.Form.Get("youtube"),
		GooglePlus:  r.Form.Get("googleplus"),
		Status:      r.Form.Get("status"),
	}
	err := EditCompany(editName, c)
	if err != nil {
		fmt.Fprintln(w, "Problem updating company: ", err.Error())
		log.Println("Problem updating company: ", err.Error())
		return
	}
	log.Println("successful edit company: ", editName)
	http.Redirect(w, r, "/admin/company", 302)
}

func adminDeleteCompanyHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if !IsAdmin(email) {
		http.Redirect(w, r, "/login", 302)
		return
	}
	r.ParseForm()
	name := r.Form.Get("deletename")
	if name == "" {
		log.Println("invalid form post: missing name")
		return
	}
	err := DeleteCompany(name)
	if err != nil {
		fmt.Fprintln(w, "Problem removing company: ", err.Error())
		log.Println("Problem removing company: ", err.Error())
		return
	}
	log.Println("successful removed company: ", name)
	http.Redirect(w, r, "/admin/company", 302)
}

// func companyListHandler(w http.ResponseWriter, r *http.Request) {
// 	// will need to cache these results
// 	t, _ := template.ParseFiles("templates/company_list.gtpl")
// 	t.Execute(w, GetCompanies())
// }

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		t, _ := template.ParseFiles("templates/contact.gtpl")
		t.Execute(w, nil)
		return
	}
	r.ParseForm()
	if r.Form.Get("app") != "syndica" {
		return
	}
	name := r.Form.Get("name")
	if name == "" {
		return
	}
	email := r.Form.Get("email")
	if email == "" {
		return
	}
	comment := r.Form.Get("comment")
	if comment == "" {
		return
	}
	SendContactUsEmail(name, email, comment)
	t, _ := template.ParseFiles("templates/contact_thanks.gtpl")
	t.Execute(w, nil)
}

func joinTeamHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if r.Method != "POST" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	// join the company team
	r.ParseForm()
	company := r.Form.Get("company")
	if company == "" {
		return
	}
	// dont allow clients to join company team
	if IsClient(email) {
		t, _ := template.ParseFiles("templates/sorry_joining_client.gtpl")
		t.Execute(w, nil)
		return
	}
	// join team
	err := JoinTeam(email, company)
	if err != nil {
		c := GetCompany(company)
		t, _ := template.ParseFiles("templates/sorry_joining.gtpl")
		t.Execute(w, c)
		return
	}
	name := GetName(email)
	AddActivity(fmt.Sprintf("%s has joined %s's advisory team", name, company), email, company)
	http.Redirect(w, r, "/company?name="+company, 302)
	//t, _ := template.ParseFiles("templates/thanks_joining.gtpl")
	//t.Execute(w, c)
}

func myRewardHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		return
	}
	company := r.URL.Query().Get("company")
	if company == "" {
		return
	}
	if IsClient(email) && company != GetCompanyName(email) {
		return
	}
	q := r.URL.Query().Get("q")
	if q == "" {
		return
	}
	questionID, _ := strconv.Atoi(q)
	if questionID == 0 {
		return
	}
	t, _ := template.ParseFiles("templates/myreward.gtpl")
	t.Execute(w, GetCompanyReward(company, questionID))
}

func myTeamHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		return
	}
	company := r.URL.Query().Get("company")
	if company == "" {
		return
	}
	if IsClient(email) && company != GetCompanyName(email) {
		return
	}
	t, _ := template.ParseFiles("templates/myteam.gtpl")
	t.Execute(w, GetTeamMembers(company))
}

func teamHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if !IsClient(email) {
		http.Redirect(w, r, "/home", 302)
		return
	}
	company := GetCompanyName(email)
	if company == "" {
		return
	}
	if company != GetCompanyName(email) {
		return
	}
	t, _ := template.ParseFiles("templates/team.gtpl")
	t.Execute(w, GetTeamMembersFull(company))
}

func companyTeamQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		return
	}
	company := GetCompanyName(email)
	if company == "" {
		return
	}
	if IsClient(email) && company != GetCompanyName(email) {
		return
	}
	t, _ := template.ParseFiles("templates/myquestions.gtpl")
	t.Execute(w, GetCompanyTeamQuestions(company))
}

func companyQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		return
	}
	company := r.URL.Query().Get("company")
	if company == "" {
		return
	}
	if IsClient(email) && company != GetCompanyName(email) {
		return
	}
	t, _ := template.ParseFiles("templates/myquestions.gtpl")
	t.Execute(w, GetCompanyQuestions(company))
}

func loadCompanyQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		return
	}
	company := r.URL.Query().Get("company")
	if company == "" {
		return
	}
	if IsClient(email) && company != GetCompanyName(email) {
		return
	}
	t, _ := template.ParseFiles("templates/loadcompanyquestions.gtpl")
	t.Execute(w, GetCompanyQuestions(company))
}

func myQuestionHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		return
	}
	company := r.URL.Query().Get("company")
	if company == "" {
		return
	}
	if IsClient(email) && company != GetCompanyName(email) {
		return
	}
	t, _ := template.ParseFiles("templates/myquestion.gtpl")
	q := r.URL.Query().Get("q")
	if q == "" || q == "0" {
		t.Execute(w, GetCompanyTopQuestion(company))
		return
	}
	qid, err := strconv.Atoi(r.URL.Query().Get("q"))
	if err != nil || qid == 0 {
		http.Redirect(w, r, "/home", 302)
		return
	}
	t.Execute(w, GetCompanyQuestion(company, qid))
	return
}

func newQuestionHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if !IsClient(email) {
		http.Redirect(w, r, "/home", 302)
		return
	}
	if r.Method != "POST" {
		t, _ := template.ParseFiles("templates/newquestion.gtpl")
		t.Execute(w, nil)
		return
	}
	r.ParseForm()
	if r.Form.Get("app") != "syndica" {
		return
	}
	question := r.Form.Get("question")
	if question == "" {
		return
	}
	body := r.Form.Get("body")
	reward := r.Form.Get("reward")
	company := GetCompanyName(email)
	q := Question{
		Question: question,
		Body:     body,
		Company:  company,
		Reward:   reward,
		Private:  false,
	}
	AddNewQuestion(q)
	AddActivity(fmt.Sprintf("%s has posted a new question", company), email, company)
	http.Redirect(w, r, "/home", 302)
}

func loadCompanyCommentsHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		return
	}
	company := r.URL.Query().Get("company")
	if company == "" {
		return
	}
	// don't allow company A to view responses from company B
	if IsClient(email) && company != GetCompanyName(email) {
		return
	}
	// identify if comment is written by self
	comments := GetCompanyComments(company)
	for i := 0; i < len(comments); i++ {
		if email == comments[i].Advisor {
			comments[i].IsSelf = true
		}
	}
	t, _ := template.ParseFiles("templates/loadcompanycomments.gtpl")
	t.Execute(w, comments)
}

func newMessageHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if !IsClient(email) {
		http.Redirect(w, r, "/home", 302)
		return
	}
	if r.Method != "POST" {
		t, _ := template.ParseFiles("templates/newmessage.gtpl")
		t.Execute(w, nil)
		return
	}
	r.ParseForm()
	if r.Form.Get("app") != "syndica" {
		return
	}
	question := r.Form.Get("question")
	if question == "" {
		return
	}
	body := r.Form.Get("body")
	reward := r.Form.Get("reward")
	company := GetCompanyName(email)
	q := Question{
		Question: question,
		Body:     body,
		Company:  company,
		Reward:   reward,
		Private:  true,
	}
	AddNewQuestion(q)
	AddActivity(fmt.Sprintf("%s has posted a new message to the team", company), email, company)
	http.Redirect(w, r, "/team", 302)
}

func questionEditHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	user := GetUser(email)
	myCompany := ""
	if user.Role != "client" {
		http.Redirect(w, r, "/home", 302)
		return
	}
	myCompany = GetCompanyName(email)
	c := r.URL.Query().Get("name")
	if c != myCompany {
		http.Redirect(w, r, "/home", 302)
		return
	}
	q, err := strconv.Atoi(r.URL.Query().Get("q"))
	if err != nil || q == 0 {
		http.Redirect(w, r, "/home", 302)
		return
	}
	if r.Method != "POST" {
		t, _ := template.ParseFiles("templates/editquestion.gtpl")
		t.Execute(w, GetCompanyQuestion(c, q))
		return
	}
	// Edit Question
	r.ParseForm()
	if r.Form.Get("app") != "syndica" {
		return
	}
	question := r.Form.Get("question")
	if question == "" {
		return
	}
	body := r.Form.Get("body")
	reward := r.Form.Get("reward")
	// private := r.Form.Get("private")
	// var p bool
	// if private == "1" {
	// 	p = true
	// } else {
	// 	p = false
	// }
	editQ := Question{
		ID:       q,
		Question: question,
		Body:     body,
		Company:  c,
		Reward:   reward,
		// Private:  p,
	}
	EditQuestion(editQ)
	// if p {
	// 	http.Redirect(w, r, "/team", 302)
	// 	return
	// }
	http.Redirect(w, r, "/home", 302)
}

func myResponsesHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		return
	}
	company := r.URL.Query().Get("company")
	if company == "" {
		return
	}
	// don't allow company A to view responses from company B
	if IsClient(email) && company != GetCompanyName(email) {
		return
	}
	q, err := strconv.Atoi(r.URL.Query().Get("q"))
	if err != nil || q == 0 {
		return
	}
	responses := GetResponses(q)
	for i := 0; i < len(responses); i++ {
		if email == responses[i].Advisor {
			responses[i].IsSelf = true
		}
	}
	t, _ := template.ParseFiles("templates/myresponses.gtpl")
	type responseTpl struct {
		QuestionID int
		Company    string
		Responses  []ResponseWithUser
	}
	t.Execute(w, responseTpl{
		QuestionID: q,
		Company:    company,
		Responses:  responses,
	})
}

func responseHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if r.Method != "POST" {
		return
	}
	r.ParseForm()
	company := r.Form.Get("company")
	if company == "" {
		return
	}
	answer := r.Form.Get("answer")
	if answer == "" {
		return
	}
	questionID, err := strconv.Atoi(r.Form.Get("question_id"))
	if err != nil || questionID == 0 {
		return
	}
	replyID, err := strconv.Atoi(r.Form.Get("reply_id"))
	if err != nil || replyID == 0 {
		return
	}
	response := Response{
		QuestionID: questionID,
		Advisor:    email,
		Answer:     answer,
		ReplyID:    replyID,
	}
	AddNewResponse(response)
	name := GetName(email)
	AddActivity(fmt.Sprintf("%s has replied to %s's question", name, company), email, company)
	http.Redirect(w, r, fmt.Sprintf("/company?name=%s&q=%d", company, questionID), 302)
}

func responseEditHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	company := r.URL.Query().Get("company")
	if company == "" {
		http.Redirect(w, r, "/home", 302)
		return
	}
	qid, err := strconv.Atoi(r.URL.Query().Get("qid"))
	if err != nil || qid == 0 {
		http.Redirect(w, r, "/home", 302)
		return
	}
	rid, err := strconv.Atoi(r.URL.Query().Get("rid"))
	if err != nil || rid == 0 {
		http.Redirect(w, r, "/home", 302)
		return
	}
	// verify if client, then can only view his own questions
	if IsClient(email) && company != GetCompanyName(email) {
		http.Redirect(w, r, "/home", 302)
		return
	}
	if !IsCompanyQuestionValid(company, qid) {
		http.Redirect(w, r, "/home", 302)
		return
	}
	if r.Method != "POST" {
		answer := GetResponseAnswer(rid, qid, email)
		t, _ := template.ParseFiles("templates/editresponse.gtpl")
		t.Execute(w, struct {
			Company string
			QID     int
			RID     int
			Answer  string
		}{
			Company: company,
			QID:     qid,
			RID:     rid,
			Answer:  answer,
		})
		return
	}
	// if Post then get form data and update response
	r.ParseForm()
	if r.Form.Get("app") != "syndica" {
		http.Redirect(w, r, "/login?error=1", 302)
		return
	}
	answer := strings.TrimSpace(r.Form.Get("answer"))
	if answer == "" {
		answer = "*deleted by user*"
	}
	response := Response{
		Answer:     answer,
		Advisor:    email,
		QuestionID: qid,
		ReplyID:    rid,
	}
	EditResponse(response)
	http.Redirect(w, r, fmt.Sprintf("/company?name=%s&q=%d", company, qid), 302)
}

func responseDeleteHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	company := r.URL.Query().Get("company")
	if company == "" {
		http.Redirect(w, r, "/home", 302)
		return
	}
	qid, err := strconv.Atoi(r.URL.Query().Get("qid"))
	if err != nil || qid == 0 {
		http.Redirect(w, r, "/home", 302)
		return
	}
	rid, err := strconv.Atoi(r.URL.Query().Get("rid"))
	if err != nil || rid == 0 {
		http.Redirect(w, r, "/home", 302)
		return
	}
	response := Response{
		Advisor:    email,
		QuestionID: qid,
		ReplyID:    rid,
	}
	DeleteResponse(response)
	http.Redirect(w, r, fmt.Sprintf("/company?name=%s&q=%d", company, qid), 302)
}

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	r.ParseForm()
	m := strings.TrimSpace(r.Form.Get("message"))
	if m == "" {
		return
	}
	message := Message{
		User:    "peter@syndica.net",
		From:    email,
		Message: m,
	}
	AddMessage(message)
	http.Redirect(w, r, "/messages", 302)
}

func adminSendMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if !IsAdmin(email) {
		http.Redirect(w, r, "/login", 302)
		return
	}
	r.ParseForm()
	to := strings.TrimSpace(r.Form.Get("to"))
	m := strings.TrimSpace(r.Form.Get("message"))
	if to == "" || m == "" {
		return
	}
	message := Message{
		User:    to,
		From:    email,
		Message: m,
	}
	AddMessage(message)
	http.Redirect(w, r, "/messages", 302)
}

func likeHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		return
	}
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	if id == 0 {
		return
	}
	AddLike(id)
	name := GetName(email)
	name2 := GetNameFromResponseID(id)
	company := GetCompanyFromResponseID(id)
	AddActivity(fmt.Sprintf("%s has acknowledged %s's response to %s's question", name, name2, company), email, company)
}

func commentHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if r.Method != "POST" {
		return
	}
	r.ParseForm()
	company := r.Form.Get("company")
	if company == "" {
		return
	}
	if IsClient(email) && company != GetCompanyName(email) {
		return
	}
	replyID, err := strconv.Atoi(r.Form.Get("reply_id"))
	if err != nil {
		return
	}
	comment := r.Form.Get("comment")
	if comment == "" {
		return
	}
	AddNewComment(Comment{
		Company: company,
		Advisor: email,
		Comment: comment,
		ReplyID: replyID,
	})
	// name := GetName(email)
	// AddActivity(fmt.Sprintf("%s has added a comment for %s", name, company), email, company)
	http.Redirect(w, r, fmt.Sprintf("/company?name=%s", company), 302)
}

func commentEditHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	company := r.URL.Query().Get("company")
	if company == "" {
		http.Redirect(w, r, "/home", 302)
		return
	}
	cid, err := strconv.Atoi(r.URL.Query().Get("cid"))
	if err != nil || cid == 0 {
		http.Redirect(w, r, "/home", 302)
		return
	}
	// verify if client, then can only view his own questions
	if IsClient(email) {
		if GetCompanyName(email) != company {
			http.Redirect(w, r, "/home", 302)
			return
		}
	}
	if r.Method != "POST" {
		comment := GetCompanyComment(cid, company, email)
		t, _ := template.ParseFiles("templates/editcomment.gtpl")
		t.Execute(w, struct {
			ID      int
			Company string
			Comment string
		}{
			ID:      cid,
			Company: company,
			Comment: comment,
		})
		return
	}
	// if Post then get form data and update comment
	r.ParseForm()
	if r.Form.Get("app") != "syndica" {
		http.Redirect(w, r, "/login?error=1", 302)
		return
	}
	comment := strings.TrimSpace(r.Form.Get("comment"))
	if comment == "" {
		comment = "*deleted by user*"
	}
	EditComment(Comment{
		ID:      cid,
		Company: company,
		Advisor: email,
		Comment: comment,
	})
	http.Redirect(w, r, fmt.Sprintf("/company?name=%s", company), 302)
}

func commentDeleteHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	company := r.URL.Query().Get("company")
	if company == "" {
		http.Redirect(w, r, "/home", 302)
		return
	}
	cid, err := strconv.Atoi(r.URL.Query().Get("cid"))
	if err != nil || cid == 0 {
		http.Redirect(w, r, "/home", 302)
		return
	}
	comment := Comment{
		ID:      cid,
		Advisor: email,
		Company: company,
	}
	DeleteComment(comment)
	http.Redirect(w, r, fmt.Sprintf("/company?name=%s", company), 302)
}

func commentLikeHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		return
	}
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	if id == 0 {
		return
	}
	AddLikeComment(id)
	// name := GetName(email)
	// name2 := GetNameFromCommentID(id)
	// company := GetCompanyFromCommentID(id)
	// AddActivity(fmt.Sprintf("%s has acknowledged %s's comment to %s", name, name2, company), email, company)
}

func newThreadHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if IsClient(email) {
		return
	}
	if r.Method != "POST" {
		t, _ := template.ParseFiles("templates/newthread.gtpl")
		t.Execute(w, nil)
		return
	}
	r.ParseForm()
	title := r.Form.Get("title")
	title = strings.TrimSpace(title)
	if title == "" {
		return
	}
	url := r.Form.Get("url")
	url = strings.TrimSpace(url)
	body := r.Form.Get("body")
	body = strings.TrimSpace(body)
	t := Thread{
		Title:   title,
		Advisor: email,
		Url:     url,
		Body:    body,
	}
	AddNewThread(t)
	http.Redirect(w, r, "/home", 302)
}

func threadHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		path := base64.URLEncoding.EncodeToString([]byte(r.URL.RequestURI()))
		http.Redirect(w, r, "/login?r="+path, 302)
		return
	}
	if IsClient(email) {
		return
	}
	tid, _ := strconv.Atoi(r.URL.Query().Get("tid"))
	if tid == 0 {
		return
	}
	t, _ := template.ParseFiles("templates/thread.gtpl")
	t.Execute(w, GetThread(tid))
	return
}

func threadReplyHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if IsClient(email) {
		return
	}
	if r.Method != "POST" {
		return
	}
	r.ParseForm()
	tid, _ := strconv.Atoi(r.URL.Query().Get("tid"))
	if tid == 0 {
		return
	}
	response := r.Form.Get("response")
	response = strings.TrimSpace(response)
	if response == "" {
		return
	}
	t := ThreadResponse{
		DiscussionID: tid,
		Advisor:      email,
		Comment:      response,
	}
	AddNewThreadResponse(t)
	http.Redirect(w, r, fmt.Sprintf("/thread?tid=%d", tid), 302)
}

func threadCommentReplyHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if IsClient(email) {
		return
	}
	if r.Method != "POST" {
		return
	}
	r.ParseForm()
	tid, _ := strconv.Atoi(r.Form.Get("discussion_id"))
	if tid == 0 {
		return
	}
	rid, _ := strconv.Atoi(r.Form.Get("reply_id"))
	comment := r.Form.Get("comment")
	comment = strings.TrimSpace(comment)
	if comment == "" {
		return
	}
	t := ThreadResponse{
		DiscussionID: tid,
		ReplyID:      rid,
		Advisor:      email,
		Comment:      comment,
	}
	AddNewThreadResponse(t)
	http.Redirect(w, r, fmt.Sprintf("/thread?tid=%d", tid), 302)
}

func threadReplyEditHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if IsClient(email) {
		return
	}
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	if id == 0 {
		return
	}
	tid, _ := strconv.Atoi(r.URL.Query().Get("tid"))
	if tid == 0 {
		return
	}
	if r.Method != "POST" {
		t, _ := template.ParseFiles("templates/edit_thread_response.gtpl")
		t.Execute(w, GetThreadComment(id, email))
		return
	}
	// if Post then get form data and update thread comment
	r.ParseForm()
	comment := strings.TrimSpace(r.Form.Get("comment"))
	if comment == "" {
		comment = "*deleted by user*"
	}
	response := ThreadResponse{
		ID:           id,
		DiscussionID: tid,
		Advisor:      email,
		Comment:      comment,
	}
	EditThreadResponse(response)
	http.Redirect(w, r, fmt.Sprintf("/thread?tid=%d", tid), 302)
}

func threadReplyDeleteHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		http.Redirect(w, r, "/login", 302)
		return
	}
	if IsClient(email) {
		return
	}
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	if id == 0 {
		return
	}
	tid, _ := strconv.Atoi(r.URL.Query().Get("tid"))
	if tid == 0 {
		return
	}
	response := ThreadResponse{
		ID:           id,
		DiscussionID: tid,
		Advisor:      email,
	}
	DeleteThreadResponse(response)
	http.Redirect(w, r, fmt.Sprintf("/thread?tid=%d", tid), 302)
}

func loadThreadCommentsHandler(w http.ResponseWriter, r *http.Request) {
	email := GetEmail(r)
	if email == "" {
		return
	}
	if IsClient(email) {
		return
	}
	tid, _ := strconv.Atoi(r.URL.Query().Get("tid"))
	if tid == 0 {
		return
	}
	comments := GetThreadComments(tid)
	for i := 0; i < len(comments); i++ {
		if email == comments[i].Advisor {
			comments[i].IsSelf = true
		}
	}
	t, _ := template.ParseFiles("templates/thread_comments.gtpl")
	t.Execute(w, comments)
	return
}
