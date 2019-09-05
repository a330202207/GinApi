var dateFilter = function () {
    var initPickers = function () {
        var options = {
            rtl: App.isRTL(),
            format: "yyyy-mm-dd",
            language: "zh-CN",          //语言
            orientation: "auto",        //方向
            clearBtn: true,            //显示清除按钮
            forceParse: true,           //是否强制转换不符合格式的字符串
            autoclose: true,            //自动关闭
            todayHighlight: true,       //今天高亮
        };
        var startTime = $("#start_time").val();
        var endTime = $("#end_time").val();
        $('#start_time').datepicker(options).on('changeDate', function () {
            $('#start_time').attr('value', startTime);
            $("#end_time").datepicker('setStartDate', startTime);
            $("#start_time").datepicker('hide');
        });

        $('#end_time').datepicker(options).on('changeDate', function () {
            $('#end_time').attr('value', endTime);
            $("#start_time").datepicker('setEndDate', endTime);
            $("#end_time").datepicker('hide');
        });

        if (startTime !=  null || endTime != null) {
            $('#start_time').datepicker('setDate', startTime);
            $('#end_time').datepicker('setDate', endTime);
        }

    }
    return {
        init: function () {
            initPickers();
        }

    };
}();


jQuery(document).ready(function () {
    dateFilter.init();
});