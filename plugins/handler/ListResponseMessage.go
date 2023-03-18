package handler

import (
	"fmt"
	"strconv"
	"wa/Enum"
	"wa/enrollment"
	"wa/locations"
	"wa/utils"

	"go.mau.fi/whatsmeow/types/events"
)

func ListResponseMessage(evt *events.Message) {
	sender := evt.Info.Chat.User
	pushName := evt.Info.PushName
	targetJID, ok := utils.ParseJID(sender)
	if !ok {
		return
	}

	ListResponse := evt.Message.GetListResponseMessage()
	id, _ := strconv.Atoi(ListResponse.SingleSelectReply.GetSelectedRowId())
	fmt.Println(id, sender)
	//responces.ListResponces(id, sender, pushName)
	_, _, _ = sender, pushName, targetJID
	switch id {
	case Enum.RequestForPostponement:
		enrollment.SendRequestForPostponement(sender)
	case Enum.Locations:
		locations.MapLink(sender)
		/*		send2, err := api.Client.SendMessage(context.Background(), targetJID, "", &waProto.Message{
					Conversation: proto.String("Check Google maps :)"),
				})
				if err != nil {
					global.Log.Errorf("Error sending message: %v", err)
				} else {
					global.Log.Infof("Message sent (server timestamp: %s)", send2)
				}
		*/
	}

	//	locations.InteractiveMap(sender)
}
