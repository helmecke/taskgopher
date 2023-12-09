package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/olivere/ndjson"

	"github.com/helmecke/taskgopher/internal/task"
)

const mode = 0o644

// A FileStorage loads and saves tasks to file
type FileStorage struct {
	Location  string
	Pending   string
	Completed string
}

// NewFileStorage creates a new file store
func NewFileStorage(location string) *FileStorage {
	return &FileStorage{
		Location:  location,
		Pending:   "pending.data",
		Completed: "completed.data",
	}
}

// Create creates data directory and files
func (f *FileStorage) Create() {
	if err := os.MkdirAll(f.Location, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	_, err := os.OpenFile(f.pending(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, mode)
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.OpenFile(f.completed(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, mode)
	if err != nil {
		log.Fatal(err)
	}
}

// Load loads tasks from file
func (f *FileStorage) Load(all bool) (tasks []*task.Item) {
	file, err := os.OpenFile(f.pending(), os.O_RDONLY, mode)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error when closing:", err)
		}
	}()

	r := ndjson.NewReader(file)

	for i := 1; r.Next(); i++ {
		var task *task.Item
		if err := r.Decode(&task); err != nil {
			log.Panicf("Decode failed: %v", err)

			return
		}
		task.ID = i
		tasks = append(tasks, task)
	}

	if err := r.Err(); err != nil {
		log.Panicf("Reader failed: %v", err)

		return
	}

	if all {
		file, err := os.OpenFile(f.completed(), os.O_RDONLY, mode)
		if err != nil {
			log.Panic(err)
		}
		defer func() {
			if err := file.Close(); err != nil {
				fmt.Println("Error when closing:", err)
			}
		}()

		r := ndjson.NewReader(file)

		for r.Next() {
			var task *task.Item
			if err := r.Decode(&task); err != nil {
				log.Panicf("Decode failed: %v", err)

				return
			}

			tasks = append(tasks, task)
		}

		if err := r.Err(); err != nil {
			log.Panicf("Reader failed: %v", err)

			return
		}
	}

	return tasks
}

// Save saves tasks to file
func (f *FileStorage) Save(tasks []*task.Item) {
	file, err := os.OpenFile(f.pending(), os.O_CREATE|os.O_WRONLY, mode)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error when closing:", err)
		}
	}()

	if err := file.Truncate(0); err != nil {
		log.Panic(err)
	}

	r := ndjson.NewWriter(file)

	for _, task := range tasks {
		if err := r.Encode(task); err != nil {
			log.Panicf("Encode failed: %v", err)

			return
		}
	}
}

// Complete completes task
// TODO: this does not belong here
func (f *FileStorage) Complete(task *task.Item) {
	file, err := os.OpenFile(f.completed(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, mode)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error when closing:", err)
		}
	}()

	r := ndjson.NewWriter(file)

	if err := r.Encode(task); err != nil {
		log.Panicf("Encode failed: %v", err)

		return
	}
}

func (f *FileStorage) pending() string {
	return f.Location + "/" + f.Pending
}

func (f *FileStorage) completed() string {
	return f.Location + "/" + f.Completed
}
