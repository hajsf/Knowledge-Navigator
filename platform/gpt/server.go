package gpt

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Postfunc(w http.ResponseWriter, r *http.Request) {

	var requestData RequestData

	err := json.NewDecoder(r.Body).Decode(&requestData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// prompt += requestData.Message // for complete
	prompt = requestData.Message // for chat
	fmt.Println(prompt)

	messages = append(messages, map[string]string{"role": "user", "content": prompt})
	//messages := []map[string]string{
	//	{"role": "user", "content": requestData.Message},
	//}

	//chatGPT := chat2(prompt)
	chatGPT := chat3(messages)

	response := ResponseData{Message: chatGPT} // strconv.FormatBool(requestData.Checked)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

/*
func Postfunc(w http.ResponseWriter, r *http.Request) {
	var requestData RequestData

	err := json.NewDecoder(r.Body).Decode(&requestData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	prompt := requestData.Message
	fmt.Println(prompt)

	messages = append(messages, map[string]string{"role": "user", "content": requestData.Message})
	//messages := []map[string]string{
	//	{"role": "user", "content": requestData.Message},
	//}
	fmt.Println("calling chat")
	// _ = chat(messages)
	chatGPT := chat2(prompt)

	for char := range chatGPT {
		w.Header().Set("Content-Type", "application/json")

		//	for i := 0; i < 5; i++ {
		//		response := ResponseData{Message: "Hello " + requestData.Message + " " + strconv.Itoa(i)}

		response := ResponseData{Message: char}

		jsonData, err := json.Marshal(response)
		if err != nil && response.Message != "" {
			fmt.Println(err)
		} else {
			fmt.Println(string(jsonData))
			json.NewEncoder(w).Encode(response)
			// Flush the data to the client
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
		}

		// time.Sleep(1 * time.Second)
	}
}
*/
