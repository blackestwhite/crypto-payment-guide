package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	toContract := "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
	// toAddress := "TVkmZB3u7HPFdDC8t42Zp9FsUrcKfLfMba"
	url := fmt.Sprintf("https://apilist.tronscan.org/api/token_trc20/transfers?limit=1&start=0&contract_address=%v&start_timestamp=&end_timestamp=&confirm=", toContract)

	clinet := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36 Edg/88.0.705.81")
	resp, _ := clinet.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	fmt.Printf("%s\n", body)

}
