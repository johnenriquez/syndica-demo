package main

import (
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	PasswordStr       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-"
	FailedLoginsMax   = 5
	FailedIPLoginsMax = 10
)

type LockItem struct {
	IP           string
	FailAttempts int
}

type LockStore map[string]*LockItem

type LockStoreIP map[string]int

var (
	LS   LockStore
	LSIP LockStoreIP
)

func init() {
	rand.Seed(time.Now().UnixNano())
	LS = make(LockStore)
	LSIP = make(LockStoreIP)
}

func GetRandomStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = PasswordStr[rand.Intn(len(PasswordStr))]
	}
	return string(b)
}

func GeneratePasswordHash(p string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return string(hash)
}

func IsPasswordHashValid(hash, p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	return err == nil
}

func IsIPLocked(ip string) bool {
	_, ok := LSIP[ip]
	if !ok {
		return false
	}
	return LSIP[ip] >= FailedIPLoginsMax
}

func LockIP(ip string) {
	if _, ok := LSIP[ip]; !ok {
		LSIP[ip] = 1
		return
	}
	LSIP[ip]++
	return
}

func UnlockIP(ip string) {
	delete(LSIP, ip)
}

func IsUserLocked(email string) bool {
	_, ok := LS[email]
	if !ok {
		return false
	}
	return LS[email].FailAttempts >= FailedLoginsMax
}

func LockUser(email, ip string) {
	if _, ok := LS[email]; !ok {
		LS[email] = &LockItem{ip, 1}
		return
	}
	LS[email].FailAttempts++
	log.Println("lock fails: ", email, LS[email].FailAttempts)
	return
}

func LockUserAdmin(email string) {
	if _, ok := LS[email]; !ok {
		LS[email] = &LockItem{"admin", FailedLoginsMax}
		return
	}
	LS[email].FailAttempts = FailedLoginsMax
	log.Println("admin locked user: ", email)
	return
}

func UnlockUser(email string) {
	delete(LS, email)
}

func UnlockUserAdmin(email string) {
	delete(LS, email)
	log.Println("admin unlocked user: ", email)
}
