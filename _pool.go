package main

// func pool(tasks []map[string]interface{}) {
// 	var wg sync.WaitGroup
// 	jobs := make(chan int, len(tasks))
// 	results := make(chan job, len(tasks))
// 	for wid := 0; wid < 8; wid++ {
// 		wg.Add(1)
// 		go func(wid int, jobs <-chan int, results chan<- job) {
// 			defer wg.Done()
// 			for i := range jobs {
// 				results <- doTask(tasks, i)
// 			}
// 		}(wid, jobs, results)
// 	}
// 	for i := 0; i < len(tasks); i++ {
// 		jobs <- i
// 	}
// 	close(jobs)

// 	for r := 0; r < len(tasks); r++ {
// 		res := <-results
// 		printResult(res)
// 	}
// 	wg.Wait()
// }
