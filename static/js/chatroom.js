const SendMessageAction = "send-message";
const JoinRoomAction = "join-room";
const LeaveRoomAction = "leave-room";

let models = {
  user:{
    isLogin:null,
    user_name:null,
    checkLogin:function(){
      return new Promise((resolve, reject)=>{
        return fetch("/api/user",{
          method:"GET"
        }).then((response)=>{
          return response.json();
        }).then((result)=>{
          if(result.data != null){
            models.user.isLogin = true;
            models.user.user_name = result.data.username;
            resolve("success");
          }
          else{
            models.user.isLogin = false;
            reject("fail");
          }
        });
      });
    },
  }
}

let viewers = {
  appendLog:function(item) {
    let  doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
    log.appendChild(item);
    if (doScroll) {
        log.scrollTop = log.scrollHeight - log.clientHeight;
    }
  },
  renderWSOpen:function(){
    let item = document.createElement("div");
    item.innerHTML = "<b>Room connected!</b>";
    this.appendLog(item);
  },
  renderMsg:function(msg){
    let  item = document.createElement("div");
    item.innerHTML = msg;
    this.appendLog(item);
  },
  renderUser:function(users){
    let user_box = document.querySelector(".client-list");
    //clear first
    while(user_box.hasChildNodes()){
      user_box.removeChild(user_box.firstChild);
    }
    for (let i = 0;i<users.length;i++) {
      let  item = document.createElement("div");
      item.innerHTML = users[i];
      user_box.appendChild(item);
    }
  },
  showRoomInfo:function(){
    let roomName = window.location.search.split("&")[0].slice(6,);
    let owner = window.location.search.split("owner=")[1];
    let roomNameitem = document.querySelector(".chatroom-name");
    let roomOwner = document.querySelector(".chatroom-owner");
    roomNameitem.innerHTML = roomName;
    roomOwner.innerHTML = owner;

  },
}

let controllers = {
  member: {
    checkLogin:function(){
      return new Promise((resolve, reject)=>{
        models.user.checkLogin().then(msg => {
          // viewers.user.isLogin();
          resolve(true);
        }).catch((msg)=>{
          alert("please login first");
          window.location.href = "/";
        });
      })
    },
  },
  ws:null,
  wsURL:"ws://https://chatroom-m6cesrpbta-de.a.run.app/chatroom/ws",
  roomName:null,
  rooms:null,
  room:null,
  newMessage:null,
  Users:null,
  connect:function(){
    this.connectToWebsocket();
  },
  connectToWebsocket:function(){
    this.roomName = window.location.search.split("&")[0].slice(6,);
    this.ws = new WebSocket(this.wsURL+ "?room=" + this.roomName);
    this.ws.addEventListener("open", (event)=>{
      this.onWebsocketOpen(event);
      this.joinRoom();
    })

    this.ws.addEventListener("message", (event)=>{
      this.handleNewMessage(event);
    })
  },
  onWebsocketOpen:function(event){
    viewers.renderWSOpen();
  },
  handleNewMessage(event){
    let msg = JSON.parse(event.data).message;
    let getroomName = JSON.parse(event.data).target;
    let sender = JSON.parse(event.data).sender === null ? "":JSON.parse(event.data).sender.Name;

    msg = msg.split(/\r?\n/);
    if (this.roomName === getroomName) {

      let action = JSON.parse(event.data).action;
      let sendMsg = "";
      switch (action) {
        case SendMessageAction:
          sendMsg = sender + ":" + msg;
          break;
        case JoinRoomAction:
          sendMsg = msg;
          this.Users = JSON.parse(event.data).user;
          break;
        case LeaveRoomAction:
          sendMsg = msg;
          this.Users = JSON.parse(event.data).user;
          break;
      }
      viewers.renderMsg(sendMsg);
      viewers.renderUser(this.Users);
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
      action: SendMessageAction,
      message: message,
      target: this.roomName,
    })
    this.ws.send(this.newMessage);
    this.newMessage.message = "";
  },
  joinRoom: function(){
    this.newMessage = JSON.stringify({
      action: JoinRoomAction,
      message: "",
      target: this.roomName,
    })
    this.ws.send(this.newMessage);
  },
  leaveRoom: function(){
    this.newMessage = JSON.stringify({
      action: LeaveRoomAction,
      message: "",
      target: this.roomName,
    });
    this.ws.send(this.newMessage);
    window.location.href ="/";
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
  LeaveRoom:function(){
    let leave_btn = document.querySelector(".leave-btn");
    leave_btn.addEventListener("click",()=>{
      if (!this.ws) {
        return false;
      }
      this.leaveRoom();
    })
  },
  init:function(){
    controllers.member.checkLogin().then(()=>{
      viewers.showRoomInfo();
      controllers.connect();
      controllers.SubmitMessage();
      controllers.LeaveRoom();
    });
  },
}

controllers.init();

