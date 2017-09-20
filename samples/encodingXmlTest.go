package main

import (
	"encoding/xml"
	"fmt" 
)

type Person struct {
	Name           string
	Age            int
	ServerResponse bool
}

func  main( ) {

	var person Person
	var body = `<Person>
	<Name>lyt</Name>
	<Age>23</Age>
	<ServerResponse>false</ServerResponse>
	</Person>`
	xml.Unmarshal([]byte(body), &person)
	
	fmt.Printf("%v\n",person)
	
	person.ServerResponse = true

	responseXML, _ := xml.Marshal(person)
	fmt.Printf( "%v\n",string(responseXML))
}

 