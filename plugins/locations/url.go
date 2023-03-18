package locations

import (
	"context"
	"wa/api"
	"wa/global"
	"wa/utils"

	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func MapLink(sender string) {
	targetJID, ok := utils.ParseJID(sender)
	if !ok {
		return
	}

	/*	content, err := ioutil.ReadFile("./documents/kottouf.png")
		if err != nil {
			fmt.Println(err)
		}
		resp, err := api.Client.Upload(context.Background(), content, whatsmeow.MediaImage)
		if err != nil {
			fmt.Println(err)
		}
	*/
	// Creating template message
	msg := &waProto.TemplateMessage{
		HydratedTemplate: &waProto.TemplateMessage_HydratedFourRowTemplate{
			/*				Title: &waProto.TemplateMessage_HydratedFourRowTemplate_ImageMessage{
								ImageMessage: &waProto.ImageMessage{
									//	Caption:  proto.String("شكرا و عيدكم مبارك"),
									Mimetype: proto.String("image/png"), // replace this with the actual mime type
									// you can also optionally add other fields like ContextInfo and JpegThumbnail here

									Url:           &resp.URL,
									DirectPath:    &resp.DirectPath,
									MediaKey:      resp.MediaKey,
									FileEncSha256: resp.FileEncSHA256,
									FileSha256:    resp.FileSHA256,
									FileLength:    &resp.FileLength,
									Height:        proto.Uint32(410),
									Width:         proto.Uint32(1200),
								},
							},
			*/
			Title: &waProto.TemplateMessage_HydratedFourRowTemplate_HydratedTitleText{
				HydratedTitleText: "مواقع الكليات",
			},
			TemplateId:          proto.String("template-id"),
			HydratedContentText: proto.String("يمكن التعرف عليها من خلال زيارة الخريطة التفاعلية التالية"),
			//	HydratedFooterText:  proto.String("يمكن التعرف عليها من خلال زيارة الخريطة التفاعلية التالية"),
			HydratedButtons: []*waProto.HydratedTemplateButton{

				// This for URL button
				{
					Index: proto.Uint32(1),
					HydratedButton: &waProto.HydratedTemplateButton_UrlButton{
						UrlButton: &waProto.HydratedTemplateButton_HydratedURLButton{
							DisplayText: proto.String("👉 أنقر هنا"),
							Url:         proto.String("https://www.google.com/maps/d/viewer?mid=1WEMnsCfckhX33_740nHuwF44NTE&hl=ar&ll=18.249015000000007%2C42.559155999999994&z=8"),
						},
					},
				},
			},
		},
	}

	// Sending message
	// WaClient.SendMessage(event.Info.Chat, "", this_message)

	send, err := api.Client.SendMessage(context.Background(), targetJID, "", &waProto.Message{
		ViewOnceMessage: &waProto.FutureProofMessage{
			Message: &waProto.Message{
				TemplateMessage: msg,
			},
		}})
	if err != nil {
		global.Log.Errorf("Error sending message: %v", err)
	} else {
		global.Log.Infof("Message sent (server timestamp: %s)", send)
	}
	/*
		send, err = global.Cli.SendMessage(jid, "", &waProto.Message{
			ImageMessage: &imageMessage,
		})

		if err != nil {
			global.Log.Errorf("Error sending message: %v", err)
		} else {
			global.Log.Infof("Message sent (server timestamp: %s)", send)
		}
	*/
}
