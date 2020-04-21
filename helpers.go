package main

import (
	"fmt"
	"time"
)

func GetTimeAgo(m string) string {
	layout := "2006-01-02 15:04:05 MST"
	t, _ := time.Parse(layout, m+" PDT")
	delta := time.Since(t)
	hoursAgo := (int)(delta.Hours())
	if hoursAgo >= 48 {
		return fmt.Sprintf("%d days ago", hoursAgo/24)
	} else if hoursAgo >= 24 {
		return fmt.Sprintf("1 day ago")
	} else if hoursAgo > 1 {
		return fmt.Sprintf("%d hours ago", hoursAgo)
	} else if hoursAgo == 1 {
		return fmt.Sprintf("1 hour ago")
	} else {
		return fmt.Sprintf("few minutes ago")
	}
}
