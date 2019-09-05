var Del = function () {
    var handleDel = function () {
        $("#table td").on("click", ".del", function() {
            var _self = $(this);
            var url = _self.attr("url");

            $.ajax({
               url : url,
               type : "get",
               dataType : "json",
               success : function(res) {
                   if (res.status == 0) {
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

    var removeCache = function () {
        $('#remove_cache').on('click', function () {
            var _self = $(this);
            var url = _self.attr("url");
            $.ajax({
                url : url,
                type : "get",
                dataType : "json",
                success : function(res) {
                    if (res.status == 0) {
                        layer.msg(res.msg, {icon: 1});
                    } else {
                        layer.msg(res.msg, {icon: 2});
                        return false;
                    }
                }
            });
        })
    }

    return {
        init: function () {
            handleDel();
            removeCache();
        }
    };

}();

jQuery(document).ready(function () {
    Del.init();
});