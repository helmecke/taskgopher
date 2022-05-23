package taskgopher

import (
	"log"
	"os"

	"github.com/olivere/ndjson"
)

const mode = 0644

// A FileStore loads and saves tasks to file
type fileStore struct {
	Location  string
	Pending   string
	Completed string
}

func newFileStore(location string) *fileStore {
	return &fileStore{
		Location:  location,
		Pending:   "pending.data",
		Completed: "completed.data",
	}
}

func (f *fileStore) init() {
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

func (f *fileStore) load(all bool) (tasks []*Task) {
	file, err := os.OpenFile(f.pending(), os.O_RDONLY, mode)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := ndjson.NewReader(file)

	for i := 1; r.Next(); i++ {
		var task *Task
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
		defer file.Close()

		r := ndjson.NewReader(file)

		for r.Next() {
			var task *Task
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

func (f *fileStore) save(tasks []*Task) {
	file, err := os.OpenFile(f.pending(), os.O_CREATE|os.O_WRONLY, mode)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

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

func (f *fileStore) complete(task *Task) {
	file, err := os.OpenFile(f.completed(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, mode)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := ndjson.NewWriter(file)

	if err := r.Encode(task); err != nil {
		log.Panicf("Encode failed: %v", err)

		return
	}
}

func (f *fileStore) pending() string {
	return f.Location + "/" + f.Pending
}

func (f *fileStore) completed() string {
	return f.Location + "/" + f.Completed
}
