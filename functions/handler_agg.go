package functions

import (
	"fmt"
	"time"
)

// input  | 1 param: time interval
// output | get the info of the next feed to fetch, prio null > oldest
func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.Arguments) < 1 {
		return fmt.Errorf("not enough arguments")
	}

	interval, err := time.ParseDuration(cmd.Arguments[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		fmt.Printf("[Interval: %s] Scraping... \n", cmd.Arguments[0])
		HandlerScrapeFeeds(s, cmd)
	}
}
