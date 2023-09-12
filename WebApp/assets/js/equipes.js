$('#nova-equipe').on("submit", criarEquipe);
$('#atualizar-tarefa-equipe').on("click", editarTarefaDeEquipe);
$('.deletar-equipe').on("click", deletarEquipe);

function criarEquipe(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/equipes",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            observacao: $('#descricao').val(),
        }
    }).done(function() {
        window.location = "/equipes";
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao criar a equipe!!!", "error");
    })
}

function editarTarefaDeEquipe(evento) {
    $(this).prop('disabled', true);

    const tarefaId = $(this).data('tarefa-equipe-id');
    
    $.ajax({
        url: `/tarefas/${tarefaId}/equipe`,
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
                window.location = "/equipe";
            });
        }).fail(function() {
            Swal.fire("Ops...", "Falha em editar a tarefa!!", "error");
    }).always(function() {
        $('#atualizar-tarefa').prop('disabled', false)
    });
    
}

function deletarEquipe(evento) {
    evento.preventDefault(); 

    Swal.fire({
        title: "Atenção!",
        text: "Deseja realmente excluir essa equipe? Essa ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao) {
        if (!confirmacao.value) return;
    
   
    const elementoClicado = $(evento.target);
    const equipe = elementoClicado.closest('div');
    const equipeId = equipe.data('equipe-id');

    $.ajax({
        url: `/equipes/${equipeId}`,
        method: "DELETE"
    }).done(function() {
        equipe.fadeOut("slow", function() {
            $(this).remove();
        });    
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao excluir a equipe", "error");
    });
})
}
