create table users
(
	id SERIAL primary key,
	name character varying,
	type character varying
);

create table relationships
(
	id SERIAL primary key,
	user_id int,
	other_id int,
	state character varying,
	type character varying
);
