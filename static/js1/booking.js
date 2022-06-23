document.querySelector('#login-btn').onclick = () => {
    window.location.href = "login"
 }
/********************图片放大镜****************** */
$(function () {
    $(".preview_img").mouseover(function (e) {
        $(this).children(".mask").show();
        $(this).children(".big").show();
    });
    $(".preview_img").mouseout(function () {
        $(this).children(".mask").hide();
        $(this).children(".big").hide();
    });
    $(".preview_img").mousemove(function (e) {
        var x = e.pageX - this.offsetLeft;
        var y = e.pageY - this.offsetTop;
        var maskX = x - parseInt($(this).children(".mask").css("width")) / 2;
        var maskY = y - parseInt($(this).children(".mask").css("height")) / 2;
        var maskMax = parseInt($(this).css("width")) - parseInt($(this).children(".mask").css("width"));
        if (maskX <= 0) {
            maskX = 0;
        } else if (maskX >= maskMax) {
            maskX = maskMax;
        }
        if (maskY <= 0) {
            maskY = 0;
        } else if (maskY >= maskMax) {
            maskY = maskMax;
        }
        $(this).children(".mask").css("left", maskX + 'px');
        $(this).children(".mask").css("top", maskY + 'px');
        var bigMax = parseInt($(".bigimg").css("width")) - parseInt($(".big").css("width"));
        var bigX = maskX * bigMax / maskMax;
        var bigY = maskY * bigMax / maskMax;
        $(".bigimg").css("left", -bigX + 'px');
        $(".bigimg").css("top", -bigY + 'px');
    });

    /**********************图片切换********************** */
    $(".list_item li").mouseover(function () {
        $(this).addClass("current").siblings().removeClass("current");
        $(".preview_img img").attr('src', $(this).children("img")[0].src);
        $(".big img").attr('src', $(this).children("img")[0].src);
    });
    // $(".choose_color a").click(function() {
    //     $(this).addClass("current").siblings().removeClass("current");
    // });
    // $(".choose_version a").click(function() {
    //     $(this).addClass("current").siblings().removeClass("current");
    // });
    // $(".choose_type a").click(function() {
    //     $(this).addClass("current").siblings().removeClass("current");
    // });

    // $(".reduce").mouseover(function() {
    //     if ($(".choose_amount input[type='text']").val() <= 1) {
    //         $(".reduce").css("cursor","not-allowed");
    //     }
    //     else {
    //         $(".reduce").css("cursor","pointer");
    //     }
    // });
    // $(".reduce").click(function() {
    //     if ($(".choose_amount input[type='text']").val() <= 1) {
    //         $(".reduce").css("cursor","not-allowed");
    //     }
    //     else {
    //         $(".reduce").css("cursor","pointer");
    //         num =  $(".choose_amount input[type='text']").val();
    //         $(".choose_amount input[type='text']").val(num-1);
    //     }
    // });
    // $(".add").click(function() {
    //     num =  $(".choose_amount input[type='text']").val();
    //     $(".choose_amount input[type='text']").val(Number(num)+1);
    // });

    /***********************侧边栏（猜你喜欢）***********************/
    $(".tab_list li").mouseover(function () {
        $(this).addClass("current").siblings().removeClass("current");
    });
    /***********************（物品详情）***********************/

    $(".detail_tab_list li").click(function () {
        var index = $(this).index();
        $(this).addClass("current").siblings().removeClass("current");
        $(".item").eq(index).show().siblings().hide();
    });
});

/***************bookingform*************** */
let bookingForm = document.querySelector('.booking-form');

document.querySelector('#booking-btn').onclick = () => {
    bookingForm.classList.add('active');
}

document.querySelector('#close-booking-form').onclick = () => {
    bookingForm.classList.remove('active');
}

let menu = document.querySelector('#menu-btn');
let navbar = document.querySelector('.header .nav');

menu.onclick = () => {
    menu.classList.toggle('fa-times');
    navbar.classList.toggle('active');
}

window.onscroll = () => {
    bookingForm.classList.remove('active');
    menu.classList.remove('fa-times');
    navbar.classList.remove('active');

    if (window.scrollY > 0) {
        document.querySelector('.header').classList.add('active');
    } else {
        document.querySelector('.header').classList.remove('active');
    }
}
/*****************日期选择限制****************8 */
var date = new Date();
var year = date.getFullYear();
var month = ('0' + (date.getMonth() + 1)).slice(-2);
var day = ('0' + (date.getDate())).slice(-2) - 1;
var day1 = ('0' + (date.getDate() + 1)).slice(-2) - 1;
var time = year + '-' + month + '-' + day;
var time1 = year + '-' + month + '-' + day1;

//限制不能选择今天之后的日期（加上属性max）
document.getElementById('time').setAttribute('min', time1);

//限制不能选择今天之后的日期（加上属性max）
document.getElementById('time1').setAttribute('min', time1);







