$(function () {


    window.onbeforeunload= function(event) {

        $.post(
            "/closepage",
        );
    };

    $(".submitBtn").click(function () {
        $(".errorinfo").html("");
        var name = $(".formname").val().trim();
        var pw = $(".formpw").val().trim();
        if(!name || !pw){

            showErrorInfo("不能为空~");
            return false ;
        }

        $.post(
            "/user/login",
            {name:name,pw:pw},
            function (data,status) {
                
            }
        );
    });
    function showErrorInfo(str){
        $(".errorinfo").html(str);
    }


});
