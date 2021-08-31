package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
	// For each file.IsRegular initiate goRoutine(sum file and send result to c)
	resultQueue := make(chan result)
	errorQueue := make(chan error, 1)

	go func() {
		var wg sync.WaitGroup

		err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}
			// We found a regular file, let's work with it
			wg.Add(1)
			go func() {
				data, err := ioutil.ReadFile(path)
				select {
				case resultQueue <- result{
					path,
					md5.Sum(data),
					err,
				}:
				case <-done:

				}
				wg.Done()
			}()

			select {
			case <-done:
				return errors.New("Walk Cancelled")
			default:
				return nil
			}
		})
		go func() {
			wg.Wait()
			close(resultQueue)
		}()

		errorQueue <- err
	}()

	return resultQueue, errorQueue
}

func md5All(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

	resultQueue, errorQueue := sumFiles(done, root)

	m := make(map[string][md5.Size]byte)
	for r := range resultQueue {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}
	if err := <-errorQueue; err != nil {
		return nil, err
	}
	return m, nil
}

func main() {
	// Calculate the MD5 sum of all files under the specified directory,
	// then print the results sorted by path name.
	m, err := md5All(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
}
