CREATE TYPE type_enum AS ENUM ('impression', 'click', 'visible');


CREATE TABLE IF NOT EXISTS event (
  type_enum type_enum,
  user_agent varchar(250) NOT NULL,
  ip varchar(250) NOT NULL,
  ts timestamp NOT NULL,
  PRIMARY KEY (ts)
);


INSERT INTO event VALUES ('impression', 'mock-user-agent', 'mock-ip', current_timestamp);