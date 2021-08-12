package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)


var upgrader = websocket.Upgrader{}

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r,"index.html")
	})

	http.HandleFunc("/v1/ws", func(w http.ResponseWriter, r *http.Request){
		var conn, _ = upgrader.Upgrade(w,r,nil)
		go func(conn *websocket.Conn){
			for {
				mType, msg, _ := conn.ReadMessage()
				err := conn.WriteMessage(mType,msg)
				if err != nil {
					log.Println(err)
				}
			}
		}(conn)
	})

	http.HandleFunc("/v2/ws", func(w http.ResponseWriter, r *http.Request){
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn){
			for {
				_, msg, _ := conn.ReadMessage()
				println(string(msg))
			}

		}(conn)
	})

	http.HandleFunc("/v3/ws", func(w http.ResponseWriter, r *http.Request){
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn){
			ch := time.Tick(5 * time.Second)
			for range ch {
				_ = conn.WriteJSON(myStruct{
					Username:  "iitheo",
					FirstName: "Theo",
					LastName:  "K",
				})
			}

		}(conn)
	})

	http.HandleFunc("/v4/ws", func(w http.ResponseWriter, r *http.Request){
		var conn, _ = upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn){
			for {
				_, _, err := conn.ReadMessage()
				if err != nil {
					_ = conn.Close()
				}
			}

		}(conn)
	})

	err := http.ListenAndServe(":3000",nil)
	if err != nil {
		log.Println("Error:", err)
	}
}

type myStruct struct {
	Username string `json:"username"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}
