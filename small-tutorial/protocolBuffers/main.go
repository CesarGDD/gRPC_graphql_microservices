package main

import (
	"fmt"
	"log"
	tutorial "test/proto"

	"google.golang.org/protobuf/proto"
)

func main() {
	// Crear un nuevo mensaje Person
	person := &tutorial.Person{
		Id:    1,
		Name:  "Juan Perez",
		Email: "juanperez@example.com",
	}

	// Serializar el mensaje a binario
	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("Error serializing person: ", err)
	}

	fmt.Println("Serialized data: ", data)

	// Deserializar el mensaje desde binario
	newPerson := &tutorial.Person{}
	err = proto.Unmarshal(data, newPerson)
	if err != nil {
		log.Fatal("Error deserializing person: ", err)
	}

	// Mostrar el mensaje deserializado
	fmt.Println("Deserialized person: ", newPerson)
}
