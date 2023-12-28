create table users
(
	id_users serial not null constraint pk_users primary key,
	email varchar(25) not null
	constraint uq_email_users unique,
	constraint ch_email_users check (email like '%@%.%'),
	password_hash varchar not null,
	is_admin boolean not null default false
);
