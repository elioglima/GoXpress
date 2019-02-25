function enviar(chamada) {
    $("#FormGeral").action = chamada;
    $("#FormGeral").submit();
}

function AcessoSistema() {

    var data_json = {
        usuario: $("#usuario").val(),
        senha:$("#senha").val()
    }
    var token = "102030";

    $.ajax({
        url: '/api/acesso',
        type: 'post',
        dataType: 'json',
        contentType: "application/json; charset=utf-8",
        traditional: true,
        success: function (data) {
            
            console.log(data.Token);
        },

        data: JSON.stringify(data_json),
        error: function (exception) { 
            alert('Exeption:' + JSON.stringify(exception)); 
        },
        beforeSend: function ( xhr ) {
            xhr.setRequestHeader( 'X-CSRF-Token', token );
          }
    });
}
