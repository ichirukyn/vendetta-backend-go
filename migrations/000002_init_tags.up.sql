BEGIN;

CREATE TABLE IF NOT EXISTS tags (
  "id" serial PRIMARY KEY,
  "name" varchar not null,
  "desc" text not null,
  "priority" smallint not null,
  "strength" double precision NOT NULL,
  "health" double precision NOT NULL,
  "speed" double precision NOT NULL,
  "dexterity" double precision NOT NULL,
  "accuracy" double precision NOT NULL,
  "intelligence" double precision NOT NULL,
  "submission" double precision NOT NULL,
  "soul" double precision NOT NULL
);

INSERT INTO tags ("name", "desc", "priority", "strength", "health", "speed", "dexterity", "accuracy", "intelligence", "submission", "soul") VALUES
('Люди', '', 3, 0.2, 0, 0, 0, 0, 0, 0, 0.2),
('Эльфы', '', 3, 0, 0, 0, 0.2, 0, 0, 0, 0.2),
('Дворфы', '', 3, 0.2, 0.2, 0, 0, 0, 0, 0, 0),
('Небожители', '', 3, 0, 0, 0, 0, 0, 0.2, 0, 0.2),
('Демоны', '', 3, 0, 0, 0, 0, 0, 0.2, 0, 0.2),
('Нежить', '', 3, 0, 0, 0, 0, 0, 0.2, 0.2, 0),
('Воин',	'',	4,	0.2,	0.25,	0.1,	0,	0,	0,	0,	0),
('Маг',	'',	4,	0,	0.2,	0,	0,	0.15,	0.2,	0,	0.2),
('Лучник',	'',	4,	0,	0.15,	0,	0.2,	0.15,	0,	0,	0);

COMMIT;