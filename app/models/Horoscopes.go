package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Horoscope map[string][]Horoscope `json:"horoscope"`
}
type Horoscope struct {
	Content string      `json:"content"`
	Item    string      `json:"item"`
	Money   int         `json:"money"`
	Total   int         `json:"total"`
	Job     int         `json:"job"`
	Color   string      `json:"color"`
	Day     interface{} `json:"day"`
	Love    int         `json:"love"`
	Rank    int         `json:"rank"`
	Sign    string      `json:"sign"`
}

const apiUrl = "http://api.jugemkey.jp/api/horoscope/free"

func GetHoroscope(year string, month string, day string) (horoscope []Horoscope) {
	url := fmt.Sprintf("%s/%s/%s/%s", apiUrl, year, month, day)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Error Requesr=%s", err.Error())
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Client.Do err=%s", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ioutil.ReadAll err=%s", err.Error())
	}
	var res Response

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatalf("json.Unmarshal err=%s", err.Error())
	}
	for _, hh := range res.Horoscope {
		for _, h := range hh {
			horoscope = append(horoscope, h)
		}
	}
	return horoscope
}
