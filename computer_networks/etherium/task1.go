package main

import (
	"context"
	firebase "firebase.google.com/go"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/websocket"
	"google.golang.org/api/option"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

type lastBlock struct {
	Number       string `json:"number"`
	Time         string `json:"time"`
	Difficulty   string `json:"difficulty"`
	Hash         string `json:"hash"`
	Transactions string `json:"transactions"`
}

var ctx = context.Background()

// configure database URL
var conf = &firebase.Config{
	DatabaseURL: "https://iu9alexeev-d3456-default-rtdb.firebaseio.com/",
}

// fetch service account key
var opt = option.WithCredentialsFile("iu9alexeev-d3456-firebase-adminsdk-ngu6u-f6b2e49f6d.json")

var addr = flag.String("addr", "localhost:8010", "http service address")
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func task1(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer conn.Close()
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/a26c02aa151e4ba59f0df5ff0c7ea13b")
	if err != nil {
		log.Fatalln(err)
	}
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}
	client_bd, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("error in creating firebase DB client: ", err)
	}
	for {
		header, err := client.HeaderByNumber(context.Background(), nil)
		if err != nil {
			fmt.Println(1)
			log.Fatal(err)
		}
		message := header.Number.String() + "<br>"
		blockNumber := big.NewInt(header.Number.Int64())
		block, err := client.BlockByNumber(context.Background(), blockNumber) //get block with this number
		if err != nil {
			fmt.Println(2)
			log.Fatal(err)
		}
		message += block.Number().String() + "<br>" + strconv.FormatUint(block.Time(), 10) + "<br>" + block.Hash().String() +
			"<br>" + block.Difficulty().String() + "<br>" + strconv.Itoa(len(block.Transactions()))
		conn.WriteMessage(websocket.TextMessage, []byte(message))
		ref := client_bd.NewRef("last_block/ ")
		b := lastBlock{
			Number:       block.Number().String(),
			Time:         strconv.FormatUint(block.Time(), 10),
			Difficulty:   block.Difficulty().String(),
			Hash:         block.Hash().String(),
			Transactions: strconv.Itoa(len(block.Transactions())),
		}
		if err := ref.Set(context.TODO(), b); err != nil {
			fmt.Println(3)
			log.Fatal(err)
		}
		time.Sleep(2 * time.Second)
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/task1", task1)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
