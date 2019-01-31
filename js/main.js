$(document).ready(function () {
    $("#submit").click(function () {
        let name = $("#name").val();
        let email = $("#email").val();
        let message = $("#message").val();
        let box = $("#container");
        $("#err").remove();
        switch ("") {
            case name:
                box.prepend("<div id=\"err\">Введите имя</div>");
                break;
            case email:
                box.prepend("<div id=\"err\">Введите Email</div>");
                break;
            case message:
                box.prepend("<div id=\"err\">Введите сообщение</div>");
                break;
            default:
                SendMail(name, email, message)
                break;
        }
    });
});

function SendMail(name, email, message) {
    $.ajax({
        url: "/proccess-form",
        method: "POST",
        data: {
            "name": name,
            "email": email,
            "message": message
        },
        success: function(answer){
            let box = $("#container");
            data = JSON.parse(answer);
            console.log(data);
            if(data['Result']){
                box.html("<div id=\"answer\">Сообщение отправлено</div>");
            }
            else{
                box.prepend("<div id=\"err\">Что то солмалось, сообщение не отправленно :(</div>");
            }
        }
    });
}