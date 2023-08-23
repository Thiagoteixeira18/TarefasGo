$('#nova-tarefa').on('submit', criarTarefa);

function criarTarefa(evento) {
    evento.preventDefault();

    $.ajax({
        url:    "/tarefas",
        method: "POST",
        data:   {
            tarefa:      $('#tarefa').val(),
            observacao:  $('#observacao').val(),
            prazo:       $('#prazo').val(),
        }
    }).done(function() {
       window.location = "/home";
    }).fail(function() {
       alert("Erro ao criar publicação!!!");
    }) 
}