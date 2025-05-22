package worker

import (
	"fmt"
	"sync"
	"time"

	"github.com/nathanberry97/rss2go/internal/database"
	"github.com/nathanberry97/rss2go/internal/rss"
	"github.com/nathanberry97/rss2go/internal/schema"
	"github.com/nathanberry97/rss2go/internal/services"
)

func ScheduleFeedUpdates(workers int) {
	go runFeedUpdate(workers)

	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			runFeedUpdate(workers)
		}
	}
}

func runFeedUpdate(workers int) {
	jobs := make(chan schema.Task, 10)
	var wg sync.WaitGroup

	for w := 1; w <= workers; w++ {
		wg.Add(1)
		go worker(jobs, &wg)
	}

	conn := database.DatabaseConnection()
	defer conn.Close()

	feeds, err := services.GetFeeds(conn)
	if err != nil {
		fmt.Println(err)
	}

	for _, url := range feeds {
		jobs <- schema.Task{
			FeedId: int64(url.ID),
			URL:    url.URL,
			Conn:   conn,
		}
	}

	close(jobs)
	wg.Wait()
}

func worker(jobs <-chan schema.Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		_, articles, err := rss.FeedHandler(job.URL)
		if err != nil {
			fmt.Println(err)
		}

		services.InsertArticles(job.Conn, articles, job.FeedId)
	}
}
