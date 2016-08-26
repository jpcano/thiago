package main

import (
	"log"
	"fmt"
	"github.com/jpcano/thiago/thiago"
)

func main() {
	info := thiago.SessionInfo{
		Host: "127.0.0.1:27017",
		Username: "",
		Password: "",
		Database: "thiago",
		Publications: "publications",
		Subscribers: "subscribers",
	}

	// Ask Alberto why if I put this inside the if does not work
	caca, err := thiago.GetSession(info)
	if  err != nil {
		log.Fatal(err)
	}

	defer caca.Close()
	
	fmt.Println("Adding publication")
	if err := caca.Publish("http://helpscout", []string{"Go", "Python"}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Publication inserted successfully!")

	
	fmt.Println("Adding petros.koklas subscribed to Go and Java")
	if err := caca.Subscribe("petros.koklas", []string{"Go", "Java"}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Subscriber inserted successfully!")


	fmt.Println("Adding jesus.cano subscribed to Go and Python")
	if err := caca.Subscribe("jesus.cano", []string{"Go", "Python"}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Subscriber inserted successfully!")

	fmt.Println("Adding Pablo subscribed to Javascript")
	if err := caca.Subscribe("pablo", []string{"Javascript"}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Subscriber inserted successfully!")


	fmt.Println("Finding subscribers subscribed to Java:")
	if names, err := caca.FindSubscriberByTags([]string{"Java"}); err == nil {
		for _, name := range names{
			fmt.Println(name)
		}
	} else {
		log.Fatal(err)
	}
	
	fmt.Println("Finding subscribers subscribed to Go:")
	if names, err := caca.FindSubscriberByTags([]string{"Go"})
	err == nil {
		for _, name := range names{
			fmt.Println(name)
		}
	} else {
		log.Fatal(err)
	}

	fmt.Println("Finding subscribers subscribed to Java or Python:")
	if names, err := caca.FindSubscriberByTags([]string{"Java", "Python"})
	err == nil {
		for _, name := range names{
			fmt.Println(name)
		}
	} else {
		log.Fatal(err)
	}
	

	fmt.Println("Finding subscribers subscribed to Javascript:")
	if names, err := caca.FindSubscriberByTags([]string{"Javascript"})
	err == nil {
		for _, name := range names{
			fmt.Println(name)
		}
	} else {
		log.Fatal(err)
	}
}
