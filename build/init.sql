CREATE TYPE type_enum AS ENUM ('impression', 'click', 'visible');

DROP TABLE IF EXISTS events;

CREATE TABLE IF NOT EXISTS events (
  id SERIAL  PRIMARY KEY,
  type_enum type_enum,
  user_agent varchar(250) NOT NULL,
  ip varchar(250) NOT NULL,
  ts timestamp NOT NULL
);


INSERT INTO events(type_enum, user_agent, ip, ts) VALUES ('impression', 'mock-user-agent', 'mock-ip', current_timestamp);
INSERT INTO events(type_enum, user_agent, ip, ts) VALUES ('click', 'mock-user-agent', 'mock-ip', current_timestamp);
INSERT INTO events(type_enum, user_agent, ip, ts) VALUES ('visible', 'mock-user-agent', 'mock-ip', current_timestamp);
INSERT INTO events(type_enum, user_agent, ip, ts) VALUES ('visible', 'mock-user-agent', 'mock-ip', '2016-06-22 19:10:25-07');
INSERT INTO events(type_enum, user_agent, ip, ts) VALUES ('visible', 'mock-user-agent', 'mock-ip', '2021-04-30 09:49:27.417975');
INSERT INTO events(type_enum, user_agent, ip, ts) VALUES ('visible', 'mock-user-agent', 'mock-ip', '2021-04-30 09:49:27.417975');
INSERT INTO events(type_enum, user_agent, ip, ts) VALUES ('visible', 'mock-user-agent', 'mock-ip', '2023-04-30 09:49:27.417975');
INSERT INTO events(type_enum, user_agent, ip, ts) VALUES ('visible', 'mock-user-agent', 'mock-ip', '2024-04-30 09:49:27.417975');
