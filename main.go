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

	resource, err := thiago.GetSession(info)
	if  err != nil {
		log.Fatal(err)
	}

	defer resource.Close()
	
	fmt.Println("Adding publication")
	if err := resource.Publish("http://helpscout", []string{"html", "css"}); err != nil {
		log.Fatal(err)
	}
		
	fmt.Println("Adding complex publication")
	if err := resource.Publish("http://helpscout", []map[string]interface{}{{"java":78, "lisp": 88}, {"elixir":77, "react": 02}}); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Adding petros.koklas subscribed to Go and Java")
	if err := resource.Subscribe("petros.koklas", []string{"Java", "go"}); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Adding jesus.cano subscribed to a complex publication")
	if err := resource.Subscribe("jesus.cano", []map[string]interface{}{{"java":78, "lisp": 88}, {"elixir":77, "react": 02}}); err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Adding Pablo subscribed to Javascript")
	if err := resource.Subscribe("pablo", []string{"Javascript"}); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Finding subscribers subscribed to Java:")
	if names, err := resource.FindSubscriberByTags([]string{"Java"}); err == nil {
		for _, name := range names{
			fmt.Println(name)
		}
	} else {
		log.Fatal(err)
	}
	
	fmt.Println("Finding subscribers subscribed to Go:")
	if names, err := resource.FindSubscriberByTags([]string{"go"})
	err == nil {
		for _, name := range names{
			fmt.Println(name)
		}
	} else {
		log.Fatal(err)
	}

	fmt.Println("Finding subscribers subscribed to Java or Python (or both):")
	if names, err := resource.FindSubscriberByTags([]string{"Java", "Python"})
	err == nil {
		for _, name := range names{
			fmt.Println(name)
		}
	} else {
		log.Fatal(err)
	}
	

	fmt.Println("Finding subscribers subscribed to Javascript:")
	if names, err := resource.FindSubscriberByTags([]string{"Javascript"})
	err == nil {
		for _, name := range names{
			fmt.Println(name)
		}
	} else {
		log.Fatal(err)
	}
	
	fmt.Println("Finding subscribers subscribed to a complex publication:")
	if names, err := resource.FindSubscriberByTags([]map[string]interface{}{{"java": 78, "lisp": 88}, {"elixir": 77, "rect": 02}})
	err == nil {
		for _, name := range names{
			fmt.Println(name)
		}
	} else {
		log.Fatal(err)
	}
}
