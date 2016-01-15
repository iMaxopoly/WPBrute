package main

import (
	"os"
	"bufio"
	"fmt"
)

/***
 *     _   _ _   _ _ _ _   _
 *    | | | | | (_) (_) | (_)
 *    | | | | |_ _| |_| |_ _  ___  ___
 *    | | | | __| | | | __| |/ _ \/ __|
 *    | |_| | |_| | | | |_| |  __/\__ \
 *     \___/ \__|_|_|_|\__|_|\___||___/
 *
 *
 */

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}

func file_loader(arg int, usr_chan, pwd_chan chan bool) {
	inFile, err := os.Open(os.Args[arg])
	check_err(err)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if arg == 2 {
			usr = append(usr, scanner.Text())
		} else {
			pwd = append(pwd, scanner.Text())
		}
	}
	if arg == 2 {
		fmt.Println("loaded: ", len(usr), "usernames")
		usr_chan <- true
	}else{
		fmt.Println("loaded: ", len(pwd), "passwords")
		pwd_chan <- true
	}
}