package utils

import (
	"context"
	"strconv"
	"wa/Enum"
	"wa/api"
	"wa/global"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func MainMenu(sender, welcome string) {

	targetJID, ok := ParseJID(sender)
	if !ok {
		return
	}
	/*
		msg1 := &waProto.ListMessage{
			Title:       proto.String(welcome),
			Description: proto.String("يرجى تحديد سبب التواصل"),
			ButtonText:  proto.String("أنقر هنا  👈"),
			ListType:    waProto.ListMessage_SINGLE_SELECT.Enum(),
			Sections: []*waProto.ListMessage_Section{
				{
					Title: proto.String("لدي إستفسار بخصوص:"),
					Rows: []*waProto.ListMessage_Row{
						{
							RowId: proto.String(strconv.Itoa(Enum.Enrollment)),
							Title: proto.String("عمادة القبول والتسجيل"),
							//	Description: proto.String("عمادة القبول والتسجيل"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.Libraries)),
							Title: proto.String("عمادة شؤون المكتبات"),
						},
						{
							RowId: proto.String(strconv.Itoa(Enum.Locations)),
							Title: proto.String("مواقع كليات وفروع جامعة الملك خالد"),
						},
					},
				},
			},
			//	ProductListInfo: &waProto.ListMessage_ProductListInfo{},
			//	FooterText:      new(string),
			//	ContextInfo:     &waProto.ContextInfo{},
		}
		send, err := api.Client.SendMessage(context.Background(), targetJID, "", &waProto.Message{
			ViewOnceMessage: &waProto.FutureProofMessage{
				Message: &waProto.Message{
					ListMessage: msg1,
				},
			}})

		if err != nil {
			global.Log.Errorf("Error sending message: %v", err)
		} else {
			global.Log.Infof("Message sent (server timestamp: %s)", send)
		}
	*/

	msg := &waProto.ButtonsMessage{
		ContentText: proto.String(welcome),
		FooterText:  proto.String("يرجى إختيار الموضوع الذي يمكننا مساعدتك به"),
		HeaderType:  waProto.ButtonsMessage_EMPTY.Enum(),
		Buttons: []*waProto.ButtonsMessage_Button{
			{
				ButtonId:       proto.String(strconv.Itoa(Enum.Enrollment)),
				ButtonText:     &waProto.ButtonsMessage_Button_ButtonText{DisplayText: proto.String("عمادة القبول والتسجيل")},
				Type:           waProto.ButtonsMessage_Button_RESPONSE.Enum(),
				NativeFlowInfo: &waProto.ButtonsMessage_Button_NativeFlowInfo{},
			},
			{
				ButtonId:       proto.String(strconv.Itoa(Enum.Libraries)),
				ButtonText:     &waProto.ButtonsMessage_Button_ButtonText{DisplayText: proto.String("عمادة شؤون المكتبات")},
				Type:           waProto.ButtonsMessage_Button_RESPONSE.Enum(), //proto.ButtonsMessage_Button_Type.Enum,
				NativeFlowInfo: &waProto.ButtonsMessage_Button_NativeFlowInfo{},
			},
			{
				ButtonId:       proto.String(strconv.Itoa(Enum.Locations)),
				ButtonText:     &waProto.ButtonsMessage_Button_ButtonText{DisplayText: proto.String("مواقع كليات وفروع جامعة الملك خالد")},
				Type:           waProto.ButtonsMessage_Button_RESPONSE.Enum(),
				NativeFlowInfo: &waProto.ButtonsMessage_Button_NativeFlowInfo{},
			},
		},
	}

	send, err := api.Client.SendMessage(context.Background(), targetJID, "", &waProto.Message{
		ViewOnceMessage: &waProto.FutureProofMessage{
			Message: &waProto.Message{
				ButtonsMessage: msg,
			},
		}})

	if err != nil {
		global.Log.Errorf("Error sending message: %v", err)
	} else {
		global.Log.Infof("Message sent (server timestamp: %s)", send)
	}

	/*	var content strings.Builder
		content.WriteString(welcome)
		content.WriteString("\n")
		content.WriteString("الرجاء إختيار خيار من أدناه")
		content.WriteString("\n")
		content.WriteString("1️⃣ إستفسارات عمادة القبول و التسجيل")
		content.WriteString("\n")
		content.WriteString("2️⃣ إستفسارات عمادة شؤون المكتبات")
		content.WriteString("\n")
		content.WriteString("3️⃣ مواقع كليات و فروع الجامعه")
		content.WriteString("\n")
		send2, err := api.Client.SendMessage(context.Background(), targetJID, "", &waProto.Message{
			Conversation: proto.String(content.String()),
		})
		if err != nil {
			global.Log.Errorf("Error sending message: %v", err)
		} else {
			global.Log.Infof("Message sent (server timestamp: %s)", send2)
		}
		fmt.Println(send2) */

}
