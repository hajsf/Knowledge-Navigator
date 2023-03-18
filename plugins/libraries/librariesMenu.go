package libraries

import (
	"context"
	"strconv"
	"wa/Enum"
	"wa/api"
	"wa/global"
	"wa/utils"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func LibrariesMenu(sender string) {

	targetJID, ok := utils.ParseJID(sender)
	if !ok {
		return
	}
	msg := &waProto.ListMessage{
		Title:       proto.String("عمادة شؤون المكتبات الخاصة بجامعة الملك خالد"),
		Description: proto.String("يرجى تحديد الإستفسار"),
		ButtonText:  proto.String("أنقر هنا  👈"),
		ListType:    waProto.ListMessage_SINGLE_SELECT.Enum(),
		Sections: []*waProto.ListMessage_Section{
			{
				Title: proto.String("لدي إستفسار بخصوص المكتبات الخاصة بجامعة الملك خالد"),
				Rows: []*waProto.ListMessage_Row{
					{
						RowId: proto.String(strconv.Itoa(Enum.Enrollment)),
						Title: proto.String("لوائح وشروط الإعارة"),
						//	Description: proto.String("عمادة القبول والتسجيل"),
					},
					{
						RowId: proto.String(strconv.Itoa(Enum.Libraries)),
						Title: proto.String("خدمات المكتبة الإلكترونية"),
					},
					{
						RowId: proto.String(strconv.Itoa(Enum.Locations)),
						Title: proto.String("ساعات العمل للمكتبات المركزية"),
					},
					{
						RowId: proto.String(strconv.Itoa(Enum.Locations)),
						Title: proto.String("أماكن تواجد المكتبات"),
					},
					{
						RowId: proto.String(strconv.Itoa(Enum.Locations)),
						Title: proto.String("التواصل مع المكتبة المركزية"),
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
				ListMessage: msg,
			},
		}})

	if err != nil {
		global.Log.Errorf("Error sending message: %v", err)
	} else {
		global.Log.Infof("Message sent (server timestamp: %s)", send)
	}
}
