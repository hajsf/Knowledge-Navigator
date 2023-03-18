package enrollment

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

func EnrollmentMenu(sender string) {

	targetJID, ok := utils.ParseJID(sender)
	if !ok {
		return
	}
	msg := &waProto.ListMessage{
		Title:       proto.String("عمادة القبول و التسجيل في بجامعة الملك خالد"),
		Description: proto.String("يرجى تحديد الإستفسار"),
		ButtonText:  proto.String("أنقر هنا  👈"),
		ListType:    waProto.ListMessage_SINGLE_SELECT.Enum(),
		Sections: []*waProto.ListMessage_Section{
			{
				Title: proto.String("لدي إستفسار بخصوص عمادة القبول و التسجيل في بجامعة الملك خالد"),
				Rows: []*waProto.ListMessage_Row{
					{
						RowId: proto.String(strconv.Itoa(Enum.Enrollment)),
						Title: proto.String("الحذف والاضافة"),
						//	Description: proto.String("عمادة القبول والتسجيل"),
					},
					{
						RowId: proto.String(strconv.Itoa(Enum.Libraries)),
						Title: proto.String("الاعتذار عن الفصل الدراسي"),
					},
					{
						RowId: proto.String(strconv.Itoa(Enum.Locations)),
						Title: proto.String("الاعتذار عن مقرر دراسي"),
					},
					{
						RowId: proto.String(strconv.Itoa(Enum.RequestForPostponement)),
						Title: proto.String("تأجيل الفصل الدراسي"),
					},
					{
						RowId: proto.String(strconv.Itoa(Enum.Locations)),
						Title: proto.String("طي القيد وإعادة القيد المطوي"),
					},
					{
						RowId: proto.String(strconv.Itoa(Enum.Locations)),
						Title: proto.String("الانسحاب من الجامعة"),
					},
					{
						RowId: proto.String(strconv.Itoa(Enum.Locations)),
						Title: proto.String("الفصل من الجامعة"),
					},
					{
						RowId: proto.String(strconv.Itoa(Enum.Locations)),
						Title: proto.String("الطالب الزائر"),
					},
					{
						RowId: proto.String(strconv.Itoa(Enum.Locations)),
						Title: proto.String("تغيير التخصص"),
					},
					{
						RowId: proto.String(strconv.Itoa(Enum.Locations)),
						Title: proto.String("المعدل الجامعي"),
					},
					{
						RowId: proto.String(strconv.Itoa(Enum.Locations)),
						Title: proto.String("المكافآت الجامعية"),
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
