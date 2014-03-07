package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var TLDS []string
var R *rand.Rand

func main() {
	num_hosts := flag.Int("n", 1, "The number of random domains to generate")
	tlds_flag := flag.String("t", "com,net,org", "The TLDs to randomly assign to domains")
	min_length := flag.Int("m", 4, "The minimum character length of domains to generate")
	max_length := flag.Int("M", 26, "The maximum character length of domains to generate")
	flag.Usage = usage
	flag.Parse()

	if *min_length > *max_length {
		panic("Minimum length cannot exceed maximum length")
	}

	TLDS = strings.Split(*tlds_flag, ",")

	// Seed prng
	R = rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < *num_hosts; i++ {
		fmt.Printf("%s\t%s\r\n", rand_ip(), rand_domain(*min_length, *max_length))
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `random_hosts generates random domains and IPs for /etc/hosts or C:\Windows\System32\drivers\etc\hosts. Set your flags and then pipe to the proper hosts file.`+"\n\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func rand_alphanum_char() string {
	rand_int := R.Intn(36)
	var rand_char byte
	if rand_int < 10 {
		rand_char = byte(rand_int + 48)
	} else {
		rand_char = byte(rand_int + 87)
	}
	return string(rand_char)
}

func rand_domain(min_length int, max_length int) string {
	rand_length := R.Intn(max_length-min_length+1) + min_length
	var domain string
	for i := 0; i < rand_length; i++ {
		domain += rand_alphanum_char()
	}
	rand_tld := R.Intn(len(TLDS))
	domain += "." + TLDS[rand_tld]
	return domain
}

func rand_ip() string {
	var octets []string
	for i := 0; i < 4; i++ {
		rand_octet_int := R.Intn(222)
		switch rand_octet_int {
		case 0:
			rand_octet_int = 222
		case 127:
			rand_octet_int = 223
		}
		octets = append(octets, strconv.Itoa(rand_octet_int))
	}
	return strings.Join(octets, ".")
}
