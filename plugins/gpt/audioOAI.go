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
	"path/filepath"
	"wa/api"
)

func AudioGPT(fileName string) string {
	url := "https://api.openai.com/v1/audio/transcriptions"
	API_KEY := ""
	var textValue string
	var result map[string]interface{}
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fmt.Println("Converting:", fileName)
	cmd := exec.Command("ffmpeg", "-y", "-i", fileName, "outputGTP3.mp3")
	err := cmd.Run()
	if err != nil {
		fmt.Println("error converting the file:", err)
	} else {
		dir, err := os.Getwd()
		if err != nil {
			// handle error
			fmt.Println("error:", err)
		}
		outputFilePath := filepath.Join(dir, "outputGPT3.mp3")
		fmt.Println("path: ", outputFilePath)

		fw, err := w.CreateFormFile("file", outputFilePath) // "outputGPT3.mp3")
		if err != nil {
			panic(err)
		}

		file2, err := os.Open("outputGPT3.mp3")
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

		req, err := http.NewRequest("POST", url, &b)
		if err != nil {
			panic(err)
		}

		//	req.Header.Add("Content-Type", "application/json")
		req.Header.Set("Content-Type", w.FormDataContentType())
		req.Header.Set("Authorization", "Bearer "+API_KEY)

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
