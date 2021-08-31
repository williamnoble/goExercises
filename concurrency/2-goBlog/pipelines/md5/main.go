package main

import (
	"crypto/md5"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	// md5 returns a map of [string]md5Sum
	md5, err := md5All(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	var paths []string
	for md5Path := range md5 {
		paths = append(paths, md5Path)
	}

	sort.Strings(paths)
	{
		for _, p := range paths {
			fmt.Printf("%x %s\n", md5[p], p)
		}
	}

}

func md5All(root string) (map[string][md5.Size]byte, error) {
	m := make(map[string][md5.Size]byte)
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		m[path] = md5.Sum(data)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return m, nil
}
