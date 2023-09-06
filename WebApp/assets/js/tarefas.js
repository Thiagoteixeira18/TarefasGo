$('#nova-tarefa').on('submit', criarTarefa);
$('.concluir-tarefa').on("click", concluirTarefa);
$('.deletar-tarefa').on("click", deletarTarefa);
$('#atualizar-tarefa').on("click", atualizarTarefa);

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
        Swal.fire("Ops...", "Erro ao criar tarefa", "error");
    }) 
}

function concluirTarefa(evento) {
    evento.preventDefault(); 

    Swal.fire({
        title: "Atenção!",
        text: "Deseja excluir essa tarefa?",
        showCancelButton: true,
        CancelButtonText: "Cancelar",
        icon: "Warning"
    }).then(function(confirmacao) {
        if (!confirmacao.value) return;

    const elementoClicado = $(evento.target);
    const tarefa = elementoClicado.closest('div');
    const tarefaId = tarefa.data('tarefa-id');

    elementoClicado.prop('disabled', true);

    $.ajax({
        url: `/tarefas/${tarefaId}`,    
        method: "DELETE"
    }).done(function() {
        tarefa.fadeOut("slow", function() {
            $(this).remove();
        });    
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao concluir a tarefa", "error");
    });
})
}

function deletarTarefa(evento) {
    evento.preventDefault(); 

    Swal.fire({
        title: "Atenção!",
        text: "Deseja realmente excluir essa tarefa?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao) {
        if (!confirmacao.value) return;
    
    const elementoClicado = $(evento.target);
    const tarefa = elementoClicado.closest('div');
    const tarefaId = tarefa.data('tarefa-id');

    $.ajax({
        url: `/tarefas/${tarefaId}`,
        method: "DELETE"
    }).done(function() {
        tarefa.fadeOut("slow", function() {
            $(this).remove();
        });    
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao excluir a tarefa", "error");
    });
})
}

function atualizarTarefa(evento) {
    $(this).prop('disabled', true);

    const tarefaId = $(this).data('tarefa-id');
    
    $.ajax({
        url: `/tarefas/${tarefaId}`,
        method: "PUT",
        data: {
            tarefa: $('#tarefa').val(),
            observacao: $('#observacao').val(),
            prazo: $('#prazo').val()
        }
    }).done(function() {
        Swal.fire( 
            'Sucesso',
            'Tarefa atualizada com sucesso!',
            'success')
            .then(function() {
                window.location = "/home";
            })
        }).fail(function() {
            Swal.fire("Ops...", "Falha em editar a tarefa!!", "error");
    }).always(function() {
        $('#atualizar-tarefa').prop('disabled', false)
    });
    
} 