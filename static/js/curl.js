var form = function () {

    var form = $("#form");
    var summit = $("#submit");

    var url = form.attr("action");
    var saveData = function () {
        summit.click(function () {
            var postData = form.serialize();
            if ($('#tree').length > 0) {
                var ref = $('#tree').jstree(true);

                var sel = ref.get_selected(false);
                console.log(sel)
                $.each(sel, function (index, value) {
                    postData += "&resource_ids=" + value
                });
            }
            console.log(postData)
            summit.button("loading");
            ajax(url, postData)
        });
    }

    var delData = function () {
        $("#table td").on("click", ".del", function () {
            var _self = $(this);
            var url = _self.attr("url");
            var id = _self.attr("data-id");
            $.ajax({
                url: url,
                type: "post",
                data: {"id": id},
                dataType: "json",
                success: function (res) {
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

    var search = function () {
        var search_summit = $("#search-button");
        var search_form = $("#search");
        var url = search_form.attr("action");

        search_summit.click(function () {
            var key = search_form.find('input[name="keyword"]').val();
            location.href = url + "?keyword=" + key;
        })
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
                        summit.button('reset');
                        location.reload()
                    }, 2000);
                } else {
                    layer.msg(res.msg, {icon: 2});
                    setTimeout(function () {
                        summit.button('reset');
                    }, 2000);
                }
            }
        });
    }

    var handleLogout = function () {
        $("#logout").click(function () {
            var url = $("#logout").attr("data-url");
            $.ajax({
                url: url,
                type: "get",
                success: function (res) {
                    console.log(res);
                    if (res.code == 200) {
                        layer.msg(res.msg, {icon: 1});
                        setTimeout(function () {
                            window.location.href = '/admin/backend_login.html';
                        }, 2000);
                    } else {
                        layer.msg(res.msg, {icon: 2});
                        setTimeout(function () {
                            location.reload()
                        }, 2000);
                        return false;
                    }
                }
            })
        })
    }

    return {
        init: function () {
            saveData();
            delData();
            search();
            handleLogout();
        }
    }
}();


jQuery(document).ready(function () {
    form.init();
});