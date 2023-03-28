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

type transBlock struct {
	ChainId  string `json:"chainId"`
	Hash     string `json:"hash"`
	Value    string `json:"value"`
	Cost     string `json:"cost"`
	To       string `json:"to"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
}

var ctx = context.Background()

// configure database URL
var conf = &firebase.Config{
	DatabaseURL: "https://iu9alexeev-d3456-default-rtdb.firebaseio.com/",
}

// fetch service account key
var opt = option.WithCredentialsFile("iu9alexeev-d3456-firebase-adminsdk-ngu6u-f6b2e49f6d.json")

var addr = flag.String("addr", "localhost:8030", "http service address")
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/task3", task3)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func task3(w http.ResponseWriter, r *http.Request) {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/a26c02aa151e4ba59f0df5ff0c7ea13b")
	if err != nil {
		log.Fatalln(err)
	}
	c, _ := upgrader.Upgrade(w, r, nil)
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
		block, err := client.BlockByNumber(context.Background(), blockNumber) //get block with this number
		if err != nil {
			log.Fatal(err)
		}
		var message string
		for i, tx := range block.Transactions() {
			message = tx.ChainId().String() + "<br>" + tx.Hash().String() + "<br>" + tx.Value().String() + "<br>" +
				tx.Cost().String() + "<br>" + tx.To().String() + "<br>" + strconv.FormatUint(tx.Gas(), 10) + "<br>" + tx.GasPrice().String() + "<br>"
			b := transBlock{
				ChainId:  tx.ChainId().String(),
				Hash:     tx.Hash().String(),
				Value:    tx.Value().String(),
				Cost:     tx.Cost().String(),
				To:       tx.To().String(),
				Gas:      strconv.FormatUint(tx.Gas(), 10),
				GasPrice: tx.GasPrice().String(),
			}
			ref := client_bd.NewRef("blocks_data/ " + fmt.Sprint(i))
			if err := ref.Set(context.TODO(), b); err != nil {
				log.Fatal(err)
			}
		}
		c.WriteMessage(websocket.TextMessage, []byte(message))
		time.Sleep(time.Second * 5)
	}

}
