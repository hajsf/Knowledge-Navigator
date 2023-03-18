package locations

import (
	"context"
	"fmt"

	"wa/api"
	"wa/global"
	"wa/utils"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func SendLocation(sender string) {
	fmt.Println("sending file")
	msg := &waProto.Message{
		LocationMessage: &waProto.LocationMessage{
			DegreesLatitude:  proto.Float64(18.2477238),
			DegreesLongitude: proto.Float64(42.5580554),
			Name:             proto.String("موقع جامعة الملك خالد"),
			/*	Address:                           new(string),
				Url:                               new(string),
				IsLive:                            new(bool),
				AccuracyInMeters:                  new(uint32),
				SpeedInMps:                        new(float32),
				DegreesClockwiseFromMagneticNorth: new(uint32),
				Comment:                           new(string),
				JpegThumbnail:                     []byte{},
				ContextInfo:                       &waProto.ContextInfo{}, */
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
