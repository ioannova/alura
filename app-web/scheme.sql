
create table if not exists produtos (
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
	quantidade integer
)

insert into produtos (nome, descricao, preco, quantidade) values
('Camisa', 'Preta', 19, 10),
('Fone', 'Bom', 99, 5)