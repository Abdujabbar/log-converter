package main

import (
	"fmt"
	"os"
)

const ftype = "1"
const stype = "2"

func main() {
	// fname := "/tmp/logs.txt"
	// fileRangeChan := make(chan watcher.FRange)
	// go func(r chan watcher.FRange) {
	// 	watcher.Process(r, fname)
	// }(fileRangeChan)

	// for {
	// 	select {
	// 	case v, ok := <-fileRangeChan:
	// 		if ok {
	// 			file, _ := os.Open(fname)
	// 			rbytes := make([]byte, v.End-v.Start)
	// 			file.ReadAt(rbytes, v.Start)
	// 			fmt.Printf("Written string: %v", string(rbytes))
	// 		}
	// 		break
	// 	}
	// }
	// args, files, err := parseInputArgs(os.Args)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(args)
	// fmt.Println(files)

	// dao := repo.DAO{
	// 	Server:   "localhost",
	// 	Database: "testdb",
	// }
	// dao.Connect()
	// records, err := dao.FindAll()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// fmt.Println(records)

	// t := time.Now()
	// raw := repo.Record{
	// 	ID:     bson.NewObjectId(),
	// 	Time:   t.Unix() + rand.Int63n(10),
	// 	Msg:    "Hello World",
	// 	Format: "first_format",
	// }
	// err := dao.Insert(raw)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// sess, err := mgo.Dial("localhost")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer sess.Close()
	// err = sess.Ping()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println("MongoDB server is healthy.")

	// // collection := sess.DB("testdb").C("logs")
	// // collection.Create()
	// // var i interface{}
	// // r := make([]*interface{}, 10)
	// // collection.Find(i).All(&r)
	// // fmt.Println(r)
	// os.Exit(0)
	// s := "Feb 1, 2018 at 3:04:05pm (UTC)"
	// fmt.Println(time.Parse("Jan 2, 2006 at 3:04:05pm (UTC)", s))
}

func parseInputArgs(args []string) ([]string, []string, error) {
	if len(args) < 3 {
		return nil, nil, fmt.Errorf("Required parameters files and format doesn't passed")
	}
	availableFormats := map[string]bool{}
	availableFormats[ftype] = true
	availableFormats[stype] = true
	if _, ok := availableFormats[args[len(args)-1]]; !ok {
		return nil, nil, fmt.Errorf("format parameter can receive only %v or %v", ftype, stype)
	}

	files := args[1 : len(args)-1]

	for _, v := range files {
		if _, err := os.Stat(v); os.IsNotExist(err) {
			return nil, nil, fmt.Errorf("%v doesn't exist, please make sure that is already exists %v and run program again", v, v)
		}
	}
	return args, files, nil
}
