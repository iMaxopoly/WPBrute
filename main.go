package main

import (
	"fmt"
	"os"
	"sync"
	"strconv"
	"math"
	"net/url"
)

/***
 *    ______            _      ______
 *    | ___ \          | |     | ___ \
 *    | |_/ /_ __ _   _| |_ ___| |_/ / __ ___  ___ ___
 *    | ___ \ '__| | | | __/ _ \  __/ '__/ _ \/ __/ __|
 *    | |_/ / |  | |_| | ||  __/ |  | | |  __/\__ \__ \
 *    \____/|_|   \__,_|\__\___\_|  |_|  \___||___/___/
 *
 *
 */

// Globals
var location string

// Create map of username list & password list
var usr []string
var pwd []string

// Goroutines manager
var wg sync.WaitGroup

func main() {
	// Adding workers, default being 10
	n_workers := func() (int) {
		if os.Args[4] == "" {
			return 10
		}
		t, err := strconv.Atoi(os.Args[4])
		check_err(err)
		return t
	}()
	wg.Add(n_workers)

	// Parse and store URL
	location = func() (string) {
		_, err := url.Parse(os.Args[1])
		check_err(err)
		return os.Args[1]
	}()

	// Load user list
	usr_chan := make(chan bool, 1)
	pwd_chan := make(chan bool, 1)
	go file_loader(2, usr_chan, pwd_chan)

	// Load pass list
	go file_loader(3, usr_chan, pwd_chan)

	// Wait till files loaded.
	<-usr_chan
	<-pwd_chan

	// Distribute workload to workers
	section := int(math.Floor(float64(len(usr)) / float64(n_workers)))
	starting_point := 0
	remainder := len(usr) - (section * n_workers)
	for x := 0; x < n_workers-1; x++ {
		go request(starting_point, starting_point+section)
		starting_point += section
	}
	go request(starting_point, starting_point+remainder)

	// Wait for tasks to finish
	wg.Wait()
	fmt.Println("Fin")
}

