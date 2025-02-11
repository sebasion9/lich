-- camel case because
drop table if exists machines;
create table machines(
    id int primary key unique not null,
    name varchar(255) not null,
    last_fetch timestamp,
    created_at timestamp default current_timestamp
);

drop table if exists resources;
create table resources(
    id int primary key unique not null,
    name varchar(255) not null,
    path varchar(255) not null,
    machine_id int not null,
    created_at timestamp default current_timestamp,
    foreign key(machine_id) references machines(id)
);

drop table if exists resourceVersions;
create table resourceVersions(
    id int primary key unique not null,
    resource_id int not null,
    num id not null,
    date timestamp default current_timestamp,
    foreign key(resource_id) references resources(id)
);
