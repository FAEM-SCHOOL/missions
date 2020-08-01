в строке 14 нужно прописать данные БД, в которой выполнены комманды:

CREATE TABLE users(
id PRIMORY KEY,
login text,
password text);


CREATE TABLE tasks(
id integer,
user_id integer,
date text,
task text);


Локальный сервер работает на порту 5678, но если у Вас он занят, то Вы можете в строке 17 его изменить.

