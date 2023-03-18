package enrollment

import (
	"context"
	"fmt"
	"os"

	"wa/api"
	"wa/global"
	"wa/utils"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
)

func SendRequestForPostponement(sender string) {
	fmt.Println("sending file")
	content, err := os.ReadFile("./documents/requestforpostponement.pdf")
	//content, err := ioutil.ReadFile("./bots/documents/Kottouf.png")
	if err != nil {
		fmt.Println(err)
	}

	resp, err := api.Client.Upload(context.Background(), content, whatsmeow.MediaDocument)
	if err != nil {
		fmt.Println(err)
	}

	msg := &waProto.DocumentMessage{
		FileName:      proto.String("طلب التأجيل"),
		Mimetype:      proto.String("application/pdf"), // replace this with the actual mime type
		Url:           &resp.URL,
		DirectPath:    &resp.DirectPath,
		MediaKey:      resp.MediaKey,
		FileEncSha256: resp.FileEncSHA256,
		FileSha256:    resp.FileSHA256,
		FileLength:    &resp.FileLength,
	}

	targetJID, ok := utils.ParseJID(sender)
	if !ok {
		return
	}
	//	send, err := global.Cli.SendMessage(jid, "", msg) // jid = recipient
	send, err := api.Client.SendMessage(context.Background(), targetJID, "", &waProto.Message{
		DocumentMessage: msg,
		//ImageMessage: msg,
	})

	if err != nil {
		global.Log.Errorf("Error sending message: %v", err)
	} else {
		global.Log.Infof("Message sent (server timestamp: %s)", send)
	}
}
