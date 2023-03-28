package main

import (
	"context"
	firebase "firebase.google.com/go"
	"flag"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/websocket"
	"google.golang.org/api/option"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

type numberBlock struct {
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

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var addr = flag.String("addr", "localhost:8020", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/task2", task2)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func task2(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
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
		blockNumber := big.NewInt(15960495)
		block, err := client.BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			log.Fatal(err)
		}
		message := block.Number().String() + "<br>" + strconv.FormatUint(block.Time(), 10) + "<br>" + block.Hash().String() +
			"<br>" + block.Difficulty().String() + "<br>" + strconv.Itoa(len(block.Transactions())) + "<br>"
		c.WriteMessage(websocket.TextMessage, []byte(message))
		ref := client_bd.NewRef("from_block_data/ ")
		b := numberBlock{
			Number:       block.Number().String(),
			Time:         strconv.FormatUint(block.Time(), 10),
			Difficulty:   block.Difficulty().String(),
			Hash:         block.Hash().String(),
			Transactions: strconv.Itoa(len(block.Transactions())),
		}
		if err := ref.Set(context.TODO(), b); err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)
	}
}
