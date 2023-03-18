package main

import (
	"context"
	"wa/api"
	"wa/global"
	"wa/server"
)

func init() {
	// create a context that we can cancel
	api.Messages = []map[string]string{ // ['system', 'assistant', 'user']
		{"role": "user", "content": `
Hi, I'm your client, and you are my medical insuranse company called HeathCare, and your name is MedBot, having this list of clients in csv format: ` + api.Clients +
			` each client is having a policy number, once the client is in need of your service, the medical center is issuing an
approval request that is having a unique number so that you can see the details, compare it with the client policy details along with the
client history, and accordingly either approve or decline the request, he is an update database of the active polices for your reference in csv format:
: ` + api.Policies + ` and here is an updated list of the historical claims for each policy: ` + api.HistoricalClaims +
			`here also you can find a csv file of the currently active claims: ` + api.ActiveClaims +
			` You need to welcome the client politly by showing your name and only at the welcoming message ask for his policy number
			after that you should be able to read all his data from the csv files provided to you, either by using his policy number, name,
			or mobile number as an identity verification, upon retrivng the client data send him a welcoming message and mention his name in it,
			do not ask him about any other identity verification, you should know his identity from the policy number only and should know 
			his name and all other info from the csv data provided to you, do not ask him for another identification if his policy number is valid,
			avoid repeating messages of your conversation with the client, you should be able to give the client any info he need from the records you have
			at the csv formated data provided to you, and do not pull any other data from the net`},
	}
	api.Counter = 0
	global.Ctx = context.Background()
	global.Ctx, global.Cancel = context.WithCancel(global.Ctx)
}

func main() {
	println("Hi there...")

	/*messages := []map[string]string{ // ['system', 'assistant', 'user']
		{"role": "system", "content": `You are a medical insuranse company, having this list of clients in csv format: ` + clients +
			` each client is having a policy number, once the client is in need of your service, the medical center is issuing an
	   approval request that is having a unique number so that you can see the details, compare it with the client policy details along with the
	   client history, and accordingly either approve or decline the request, he is an update database of the active polices for your reference in csv format:
	   : ` + policies + ` and here is an updated list of the historical claims for each policy: ` + historicalClaims +
			`here also you can find a csv file of the currently active claims: ` + activeClaims +
			` by the way do not forget to welcome the client by his name, you can ask for his number to confirm his identity.
	   	you need to stick to this data and task, and do not ask or provide any info outside these csv files`},
	} */

	/*	api.Content = `I'm your client, and you are my medical insuranse company, having this of clients in csv format: ` + clients +
			` each client is having a policy number, once the client is in need of your service, the medical center is issuing an
		approval request that is having a unique number so that you can see the details, compare it with the client policy details along with the
		client history, and accordingly either approve or decline the request, he is an update database of the active polices for your reference in csv format:
		: ` + policies + ` and here is an updated list of the historical claims for each policy: ` + historicalClaims +
			`here also you can find a csv file of the currently active claims: ` + activeClaims +
			` by the way do not forget to welcome the client by his name, you can ask for his number to confirm his identity.
			you need to stick to this data and task, and do not ask or provide any info outside these csv files`
	*/
	//chatGPT := chat2(prompt)

	//_ = gpt.Chat3(messages)

	// a WaitGroup for the goroutines to tell us they've stopped
	/*	wg := sync.WaitGroup{}

		wg.Add(1)
		go ui.Init(global.Ctx, &wg)

		go func() {
	*/server.Run(global.Ctx, ":1235")
	/*
	   		println("server closed")
	   	}()

	   println("server started")

	   // listen for Ctrl+C
	   c := make(chan os.Signal, 1)
	   signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	   select {
	   case <-global.Ctx.Done():

	   	fmt.Println("Context had been cancelled")
	   	global.Cancel()

	   case <-c:

	   		fmt.Println("Server aborted")

	   		// destroy the ui if it is opened
	   		if ui.Wv != nil {
	   			fmt.Println("Destroying the ui")
	   			wg.Done()
	   			ui.Wv.Destroy()
	   		}

	   		fmt.Println("Destroying the wa client")
	   		if api.Client.IsConnected() {
	   			api.Passer.Data <- api.SSEData{
	   				Event:   "notification",
	   				Message: "Server is shut down at the host machine...",
	   			}
	   			api.Client.Disconnect()
	   		}
	   	}

	   // and wait for them both to reply back
	   wg.Wait()
	   fmt.Println("main: all goroutines have told us they've finished")
	*/
}
