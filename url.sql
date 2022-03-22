CREATE TABLE Urls (
	id serial primary key,
	original_url varchar not null,
	expired_date timestamp not null
)

select * from urls