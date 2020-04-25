package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	dir    string // osu! songs folder
	search string // search keywords
	exist  map[int]bool
	ban    map[int]bool
)

func main() {
	var modes, stats []int
	dir, modes, stats = loadConfig()
	exist = loadExist(dir)
	ban = loadBan()
	page := 1

	fmt.Printf("osu! Songs folder: %s\n", dir)
	fmt.Printf("Search keywords: %s\n", search)
	fmt.Printf("Modes: %+v (0: Standard 1: Taiko 2: Catch 3: Mania)\n", modes)
	fmt.Printf("Status: %+v (0: unranked 1: ranked 2: approved 3: qualified 4: loved)\n", stats)
	time.Sleep(time.Second)
	fmt.Print("\nStart Download? (y/n) ")

	var yes string
	_, err := fmt.Scan(&yes)
	check(err)
	yes = strings.TrimSpace(yes)
	yes = strings.ToLower(yes)
	fmt.Println(len(yes))
	if yes != "y" {
		os.Exit(1)
	}

	u, err := url.Parse("https://bloodcat.com/osu/")
	check(err)
	params := url.Values{}
	params.Add("mod", "json")
	params.Add("q", search)
	params.Add("m", joinInts(modes))
	params.Add("s", joinInts(stats))
	params.Add("p", strconv.Itoa(page))
	u.RawQuery = params.Encode()
	for {
		// request map list
		fmt.Printf("Trawling page %d...\n", page)
		tasks := make([]map[string]interface{}, 0, 100)
		resp, err := http.Get(u.String())
		check(err)
		js, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		check(err)
		err = json.Unmarshal(js, &tasks)
		check(err)
		if len(tasks) == 0 {
			break
		}

		// do the task
		for i := range tasks {
			res := doTask(tasks, i)
			printResult(res)
		}

		// ready for next request
		page++
		params.Set("p", strconv.Itoa(page))
		u.RawQuery = params.Encode()
	}
	fmt.Println("Finish!")
}
