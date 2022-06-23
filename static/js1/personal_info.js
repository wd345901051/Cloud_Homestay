function checkUserName() {
    var username = document.myform.username.value;
    if (username.match(/^\w{6,18}$/) == null) {
        alert("请输入：6-18位有效字符（字母、数字、下划线）");
        return false;
    }
    return true;
}

function checkPSW1() {
    var psw1 = document.myform.psw1.value;
    if (psw1.length < 6 || psw1.length > 18) {
        alert("请输入：6-18位任意字符");
        return false;
    }
    return true;
}

function checkPSW2() {
    var psw1 = document.myform.psw1.value;
    var psw2 = document.myform.psw2.value;
    if (psw2 == "") {
        alert("请与第一次输入的密码一致");
    } else if (psw1 == psw2) {
        return true;
    } else {
        alert("前后两次输入密码不一致，请重新输入");
        return false;
    }

    if (psw2.length < 6 || psw2.length > 18) {
        alert("请输入：6-18位任意字符");
        return false;
    }

}

function checkAge() {
    //大于0的任意两位整数
    var age = document.myform.age.value;
    if (age.match(/^[0-9]{2}$/) == null || age == "00") {
        alert("请输入：大于0的任意两位整数");
        return false;
    }
    return true;
}

function checkPhone() {
    //由1开头的11位整数
    var phone = document.myform.phone.value;
    if (phone.match(/1[0-9]{10}/)) {
        return true;
    } else {
        alert("请输入：由1开头的11位整数");
        return false;
    }

}

function checkEmail() {
    //有效的Email地址
    //以数字字母开头，中间可以是多个数字字母下划线或者-
    //然后是@，后面是数字字母
    //然后是.加2-4个字母结尾
    var email = document.myform.email.value;
    if (email.match(/\w+@[a-zA-Z0-9]+\.([a-zA-Z]{3,})/) == null) {
        alert("请输入：有效的Email地址");
        return false;
    }
    return true;
}

function doSubmit() {
    if (checkUserName() && checkPSW1() && checkPSW2() && checkAge() && checkPhone() && checkEmail()) {
        alert("提交成功");
    } else {
        alert("提交失败");
    }
}