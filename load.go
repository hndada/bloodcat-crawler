package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func loadConfig() {
	f, err := os.Open("config.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var values []string
	var text string
	raw := make(map[string]bool)
	for scanner.Scan() {
		text = scanner.Text()
		if text == "" || strings.HasPrefix(text, "//") {
			continue
		}
		values = strings.SplitN(text, ":", 2)
		if len(values) < 2 {
			continue
		}
		switch values[0] {
		case "Songs":
			dir = values[1]
			if _, err := os.Stat(dir); err != nil {
				log.Fatalf("invalid: %s is not a valid directory.", dir)
			}
		case "Search":
			search = values[1]
		default:
			switch values[1] {
			case "1":
				raw[values[0]] = true
			case "0":
				raw[values[0]] = false
			default:
				log.Fatal("invalid: value should be 1 or 0.")
			}
		}
	}

	for i, mode := range []string{"Standard", "Taiko", "Catch", "Mania"} {
		if on, ok := raw[mode]; on && ok {
			modes = append(modes, i)
		}
	}
	for i, stat := range []string{"Unranked", "Ranked", "Approved", "Qualified", "Loved"} {
		if on, ok := raw[stat]; on && ok {
			stats = append(stats, i)
		}
	}
}

func loadExist(root string) map[int]bool {
	fs, err := ioutil.ReadDir(root)
	check(err)
	mapSetIDs := make(map[int]bool)
	var s string
	var id int
	for _, f := range fs {
		s = strings.SplitN(f.Name(), " ", 2)[0]
		id, err = strconv.Atoi(s)
		if err != nil {
			continue
		}
		mapSetIDs[id] = true
	}
	return mapSetIDs
}

func loadBan() map[int]bool {
	var s string
	var id int
	ban := make(map[int]bool)

	f, err := os.Open("ban.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s = scanner.Text()
		if s == "" {
			continue
		}
		id, err = strconv.Atoi(s)
		if err != nil {
			continue
		}
		ban[id] = true
	}
	return ban
}
