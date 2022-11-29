package connect

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Value struct {
	ProjectID   string `json:"ИДПроекта"`
	ProjectName string `json:"ИмяПроекта"`
	TaskName    string `json:"НазваниеЗадачи"`
	StartDate   string `json:"ДатаНачалаЗадачи"`
	EndDate     string `json:"ДатаОкончанияЗадачи"`
}

type Projects struct {
	Odata_link     string  `json:"odata.metadata"`
	Values         []Value `json:"value"`
	Odata_nextlink string  `json:"odata.nextLink"`
}

func AnoData() {
	count := 0
	baseUrl := "https://spanorsi.lancloud.ru/pwa/_api/ProjectData/"
	queryUrl := "Задачи?$select=ИДПроекта,ИмяПроекта,НазваниеЗадачи,ДатаНачалаЗадачи,ДатаОкончанияЗадачи"
	jsonUrl := "&$format=json"
	Url := baseUrl + queryUrl + jsonUrl

	for {

		client := http.Client{}
		req, err := http.NewRequest("GET", Url, nil)
		if err != nil {
			log.Fatal("Сработал в запросе")

		}

		req.SetBasicAuth("conteq_1@ano-rsi.ru", "BesB8288")
		resp, err := client.Do(req)
		bodyText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Сработал в авторизации")
		}

		defer resp.Body.Close()

		pro := Projects{}

		json_error := json.Unmarshal(bodyText, &pro)

		if json_error != nil {
			log.Fatal()
		}

		db(&pro)
		count += len(pro.Values)

		if &pro.Odata_nextlink == nil {
			break
		}
		Url = baseUrl + *&pro.Odata_nextlink + jsonUrl
	}
}
