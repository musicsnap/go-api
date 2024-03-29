package utils

import (
	"fmt"
	"go-api/middlewares"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

func GetHash(n int) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyz")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func MakeTempDir() (string, error) {
	dir := "./upload/" + GetHash(10)
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return "", err
	}
	return dir, nil
}

func GetFilenameExt(url string) string {
	_, filename := path.Split(url)
	return filename
}

func GetFilename(url string) string {
	_, filename := path.Split(url)
	i := strings.LastIndex(filename, ".")
	return filename[:i]
}

func GetFileExt(url string) string {
	_, filename := path.Split(url)
	i := strings.LastIndex(filename, ".")
	return filename[i+1:]
}

func DownloadRemoteFile(url string, saveDir string, style string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	filename := GetFilenameExt(url)

	if style != "" {
		filename = strings.Replace(filename, style, "", -1)
	}

	localPath := saveDir + "/" + filename
	out, err := os.Create(localPath) //不存在就删除
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return localPath, err
	}
	return localPath, nil
}

func Warning(err error, message string) {
	if err != nil {
		log.Println(err)
		log.Println(message)
		middlewares.PushError(message)
		log.Println(middlewares.GetErrors())
	}
}

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func Gethttp(url string, num int) string {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < num; i++ {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func() {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			resp, _ := http.Get(url)
			defer resp.Body.Close()
			//fmt.Println(resp.Status)
		}()
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
	fmt.Println(time.Now().Sub(start).Seconds())
	return ""
}
