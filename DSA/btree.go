package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//if err != nil {
//	if os.IsExist(err) {
//		log.Println("file already exists")
//		return
//	}
//	log.Fatal(err)
//}
//

func main() {
	db_path := "btree/wena.db"

	clear_btree(db_path)
	for i := 0; i < 100; i++ {
		add_element(db_path, strconv.Itoa(i), []byte("value"+strconv.Itoa(i)))

	}

	key := "33"
	id, value := search_element(db_path, key)
	println("key: %s, value: %s", id, value)

}

func search_element(db string, key string) (string, string) {
	var _iterations uint8 = 0

	file, _ := os.OpenFile(db, os.O_RDONLY, 0)

	file.Seek(0, 0)
	reader := bufio.NewReader(file)

	var id, value string = "", ""
	for _iterations < 10 {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Iterations number: %u", _iterations)
		fmt.Println(line)
		fmt.Println("-------------")
		parts := strings.Split(line, ":")
		id, value = parts[0], parts[1]

		if id == key {
			break
		}

		_iterations++
	}

	return id, value
}

func add_element(db string, key string, val []byte) {
	file, _ := os.OpenFile(
		db,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)

	fmt.Fprintf(file, "%s:%s\n", key, val)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
}

func clear_btree(path string) {
	os.RemoveAll(path)
}
