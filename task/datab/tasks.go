package datab

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskbucket = []byte("tasks")
var db *bolt.DB

//Task is a struct ...
type Task struct {
	key   int
	value string
}

//Init is used to open a database
func Init(dbpath string) error {
	var err error
	db, err = bolt.Open(dbpath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskbucket)
		return err
	})
}

//CreateTask is a function used to create tasks in database
func CreateTask(Task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskbucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(Task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return (int(binary.BigEndian.Uint64(b)))
}

//Alltasks return the slice of tasks
func Alltasks() ([]Task, error) {
	var task []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskbucket)

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			task = append(task, Task{
				key:   btoi(k),
				value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return task, nil
}

//Deletetask deletes the corresponding key and value pairsS
func Deletetask(key int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskbucket)
		return b.Delete(itob(key))

	})
	return err
}
