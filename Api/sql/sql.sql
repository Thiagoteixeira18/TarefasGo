CREATE DATABASE IF NOT EXISTS todo;
USE todo;

DROP TABLE IF EXISTS usuarios;
DROP TABLE IF EXISTS tarefas;
DROP TABLE IF EXISTS equipes;

CREATE TABLE usuarios (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    senha VARCHAR(100) NOT NULL
) ENGINE=InnoDB;


CREATE TABLE tarefas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    tarefa VARCHAR(100) NOT NULL, 
    observacao VARCHAR(300),

    autor_id INT NOT NULL,
    FOREIGN KEY (autor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE, 

    prazo VARCHAR(50) NOT NULL
) ENGINE=InnoDB;

CREATE TABLE equipes (
    id INT AUTO_INCREMENT PRIMARY KEY, 
    nome VARCHAR(50) NOT NULL UNIQUE,
    descricao VARCHAR(100) NOT NULL,

    autor_id INT NOT NULL,
    FOREIGN KEY (autor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE, 

    tarefas_id INT, 
    FOREIGN KEY (tarefas_id) 
    REFERENCES tarefas(id)
    ON DELETE CASCADE,
   
    participantes_id INT ,
    FOREIGN KEY (participantes_id)
    REFERENCES usuarios(id) 
    ON DELETE CASCADE
) ENGINE=InnoDB;
