$(function () {
    baseApp.init();
})
var baseApp = {
    init: function () {
        this.initAside()
        this.confirmDelete()
        this.resizeIframe()
        this.changeStatus()
        this.changeNum()
    },
    initAside: function () {
        $('.aside>li:nth-child(1) ul,.aside>li:nth-child(2) ul').hide()
        $('.aside h4').click(function () {
            $(this).siblings('ul').slideToggle();
        })
    },
    //设置iframe的高度
    resizeIframe: function () {
        $("#rightMain").height($(window).height() - 80)
    },
    // 删除提示
    confirmDelete: function () {
        $(".delete").click(function () {
            var flag = confirm("您确定要删除吗?")
            return flag
        })
    },
    changeStatus: function () {
        $(".chStatus").click(function () {
            var id = $(this).attr("data-id")
            var table = $(this).attr("data-table")
            var field = $(this).attr("data-field")
            var el = $(this)
            $.post("/admin/changeStatus", {id: id, table: table, field: field}, function (response) {
                if (response.success) {
                    if (el.attr("src").indexOf("yes") != -1) {
                        el.attr("src", "/static/admin/images/no.gif")
                    } else {
                        el.attr("src", "/static/admin/images/yes.gif")
                    }
                }
            })
        })
    },
    changeNum: function () {
        $(".chSpanText").click(function () {
            // 1、获取el 以及el里面的属性值
            var id = $(this).attr("data-id")
            var table = $(this).attr("data-table")
            var field = $(this).attr("data-field")
            var value = $(this).html().trim()
            var spanEl = $(this)
            //2、创建一个input的dom节点   var input=$("<input value='' />");
            var input = $("<input style='width:60px'  value='' />");
            // 3、把input放在el里面   $(this).html(input);
            $(this).html(input);
            //4、让input获取焦点  给input赋值    $(input).trigger('focus').val(val);
            $(input).trigger("focus").val(value);
            //5、点击input的时候阻止冒泡
            $(input).click(function (e) {
                e.stopPropagation();
            })
            //6、鼠标离开的时候给span赋值,并触发ajax请求
            $(input).blur(function () {
                var inputValue = $(this).val()
                spanEl.html(inputValue)
                //触发ajax请求
                $.post("/admin/changeCellValue", {id: id, table: table, field: field, value: inputValue}, function (response) {
                    console.log(response)
                })
            })
        })

    }
}