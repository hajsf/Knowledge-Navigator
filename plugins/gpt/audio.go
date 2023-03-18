package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"wa/api"
)

func Audio3(fileName string) string {
	url := "https://experimental.willow.vectara.io/v1/audio/transcriptions"
	var result map[string]interface{}
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fmt.Println("Converting:", fileName)
	cmd := exec.Command("ffmpeg", "-y", "-i", fileName, "output.mpeg")
	err := cmd.Run()
	if err != nil {
		fmt.Println("error converting the file:", err)
	} else {
		fw, err := w.CreateFormFile("file", "output.mpeg")
		if err != nil {
			panic(err)
		}

		file2, err := os.Open("output.mpeg")
		if err != nil {
			panic(err)
		}
		defer file2.Close()

		if _, err = io.Copy(fw, file2); err != nil {
			panic(err)
		}

		if fw, err = w.CreateFormField("model"); err != nil {
			panic(err)
		}
		if _, err = fw.Write([]byte("whisper-1")); err != nil {
			panic(err)
		}

		w.Close()
		fmt.Println("Audio file is ready, sending it to chatGPT")
		req, err := http.NewRequest("POST", url, &b)
		if err != nil {
			panic(err)
		}

		//	req.Header.Add("Content-Type", "application/json")
		req.Header.Set("Content-Type", w.FormDataContentType())
		req.Header.Add("customer-id", "")
		req.Header.Add("x-api-key", "")

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)

		//	fmt.Println(res)
		text := string(body)
		//	fmt.Println(text)

		json.Unmarshal([]byte(text), &result)

		fmt.Println("result:", result)
	}

	textValue, ok := result["text"].(string)
	if !ok {
		// handle error
		fmt.Println("Sorry can not get the text value:")
	}

	api.Content = textValue
	return textValue
}
