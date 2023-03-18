package handler

import (
	"encoding/json"
	"fmt"
	"wa/api"
	"wa/translation"
	"wa/utils"

	"github.com/abadojack/whatlanggo"
	"go.mau.fi/whatsmeow/types/events"
)

func ExtendedTextMessage(evt *events.Message) {
	sender := evt.Info.Chat.User
	pushName := evt.Info.PushName

	info, err := json.MarshalIndent(evt.Message.ExtendedTextMessage.GetText(), "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	msgReceived := string(info)
	data, _ := utils.PrepareModel(evt.Info.Chat.User,
		sender, pushName, evt.Info.Timestamp.Local().Format("Mon 02-Jan-2006 15:04"),
		evt.Info.ID, "text", msgReceived, "", "")
	api.Passer.Data <- api.SSEData{
		Event:   "message", // default: source.onmessage = function (event) {}
		Message: data,
	}

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
	fmt.Println("Hi Extended")
	welcome := translation.HelloPerson(lang, pushName)
	fmt.Println(welcome)

	if !evt.Info.IsFromMe {
		utils.MainMenu(sender, welcome)
	}
}
