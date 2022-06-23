// let loginForm = document.querySelector('.login-form');

document.querySelector('#login-btn').onclick = () => {
   window.location.href = "login"
}
   
// document.querySelector('#close-login-form').onclick = () =>{
//    loginForm.classList.remove('active');
// }

let menu = document.querySelector('#menu-btn');
let navbar = document.querySelector('.header .nav');

menu.onclick = () => {
   menu.classList.toggle('fa-times');
   navbar.classList.toggle('active');
}

window.onscroll = () => {
   // loginForm.classList.remove('active');
   menu.classList.remove('fa-times');
   navbar.classList.remove('active');

   if (window.scrollY > 0) {
      document.querySelector('.header').classList.add('active');
   } else {
      document.querySelector('.header').classList.remove('active');
   }
}

/******************排行榜************* */
$(function () {
   $("li").hover(function () {
      $(this).addClass("current")
   }, function () {
      $(this).removeClass("current")
   })
});

/*****************排序******************* */
$(".tab_list li").click(function () {
   var index = $(this).index();
   $(this).addClass("current").siblings().removeClass("current");
   $(".box-container").eq(index).show().siblings().hide();
});

/******************iframe************************ */
// var ifm=document.getElementById("iframe1Id");
// ifm.height=document.documentElement.scrollHeight;
function partRefresh1() {
   document.getElementById("iframe1Id").src = "room1"; // 方法一: 通过和替换iframe的src来实现局部刷新
}
function partRefresh2() {
   document.getElementById("iframe2Id").src = "room2";
}
function partRefresh3() {
   document.getElementById("iframe3Id").src = "room3";
}









