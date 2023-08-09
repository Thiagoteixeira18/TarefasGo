CREATE DATABASE IF NOT EXISTS todo;
USE todo;

DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS tarefas;
DROP TABLE IF EXISTS equipes;
DROP TABLE IF EXISTS tarefas_equipe;
DROP TABLE IF EXISTS usuarios_equipe;

CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    prazo VARCHAR(50) NOT NULL,
) ENGINE=InnoDB;

CREATE TABLE equipes (
    id INT AUTO_INCREMENT PRIMARY KEY, 
    nome VARCHAR(50) NOT NULL UNIQUE,
    descricao VARCHAR(100) NOT NULL,

    autor_id INT NOT NULL,
    FOREIGN KEY (autor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE, 
) ENGINE=InnoDB;

CREATE TABLE tarefas_equipe (
    id INT AUTO_INCREMENT PRIMARY KEY,
    tarefa VARCHAR(100) NOT NULL, 
    observacao VARCHAR(300),

    autor_id INT NOT NULL,
    FOREIGN KEY (autor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE, 

    equipes_id INT NOT NULL,
    FOREIGN KEY (equipes_id)
    REFERENCES equipes(id)
    ON DELETE CASCADE, 

    prazo VARCHAR(50) NOT NULL
) ENGINE=InnoDB;

CREATE TABLE usuarios_equipe (
    PRIMARY KEY (equipes_id, usuario_id),
    
    equipes_id INT NOT NULL,
    FOREIGN KEY (equipes_id)
    REFERENCES equipes(id)
    ON DELETE CASCADE,
    
    usuario_id INT NOT NULL,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE
    
) ENGINE=InnoDB;
