let models = {
  user:{
    loginSuccess:null,
    loginResponse:null,
    Login:function(){
      return new Promise((resolve, reject)=>{
        //reset registerSuccess
        models.user.loginSuccess = null;
        let email = document.querySelector(".login-email").value;
        let password = document.querySelector(".login-password").value;
        let data = {
          "email":email,
          "password":password
        };
        // console.log(email,password);
        return fetch("/api/user",{
          method:'PATCH',
          headers: {
            "Content-type":"application/json",
          },
          body: JSON.stringify(data),
        }).then((response)=>{
          return response.json();
        }).then((result)=>{
          // result = JSON.parse(result);
          models.user.loginResponse = result.errorMessage;
          if(result.errorExist === false){
            models.user.loginSuccess = true;
          }else{
            models.user.loginSuccess = false;
          }
          resolve(true);
        });
      });
    },
    isLogin:null,
    user_name:null,
    checkLogin:function(){
      return new Promise((resolve, reject)=>{
        return fetch("/api/user",{
          method:"GET"
        }).then((response)=>{
          return response.json();
        }).then((result)=>{
          console.log(result)
          if(result.data != null){
            models.user.isLogin = true;
            models.user.user_name = result.data.username;
            console.log(models.user.user_name);
            // models.user.user_name = JSON.parse(result).data.name;
          }
          else{
            models.user.isLogin = false;
          }
          resolve(true);
        });
      });
    },
    Logout:function(){
      return new Promise((resolve, reject)=>{
        return fetch("/api/user",{
          method:"DELETE"
        }).then((response)=>{
          return response.json();
        }).then((result)=>{
          console.log(result);
          models.user.isLogin = null;
          resolve(true);
        });
      });
    },
    registerSuccess:null,
    registerResponse:null,
    Register:function(){
      return new Promise((resolve, reject)=>{
        //reset registerSuccess
        models.user.registerSuccess = null;
        let formElement = document.querySelector("#register-form");
        let name = formElement.name.value;
        let email = formElement.email.value;
        let password = formElement.password.value;
        let repassword = formElement.repassword.value;
        let data = {
            name:name.toString(),
            email:email.toString(),
            password:password.toString(),
            repassword:repassword.toString()
          };
        console.log(data);
        return fetch("/api/user",{
          method:"POST",
          headers: {
            "Content-type":"application/json",
          },
          body: JSON.stringify(data)
        }).then((response)=>{
          return response.json();
        }).then((result)=>{
          console.log(result);
          // result = JSON.parse(result);
          models.user.registerResponse = result.errorMessage;
          if(result.errorExist === false){
            models.user.registerSuccess = true;
          }else{
            models.user.registerSuccess = false;
          }
          resolve(true);
        });
      })
    },
  },
  room:{
    createStatus:null,
    createRespsonse:null,
    createRoom:function(){
      return new Promise((resolve, reject)=>{
        let formElement = document.querySelector("#create-form");
        let chatroomname = formElement.roomname.value;
        let data = {
          roomname:chatroomname.toString(),
          };
        console.log(data);
        return fetch("/api/room",{
          method:"POST",
          headers: {
            "Content-type":"application/json",
          },
          body: JSON.stringify(data)
        }).then((response)=>{
          return response.json();
        }).then((result)=>{
          console.log(result);
          // result = JSON.parse(result);
          if(result.errorExist === false){
            models.room.createStatus = true;
          }else{
            models.room.createStatus = false;
          }
          models.room.createRespsonse = result.errorMessage;
          resolve(true);
        });
      })
    },
    page:1,
    AllRoomList:null,
    getRoomList:function(){
      return new Promise((resolve, reject)=>{
        let url = "/api/rooms?page=" + (models.room.page-1);
        console.log(url);
        return fetch(url,{
          method:"GET"
        }).then((response)=>{
          return response.json();
        }).then((result)=>{
          console.log(result);
          models.room.AllRoomList = result.data;
          resolve(true);
        });
      });
    },
  }
};

let views = {
  user:{
    registerStatus:function(){
      let register_status = document.querySelector(".register-status");
      register_status.style.display = "flex";
      register_status.style.color = "blue";
      if(models.user.registerSuccess){ // register success
        register_status.innerHTML = models.user.registerResponse;

        //清除註冊資訊
        let formElement = document.querySelector("#register-form");
        formElement.name.value = "";
        formElement.email.value = "";
        formElement.password.value = "";
        formElement.repassword.value = "";

      }else{
        // register fail
        register_status.innerHTML = models.user.registerResponse;
        register_status.style.color = "red";

      }
    },
    loginStatus:function(){
      let login_status = document.querySelector(".login-status");
      login_status.style.display = "flex";
      if(models.user.loginSuccess){ // register success
        login_status.innerHTML = models.user.loginResponse;
        login_status.style.color = "blue";

        //清除登入資訊
        document.querySelector(".login-email").value = "";
        document.querySelector(".login-password").value = "";
        // 重新導向 "/"
        window.location.replace('/');

      }else{ // register fail
        login_status.innerHTML = models.user.loginResponse;
        login_status.style.color = "red";
      }
    },
    isLogin:function(){
      //判斷已經登入
      if(models.user.isLogin){
        ///已登入 顯示chatroom & logout bar & username
        let chatroom_box = document.querySelector(".chatroom-main");
        chatroom_box.style.display = "block";

        let logout_bar = document.querySelector(".logout-bar");
        logout_bar.style.display = "flex";

        let username = document.querySelector(".logout-bar-username")
        username.innerHTML = models.user.user_name + ",";
        //隱藏login register box
        let login_box = document.querySelector(".login-box");
        login_box.style.display = "none";

        let register_box = document.querySelector(".register-box");
        register_box.style.display = "none";

      }else{
        //未登入 顯示登入box
        let login_box = document.querySelector(".login-box");
        login_box.display = "block";

        //隱藏註冊欄
        let register_box = document.querySelector(".register-box");
        register_box.display = "none";
      }
    },
    Logout:function(){
      //判斷已經登出
      if(models.user.isLogin == null){
        //顯示登入box
        let login_box = document.querySelector(".login-box");
        login_box.style.display = "block";
        //隱藏註冊box, chatroom and logout bar
        let register_box = document.querySelector(".register-box");
        register_box.style.display = "none";

        let chatroom_box = document.querySelector(".chatroom-main");
        chatroom_box.style.display = "none";

        let logout_bar = document.querySelector(".logout-bar");
        logout_bar.style.display = "none";
      }
    },
    renderLogin:function(){
      //display login-box
      let login_box = document.querySelector(".login-box");
      login_box.style.display = "block";
      //hide register-box
      let register_box = document.querySelector(".register-box");
      register_box.style.display = "none";
    },
    renderRegister:function(){
      //display register-box
      let register_box = document.querySelector(".register-box");
      register_box.style.display = "block";
      //hide login-box
      let login_box = document.querySelector(".login-box");
      login_box.style.display = "none";
      
    }
  },
  room:{
    showCreateRoomBox:function(){
      // hide chatroom list
      document.querySelector(".chatroom-main").style.display = "none";
      // show create box
      document.querySelector(".chatroom-create-main").style.display = "block";
    },
    createRoom:function(){
      // show status 
      document.querySelector(".create-status").innerHTML = models.room.createRespsonse;
      if(models.room.createStatus){
        document.querySelector(".create-status").style.color = "blue";
      }else{
        document.querySelector(".create-status").style.color = "red";
      }
    },
    renderRoomList:function(){
      let roomParent = document.querySelector(".chatroom-name-list-title");
      let ownerParent = document.querySelector(".chatroom-leader-list-title");
      let ManiParent = document.querySelector(".chatroom-click-list-title");
      //clear div
      while(roomParent.hasChildNodes()){
        roomParent.removeChild(roomParent.firstChild)
      }
      while(ownerParent.hasChildNodes()){
        ownerParent.removeChild(ownerParent.firstChild)
      }
      while(ManiParent.hasChildNodes()){
        ManiParent.removeChild(ManiParent.firstChild)
      }
      // Add div
      for(let i = 0; i < models.room.AllRoomList.length;i++) {
        let roomName = models.room.AllRoomList[i].chatroomName;
        let ownerName = models.room.AllRoomList[i].owner;
        // Add Room 
        let newRoomDiv = document.createElement("div");
        newRoomDiv.textContent = roomName;
        newRoomDiv.className = "room-" + (i+1);
        roomParent.appendChild(newRoomDiv);
        //Add Owner
        let newOwnerDiv = document.createElement("div");
        newOwnerDiv.textContent = ownerName;
        newOwnerDiv.className = "owner-" + (i+1);
        ownerParent.appendChild(newOwnerDiv);
        //Add Manipulate
        let newManiDiv = document.createElement("div");
        newManiDiv.textContent = "Enter";
        newManiDiv.className = "Mani-item";
        newManiDiv.setAttribute("id","Mani"+(i+1))
        ManiParent.appendChild(newManiDiv);
      }
      // rander page
      page = document.querySelector(".chatroom-page-number");
      page.innerHTML = models.room.page;
    }
  },
};

let controllers = {
  room:{
    showCreateRoomBox:function(){
      let create_btn = document.querySelector(".chatroom-create");
      create_btn.addEventListener("click", ()=>{
        views.room.showCreateRoomBox();
      }
      )
    },
    backtoRoomList:function(){
      let backtoRoomList_btn = document.querySelector(".back-to-chatroom-list-btn");
      backtoRoomList_btn.addEventListener("click",()=>{
        window.location.reload();
      })
    },
    createRoom:function(){
      let create_btn = document.querySelector(".room-create-btn");
      create_btn.addEventListener("click",()=>{
        console.log("click create")
        let chatroom = document.querySelector("#create-form").roomname.value;
        console.log(chatroom);
        if( chatroom.length < 4){
          // show error alert 
          let create_status = document.querySelector(".create-status");
          create_status.innerHTML = "name length < 4";
          create_status.style.color = "red";

        }else {
          models.room.createRoom().then(()=>{
            views.room.createRoom();
          })
        }
      })
    },
    getRoomList:function(){
      models.room.getRoomList().then(()=>{
        views.room.renderRoomList();
      })
    },
    nextPage:function(){
      let nextPage = document.querySelector(".chatroom-page-next");
      nextPage.addEventListener("click",()=>{
        if(models.room.AllRoomList.length !== 0){
          models.room.page += 1;
        }
        this.getRoomList();
      })
    },
    prevPage:function(){
      let prevPage = document.querySelector(".chatroom-page-previous");
      prevPage.addEventListener("click",()=>{
        models.room.page -= 1;
        if(models.room.page < 1){
          models.room.page = 1;
        }
        this.getRoomList();
      })
    },
  },
  member: {
    checkLogin:function(){
      return new Promise((resolve, reject)=>{
        models.user.checkLogin().then(()=>{
          views.user.isLogin();
          resolve(true);
        });
      })
    },
    logout:function(){
      return new Promise((resolve, reject)=>{
        let logout_btn = document.querySelector(".logout-bar-btn");
        logout_btn.addEventListener("click", ()=>{
          models.user.Logout().then(()=>{
            views.user.Logout();
            resolve(true);
          });
        });
      })
    },
    register:function(){
        let register_btn = document.querySelector(".register-btn");
        register_btn.addEventListener("click", ()=>{
          console.log("click register");
          //判斷規則
          let formElement = document.querySelector("#register-form");
          // let name = formElement.name.value;
          let email = formElement.email.value;
          let password = formElement.password.value;
          let passwordReconfirm = formElement.repassword.value;

          // regular rules
          let emailRule = /^\w+((-\w+)|(\.\w+))*\@[A-Za-z0-9]+((\.|-)[A-Za-z0-9]+)*\.[A-Za-z]+$/;
          let emailCheck = (email.search(emailRule) == 0) ? (true):(false);
          // let nameCheck = (name.length >= 4) ? (true):(false);
          let passwordRule = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,}$/;
          let passwordCheck = (password.search(passwordRule) == 0 ) ? (true):(false);
          // models.user.registerSuccess = emailCheck&&passwordCheck;

          let register_status = document.querySelector(".register-status");
          register_status.style.display = "flex";
            register_status.style.color = "red";
          if(!emailCheck){
            register_status.innerHTML = "please confirm email format";
          }
          else if(!passwordCheck){
            register_status.innerHTML = "please confirm password format";
          }
          else if(password !== passwordReconfirm){
            register_status.innerHTML = "please confirm re-entered password";
          }
          else{
            models.user.Register().then(()=>{
              views.user.registerStatus();
            });
          }
        });
    },
    login:function(){
        let login_btn = document.querySelector(".login-btn");
        login_btn.addEventListener("click", ()=>{
          
          let email = document.querySelector(".login-email").value;
          let password = document.querySelector(".login-password").value;
          if (email.length === 0 || password.length === 0) {
            let login_status = document.querySelector(".login-status");
            login_status.style.display = "flex";
            login_status.style.color = "red";
            login_status.innerHTML ="please enter your email or password"
          }else{
            models.user.Login().then(()=>{
              console.log("login");
              views.user.loginStatus();
            });
          }
        });
    },
    switchtoLogin:function(){
      let switchtoLogin_btn = document.querySelector(".register-login");
      switchtoLogin_btn.addEventListener("click",()=>{
        views.user.renderLogin();
      })
    },
    switchtoRegister:function(){
      let switchtoRegister_btn = document.querySelector(".login-register");
      switchtoRegister_btn.addEventListener("click",()=>{
        console.log("switch to Register")
        views.user.renderRegister();
      })
    },
  },

  init:function(){
    controllers.member.checkLogin().then(()=>{
      controllers.member.register();
      controllers.member.login();
      controllers.member.logout();
      controllers.room.showCreateRoomBox();
      controllers.room.backtoRoomList();
      controllers.room.createRoom();
      controllers.room.getRoomList();
      controllers.room.nextPage();
      controllers.room.prevPage();
    });
    // controllers.member.register();
    // controllers.member.login();
    // controllers.member.logout();
    // controllers.member.logout();
    controllers.member.switchtoLogin();
    controllers.member.switchtoRegister();
  },
};

controllers.init();