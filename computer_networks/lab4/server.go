package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

type Node struct {
	Self *Address
	Next *Address
}

type Address struct {
	IPv4 string
	Port string
}

type Package struct {
	From    *Address
	Command string
	Data    []Blog
}

type Blog struct {
	Title   string
	Content string
	Tags    []string
}

var logger = log.New(os.Stdout, "", 2)

func isSubarray(arr1, arr2 []string) bool {
	for _, elem := range arr1 {
		flag := false
		for _, elem2 := range arr2 {
			if elem == elem2 {
				flag = true
				break
			}
		}
		if flag == false {
			return false
		}
	}
	return true
}

func printData(data []Blog) string {
	res := ""
	for _, elem := range data {
		res += elem.Title + "\n" + elem.Content + "\n" + strings.Join(elem.Tags, ", ") + "\n"
	}
	return res
}

func newNode(self, next string) *Node {
	return &Node{
		Self: splitAddress(self),
		Next: splitAddress(next),
	}
}

func splitAddress(address string) *Address {
	splitted := strings.Split(address, ":")
	if len(splitted) != 2 {
		logger.Fatal("Некорректный адрес! Формат: <IPv4>:<порт>")
	}
	return &Address{
		IPv4: splitted[0],
		Port: ":" + splitted[1],
	}
}

func (node *Node) HandleServer() {
	listen, err := net.Listen("tcp", node.Self.Port)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("Сервер запущен на порте " + node.Self.Port)
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {

		}
	}(listen)
	for {
		conn, err := listen.Accept()
		if err != nil {
			logger.Println(err)
			err := conn.Close()
			if err != nil {
				return
			}
			continue
		}
		go node.HandleConnection(conn)
	}
}

func (node *Node) HandleConnection(conn net.Conn) {
	logger.Println("Обрабатывается новое подключение...")
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	buffer := make([]byte, 1024*8)
	length, err := conn.Read(buffer)
	if length == 0 || err != nil {
		logger.Fatal(err)
	}
	var pack Package
	err = json.Unmarshal(buffer[:length], &pack)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("Запрос: " + pack.Command + ", адрес отправителя: " + pack.From.IPv4 + pack.From.Port)
	switch pack.Command {
	case "/read":
		node.HandleReading(conn, &pack)
	case "/write":
		node.HandleWriting(conn, &pack)
	default:
		logger.Println("Запрос нераспознан!")
	}
}

func (node *Node) HandleReading(conn net.Conn, pack *Package) {
	inputTags := pack.Data[0].Tags
	logger.Println("Обрабатывается запрос на чтение...")
	var newPack = Package{
		From:    node.Self,
		Command: pack.Command,
		Data:    []Blog{},
	}
	file, err := os.Open("./data/data.txt")
	if err != nil {
		logger.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		title := scanner.Text()
		scanner.Scan()
		content := scanner.Text()
		scanner.Scan()
		tags := scanner.Text()
		tagsArray := strings.Split(tags, ", ")
		data := Blog{
			Title:   title,
			Content: content,
			Tags:    tagsArray,
		}
		if isSubarray(inputTags, tagsArray) {
			newPack.Data = append(newPack.Data, data)
		}
	}
	logger.Println("Считаны блоги: \"" + printData(newPack.Data) + "\"")
	jsonPack, err := json.Marshal(newPack)
	if err != nil {
		logger.Fatal(err)
	}
	conn.Write(jsonPack)
	logger.Println("Отправлен ответ.")
}

func (node *Node) HandleWriting(conn net.Conn, pack *Package) {
	logger.Println("Обрабатывается запрос на запись...")
	//err := os.Remove("./data/data.txt")
	//file, err := os.Create("./data/data.txt")

	file, err := os.OpenFile("./data/data.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		logger.Fatal("Файл для записи не найден")
	}
	_, err = file.WriteString(pack.Data[0].Title + "\n" + pack.Data[0].Content + "\n" + strings.Join(pack.Data[0].Tags, ", ") + "\n")
	if err != nil {
		logger.Fatal(err)
		return
	}
	logger.Println("Запись завершена.")
}

func (node *Node) HandleClient() {
	for {
		command := inputString()
		switch command {
		case "/read":
			node.RequestReading()
		case "/write":
			node.RequestWriting()
		case "/exit":
			os.Exit(1)
		case "/help":
			showCommands()
		default:
			fmt.Println("Введите /help для просмотра доступных команд.")
		}
	}
}

func inputString() string {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.Replace(str, "\n", "", -1)
}

func (node *Node) RequestReading() {
	logger.Println("Проверка связи с сервером...")
	conn, err := net.Dial("tcp", node.Next.IPv4+node.Next.Port)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("Подключение установлено.")
	defer conn.Close()
	logger.Println("Введите теги через запятую:")
	inputTags := strings.Split(inputString(), ", ")
	var pack = Package{
		From:    node.Self,
		Command: "/read",
		Data: []Blog{
			Blog{
				Title:   "",
				Content: "",
				Tags:    inputTags,
			},
		},
	}
	jsonPack, err := json.Marshal(pack)
	if err != nil {
		logger.Fatal(err)
	}
	conn.Write(jsonPack)
	logger.Println("Отправлен запрос на чтение...")
	buffer := make([]byte, 1024)
	var message string
	for {
		length, err := conn.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			logger.Fatal(err)
		}
		message += string(buffer[:length])
	}
	err = json.Unmarshal([]byte(message), &pack)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("Получен и расшифрован ответ.")
	for _, elem := range pack.Data {
		fmt.Println("Title: " + elem.Title + "\nContent: " + elem.Content + "\nTags: " + strings.Join(elem.Tags, ", "))
	}
}

func (node *Node) RequestWriting() {
	fmt.Println("Введите заголовок блога:")
	title := inputString()
	fmt.Println("Введите содержимое блога:")
	content := inputString()
	fmt.Println("Введите теги через запятую:")
	tags := inputString()
	tagsArray := strings.Split(tags, ", ")
	var data []Blog
	data = append(data, Blog{
		Title:   title,
		Content: content,
		Tags:    tagsArray,
	})
	logger.Println("Проверка связи с сервером...")
	conn, err := net.Dial("tcp", node.Next.IPv4+node.Next.Port)
	if err != nil {
		logger.Println(err)
		node.RequestWriting()
	}
	logger.Println("Подключение установлено.")
	defer conn.Close()
	var pack = Package{
		From:    node.Self,
		Command: "/write",
		Data:    data,
	}
	jsonPack, err := json.Marshal(pack)
	if err != nil {
		logger.Println(err)
	}
	conn.Write(jsonPack)
	logger.Println("Отправлен запрос на запись.")
}

func showCommands() {
	fmt.Println(`ДОСТУПНЫЕ КОМАНДЫ:
					/read - вывести все блоги с указанными тегами.
					/write - добавить блог.
					/exit - завершить работу.
					/help - отобразить список команд.`)
}

func init() {
	if len(os.Args) != 3 {
		logger.Fatal("Некорректный запуск! Формат: ./server <адрес-текущего-узла>" +
			"<адрес-следующего-узла>")
	}
}

func main() {
	node := newNode(os.Args[1], os.Args[2])
	go node.HandleServer()
	node.HandleClient()
}
