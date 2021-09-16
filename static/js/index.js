let models = {
  user:{
    loginSuccess:null,
    useGoogleLogin:false,
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
          if(result.ok){
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
          if(result != null){
            models.user.isLogin = true;
            models.user.user_name = JSON.parse(result).data.name;
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
    Register:function(){
      return new Promise((resolve, reject)=>{
        //reset registerSuccess
        models.user.registerSuccess = null;
        let formElement = document.querySelector("#register-form");
        let name = formElement.name.value;
        let email = formElement.email.value;
        let password = formElement.password.value;
        let data = {
            name:name.toString(),
            email:email.toString(),
            password:password.toString()
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
          result = JSON.parse(result);
          if(result.ok){
            models.user.registerSuccess = true;
          }else{
            models.user.registerSuccess = false;
          }
          // console.log(result);
          // console.log(models.user.registerSuccess);
          resolve(true);
        });
      })
    },
  },
};

let views = {
  user:{
    registerStatus:function(){
      let register_status = document.querySelector(".register-status");
      register_status.style.display = "flex";
      if(models.user.registerSuccess){ // register success
        register_status.innerHTML = "註冊成功，請登入";
        register_status.style.color = "blue";

        //清除註冊資訊
        let formElement = document.querySelector("#register-form");
        formElement.name.value = "";
        formElement.email.value = "";
        formElement.password.value = "";

      }else{
        // register fail
        let formElement = document.querySelector("#register-form");
        let name = formElement.name.value;
        let email = formElement.email.value;
        let password = formElement.password.value;
        //其中為空
        if(name == "" || email == "" || password == ""){
          register_status.innerHTML = "註冊失敗，請確認輸入";
          register_status.style.color = "red";
        }
        else{
          register_status.innerHTML = "註冊失敗，電子信箱已被註冊";
          register_status.style.color = "red";
        }
      }
    },
    loginStatus:function(){
      let login_status = document.querySelector(".login-status");
      login_status.style.display = "flex";
      if(models.user.loginSuccess){ // register success
        login_status.innerHTML = "登入成功";
        login_status.style.color = "blue";

        //清除登入資訊
        document.querySelector(".login-email").value = "";
        document.querySelector(".login-password").value = "";
        // 重新導向 "/"
        window.location.replace('/');

      }else{ // register fail
        login_status.innerHTML = "登入失敗，帳號或密碼錯誤";
        login_status.style.color = "red";
      }
    },
    isLogin:function(){
      //判斷已經登入
      if(models.user.isLogin){
        ///已登入 顯示chatroom

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
        login_box.style.display = "flex";
        //隱藏註冊box
        let register_box = document.querySelector(".register-box");
        register_box.style.display = "none";
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
};

let controllers = {
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
        let logout_btn = document.querySelector("#logout-btn");
        logout_btn.addEventListener("click", ()=>{
          //Goole logout
          var auth2 = gapi.auth2.getAuthInstance();
          auth2.signOut().then(function () {
            console.log('User signed out.');
          });
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
          let name = formElement.name.value;
          let email = formElement.email.value;
          let password = formElement.password.value;

          // regular rules
          let emailRule = /^\w+((-\w+)|(\.\w+))*\@[A-Za-z0-9]+((\.|-)[A-Za-z0-9]+)*\.[A-Za-z]+$/;
          let emailCheck = (email.search(emailRule) == 0) ? (true):(false);
          // let nameCheck = (name.length >= 4) ? (true):(false);
          let passwordCheck = (password.length > 6) ? (true):(false);
          models.user.registerSuccess = emailCheck&&passwordCheck;
          if(!models.user.registerSuccess){
            let register_status = document.querySelector(".register-status");
            register_status.style.display = "flex";
            register_status.innerHTML = "請確認信箱格式或密碼長度小於6";
            register_status.style.color = "red";
          }else{
            models.user.Register().then(()=>{
              console.log("tstet");
              views.user.registerStatus();
            });
          }
        });
    },
    login:function(){
        let login_btn = document.querySelector(".login-btn");
        login_btn.addEventListener("click", ()=>{
          models.user.Login().then(()=>{
            console.log("login");
            views.user.loginStatus();
          });
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
    // controllers.member.checkLogin().then(()=>{
    //   controllers.member.register();
    //   controllers.member.login();
    //   controllers.member.logout();
    // });
    controllers.member.register();
    controllers.member.login();
    controllers.member.logout();
    controllers.member.switchtoLogin();
    controllers.member.switchtoRegister();
  },
};

controllers.init();