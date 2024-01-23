create table stocks (
	id UUID not null,
	name varchar(255) not null,
	is_available boolean not null,
    primary key(id)
)