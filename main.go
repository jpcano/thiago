package main

import (
	"fmt"
	"github.com/jpcano/thiago/thiago"
	"log"
)

func main() {
	defer thiago.Close()
	var err error
	var names []string
	
	fmt.Println("Adding publication")
	if err = thiago.Publish("http://helpscout", []string{"Go", "Python"}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Publication inserted successfully!")

	
	fmt.Println("Adding petros.koklas subscribed to Go and Java")
	if err = thiago.Subscribe("petros.koklas", []string{"Go", "Java"}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Subscriber inserted successfully!")


	fmt.Println("Adding jesus.cano subscribed to Go and Python")
	if err = thiago.Subscribe("jesus.cano", []string{"Go", "Python"}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Subscriber inserted successfully!")

	fmt.Println("Adding Pablo subscribed to Javascript")
	if err = thiago.Subscribe("pablo", []string{"Javascript"}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Subscriber inserted successfully!")


	fmt.Println("Finding subscribers subscribed to Java:")
	names, err = thiago.FindSubscriberByTags([]string{"Java"})
	if err != nil {
		log.Fatal(err)
	}
	
	for _, name := range names{
		fmt.Println(name)
	}

	fmt.Println("Finding subscribers subscribed to Go:")
	names, err = thiago.FindSubscriberByTags([]string{"Go"})
	if err != nil {
		log.Fatal(err)
	}
	
	for _, name := range names{
		fmt.Println(name)
	}

	fmt.Println("Finding subscribers subscribed to Java or Python:")
	names, err = thiago.FindSubscriberByTags([]string{"Java", "Python"})
	if err != nil {
		log.Fatal(err)
	}
	
	for _, name := range names{
		fmt.Println(name)
	}

	fmt.Println("Finding subscribers subscribed to Javascript:")
	names, err = thiago.FindSubscriberByTags([]string{"Javascript"})
	if err != nil {
		log.Fatal(err)
	}
	
	for _, name := range names{
		fmt.Println(name)
	}

}
