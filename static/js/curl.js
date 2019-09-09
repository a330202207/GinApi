var form = function () {

    var form = $("#form");
    var sumit = $("#submit");

    var url = form.attr("action");
    var saveData = function () {
        sumit.click(function () {
            var postData = form.serialize();
            sumit.button("loading");
            ajax(url, postData)
        });
    }

    var delData = function () {
        $("#table td").on("click", ".del", function() {
            var _self = $(this);
            var url = _self.attr("url");
            var id = _self.attr("data-id");
            $.ajax({
                url : url,
                type : "post",
                data:{"id":id},
                dataType : "json",
                success : function(res) {
                    if (res.code == 200) {
                        layer.msg(res.msg, {icon: 1});
                        _self.parent('td').parent('tr').remove();
                    } else {
                        layer.msg(res.msg, {icon: 2});
                        return false;
                    }
                }
            });
        });
    }

    var ajax = function (url, postData) {
        $.ajax({
            url: url,
            type: "post",
            dataType: "json",
            data: postData,
            success: function (res) {
                if (res.code == 200) {
                    layer.msg(res.msg, {icon: 1})
                    setTimeout(function () {
                        sumit.button('reset');
                        location.reload()
                    }, 2000);
                } else {
                    layer.msg(res.msg, {icon: 2});
                    setTimeout(function () {
                        sumit.button('reset');
                    }, 2000);
                }
            }
        });
    }



    return {
        init: function () {
            saveData();
            delData();
        }
    }
}();


jQuery(document).ready(function () {
    form.init();
});