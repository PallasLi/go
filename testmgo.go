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

  session, err := mgo.Dial(URL)  //连接数据库
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)

  db := session.DB("mydb")	 //数据库名称
  collection := db.C("person") //如果该集合已经存在的话，则直接返回


  //*****集合中元素数目********
  countNum, err := collection.Count()
  if err != nil {
    panic(err)
  }
  fmt.Println("Things objects count: ", countNum)

  //*******插入元素*******
  temp := &Person{
    PHONE: "18811577546",
    NAME:  "zhangzheHero"
  }
    //一次可以插入多个对象 插入两个Person对象
  err = collection.Insert(&Person{"Ale", "+55 53 8116 9639"}, temp)
  if err != nil {
    panic(err)
  }

  //*****查询单条数据*******
  result := Person{}
  err = collection.Find(bson.M{"phone": "456"}).One(&result)
  fmt.Println("Phone:", result.NAME, result.PHONE)

  //*****查询多条数据*******
  var personAll Men  //存放结果
  iter := collection.Find(nil).Iter()
  for iter.Next(&result) {
    fmt.Printf("Result: %v\n", result.NAME)
    personAll.Persons = append(personAll.Persons, result)
  }

  //*******更新数据**********
  err = collection.Update(bson.M{"name": "ccc"}, bson.M{"$set": bson.M{"name": "ddd"}})
  err = collection.Update(bson.M{"name": "ddd"}, bson.M{"$set": bson.M{"phone": "12345678"}})
  err = collection.Update(bson.M{"name": "aaa"}, bson.M{"phone": "1245", "name": "bbb"})

  //******删除数据************
  _, err = collection.RemoveAll(bson.M{"name": "Ale”})
}