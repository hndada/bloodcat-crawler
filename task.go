package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	errDownload = -(iota + 1)
	mapExist
	mapBanned
)

type job struct {
	result int
	id     int
	artist string
	title  string
}

func doTask(tasks []map[string]interface{}, i int) job {
	v := tasks[i]["id"].(string)
	setID, err := strconv.Atoi(v)
	artist, title := tasks[i]["artist"].(string), tasks[i]["title"].(string)

	res := job{i, setID, artist, title}
	switch {
	case err != nil:
		res.result = errDownload
	case exist[setID]:
		res.result = mapExist
	case ban[setID]:
		res.result = mapBanned
	default:
		if err = download(setID, artist, title); err != nil {
			res.result = errDownload
		}
	}
	return res
}

func download(setID int, creator, title string) error {
	dl := fmt.Sprintf("https://bloodcat.com/osu/s/%d", setID)
	resp, err := http.Get(dl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fname := getFname(setID, creator, title)
	out, err := os.Create(filepath.Join(dir, fname))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func printResult(res job) {
	fname := getFname(res.id, res.artist, res.title)
	switch res.result {
	case errDownload:
		fmt.Printf("error: %s\n", fname)
	case mapExist:
		fmt.Printf("exist: %s\n", fname)
	case mapBanned:
		fmt.Printf("banned: %s\n", fname)
	default:
		fmt.Printf("success: %s\n", fname)
	}
}

func getFname(setID int, artist, title string) string {
	name := fmt.Sprintf("%d %s - %s.osz", setID, artist, title)
	for _, letter := range []string{"<", ">", ":", "\"", "/", "\\", "|", "?", "*"} {
		name = strings.ReplaceAll(name, letter, "-")
	}
	return name
}
