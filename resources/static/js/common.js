$("#signIn").click(function () {
    var $username = $("#inputEmail").val();
    var $password = $("#inputPassword").val();
    console.log("name" + $username)
    $.ajax({
        url: "/login",
        type: "post",
        data: "username=" + $username + "&password=" + $password,
        success: function (info) {
            console.log(info)
            if (info == 1) {
                alert("登录成功")
                location.href = "home"
            } else {
                alert("登录失败")
            }
        },
        error: function () {
            alert("系统异常")
        }
    });
})