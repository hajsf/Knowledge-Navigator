package locations

import (
	"context"
	"wa/api"
	"wa/global"
	"wa/utils"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func InteractiveMap(sender string) {
	//	msg := &waProto.Message{
	//		Conversation: proto.String("*مع السلامه*")}

	msg := &waProto.Message{
		ExtendedTextMessage: &waProto.ExtendedTextMessage{
			Title: proto.String("يمكن التعرف عليها من خلال زيارة الخريطة التفاعلية التالية"),
			Text:  proto.String("https://www.google.com/maps/d/viewer?mid=1WEMnsCfckhX33_740nHuwF44NTE&hl=ar&ll=18.249015000000007%2C42.559155999999994&z=8"),
			//	CanonicalUrl: proto.String("https://forms.gle/CUoZmGpMUgBiqXWH7"),
			MatchedText: proto.String("https://www.google.com/maps/d/viewer?mid=1WEMnsCfckhX33_740nHuwF44NTE&hl=ar&ll=18.249015000000007%2C42.559155999999994&z=8"),
			//		JpegThumbnail: thumb,
			//	Description: proto.String("https://forms.gle/CUoZmGpMUgBiqXWH7"),
		},
	}

	targetJID, ok := utils.ParseJID(sender)
	if !ok {
		return
	}
	send, err := api.Client.SendMessage(context.Background(), targetJID, "", msg) // jid = recipient

	if err != nil {
		global.Log.Errorf("Error sending message: %v", err)
	} else {
		global.Log.Infof("Message sent (server timestamp: %s)", send)
	}
}
