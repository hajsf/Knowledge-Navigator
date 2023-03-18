package api

import (
	"fmt"
	"net/http"
)

type SSEData struct {
	Event, Message string
}
type DataPasser struct {
	Data       chan SSEData
	Logs       chan string
	Connection chan struct{} // To control maximum allowed clients connections
}

var Passer DataPasser
var Content string
var Counter int
var Messages = []map[string]string{}

const Clients = `
Policy Number,Name,Mobile Number
12345,"John Doe","555-1234"
67890,"Jane Smith","555-5678"
24680,"Bob Johnson","555-2468"
13579,"Alice Brown","555-1357"
98765,"Charlie Davis","555-9876"
`
const Policies = `
Policy Number,Surgical Coverage,Dental Coverage
12345,$5000,$1000
67890,$6000,$1200
24680,$7000,$1400
13579,$8000,$1600
98765,$9000,$1800
`
const HistoricalClaims = `
Policy Number,Claim Date,Surgical Coverage Used,Dental Coverage Used,Medical Center
12345,2022-01-01,$5000,$0,"St. Mary's Hospital"
67890,2022-02-15,$0,$1200,"City Dental Clinic"
24680,2022-03-20,$7000,$1400,"General Hospital"
13579,2022-04-10,$8000,$1600,"Downtown Medical Center"
98765,2022-05-05,$9000,$1800,"Westside Clinic"
`

const ActiveClaims = `
Policy Number,Claim Date,Claim Type,Coverage Used,Medical Center
12345,2022-06-01,Surgical,$6000,"St. Mary's Hospital"
67890,2022-07-15,Dental,$1500,"City Dental Clinic"
24680,2022-08-20,Surgical,$8000,"General Hospital"
13579,2022-09-10,Dental,$2000,"Downtown Medical Center"
98765,2022-10-05,Surgical,$10000,"Westside Clinic"
`

func (p DataPasser) HandleSignal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Cache-Control", "no-cache")
	// Allow cross origin  if required
	setupCORS(&w, r)

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Internal error", 500)
		return
	}

	fmt.Println("Client connected from IP:", r.RemoteAddr)
	// fmt.Println(len(p.connection), "new connection recieved")
	if len(p.Connection) > 0 {
		fmt.Fprint(w, "event: notification\ndata: Connection is opened in another browser/tap ...\n\n")
		flusher.Flush()
	}
	p.Connection <- struct{}{}

	fmt.Fprint(w, "event: notification\ndata: Connecting to WhatsApp server ...\n\n")
	flusher.Flush()

	// Connect to the WhatsApp client
	go p.Connect()

	for {
		select {
		case data := <-p.Data:
			// fmt.Println("SSE data recieved")
			//fmt.Println("msg recieved:", data.Message)
			switch {
			case len(data.Event) > 0:
				fmt.Fprintf(w, "event: %v\ndata: %v\n\n", data.Event, data.Message)
			case len(data.Event) == 0:
				fmt.Fprintf(w, "data: %v\n\n", data.Message)
			}
			flusher.Flush()
		case <-r.Context().Done():
			<-p.Connection
			fmt.Println("Connection closed from IP:", r.RemoteAddr)
			return
		}
	}
}

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

/*	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Resource Not Found"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
*/
