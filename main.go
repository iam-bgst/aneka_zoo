package main

import (
	"context"
	"fmt"
	"log"
	
	"zoo/application"
	"zoo/infrastructures"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	app, err := application.NewApp(ctx)
	if err != nil {
		log.Print("Failed to initialize app. Error: ", err)
		panic(err)
	}
	
	defer app.Close(ctx)
	
	go func() {
		data := <-app.TerminalHandler
		
		msgStr := fmt.Sprintf("system call: %+v", data)
		app.Logger.InfoWithContext(ctx, "", msgStr, "err", data)
		
		cancel()
	}()
	
	go infrastructures.NewHTTPServer(app)
	<-ctx.Done()
}
