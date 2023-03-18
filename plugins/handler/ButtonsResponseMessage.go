package handler

import (
	"fmt"
	"strconv"
	"wa/Enum"
	"wa/enrollment"
	"wa/libraries"
	"wa/locations"

	"go.mau.fi/whatsmeow/types/events"
)

func ButtonsResponseMessage(evt *events.Message) {
	sender := evt.Info.Chat.User
	pushName := evt.Info.PushName

	_, _ = sender, pushName

	fmt.Println("Button responce pressed")
	ButtonResponse := evt.Message.GetButtonsResponseMessage()
	id, _ := strconv.Atoi(ButtonResponse.GetSelectedButtonId())
	switch id {
	case Enum.Locations:
		locations.MapLink(sender)
		locations.InteractiveMap(sender)
		locations.SendLocation(sender)
	case Enum.Libraries:
		libraries.LibrariesMenu(sender)
	case Enum.Enrollment:
		enrollment.EnrollmentMenu(sender)
	}
}
