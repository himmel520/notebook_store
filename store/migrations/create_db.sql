create table systems 
(
	id_systems serial not null constraint pk_systems primary key,
	name varchar(30) not null
	constraint uq_name_systems unique
);

create table screens
(
	id_screens serial not null constraint pk_screens primary key,
	size_inches numeric(4, 2) not null
	constraint ch_range_screens check (size_inches > 0 and size_inches < 100),
	resolution varchar(9) not null
);

create table processors
(
	id_processors serial not null constraint pk_processors primary key,
	model varchar(30) not null,
	speed_ghz numeric(3,2) not null
	constraint ch_positive_processors check (speed_ghz > 0)
);

create table storages
(
	id_storages serial not null constraint pk_storages primary key,
	type_storage varchar(10) not null,
	size_gb int not null
	constraint ch_positive_storages check (size_gb > 0)
);

create table rams
(
	id_rams serial not null constraint pk_rams primary key,
	size_gb int not null
	constraint ch_positive_rams check (size_gb > 0)
);

create table notebooks
(
	id_notebooks serial not null constraint pk_notebooks primary key,
	systems_id int not null references systems(id_systems),
	screens_id int not null references screens(id_screens),
	processors_id int not null references processors(id_processors),
	storages_id int not null references storages(id_storages),
	rams_id int not null references rams(id_rams),
	model varchar(30) not null,
	description varchar not null,
	price numeric(10, 2) not null
	constraint ch_positive_notebooks check (price > 0)
);