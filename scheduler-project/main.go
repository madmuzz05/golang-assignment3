package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func callApi(water, wind int) {
	var reqData = struct {
		water int
		wind  int
	}{
		water: water,
		wind:  wind,
	}

	fmt.Printf("%+v\n", reqData)

	data := map[string]interface{}{
		"wind":  reqData.wind,
		"water": reqData.water,
	}

	request, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{}

	model, err := http.NewRequest("POST", "http://localhost:8080/log/update_status", bytes.NewBuffer(request))
	model.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(model)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode == http.StatusCreated {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		var status = struct {
			StatusWind  string `json:"status_wind"`
			StatusWater string `json:"status_water"`
		}{}
		err = json.Unmarshal(body, &status)
		fmt.Println("status water :", status.StatusWater)
		fmt.Println("status wind :", status.StatusWind)
	}
}

func main() {
	for i := 1; ; i++ {
		water := rand.Intn(100)
		wind := rand.Intn(100)
		if water != 0 && wind != 0 {
			callApi(water, wind)
			time.Sleep(15 * time.Second)
		}
	}
}
