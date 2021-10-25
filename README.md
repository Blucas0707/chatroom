# Golang chatroom
## Introduction
* 用golang + websocket打造多人聊天室<br>
* Framework: echo<br>
* Deploy: (GCP)Cloud Build + Cloud Run + Cloud SQL<br>

部分實作參考文章：
1. https://github.com/gorilla/websocket/tree/master/examples/chat
2. https://dev.to/jeroendk/building-a-simple-chat-application-with-websockets-in-go-and-vue-js-gao

## Demo
```
test account1: 123@gmail.com
test password1: 123QWEasd

test account2: 1234@gmail.com
test password2: 123QWEasd
```

## Main concept

* 一個房間：
  1. client會透過websocket機制與room server 建立handshake, 各別用2 個goroutine - read & write來作為與room server 傳遞與接收訊息。
  2. 同時，client 會被註冊到room server 的 register channel裡，來記錄此防參與者
  3. 之後，此房的參與者只要message channel 有訊息從client的read goroutine進來，就會透過message channel 將訊息傳遞給room server 中register的所有client, 並透過write goroutine去送出訊息至client
  
* 多個房間：
  1. 原理與一個房間類似，但要多一個 multi room server 來管理所有的room server 
  2. 因此建立新房間時，必須先確認multi-room server 中是否存在，若否，則新增room server; 若存在，則將此client直接register 給該room server.

## Deploy
* 在SQL的選擇上，若是使用local or GCE 則可以用docker-compose 將server ＆ SQL的container一次建出來，可以透過docker network 使兩個container 互通，若是想要保存每次的SQL資料，記得指定docker volume，否則每次rebuild container，資料會消失
  
* 這邊是使用GCP 的Cloud Build + Cloud Run + Cloud SQL，直接部署一個Cloud SQL並使Cloud Run 的server連上，後續透過Cloud Build的設定，只要GitHub branch更新，就會redeploy一個新的image