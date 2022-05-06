

create table users{
    id serial primary key,
    name varchar not null,
    email varchar not null,
    password varchar not null,
}

INSERT INTO user (name, email, password) VALUES 
('Ronaldo', 'teste@gmail.com' , '123456'),
('Everton', 'teste@gmail.com', '123456'),
('Neymar', 'teste@gmail.com', '123456'),
('Messi', 'teste@gmail.com', '123456'),
('Cristiano', 'teste@gmail.com', '123456'),
('Zidane', 'teste@gmail.com', '123456');
