```
package main

import (
	"fmt"
	"time"
)

/*
https://gobyexample.com/time
https://www.digitalocean.com/community/tutorials/how-to-use-dates-and-times-in-go
*/

func main() {
	p := fmt.Println

	// Getting time
	now := time.Now()                                    // current local time
	epoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC) // specific time

	// Parsing time from string (pkg.go.dev/time#pkg-constants)
	commitStr := "Sun Mar 2 20:47:34 2008 -0800" // research.swtch.com/govcs
	commit, _ := time.Parse("Mon Jan 2 15:04:05 2006 -0700", commitStr)
	commitStr = commit.Format(time.RFC3339Nano)
	commit, _ = time.Parse(time.RFC3339Nano, commitStr)

	// Formatting time (pkg.go.dev/time#pkg-constants)
	p("custom\t", epoch.Format("2.1.2006 03:04:05"))
	p("RFC3339\t", epoch.Format(time.RFC3339))
	p("Unix\t", epoch.Format(time.UnixDate))

	p("---")

	// Comparing two times
	p("epoch before commit?", epoch.Before(commit))
	p("commit after epoch?", commit.After(epoch))
	p("duration between commit and epoch:", commit.Sub(epoch))

	p("---")

	// Adding and substacting times
	p("now\t", now)
	p("now+1h\t", now.Add(time.Hour))
	p("now-1h\t", now.Add(-time.Hour)) // Sub is used to get diff
}
```
