-- Create a new database called 'messages'
CREATE DATABASE messages

-- Get a list of databases
SELECT datname FROM pg_database
WHERE datistemplate = false

CREATE TABLE message(
   ID SERIAL PRIMARY KEY     NOT NULL,
   From_Msg           VARCHAR(50)    NOT NULL,
   MSG            VARCHAR(256)     NOT NULL,
   RoomId        INT
);