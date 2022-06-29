
//通过id属性值获取密码文本框
let pwd = document.getElementById('password');
//通过id属性值获取显示密码复选框
let chk = document.getElementById('showPwd');
//显示密码
chk.onchange = function (e) {
	try {
		//通过控件节点的checked属性判断id为showPwd的复选框是否被选中
		if (e.target.checked) {
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

//获取验证码
var getCode = document.getElementById('getCode');
var wait = 60;
function time(btn) {
	if (wait === 0) {
		btn.removeAttribute("disabled");
		btn.innerHTML = "获取验证码";
		wait = 60;
	} else {
		btn.setAttribute("disabled", true);
		btn.setAttribute("style", "pointer-events: none;opacity: 0.6");
		btn.innerHTML = wait + "秒后重试";
		wait--;
		setTimeout(function () {
			time(btn);
		}, 1000);
	}
}
getCode.onclick = function () {
	time(this);
};

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

firstForm.addEventListener("submit", (e) => e.preventDefault());
secondForm.addEventListener("submit", (e) => e.preventDefault());




