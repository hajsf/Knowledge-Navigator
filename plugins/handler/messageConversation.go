package handler

import (
	"fmt"
	"wa/api"
	"wa/gpt"
	"wa/helpers"
	"wa/utils"

	"go.mau.fi/whatsmeow/types/events"
)

func Conversation(evt *events.Message) {
	var msgReceived string

	sender := evt.Info.Chat.User
	pushName := evt.Info.PushName
	received := evt.Message.GetConversation()
	// convert numbers in Arabic scrtip to numbers in latin script
	for _, e := range received {
		if e >= 48 && e <= 57 {
			//	fmt.Println("Number in english script number")
			msgReceived = fmt.Sprintf("%s%v", msgReceived, string(e))
		} else if e >= 1632 && e <= 1641 {
			//	fmt.Println("It is Arabic script")
			msgReceived = fmt.Sprintf("%s%v", msgReceived, helpers.NormalizeNumber(e))
		} else {
			//	fmt.Println("Dose not looks to be a number")
			msgReceived = fmt.Sprintf("%s%v", msgReceived, string(e))
		}
	}

	/*	fmt.Println("Received a message!", evt.Message.GetConversation()) */
	message := map[string]string{"role": "user", "content": msgReceived}
	// chatGPT3 := gpt.ChatGPT3(message)
	chatGPT3 := gpt.Chat3(message)
	fmt.Println("chatGPT:", chatGPT3)
	fmt.Println("msgReceived:", msgReceived)
	data, err := utils.PrepareModel(evt.Info.Chat.User,
		sender, pushName, evt.Info.Timestamp.Local().Format("Mon 02-Jan-2006 15:04"),
		evt.Info.ID, "text", msgReceived, chatGPT3, "")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("data:", data)
	}

	api.Passer.Data <- api.SSEData{
		Event:   "message", // default: source.onmessage = function (event) {}
		Message: data,
	}
	/*
	   msg := whatlanggo.Detect(msgReceived)
	   fmt.Println("Language:", msg.Lang.String(), " Script:", whatlanggo.Scripts[msg.Script], " Confidence: ", msg.Confidence)

	   //	name := "Hasan"
	   var lang string
	   switch whatlanggo.Scripts[msg.Script] {
	   case "Arabic":

	   	// go WelcomeMessage(sender, pushName)
	   	lang = "ar"

	   case "Latin":

	   		lang = "en"
	   		//	go WelcomeMessageLatin(sender, pushName)
	   	}

	   _ = lang
	   fmt.Println("Hi conversation")
	   welcome := translation.HelloPerson(lang, pushName)
	   fmt.Println(welcome)

	   	if !evt.Info.IsFromMe {
	   		// utils.MainMenu(sender, welcome)
	   	}
	*/
}
