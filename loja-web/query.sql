create table produtos(
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
	quantidade integer
);

INSERT INTO public.produtos(
	nome, descricao, preco, quantidade)
	VALUES ('Camiseta', 'camiseta dos guri', 50.7, 10);

INSERT INTO public.produtos(
	nome, descricao, preco, quantidade)
	VALUES ('Terno', 'do jamersson bonde', 100, 1);

INSERT INTO public.produtos(
	nome, descricao, preco, quantidade)
	VALUES ('Sapato', 'crazy frog', 201, 2);