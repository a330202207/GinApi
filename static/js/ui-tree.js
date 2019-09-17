var UITree = function () {

    var getTreeData = function () {
        $.ajax({
            url: "/admin/resource/resources",
            type: "get",
            success: function (res) {
                if (res) {
                    data = res
                } else {
                    data = []
                }
                loadTree(res.data)
            }
        })
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
        }).on('ready.jstree', function(e, data){
            var id = $("#id").val();
            $.ajax({
                url: "/admin/role/myResources?id=" + id,
                type: "get",
                success: function (res) {
                    console.log(res.data);
                    $('#tree').jstree('open_all');
                    $.each(res.data,function(index,data){//遍历数据
                        var node = $('#tree').jstree("get_node", data.resource_id);
                        var isLeaf = $('#tree').jstree("is_leaf", node);
                        if(isLeaf){
                            $('#tree').jstree('check_node', data.resource_id);
                        }
                    });
                }
            })
        });
    };

    return {
        //main function to initiate the module
        init: function () {
            getTreeData();
        }
    };

}();

if (App.isAngularJsApp() === false) {
    jQuery(document).ready(function () {
        UITree.init();
    });
}