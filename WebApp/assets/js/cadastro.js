$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault();
    console.log("Dentro da função usuario!");


if ($('#senha').val() !== $('#confirmar-senha').val()) {
    Swal.fire("Ops...", "As senhas não coincidem", "error");
    return;
}

$.ajax({
    url: "/usuarios",
    method: "POST",
    data: {
        nome: $('#nome').val(),
        nick: $('#nick').val(),
        email: $('#email').val(),
        senha: $('#senha').val()
    
}
}).done(function() {
   alert("Usuario cadastrado com sucesso!"); 
}).fail(function() {
    console.log(erro);
  alert("erro ao cadastrar usuario");
});
}