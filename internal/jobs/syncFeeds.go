package jobs

import (
	"log"
	"sync"
	"time"

	"github.com/nathanberry97/rss2go/internal/database"
	"github.com/nathanberry97/rss2go/internal/rss"
	"github.com/nathanberry97/rss2go/internal/schema"
	"github.com/nathanberry97/rss2go/internal/services"
)

func SyncFeeds(workers int) {
	go runFeedUpdate(workers)

	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		runFeedUpdate(workers)
	}
}

func runFeedUpdate(workers int) {
	jobs := make(chan schema.Task, 10)
	var wg sync.WaitGroup

	for range workers {
		wg.Add(1)
		go worker(jobs, &wg)
	}

	conn := database.DatabaseConnection()
	defer conn.Close()

	feeds, err := services.GetFeeds(conn)
	if err != nil {
		log.Printf("Error getting feeds: %v", err)
		return
	}

	log.Printf("Starting feed update for %d feeds with %d workers", len(feeds), workers)

	for _, url := range feeds {
		jobs <- schema.Task{
			FeedId: int64(url.ID),
			URL:    url.URL,
			Conn:   conn,
		}
	}

	close(jobs)
	wg.Wait()
	log.Println("Feed update completed")
}

func worker(jobs <-chan schema.Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		_, articles, err := rss.FeedHandler(job.URL)
		if err != nil {
			log.Printf("Error handling feed %s: %v", job.URL, err)
			continue
		}

		if err := services.InsertArticles(job.Conn, articles, job.FeedId); err != nil {
			log.Printf("Error inserting articles for feed %s: %v", job.URL, err)
		}
	}
}
