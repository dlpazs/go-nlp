package controller

import(
	//"bytes"
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"net/url"
)

func SummarizationHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	
	formData := url.Values{
		"query": {r.FormValue("summarize")},
	}
	
	url := "http://localhost:8080/summarize"
	resp, _ := http.PostForm(url, formData)
	
	body, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	
	fmt.Fprintf(w, "You said: %s, the response was\n\n %s", r.FormValue("summarize"), bodyString)
	
}