var Login = function () {

    var checkToken = function () {
        token = sessionStorage.getItem("token");
        if (token) {
            window.location.href = '/admin/index.html';
        }
    }

    var handleLogin = function () {
        checkToken();
        var form = $("#form");
        var sumit = $("#submit");
        sumit.click(function () {
            sumit.button("loading");
            var username = $("#username").val();
            var password = $("#password").val();

            if (username == '') {
                layer.msg('请填写用户名!', {icon: 2});
                sumit.button('reset');
                return
            }

            if (password == '') {
                layer.msg('请填写密码!', {icon: 2});
                sumit.button('reset');
                return
            }

            $.ajax({
                url: "/admin/index.html",
                type: "post",
                dataType: "json",
                data: {
                    username: username,
                    password: password
                },
                success: function (res) {
                    console.log(res);
                    if (res.code == 200) {
                        token = res.data.token;
                        sessionStorage.setItem("token", token);
                        window.location.href = '/admin/index.html?';
                    } else {
                        layer.msg(res.msg, {icon: 2});
                        setTimeout(function () {
                            location.reload()
                        }, 2000);
                        return false;
                    }
                }
            });
        });
    }

    return {
        init: function () {
            handleLogin();
        }
    };

}();

jQuery(document).ready(function () {
    Login.init();
});