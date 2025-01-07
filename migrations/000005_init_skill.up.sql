BEGIN;

CREATE TABLE IF NOT EXISTS skills (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "desc" text NOT NULL,
  "desc_short" text NULL,
  "user_id" integer NULL, -- author

  "rank" smallint NULL,
  "damage" double precision NOT NULL,
  "element" text NOT NULL,
  "distance" text NOT NULL, -- 'melee', 'distant'
  "type" text NOT NULL, -- 'support', 'attack'
  "main_stat" text NOT NULL, -- 'strength', 'intelligence', 'dexterity'
  "cooldown" smallint NOT NULL,

  "class_id" integer NULL,
  "race_id" integer NULL,

  "is_stack" boolean NOT NULL,
  "hidden" boolean NOT NULL,

  FOREIGN KEY (class_id) REFERENCES CLASSES(id) ON DELETE CASCADE,
  FOREIGN KEY (race_id) REFERENCES RACES(id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES USERS(id) ON DELETE CASCADE
);


INSERT INTO "skills" ("name", "desc", "damage", "element", "distance", "is_stack", "class_id", "race_id", "desc_short", "type", "cooldown", "main_stat", "user_id", "hidden", "rank") VALUES
('Удар',   '',    1,    'phys_damage',  'melee',    'f',    1,  1,  '',    'attack',   3,  'strength', NULL,   't',    1),
('Оглушительный удар',   'Оглушительный удар',   0.5,    'phys_damage',  'melee',    'f',    NULL,   NULL,   'Оглушительный удар',   'attack',   2,  'strength', NULL,   't',    1),
('Поджигающий удар', 'Поджигающий удар', 1,  'fire_damage',  'melee',    'f',    NULL,   NULL,   'Поджигающий удар', 'attack',   3,  'strength', NULL,   't',    1),
('Исцеление',    'Исцеление',    0,  'light_damage', 'distant',  'f',    NULL,   NULL,   'Исцеление',    'support',  3,  'intelligence', NULL,   't',    1);



CREATE TABLE IF NOT EXISTS skills_effects (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,

  "type" text NOT NULL, -- percent, number, coast, shield...
  "attribute" text NOT NULL, -- strength, hp, hp_max
  "value" double precision NOT NULL, -- 1, 0, 0.5, -100

  "condition_attribute" text NOT NULL, -- hp_percent
  "condition" text NOT NULL, -- >, >=, <=...
  "condition_value" double precision NOT NULL, -- 0.5

  "direction" text NOT NULL, -- my, enemy, teammate
  "duration" smallint NOT NULL,

  "skill_id" integer NOT NULL,

  "is_single" boolean NOT NULL, -- for heal, shield
  "is_every_turn" boolean NOT NULL,
  FOREIGN KEY (skill_id) REFERENCES SKILLS(id) ON DELETE CASCADE
);

INSERT INTO "skills_effects" ("skill_id", "type", "attribute", "value", "condition_attribute", "condition", "condition_value", "direction", "duration", "name", "is_single", "is_every_turn") VALUES
(1,	'control',	'stun',	0.2,	'',	'',	0,	'enemy',	1,	'Оглушение',	't',	'f'),
(2,	'control',	'stun',	0.05,	'',	'',	0,	'enemy',	1,	'Оглушение',	'f',	'f'),
(3,	'period',	'fire_damage',	1,	'',	'',	0,	'enemy',	2,	'Горение',	'f',	'f'),
(4,	'percent',	'hp',	0.3,	'',	'',	0,	'teammate',	0,	'Исцеление',	't',	'f');


COMMIT;
