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
        email: $('#email').val(),
        nick: $('#nick').val(),
        senha: $('#senha').val()
    },
})
}