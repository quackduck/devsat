package main

import (
	api "github.com/quackduck/devzat/devzatapi"
)

func main() {
	//s, err := api.NewSession("devzat.hackclub.com:5556", os.Getenv("DEVZAT_TOKEN"))
	s, err := api.NewSession("localhost:5556", "dvz@R1HkT12cOvzc4nbPNUnE1xEdkHIhtnKGQkotDPzCOAc=")
	if err != nil {
		panic(err)
	}
	messageChan, middlewareResponseChan, err := s.RegisterListener(true, false, "orange|apple")
	if err != nil {
		panic(err)
	}
	for {
		select {
		case err = <-s.ErrorChan:
			panic(err)
		case msg := <-messageChan:
			if msg.Data == "orange" {
				middlewareResponseChan <- "🍊"
				err = s.SendMessage(api.Message{Room: msg.Room, From: "citrusbot", Data: "mmm... citrusy"})
			} else { // has to be apple because of the regex we set
				middlewareResponseChan <- "🍎"
				err = s.SendMessage(api.Message{Room: msg.Room, From: "applebot", Data: "mmm... appley"})
			}
			if err != nil {
				panic(err)
			}
		}
	}
}
