package main

import (
	"database/sql"
	"fmt"
)

type Event struct {
	Task string
	Time string
}

var (
	db, _ = sql.Open("sqlite3", "cache/schedule.db")
	createDailyDB = "CREATE TABLE IF NOT EXISTS daily (ID INTEGER PRIMARY KEY AUTOINCREMENT, task TEXT NOT NULL, time_elapsed TIME)"
	createWeeklyDB = "CREATE TABLE IF NOT EXISTS weekly (ID INTEGER PRIMARY KEY AUTOINCREMENT, task TEXT NOT NULL, wday TINYINT NOT NULL, time TIME, time_elapsed TIME)"
	createOnceDB = "CREATE TABLE IF NOT EXISTS once (ID INTEGER PRIMARY KEY AUTOINCREMENT, task TEXT NOT NULL, date DATE NOT NULL, time TIME, time_elapsed TIME)"
)

func getDailyTasks() (*Event, error) {
	var task string
	var time string
	q, err := db.Query("SELECT task, time_elapsed FROM daily")
	if err != nil {
		return nil, err
	}
	for q.Next() {
		q.Scan(&task, &time)
	}
	return &Event{Task: task, Time: time}, nil
}

func main() {
	event, _ := getDailyTasks()
	fmt.Println(event.Task)
}