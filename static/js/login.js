	//通过name属性值获取密码文本框
	let pwd = document.getElementById('password');
	//通过id属性值获取显示密码复选框
	let chk = document.getElementById('showPwd');
	//显示密码
	chk.onchange = function (e) {
	  try {
		//通过控件节点的checked属性判断id为showPwd的复选框是否被选中
		if(e.target.checked) {
		  //将密码框的type属性值改为：text
		  pwd.type = 'text';
		} else {
		  //将密码框的type属性值改为：password
		  pwd.type = 'password';
		}
	  } catch (error) {
		alert('这个浏览器不能切换控件类型');
	  }
	};

function ChickCode() {
	//获取验证码
	var getCode = document.getElementById('getCode');
	var wait = 60;
	function time(btn){
	if (wait===0) {
		btn.removeAttribute("disabled");
		btn.innerHTML = "Get Code";
		btn.removeAttribute("style")
		wait = 60;
	}else{
		btn.setAttribute("disabled",true);
		btn.setAttribute("style","pointer-events: none;opacity: 0.6");

		btn.innerHTML = "wait "+ wait + " s";
		wait--;
		setTimeout(function(){
		time(btn);
		},1000);
		}
	}
	time(getCode);
}

// function getRandomSixNum() {
// 	let RandomSixStr = ''
// 	for(let i = 0;i < 6; i++) {
// 	  RandomSixStr += String(Math.floor(Math.random()*10))
// 	}
// 	return RandomSixStr
//   }


		document.getElementById('getCode').addEventListener('click',function(){
			if (document.getElementById("email1").value==""){
				alert("请输入邮箱地址")
			}else{
			ChickCode()
			var email = document.getElementById("email1").value
			document.getElementById("email2").value=email
			document.getElementById("submit").click() 
	}
})


const signInBtn = document.getElementById("signIn");
const signUpBtn = document.getElementById("signUp");
const firstForm = document.getElementById("form1");
const secondForm = document.getElementById("form2");
const container = document.querySelector(".container");

signInBtn.addEventListener("click", () => {
	container.classList.remove("right-panel-active");
});

signUpBtn.addEventListener("click", () => {
	container.classList.add("right-panel-active");
});









// firstForm.addEventListener("submit", (e) => e.preventDefault());
// secondForm.addEventListener("submit", (e) => e.preventDefault());


