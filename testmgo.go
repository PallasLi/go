package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("server1.example.com,server2.example.com")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
		&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		panic(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Phone:", result.Phone)
}




package main

import (
  "fmt"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

type Person struct {
  NAME  string
  PHONE string
}
type Men struct {
  Persons []Person
}
const = (
  URL = "192.168.2.175:27017"
)
func main() {

  session, err := mgo.Dial(URL)  //�������ݿ�
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)

  db := session.DB("mydb")	 //���ݿ�����
  collection := db.C("person") //����ü����Ѿ����ڵĻ�����ֱ�ӷ���


  //*****������Ԫ����Ŀ********
  countNum, err := collection.Count()
  if err != nil {
    panic(err)
  }
  fmt.Println("Things objects count: ", countNum)

  //*******����Ԫ��*******
  temp := &Person{
    PHONE: "18811577546",
    NAME:  "zhangzheHero"
  }
    //һ�ο��Բ��������� ��������Person����
  err = collection.Insert(&Person{"Ale", "+55 53 8116 9639"}, temp)
  if err != nil {
    panic(err)
  }

  //*****��ѯ��������*******
  result := Person{}
  err = collection.Find(bson.M{"phone": "456"}).One(&result)
  fmt.Println("Phone:", result.NAME, result.PHONE)

  //*****��ѯ��������*******
  var personAll Men  //��Ž��
  iter := collection.Find(nil).Iter()
  for iter.Next(&result) {
    fmt.Printf("Result: %v\n", result.NAME)
    personAll.Persons = append(personAll.Persons, result)
  }

  //*******��������**********
  err = collection.Update(bson.M{"name": "ccc"}, bson.M{"$set": bson.M{"name": "ddd"}})
  err = collection.Update(bson.M{"name": "ddd"}, bson.M{"$set": bson.M{"phone": "12345678"}})
  err = collection.Update(bson.M{"name": "aaa"}, bson.M{"phone": "1245", "name": "bbb"})

  //******ɾ������************
  _, err = collection.RemoveAll(bson.M{"name": "Ale��})
}