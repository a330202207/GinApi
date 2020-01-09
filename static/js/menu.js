var Menu = function () {
    var getData = function () {
        $.ajax({
            url: "/admin/menu/menus",
            type: "get",
            success: function (res) {
                //获取菜单
                getMenu(res.data);
                loadTree(res.data)
            }
        })
    };

    var getMenu = function (data) {
        var data = JSON.parse(data);
        var now_url = window.location.pathname;
        getMenuHtml(data, now_url);
        checkMenu(data, now_url);
    };

    var getMenuHtml = function (data, url) {
        var str = "";
        $.each(data, function (k, v) {
            if (v.children.length > 0) {
                str += `<li class="nav-item" id="top-menu-` + v.id + `" ><a href="javascript:;" class="nav-link nav-toggle">
                            <i class="icon-settings"></i>
                            <span class="title">` + v.text + `</span>
                            <span class="arrow"></span>
                        </a>`;
                str += `<ul class="sub-menu" id="sub-menu-` + v.id + `">`;
                $.each(v.children, function (i, item) {
                    if (item.menu_router == url) {
                        str += `<li class="nav-item start active open" id=menu-` + item.id + `>`
                    } else {
                        str += `<li class="nav-item" id=menu-` + item.id + `>`
                    }
                    str += `<a href="` + item.menu_router + `" class="nav-link ">
                            <i class="icon-list"></i>
                            <span class="title">` + item.text + `</span>
                        </a>
                    </li>`
                });
                str += `</ul></li>`
            }
        });
        $("#menu").append(str);
    };

    var checkMenu = function (data, url) {
        $.each(data, function (k, v) {
            if (v.children.length > 0) {
                checkMenu(v.children, url)
            } else {
                if (v.menu_router == url) {
                    $("#menu-" + v.parent_id).addClass("start active open");
                    $("#menu-" + v.parent_id).parent().css('display', 'block');
                    $("#menu-" + v.parent_id).parent().parent().addClass("start active open");
                }
            }

            if (v.menu_router == url) {
                $("#top-menu-" + v.parent_id).addClass("start active open");
                $("#sub-menu-" + v.parent_id).css('display', 'block');
            }

        });
    };

    var loadTree = function (data) {
        $('#tree').jstree({
            'plugins': ["wholerow", "checkbox"],
            'core': {
                "themes": {
                    "icons": false,
                },
                "data": JSON.parse(data)
            },
        }).on('ready.jstree', function (e, data) {
            var id = $("#id").val();
            $.ajax({
                url: "/admin/role/myMenus?id=" + id,
                type: "get",
                success: function (res) {
                    $('#tree').jstree('open_all');
                    $.each(res.data, function (index, data) {//遍历数据
                        var node = $('#tree').jstree("get_node", data.menu_id);
                        var isLeaf = $('#tree').jstree("is_leaf", node);
                        if (isLeaf) {
                            $('#tree').jstree('check_node', data.menu_id);
                        }
                    });
                }
            })
        });
    };

    return {
        init: function () {
            getData();
        }
    }
}();

jQuery(document).ready(function () {
    Menu.init();
});