var Submit = function () {
    var handleSubmit = function () {
        $("#submit").click(function () {
            var form = $('form').serialize();
            var url = $("#save").attr('url');
            $.ajax({
                url : url,
                type : "post",
                dataType : "json",
                data: form,
                success : function(res) {
                    if (res.status == 0) {
                        layer.msg(res.msg, {icon: 1});
                        setTimeout(function () {
                            window.location.href = res.data.url;
                        }, 3000);

                    } else {
                        layer.msg(res.msg, {icon: 2});
                        return false;
                    }
                }
            });
        });
    }

    return {
        init: function () {
            handleSubmit();
        }
    };

}();

jQuery(document).ready(function () {
    Submit.init();
});