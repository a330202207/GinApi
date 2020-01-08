var Menu = function () {


    var arr = [];
    var id = '';

    var getData = function () {
        $.ajax({
            url: "/admin/menu/menus",
            type: "get",
            success: function (res) {

                //获取菜单
                getMenu(res.data)
                // arr = getUrl(JSON.parse(res.data), 0)
            }
        })
    };

    var getMenu = function (data) {
        var data = JSON.parse(data);
        console.log(data)
        var now_url = window.location.pathname;
        getMenuHtml(data, now_url);
        var urlData = getUrlData(data, 0);
        // checkMenu(urlData, now_url);
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

    var getUrlData = function (data, pid) {

        $.each(data, function (k, v) {
            if (v.parent_id == 0) {
                id = v.id;
                arr[id] = []
            } else {
                id = pid;
            }

            if (v.menu_router != "") {
                if (id != pid) {
                    arr[id][pid].push(v.menu_router)
                } else {
                    arr[id].push(v.menu_router)
                }


            }
            // console.log(arr)
            // console.log(v.children.length)
            return false;
            if (v.children.length > 0) {
                getUrlData(v.children, id)
            }

        });
        return arr
    };

    var checkMenu = function (data, url) {


        $.each(data, function (k, v) {
            if (v.children.length > 0) {
                checkMenu(v.children, url)

                // $.each(v.children, function (i, item) {
                //     console.log(item)
                // })

            } else {
                console.log(url)
                console.log(v.menu_router)
                if (v.menu_router == url) {

                    $("#menu-" + v.parent_id).addClass("start active open");
                    $("#menu-" + v.parent_id).parent().css('display', 'block');
                    $("#menu-" + v.parent_id).parent().parent().addClass("start active open");
                }

            }

            if (v.menu_router == url) {
                $("#top-menu-" + v.parent_id).addClass("start active open");
                $("#sub-menu-" + v.parent_id).css('display', 'block');
                console.log(v)
            }

        });


        // for (i = 0; i < data.length; i++) {
        //     if (typeof data[i] !== 'undefined' && $.inArray(url, data[i]) == 1) {
        //         $("#top-menu-" + i).addClass("start active open");
        //         $("#sub-menu-" + i).css('display', 'block');
        //         // $("#sub-menu-"+ i).next().addClass("start active open");
        //     }
        // }
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