
CREATE TABLE Users (
  id serial not null unique,
  email VARCHAR(100) NOT NULL,
  name VARCHAR(100) NOT NULL,
  password VARCHAR(70) NOT NULL
);


CREATE TABLE Users_0 (
CHECK ( id%2=0 )
) INHERITS (Users) ;

CREATE TABLE Users_1 (
CHECK ( id%2=1 )
) INHERITS (Users) ;


CREATE OR REPLACE FUNCTION users_insert_trigger()
RETURNS TRIGGER AS $$
BEGIN
IF ( NEW.id % 2 = 0) THEN
INSERT INTO Users_0 VALUES (NEW._);
ELSIF ( NEW.id % 2 = 1 ) THEN
INSERT INTO Users_1 VALUES (NEW._);
END IF;
RETURN NULL;
END;

$$
LANGUAGE plpgsql;

CREATE TRIGGER insert_users_trigger
    BEFORE INSERT ON Users
    FOR EACH ROW EXECUTE FUNCTION users_insert_trigger();

