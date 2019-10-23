package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const base = "http://127.0.0.1:8500/v1/agent"

func main() {
	rsp, err := http.Get(base + "/services")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = rsp.Body.Close()
	}()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}

	services := regexp.MustCompile(`"ID":"(.*?)"`).FindAllSubmatch(body, -1)
	for _, v := range services {
		log.Printf("%s", "deregister service: "+string(v[1]))

		url := base + "/service/deregister/" + string(v[1])

		req, err := http.NewRequest("PUT", url, nil)
		if err != nil {
			log.Fatal(err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		_ = res.Body.Close()
	}
}
