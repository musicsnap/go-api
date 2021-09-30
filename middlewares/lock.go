package middlewares

import (
	"github.com/gofrs/flock"
	"log"
)

func Lock(lockName string) (bool, *flock.Flock, error) {
	fileLock := flock.New("./locks/" + lockName + ".lock")

	locked, err := fileLock.TryLock()

	if err != nil {
		return false, nil, err
	}
	log.Println("lock ", lockName)
	return locked, fileLock, nil
}

func Unlock(l *flock.Flock) error {
	err := l.Unlock()
	log.Println("unlock")
	return err
}
