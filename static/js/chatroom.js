let viewers = {
  appendLog:function(item) {
    var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
    log.appendChild(item);
    if (doScroll) {
        log.scrollTop = log.scrollHeight - log.clientHeight;
    }
  },
  renderWSOpen:function(){
    var item = document.createElement("div");
    item.innerHTML = "<b>Room connected!</b>";
    this.appendLog(item);
  }
}

let controllers = {
  ws:null,
  wsURL:"ws://localhost:1323/chatroom/ws",
  roomName:null,
  rooms:null,
  room:null,
  newMessage:null,
  connect:function(){
    this.connectToWebsocket();
  },
  connectToWebsocket:function(){
    this.roomName = window.location.search.slice(6,);
    this.ws = new WebSocket(this.wsURL+ "?room=" + this.roomName);
    this.ws.addEventListener("open", (event)=>{
      this.onWebsocketOpen(event);
      this.joinRoom();
    })

    // this.ws.addEventListener("message", (event)=>{
    //   this.handleNewMessage(event);
    // })
  },
  onWebsocketOpen:function(event){
    viewers.renderWSOpen();
  },
  handleNewMessage(event){
    let data = event.data;
    data = data.split(/\r?\n/);
    for(let i=0;i<data.length;i++){
      let msg = JSON.parse(data[i])
      this.room = this.findRoom(msg);
      if(typeof this.room !== 'undefined'){
        this.room.messages.push(msg);
      }
    }
  },
  findRoom: function(roomName){
    for (let i = 0; i <this.rooms.length; i++) {
      if (this.rooms[i].name === roomName) {
        return this.rooms[i];
      }
    }
  },
  sendMessage: function(message){
    this.newMessage = JSON.stringify({
      action: 'send-message',
      message: message,
      target: this.roomName,
    })
    console.log("send msg: ", this.newMessage);
    this.ws.send(this.newMessage);
    this.newMessage.message = "";
    console.log("send done");
  },
  joinRoom: function(){
    this.newMessage = JSON.stringify({
      action: 'join-room',
      message: "",
      target: this.roomName,
    })
    this.ws.send(this.newMessage);
    console.log(this.newMessage);
    console.log("Join done!");
    // this.message = [];
    // this.room.push({
    //   "roomname": this.roomName,
    //   "message": [],
    // });
    // this.roomName = "";
  },
  leaveRoom: function(){
    this.ws.send(JSON.stringify({
      action: "leave-room",
      message: this.room.name,
    }));
    for (let i = 0; i < this.rooms.length; i++){
      if (this.rooms[i].name === room.name){
        this.rooms.splice(i, 1);
        break;
      }
    }
  },
  SubmitMessage:function(){
    let msg = document.getElementById("msg-input");
    let submit_btn = document.querySelector(".submit-btn");
    submit_btn.addEventListener("click",()=>{
      if (!this.ws) {
        return false;
      }
      if (!msg.value) {
          return false;
      }
      // this.newMessage.message = msg.value;
      // this.ws.send(msg.value);
      this.sendMessage(msg.value);
      msg.value = "";
      return false;
    })
  },
  init:function(){
    controllers.connect();
    // controllers.joinRoom();
    controllers.SubmitMessage();
  },
}

controllers.init();

/*
window.onload = function () {
  var conn;
  var msg = document.getElementById("msg");
  var log = document.getElementById("log");

  function appendLog(item) {
      var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
      log.appendChild(item);
      if (doScroll) {
          log.scrollTop = log.scrollHeight - log.clientHeight;
      }
  }

  document.getElementById("form").onsubmit = function () {
      if (!conn) {
          return false;
      }
      if (!msg.value) {
          return false;
      }
      console.log("send msg: ", msg.value);
      conn.send(msg.value);
      msg.value = "";
      return false;
  };

  if (window["WebSocket"]) {
      roomName = window.location.search.slice(6,)
      wsUrl = "ws://" + document.location.host + "/chatroom/ws?room=" + roomName;
      console.log(wsUrl);
      conn = new WebSocket(wsUrl);
      conn.onopen = function (evt){
        var item = document.createElement("div");
        item.innerHTML = "<b>Room connected!</b>";
        appendLog(item);
      };
      conn.onclose = function (evt) {
          var item = document.createElement("div");
          item.innerHTML = "<b>Connection closed.</b>";
          appendLog(item);
      };
      conn.onmessage = function (evt) {
          var messages = evt.data.split('\r?\n');
          console.log("received msg: ", messages);
          for (var i = 0; i < messages.length; i++) {
            let msg = JSON.parse(messages[i]);
            var item = document.createElement("div");
              item.innerText = messages[i];
              appendLog(item);
          }
      };
  } else {
      var item = document.createElement("div");
      item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
      appendLog(item);
  }

};

*/