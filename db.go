package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBConnStr = "appsyndica:RqWkak26Zv_9@tcp(syndicadb.cgysf0dosjda.us-west-1.rds.amazonaws.com:3306)/syndica?tls=skip-verify"
)

var (
	dbConn *sql.DB
	dbLock sync.Mutex
)

type User struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Email2     string `json:"email2"`
	Password   string `json:"password"`
	Year       int    `json:"year"`
	Role       string `json:"role"`
	Status     string `json:"status"`
	Profile    string `json:"profile"`
	Job        string `json:"job"`
	Title      string `json:"title"`
	Experience string `json:"experience"`
	Hash       string `json:"hash"`
}

type Company struct {
	Name        string `json:"name"`
	Contact     string `json:"contact"`
	Email       string `json:"email"`
	Website     string `json:"website"`
	Short       string `json:"short"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	PrimaryPic  string `json:"primarypic"`
	Founders    string `json:"founders"`
	Category    string `json:"category"`
	Stage       string `json:"stage"`
	Twitter     string `json:"twitter"`
	Facebook    string `json:"facebook"`
	AngelList   string `json:"angellist"`
	LinkedIn    string `json:"linkedin"`
	Instagram   string `json:"instagram"`
	Youtube     string `json:"youtube"`
	GooglePlus  string `json:"googleplus"`
	Status      string `json:"status"`
	ShowOrder   int    `json:"showorder"`
}

type Team struct {
	Advisor string `json:"advisor"`
	Company string `json:"company"`
}

type Question struct {
	ID       int
	Question string
	Company  string
	Body     string
	Reward   string
	Private  bool
}

type Response struct {
	ID         int
	QuestionID int
	Advisor    string
	Answer     string
	Likes      int
	ReplyID    int
}

type ResponseWithUser struct {
	ID         int
	QuestionID int
	Advisor    string
	Answer     string
	Likes      int
	ReplyID    int
	Name       string
	Title      string
	Job        string
	Experience string
	IsSelf     bool
}

type Activity struct {
	ID       int
	Activity string
	Time     string
	User     string
	Company  string
}

type Comment struct {
	ID      int
	Company string
	Advisor string
	Comment string
	Likes   int
	ReplyID int
}

type CommentWithUser struct {
	ID         int
	Company    string
	Advisor    string
	Comment    string
	Likes      int
	ReplyID    int
	Name       string
	Title      string
	Job        string
	Experience string
	IsSelf     bool
}

type Thread struct {
	ID      int
	Title   string
	Advisor string
	Count   int
	Created string
	TimeAgo string
	Url     string
	Body    string
}

type ThreadWithUser struct {
	Thread
	Name          string
	IsSelf        bool
	IsCompanyLink bool
}

type ThreadResponse struct {
	ID           int
	DiscussionID int
	ReplyID      int
	Advisor      string
	Comment      string
	Created      string
	TimeAgo      string
}

type ThreadResponseWithUser struct {
	ThreadResponse
	Name   string
	IsSelf bool
}

func GetDB() *sql.DB {
	var err error
	// lock and try to get connection
	dbLock.Lock()
	// if no connection then reset connection
	if dbConn == nil {
		dbConn, err = sql.Open("mysql", DBConnStr)
		if err != nil {
			log.Println(err.Error())
			dbLock.Unlock()
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			return GetDB()
		}
	}
	// if yes connection but no ping, then reconnect
	err = dbConn.Ping()
	if err != nil {
		dbConn, err = sql.Open("mysql", DBConnStr)
		if err != nil {
			log.Println(err.Error())
			dbLock.Unlock()
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			return GetDB()
		}
		// try one more time to ping, then give up
		err = dbConn.Ping()
		if err != nil {
			log.Println(err.Error())
			dbLock.Unlock()
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			return GetDB()
		}
	}
	dbLock.Unlock()
	return dbConn
}

func GetUser(email string) User {
	var u User
	var email2 sql.NullString
	var status sql.NullString
	var profile sql.NullString
	var year sql.NullInt64
	var job sql.NullString
	var title sql.NullString
	var experience sql.NullString
	row := GetDB().QueryRow("SELECT name,email,email2,year,role,status,profile,job,title,experience FROM users WHERE email=?", email)
	err := row.Scan(&u.Name, &u.Email, &email2, &year, &u.Role, &status, &profile, &job, &title, &experience)
	if err != nil {
		log.Println(err.Error())
		return u
	}
	if email2.Valid {
		u.Email2 = email2.String
	}
	if status.Valid {
		u.Status = status.String
	}
	if profile.Valid {
		u.Profile = profile.String
	}
	if year.Valid {
		u.Year = int(year.Int64)
	}
	if job.Valid {
		u.Job = job.String
	}
	if title.Valid {
		u.Title = title.String
	}
	if experience.Valid {
		u.Experience = experience.String
	}
	return u
}

func GetUsers() []User {
	rows, err := GetDB().Query("SELECT name,email,email2,year,role,status FROM users")
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()

	var uSet []User
	for rows.Next() {
		var u User
		var email2 sql.NullString
		var status sql.NullString
		var year sql.NullInt64
		err := rows.Scan(&u.Name, &u.Email, &email2, &year, &u.Role, &status)
		if err != nil {
			log.Println(err.Error())
			break
		}
		if email2.Valid {
			u.Email2 = email2.String
		}
		if status.Valid {
			u.Status = status.String
		}
		if year.Valid {
			u.Year = int(year.Int64)
		}
		uSet = append(uSet, u)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return uSet
}

func IsMBA(email string) bool {
	return strings.HasSuffix(email, "@anderson.ucla.edu")
}

func ValidInitialUser(email, hash string) bool {
	var e string
	err := GetDB().QueryRow("SELECT email FROM users WHERE email=? AND hash=? AND status='pending'", email, hash).Scan(&e)
	if err == nil {
		return true
	}
	if err == sql.ErrNoRows {
		return false
	}
	log.Println(err.Error())
	return false
}

func CreateUserInitial(u User) error {
	if IsMBA(u.Email) {
		u.Name = "MBA Advisor"
		u.Role = "advisor"
	} else {
		u.Name = "Client"
		u.Role = "client"
	}
	insert, err := GetDB().Query("INSERT INTO users (name,email,password,role,hash,status) VALUES (?,?,?,?,?,?)",
		u.Name, u.Email, GeneratePasswordHash(u.Hash), u.Role, u.Hash, "pending")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	insert.Close()
	return nil
}

func CreateUser(u User) error {
	insert, err := GetDB().Query("INSERT INTO users (name,email,email2,password,year,role,status) VALUES (?,?,?,?,?,?,?)",
		u.Name, u.Email, u.Email2, GeneratePasswordHash(u.Password), u.Year, u.Role, u.Status)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	insert.Close()
	update, err := GetDB().Query("UPDATE users SET profile=? WHERE email=?", truncStr(u.Profile), u.Email)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	update.Close()
	return nil
}

func CreateCompany(c Company) error {
	insert, err := GetDB().Query("INSERT INTO companies (name,contact,email,website,logo,primarypic,short,founders,category,stage,twitter,facebook,angellist,linkedin,instagram,youtube,googleplus,status) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		c.Name, c.Contact, c.Email, c.Website, c.Logo, c.PrimaryPic, truncStr(c.Short), c.Founders, c.Category, c.Stage, c.Twitter, c.Facebook, c.AngelList, c.LinkedIn, c.Instagram, c.Youtube, c.GooglePlus, c.Status)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	insert.Close()
	if c.Description != "" {
		update, err := GetDB().Query("UPDATE companies SET description=? WHERE name=?", truncStr(c.Description), c.Name)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		update.Close()
	}
	return nil
}

func UpdateUser(u User) error {
	update, err := GetDB().Query("UPDATE users SET name=?,password=?,year=?,status=? WHERE email=?",
		u.Name, GeneratePasswordHash(u.Password), u.Year, u.Status, u.Email)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	update.Close()
	return nil
}

func EditProfile(u User) error {
	if u.Password == "" {
		update, err := GetDB().Query("UPDATE users SET name=?,year=?,job=?,title=?,experience=? WHERE email=?",
			u.Name, u.Year, u.Job, u.Title, u.Experience, u.Email)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		update.Close()
	} else {
		update, err := GetDB().Query("UPDATE users SET name=?,password=?,year=?,job=?,title=?,experience=? WHERE email=?",
			u.Name, GeneratePasswordHash(u.Password), u.Year, u.Job, u.Title, u.Experience, u.Email)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		update.Close()
	}
	if u.Profile != "" {
		update2, err := GetDB().Query("UPDATE users SET profile=? WHERE email=?", truncStr(u.Profile), u.Email)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		update2.Close()
	}
	return nil
}

func EditClientProfile(u User) error {
	if u.Password == "" {
		update, err := GetDB().Query("UPDATE users SET name=? WHERE email=?",
			u.Name, u.Email)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		update.Close()
	} else {
		update, err := GetDB().Query("UPDATE users SET name=?,password=? WHERE email=?",
			u.Name, GeneratePasswordHash(u.Password), u.Email)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		update.Close()
	}
	return nil
}

func EditUser(editEmail string, u User) error {
	if u.Password == "" {
		update, err := GetDB().Query("UPDATE users SET name=?, email=?, email2=?, year=?, status=? WHERE email=?",
			u.Name, u.Email, u.Email2, u.Year, u.Status, editEmail)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		update.Close()
	} else {
		update, err := GetDB().Query("UPDATE users SET name=?, email=?, email2=?, password=?, year=?, status=? WHERE email=?",
			u.Name, u.Email, u.Email2, GeneratePasswordHash(u.Password), u.Year, u.Status, editEmail)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		update.Close()
	}
	if u.Role != "" && u.Role != "none" {
		update2, err := GetDB().Query("UPDATE users SET role=? WHERE email=?", u.Role, u.Email)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		update2.Close()
	}
	update3, err := GetDB().Query("UPDATE users SET profile=? WHERE email=?", truncStr(u.Profile), u.Email)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	update3.Close()
	return nil
}

func DeleteUser(email string) error {
	delete, err := GetDB().Query("DELETE FROM users WHERE email=?", email)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	delete.Close()
	return nil
}

func GetName(email string) string {
	var n string
	err := GetDB().QueryRow("SELECT name FROM users WHERE email=?", email).Scan(&n)
	if err != nil {
		log.Println("GetName: Missing name: ", err.Error())
		return ""
	}
	return n
}

func GetNameFromResponseID(id int) string {
	var n string
	err := GetDB().QueryRow("SELECT users.name FROM users JOIN responses ON users.email=responses.advisor WHERE responses.id=?", id).Scan(&n)
	if err != nil {
		log.Println(fmt.Sprintf("GetNameFromResponseID: Response #%d: Missing name: %s", id, err.Error()))
		return ""
	}
	return n
}

func GetCompanyFromResponseID(id int) string {
	var n string
	err := GetDB().QueryRow("SELECT questions.company FROM questions JOIN responses ON questions.id=responses.question_id WHERE responses.id=?", id).Scan(&n)
	if err != nil {
		log.Println(fmt.Sprintf("GetCompanyFromResponseID: Response #%d: Missing name: %s", id, err.Error()))
		return ""
	}
	return n
}

func GetNameFromCommentID(id int) string {
	var n string
	err := GetDB().QueryRow("SELECT users.name FROM users JOIN comments ON users.email=comments.advisor WHERE comments.id=?", id).Scan(&n)
	if err != nil {
		log.Println(fmt.Sprintf("GetNameFromCommentID: Comment #%d: Missing name: %s", id, err.Error()))
		return ""
	}
	return n
}

func GetCompanyFromCommentID(id int) string {
	var n string
	err := GetDB().QueryRow("SELECT comments.company FROM comments WHERE comments.id=?", id).Scan(&n)
	if err != nil {
		log.Println(fmt.Sprintf("GetCompanyFromCommentID: Comment #%d: Missing name: %s", id, err.Error()))
		return ""
	}
	return n
}

func SetUserNonce(email string, nonce string) error {
	update, err := GetDB().Query("UPDATE users SET nonce=? WHERE email=?", nonce, email)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	update.Close()
	return nil
}

func UserNonceMatches(nonce string, email string) bool {
	var e string
	err := GetDB().QueryRow("SELECT email FROM users WHERE email=? and nonce=?", email, nonce).Scan(&e)
	if err == nil {
		return true
	}
	if err == sql.ErrNoRows {
		return false
	}
	log.Println(err.Error())
	return false
}

func IsEmailExists(newEmail string) bool {
	var e string
	err := GetDB().QueryRow("SELECT email FROM users WHERE email=?", newEmail).Scan(&e)
	if err == nil {
		return true
	}
	if err == sql.ErrNoRows {
		return false
	}
	log.Println(err.Error())
	return true
}

func IsCompanyExists(newName string) bool {
	var n string
	err := GetDB().QueryRow("SELECT name FROM companies WHERE name=?", newName).Scan(&n)
	if err == nil {
		return true
	}
	if err == sql.ErrNoRows {
		return false
	}
	log.Println(err.Error())
	return true
}

func IsAdmin(email string) bool {
	var e string
	err := GetDB().QueryRow("SELECT email FROM users WHERE email=? and role='admin'", email).Scan(&e)
	if err == nil {
		return true
	}
	if err == sql.ErrNoRows {
		return false
	}
	log.Println(err.Error())
	return false
}

func IsAdvisor(email string) bool {
	var e string
	err := GetDB().QueryRow("SELECT email FROM users WHERE email=? and role='advisor'", email).Scan(&e)
	if err == nil {
		return true
	}
	if err == sql.ErrNoRows {
		return false
	}
	log.Println(err.Error())
	return false
}

func IsClient(email string) bool {
	var e string
	err := GetDB().QueryRow("SELECT email FROM users WHERE email=? and role='client'", email).Scan(&e)
	if err == nil {
		return true
	}
	if err == sql.ErrNoRows {
		return false
	}
	log.Println(err.Error())
	return false
}

func CanLogin(email, password, ip string) bool {
	if IsIPLocked(ip) {
		log.Println("IP is locked: ", ip)
		return false // cannot log in
	}
	if IsUserLocked(email) {
		log.Println("User is locked: ", email)
		return false // cannot log in
	}
	// check if user exists
	var hash string
	var fails int
	err := GetDB().QueryRow("SELECT password, loginfailed FROM users WHERE status<>'pending' AND email=?", email).Scan(&hash, &fails)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("User not found: ", email)
		} else {
			log.Println("Error with DB: ", err.Error())
		}
		return false
	}
	// check if user login correct password
	if !IsPasswordHashValid(hash, password) {
		log.Println("Invalid password for user: ", email)
		MarkFailLogin(email)
		LockUser(email, ip)
		LockIP(ip)
		return false
	}
	log.Printf("Login success for user: %s [prior failed logins: %d]\n", email, fails)
	// UnlockUser(email)
	// ClearFailLogin(email)
	SetLoginLast(email)
	return true
}

func MarkFailLogin(email string) error {
	update, err := GetDB().Query("UPDATE users SET loginfailed = loginfailed + 1 WHERE email=?", email)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	update.Close()
	return nil
}

func ClearFailLogin(email string) error {
	update, err := GetDB().Query("UPDATE users SET loginfailed = 0 WHERE email=?", email)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	update.Close()
	return nil
}

func SetLoginLast(email string) error {
	update, err := GetDB().Query("UPDATE users SET loginlast = NOW() WHERE email=?", email)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	update.Close()
	return nil
}

func truncStr(str string) string {
	num := 100000
	if len(str) > num {
		s := str[0:num] + "..."
		return s
	}
	return str
}

func GetAdvisor(email string) User {
	var u User
	var email2 sql.NullString
	var year sql.NullInt64
	var status sql.NullString
	var profile sql.NullString
	var job sql.NullString
	var title sql.NullString
	var experience sql.NullString
	row := GetDB().QueryRow("SELECT name,email,email2,year,status,profile,job,title,experience FROM users WHERE email=?", email)
	err := row.Scan(&u.Name, &u.Email, &email2, &year, &status, &profile, &job, &title, &experience)
	if err != nil {
		log.Println(err.Error())
		return u
	}
	if email2.Valid {
		u.Email2 = email2.String
	}
	if year.Valid {
		u.Year = int(year.Int64)
	}
	if status.Valid {
		u.Status = status.String
	}
	if profile.Valid {
		u.Profile = profile.String
	}
	if job.Valid {
		u.Job = job.String
	}
	if title.Valid {
		u.Title = title.String
	}
	if experience.Valid {
		u.Experience = experience.String
	}
	return u
}

func GetAdvisors() []User {
	rows, err := GetDB().Query("SELECT name,email,email2,year FROM users WHERE role='advisor'")
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()

	var uSet []User
	for rows.Next() {
		var u User
		var email2 sql.NullString
		var year sql.NullInt64
		err := rows.Scan(&u.Name, &u.Email, &email2, &year)
		if err != nil {
			log.Println(err.Error())
			break
		}
		if email2.Valid {
			u.Email2 = email2.String
		}
		if year.Valid {
			u.Year = int(year.Int64)
		}
		uSet = append(uSet, u)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return uSet
}

func GetCompany(name string) Company {
	var c Company
	var logo sql.NullString
	var primarypic sql.NullString
	var short sql.NullString
	var description sql.NullString
	var founders sql.NullString
	var category sql.NullString
	var stage sql.NullString
	var twitter sql.NullString
	var facebook sql.NullString
	var angellist sql.NullString
	var linkedin sql.NullString
	var instagram sql.NullString
	var youtube sql.NullString
	var googleplus sql.NullString
	row := GetDB().QueryRow("SELECT name,contact,email,website,short, description,logo, primarypic, founders, category, stage, twitter, facebook, angellist, linkedin, instagram, youtube, googleplus, status FROM companies WHERE name=?", name)
	err := row.Scan(&c.Name, &c.Contact, &c.Email, &c.Website, &short, &description, &logo, &primarypic, &founders, &category, &stage, &twitter, &facebook, &angellist, &linkedin, &instagram, &youtube, &googleplus, &c.Status)
	if err != nil {
		log.Println(err.Error())
		return c
	}
	if short.Valid {
		c.Short = short.String
	}
	if description.Valid {
		c.Description = description.String
	}
	if logo.Valid {
		c.Logo = logo.String
	}
	if primarypic.Valid {
		c.PrimaryPic = primarypic.String
	}
	if founders.Valid {
		c.Founders = founders.String
	}
	if category.Valid {
		c.Category = category.String
	}
	if stage.Valid {
		c.Stage = stage.String
	}
	if twitter.Valid {
		c.Twitter = twitter.String
	}
	if facebook.Valid {
		c.Facebook = facebook.String
	}
	if angellist.Valid {
		c.AngelList = angellist.String
	}
	if linkedin.Valid {
		c.LinkedIn = linkedin.String
	}
	if instagram.Valid {
		c.Instagram = instagram.String
	}
	if youtube.Valid {
		c.Youtube = youtube.String
	}
	if googleplus.Valid {
		c.GooglePlus = googleplus.String
	}
	return c
}

func GetMyCompany(email string) Company {
	var c Company
	var logo sql.NullString
	var primarypic sql.NullString
	var short sql.NullString
	var description sql.NullString
	var founders sql.NullString
	var category sql.NullString
	var stage sql.NullString
	var twitter sql.NullString
	var facebook sql.NullString
	var angellist sql.NullString
	var linkedin sql.NullString
	var instagram sql.NullString
	var youtube sql.NullString
	var googleplus sql.NullString
	row := GetDB().QueryRow("SELECT name,contact,email,website,short,description,logo,primarypic, founders, category, stage, twitter, facebook, angellist, linkedin, instagram, youtube, googleplus FROM companies WHERE email=?", email)
	err := row.Scan(&c.Name, &c.Contact, &c.Email, &c.Website, &short, &description, &logo, &primarypic, &founders, &category, &stage, &twitter, &facebook, &angellist, &linkedin, &instagram, &youtube, &googleplus)
	if err != nil {
		log.Println(err.Error())
		return c
	}
	if short.Valid {
		c.Short = short.String
	}
	if description.Valid {
		c.Description = description.String
	}
	if logo.Valid {
		c.Logo = logo.String
	}
	if primarypic.Valid {
		c.PrimaryPic = primarypic.String
	}
	if founders.Valid {
		c.Founders = founders.String
	}
	if category.Valid {
		c.Category = category.String
	}
	if stage.Valid {
		c.Stage = stage.String
	}
	if twitter.Valid {
		c.Twitter = twitter.String
	}
	if facebook.Valid {
		c.Facebook = facebook.String
	}
	if angellist.Valid {
		c.AngelList = angellist.String
	}
	if linkedin.Valid {
		c.LinkedIn = linkedin.String
	}
	if instagram.Valid {
		c.Instagram = instagram.String
	}
	if youtube.Valid {
		c.Youtube = youtube.String
	}
	if googleplus.Valid {
		c.GooglePlus = googleplus.String
	}
	return c
}

func GetMyCompanies(advisor string) []Company {
	rows, err := GetDB().Query("SELECT name,logo,primarypic,short FROM companies WHERE name in (SELECT company FROM teams WHERE advisor=?) order by name asc", advisor)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err.Error())
		}
		return nil
	}
	defer rows.Close()

	var cSet []Company
	for rows.Next() {
		var c Company
		var logo sql.NullString
		var primarypic sql.NullString
		var short sql.NullString
		err := rows.Scan(&c.Name, &logo, &primarypic, &short)
		if err != nil {
			log.Println(err.Error())
			break
		}
		if logo.Valid {
			c.Logo = logo.String
		}
		if c.Logo == "" {
			c.Logo = "/empty.png"
		}
		if primarypic.Valid {
			c.PrimaryPic = primarypic.String
		}
		if c.PrimaryPic == "" {
			c.PrimaryPic = "/empty.png"
		}
		if short.Valid {
			c.Short = short.String
		}
		cSet = append(cSet, c)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return cSet
}

func GetCompanies() []Company {
	rows, err := GetDB().Query("SELECT name,contact,email,website,logo,primarypic,short,description FROM companies WHERE status<>'pending'")
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()

	var cSet []Company
	for rows.Next() {
		var c Company
		var logo sql.NullString
		var primarypic sql.NullString
		var short sql.NullString
		var description sql.NullString
		err := rows.Scan(&c.Name, &c.Contact, &c.Email, &c.Website, &logo, &primarypic, &short, &description)
		if err != nil {
			log.Println(err.Error())
			break
		}
		if logo.Valid {
			c.Logo = logo.String
		}
		if c.Logo == "" {
			c.Logo = "/empty.png"
		}
		if primarypic.Valid {
			c.PrimaryPic = primarypic.String
		}
		if c.PrimaryPic == "" {
			c.PrimaryPic = "/empty.png"
		}
		if short.Valid {
			c.Short = short.String
		}
		if description.Valid {
			c.Description = description.String
		}
		cSet = append(cSet, c)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return cSet
}

func GetCompaniesSorted() []Company {
	rows, err := GetDB().Query("SELECT name,contact,email,website,logo,primarypic,short,description FROM companies WHERE status<>'pending' ORDER BY showorder DESC")
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()

	var cSet []Company
	for rows.Next() {
		var c Company
		var logo sql.NullString
		var primarypic sql.NullString
		var short sql.NullString
		var description sql.NullString
		err := rows.Scan(&c.Name, &c.Contact, &c.Email, &c.Website, &logo, &primarypic, &short, &description)
		if err != nil {
			log.Println(err.Error())
			break
		}
		if logo.Valid {
			c.Logo = logo.String
		}
		if c.Logo == "" {
			c.Logo = "/empty.png"
		}
		if primarypic.Valid {
			c.PrimaryPic = primarypic.String
		}
		if c.PrimaryPic == "" {
			c.PrimaryPic = "/empty.png"
		}
		if short.Valid {
			c.Short = short.String
		}
		if description.Valid {
			c.Description = description.String
		}
		cSet = append(cSet, c)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return cSet
}

func GetCompaniesAdmin() []Company {
	rows, err := GetDB().Query("SELECT name,contact,email,website,logo,primarypic,short,description,status FROM companies")
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()

	var cSet []Company
	for rows.Next() {
		var c Company
		var logo sql.NullString
		var primarypic sql.NullString
		var short sql.NullString
		var description sql.NullString
		err := rows.Scan(&c.Name, &c.Contact, &c.Email, &c.Website, &logo, &primarypic, &short, &description, &c.Status)
		if err != nil {
			log.Println(err.Error())
			break
		}
		if logo.Valid {
			c.Logo = logo.String
		}
		if c.Logo == "" {
			c.Logo = "/empty.png"
		}
		if primarypic.Valid {
			c.PrimaryPic = primarypic.String
		}
		if c.PrimaryPic == "" {
			c.PrimaryPic = "/empty.png"
		}
		if short.Valid {
			c.Short = short.String
		}
		if description.Valid {
			c.Description = description.String
		}
		cSet = append(cSet, c)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return cSet
}

func EditCompany(editName string, c Company) error {
	update, err := GetDB().Query("UPDATE companies SET contact=?,email=?,website=?,logo=?,primarypic=?,short=?,founders=?,category=?,stage=?,twitter=?,facebook=?,angellist=?,linkedin=?,instagram=?,youtube=?,googleplus=?,status=? WHERE name=?",
		c.Contact, c.Email, c.Website, c.Logo, c.PrimaryPic, truncStr(c.Short), c.Founders, c.Category, c.Stage, c.Twitter, c.Facebook, c.AngelList, c.LinkedIn, c.Instagram, c.Youtube, c.GooglePlus, c.Status, editName)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	update.Close()
	if c.Description != "" {
		update2, err := GetDB().Query("UPDATE companies SET description=? WHERE name=?", truncStr(c.Description), c.Name)
		if err != nil {
			log.Println(err.Error())
			return err
		}
		update2.Close()
	}
	return nil
}

func DeleteCompany(name string) error {
	delete, err := GetDB().Query("DELETE FROM companies WHERE name=? LIMIT 1", name)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	delete.Close()
	return nil
}

func DeleteQuestion(company string, q int) error {
	delete, err := GetDB().Query("DELETE FROM questions WHERE company=? AND id=? LIMIT 1", company, q)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	delete.Close()
	return nil
}

func EditQuestion(q Question) error {
	update, err := GetDB().Query("UPDATE questions SET question=?,body=?,reward=? WHERE company=? AND id=? LIMIT 1", q.Question, q.Body, q.Reward, q.Company, q.ID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	update.Close()
	return nil
}

func AddNewQuestion(q Question) error {
	insert, err := GetDB().Query("INSERT INTO questions (question, company, body, reward, private) VALUES (?,?,?,?,?)", q.Question, q.Company, q.Body, q.Reward, q.Private)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	insert.Close()
	return nil
}

func DeleteResponse(r Response) error {
	answer := "*deleted by user*"
	delete, err := GetDB().Query("UPDATE responses SET answer=? WHERE id=? AND question_id=? AND advisor=? LIMIT 1", answer, r.ReplyID, r.QuestionID, r.Advisor)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	delete.Close()
	return nil
}

func EditResponse(r Response) error {
	update, err := GetDB().Query("UPDATE responses SET answer=? WHERE id=? AND question_id=? AND advisor=? LIMIT 1", r.Answer, r.ReplyID, r.QuestionID, r.Advisor)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	update.Close()
	return nil
}

func AddNewResponse(r Response) error {
	insert, err := GetDB().Query("INSERT INTO responses (question_id, advisor, answer, likes, reply_id) VALUES (?,?,?,?,?)", r.QuestionID, r.Advisor, r.Answer, r.Likes, r.ReplyID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	insert.Close()
	return nil
}

func DeleteComment(c Comment) error {
	comment := "*deleted by user*"
	delete, err := GetDB().Query("UPDATE comments SET comment=? WHERE id=? AND company=? AND advisor=? LIMIT 1", comment, c.ID, c.Company, c.Advisor)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	delete.Close()
	return nil
}

func EditComment(c Comment) error {
	update, err := GetDB().Query("UPDATE comments SET comment=? WHERE id=? AND company=? AND advisor=? LIMIT 1", c.Comment, c.ID, c.Company, c.Advisor)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	update.Close()
	return nil
}

func AddNewComment(c Comment) error {
	insert, err := GetDB().Query("INSERT INTO comments (company, advisor, comment, likes, reply_id) VALUES (?,?,?,?,?)", c.Company, c.Advisor, c.Comment, c.Likes, c.ReplyID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	insert.Close()
	return nil
}

func JoinTeam(email, company string) error {
	insert, err := GetDB().Query("INSERT INTO teams (advisor,company) VALUES (?,?)", email, company)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	insert.Close()
	return nil
}

func IsOnTeam(email, company string) bool {
	var advisor string
	row := GetDB().QueryRow("SELECT advisor FROM teams WHERE advisor=? AND company=?", email, company)
	err := row.Scan(&advisor)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err.Error())
		}
		return false
	}
	return true
}

func GetTeamMembers(company string) []string {
	rows, err := GetDB().Query("SELECT users.name FROM users JOIN teams ON users.email=teams.advisor WHERE company=?", company)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err.Error())
		}
		return nil
	}
	defer rows.Close()

	var advisors []string
	for rows.Next() {
		var a string
		err := rows.Scan(&a)
		if err != nil {
			log.Println(err.Error())
			break
		}
		advisors = append(advisors, a)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return advisors
}

func GetTeamMembersFull(company string) []User {
	rows, err := GetDB().Query("SELECT users.name, users.email FROM users JOIN teams ON users.email=teams.advisor WHERE company=?", company)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err.Error())
		}
		return nil
	}
	defer rows.Close()

	var advisors []User
	var advisor User
	for rows.Next() {
		var a string
		var e string
		err := rows.Scan(&a, &e)
		if err != nil {
			log.Println(err.Error())
			break
		}
		advisor = User{
			Name:  a,
			Email: e,
		}
		advisors = append(advisors, advisor)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return advisors
}

func GetCompanyReward(company string, q int) string {
	var reward string
	row := GetDB().QueryRow("SELECT reward FROM questions WHERE company=? AND id=?", company, q)
	err := row.Scan(&reward)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return reward
}

func GetCompanyName(email string) string {
	var name string
	row := GetDB().QueryRow("SELECT name FROM companies WHERE email=?", email)
	err := row.Scan(&name)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return name
}

func GetCompanyQuestions(company string) []Question {
	rows, err := GetDB().Query("SELECT id, question, company, body, reward, private from questions WHERE company=? AND private=0 order by id desc", company)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err.Error())
		}
		return nil
	}
	defer rows.Close()

	var questions []Question
	for rows.Next() {
		var q Question
		var qBody sql.NullString
		var qReward sql.NullString
		var qPrivate sql.NullBool
		err := rows.Scan(&q.ID, &q.Question, &q.Company, &qBody, &qReward, &qPrivate)
		if err != nil {
			log.Println(err.Error())
			break
		}
		if qBody.Valid {
			q.Body = qBody.String
		}
		if qReward.Valid {
			q.Reward = qReward.String
		}
		if qPrivate.Valid {
			q.Private = qPrivate.Bool
		}
		questions = append(questions, q)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return questions
}

func GetCompanyTeamQuestions(company string) []Question {
	rows, err := GetDB().Query("SELECT id, question, company, body, reward, private from questions WHERE company=? AND private=1 order by id desc", company)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()

	var questions []Question
	for rows.Next() {
		var q Question
		var qBody sql.NullString
		var qReward sql.NullString
		var qPrivate sql.NullBool
		err := rows.Scan(&q.ID, &q.Question, &q.Company, &qBody, &qReward, &qPrivate)
		if err != nil {
			log.Println(err.Error())
			break
		}
		if qBody.Valid {
			q.Body = qBody.String
		}
		if qReward.Valid {
			q.Reward = qReward.String
		}
		if qPrivate.Valid {
			q.Private = qPrivate.Bool
		}
		questions = append(questions, q)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return questions
}

func GetCompanyTopQuestion(company string) Question {
	var q Question
	var reward sql.NullString
	row := GetDB().QueryRow("SELECT id, question, body, reward FROM questions WHERE company=? and private=0 order by id desc", company)
	err := row.Scan(&q.ID, &q.Question, &q.Body, &reward)
	if err != nil {
		log.Println(err.Error())
		return Question{}
	}
	if reward.Valid {
		q.Reward = reward.String
	}
	return q
}

func GetCompanyQuestion(company string, qid int) Question {
	var q Question
	var reward sql.NullString
	row := GetDB().QueryRow("SELECT id, company, question, body, reward FROM questions WHERE company=? and id=?", company, qid)
	err := row.Scan(&q.ID, &q.Company, &q.Question, &q.Body, &reward)
	if err != nil {
		log.Println(err.Error())
		return Question{}
	}
	if reward.Valid {
		q.Reward = reward.String
	}
	return q
}

func IsCompanyQuestionValid(company string, qid int) bool {
	var c sql.NullString
	row := GetDB().QueryRow("SELECT company FROM questions WHERE id=? AND company=?", qid, company)
	err := row.Scan(&c)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func GetResponseAnswer(rid int, qid int, email string) string {
	var answer sql.NullString
	row := GetDB().QueryRow("SELECT answer from responses WHERE id=? AND question_id=? AND advisor=?", rid, qid, email)
	err := row.Scan(&answer)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return answer.String
}

func GetResponses(questionID int) []ResponseWithUser {
	rows, err := GetDB().Query("SELECT r.id, r.question_id, r.advisor, r.answer, r.likes, r.reply_id, u.name, u.title, u.job, u.experience FROM responses r JOIN users u ON r.advisor=u.email WHERE r.question_id=? order by id desc", questionID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err.Error())
		}
		return nil
	}
	defer rows.Close()

	var responses []ResponseWithUser
	for rows.Next() {
		var r ResponseWithUser
		var rTitle sql.NullString
		var rJob sql.NullString
		var rExp sql.NullString

		err := rows.Scan(&r.ID, &r.QuestionID, &r.Advisor, &r.Answer, &r.Likes, &r.ReplyID, &r.Name, &rTitle, &rJob, &rExp)
		if err != nil {
			log.Println(err.Error())
			break
		}
		if rTitle.Valid {
			r.Title = rTitle.String
		}
		if rJob.Valid {
			r.Job = rJob.String
		}
		if rExp.Valid {
			r.Experience = rExp.String
		}
		responses = append(responses, r)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return responses
}

func AddLike(id int) {
	stmt, err := GetDB().Prepare("UPDATE responses SET likes=likes+1 WHERE id=?")
	if err != nil {
		return
	}
	stmt.Exec(id)
}

func AddLikeComment(id int) {
	stmt, err := GetDB().Prepare("UPDATE comments SET likes=likes+1 WHERE id=?")
	if err != nil {
		return
	}
	stmt.Exec(id)
}

func GetActivitiesForUser(email string) (responses []Activity) {
	var rows *sql.Rows
	var err error
	if IsAdvisor(email) {
		rows, err = GetDB().Query("SELECT activity, time, user, company FROM activity ORDER BY time desc LIMIT 100")
		if err != nil {
			log.Println(err.Error())
			return nil
		}
	} else if IsClient(email) {
		rows, err = GetDB().Query("SELECT activity.activity, activity.time, activity.user, activity.company FROM activity JOIN companies ON activity.company=companies.name WHERE companies.email=? ORDER BY time desc LIMIT 100", email)
		if err != nil {
			log.Println(err.Error())
			return nil
		}
	} else {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var a Activity
		err := rows.Scan(&a.Activity, &a.Time, &a.User, &a.Company)
		if err != nil {
			log.Println(err.Error())
			break
		}
		tval, err := time.Parse("2006-01-02 15:04:05", a.Time)
		if err != nil {
			log.Println(err.Error())
			return nil
		}
		uswest, _ := time.LoadLocation("America/Los_Angeles")
		tval = tval.In(uswest)
		a.Time = tval.Format("Jan 2, 3:04PM")
		responses = append(responses, a)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return responses
}

func AddActivity(activity, email, company string) error {
	insert, err := GetDB().Query("INSERT INTO activity (activity, user, company) VALUES (?,?,?)", activity, email, company)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	insert.Close()
	return nil
}

func GetDiscussions() []Question {
	rows, err := GetDB().Query("SELECT id, company, question FROM questions WHERE private=0 ORDER BY id desc limit 100")
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()

	var qSet []Question
	for rows.Next() {
		var q Question
		err := rows.Scan(&q.ID, &q.Company, &q.Question)
		if err != nil {
			log.Println(err.Error())
			break
		}
		qSet = append(qSet, q)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return qSet
}

func GetAdvisorTeamsQuestions(advisor string) []Question {
	rows, err := GetDB().Query("SELECT q.id, q.company, q.question FROM questions q JOIN teams t ON q.company=t.company WHERE t.advisor=? AND q.private=1 ORDER BY q.id desc limit 100", advisor)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err.Error())
		}
		return nil
	}
	defer rows.Close()

	var qSet []Question
	for rows.Next() {
		var q Question
		err := rows.Scan(&q.ID, &q.Company, &q.Question)
		if err != nil {
			log.Println(err.Error())
			break
		}
		qSet = append(qSet, q)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return qSet
}

func GetCompanyComments(company string) []CommentWithUser {
	rows, err := GetDB().Query("SELECT c.id, c.company, c.advisor, c.comment, c.likes, c.reply_id, u.name, u.title, u.job, u.experience FROM comments c JOIN users u ON c.advisor=u.email WHERE c.company=? order by c.id desc", company)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err.Error())
		}
		return nil
	}
	defer rows.Close()

	var comments []CommentWithUser
	for rows.Next() {
		var c CommentWithUser
		var cTitle sql.NullString
		var cJob sql.NullString
		var cExp sql.NullString

		err := rows.Scan(&c.ID, &c.Company, &c.Advisor, &c.Comment, &c.Likes, &c.ReplyID, &c.Name, &cTitle, &cJob, &cExp)
		if err != nil {
			log.Println(err.Error())
			break
		}
		if cTitle.Valid {
			c.Title = cTitle.String
		}
		if cJob.Valid {
			c.Job = cJob.String
		}
		if cExp.Valid {
			c.Experience = cExp.String
		}
		comments = append(comments, c)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return comments
}

func GetCompanyComment(cid int, company string, advisor string) string {
	var comment sql.NullString
	row := GetDB().QueryRow("SELECT comment from comments WHERE id=? AND company=? AND advisor=?", cid, company, advisor)
	err := row.Scan(&comment)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return comment.String
}

func GetTopThreads(n int) []ThreadWithUser {
	rows, err := GetDB().Query("SELECT d.id,d.title,d.advisor,d.url,d.count,d.created,u.name FROM discussions d JOIN users u ON d.advisor=u.email order by id desc limit ?", n)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err.Error())
		}
		return nil
	}
	defer rows.Close()

	var threads []ThreadWithUser
	for rows.Next() {
		var t ThreadWithUser
		err := rows.Scan(&t.ID, &t.Title, &t.Advisor, &t.Url, &t.Count, &t.Created, &t.Name)
		if err != nil {
			log.Println(err.Error())
			break
		}
		t.TimeAgo = GetTimeAgo(t.Created)
		t.IsCompanyLink = strings.HasPrefix(t.Url, "/company?name=")
		threads = append(threads, t)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return threads
}

func AddNewThread(t Thread) error {
	insert, err := GetDB().Query("INSERT INTO discussions (title,advisor,url,body) VALUES (?,?,?,?)", t.Title, t.Advisor, t.Url, t.Body)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	insert.Close()
	return nil
}

func GetThread(tid int) ThreadWithUser {
	var t ThreadWithUser
	row := GetDB().QueryRow("SELECT d.id,d.title,d.advisor,d.count,d.created,d.url,d.body,u.name FROM discussions d JOIN users u ON d.advisor=u.email WHERE d.id=?", tid)
	err := row.Scan(&t.ID, &t.Title, &t.Advisor, &t.Count, &t.Created, &t.Url, &t.Body, &t.Name)
	if err != nil {
		log.Println(err.Error())
		return ThreadWithUser{}
	}
	t.TimeAgo = GetTimeAgo(t.Created)
	return t
}

func AddNewThreadResponse(t ThreadResponse) error {
	insert, err := GetDB().Query("INSERT INTO threads (discussion_id, reply_id, advisor, comment) VALUES (?,?,?,?)", t.DiscussionID, t.ReplyID, t.Advisor, t.Comment)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	insert.Close()
	insert, err = GetDB().Query("UPDATE discussions SET count=count+1 WHERE id=?", t.DiscussionID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	insert.Close()
	return nil
}

func GetThreadComments(tid int) []ThreadResponseWithUser {
	rows, err := GetDB().Query("SELECT r.id,r.discussion_id,r.reply_id,r.advisor,r.comment,r.created,u.name from threads r join users u on r.advisor=u.email where r.discussion_id=? order by id asc", tid)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err.Error())
		}
		return nil
	}
	defer rows.Close()

	var comments []ThreadResponseWithUser
	for rows.Next() {
		var c ThreadResponseWithUser
		err := rows.Scan(&c.ID, &c.DiscussionID, &c.ReplyID, &c.Advisor, &c.Comment, &c.Created, &c.Name)
		if err != nil {
			log.Println(err.Error())
			break
		}
		c.TimeAgo = GetTimeAgo(c.Created)
		comments = append(comments, c)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return comments
}

func GetThreadComment(id int, advisor string) ThreadResponse {
	var t ThreadResponse
	row := GetDB().QueryRow("SELECT id,discussion_id,reply_id,advisor,comment,created from threads WHERE id=? AND advisor=?", id, advisor)
	err := row.Scan(&t.ID, &t.DiscussionID, &t.ReplyID, &t.Advisor, &t.Comment, &t.Created)
	if err != nil {
		log.Println(err.Error())
	}
	t.TimeAgo = GetTimeAgo(t.Created)
	return t
}

func DeleteThreadResponse(t ThreadResponse) error {
	comment := "*deleted by user*"
	delete, err := GetDB().Query("UPDATE threads SET comment=? WHERE id=? AND advisor=? LIMIT 1", comment, t.ID, t.Advisor)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	delete.Close()
	return nil
}

func EditThreadResponse(t ThreadResponse) error {
	update, err := GetDB().Query("UPDATE threads SET comment=? WHERE id=? AND advisor=? LIMIT 1", t.Comment, t.ID, t.Advisor)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	update.Close()
	return nil
}
