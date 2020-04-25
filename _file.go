package main

// func loadExistMap(root string) map[int]bool {
// 	mapIDs := make(map[int]bool)
// 	var id int
// 	err := filepath.Walk(root,
// 		func(path string, info os.FileInfo, err error) error {
// 			if err != nil {
// 				return err
// 			}
// 			if filepath.Ext(path) == ".osu" {
// 				if id = getID(path); id != -1 {
// 					mapIDs[id] = true
// 				}
// 			}
// 			return nil
// 		})
// 	check(err)
// 	return mapIDs
// }

// func getID(path string) int {
// 	f, err := os.Open(path)
// 	check(err)
// 	defer f.Close()
// 	scanner := bufio.NewScanner(f)
// 	var s string
// 	var id int
// 	for scanner.Scan() {
// 		if strings.HasPrefix(scanner.Text(), "BeatmapID:") {
// 			s = strings.Split(scanner.Text(), ":")[1]
// 			id, err = strconv.Atoi(s)
// 			if err != nil {
// 				return -1
// 			}
// 		}
// 		return id
// 	}
// 	return -1
// }

// func getMd5(path string) [16]byte {
// 	content, err := ioutil.ReadFile(path)
// 	check(err)
// 	return md5.Sum(content)
// }
