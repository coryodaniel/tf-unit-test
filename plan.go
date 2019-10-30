package main
import "encoding/json"
import "fmt"
import "io/ioutil"
import "os"

type Plan struct {
	Errors           []interface{} `json:"errors"`
	ChangedResources []struct {
		Action            string `json:"action"`
		Type              string `json:"type"`
		Name              string `json:"name"`
		Path              string `json:"path"`
		ChangedAttributes struct {
			ID struct {
				New struct {
					Type string `json:"type"`
				} `json:"new"`
			} `json:"id"`
			Content struct {
				New struct {
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"new"`
			} `json:"content"`
			Filename struct {
				New struct {
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"new"`
			} `json:"filename"`
		} `json:"changedAttributes"`
		Module string `json:"module"`
	} `json:"changedResources"`
	ChangedDataSources []interface{} `json:"changedDataSources"`
}

func GetPlan() Plan {
  const plan = "./plan.json"
  jsonFile, err := os.Open(plan)
  defer jsonFile.Close()

  if err != nil {
    fmt.Println(err)
  }

  byteValue, _ := ioutil.ReadAll(jsonFile)

  var result Plan
  json.Unmarshal([]byte(byteValue), &result)

  return result
}
