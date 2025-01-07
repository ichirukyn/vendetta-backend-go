BEGIN;

CREATE TABLE IF NOT EXISTS classes (
  "id" serial PRIMARY KEY,
  "name" varchar not null,
  "desc" text not null,
  "desc_short" text not null,
  "main_attr" text not null,
  "type" text not null,
  "hidden" boolean not null,
  "race_id" integer not null,
  "tag_id" integer not null,

  FOREIGN KEY (tag_id) REFERENCES TAGS(id) ON DELETE CASCADE,
  FOREIGN KEY (race_id) REFERENCES RACES(id) ON DELETE CASCADE
);

INSERT INTO "classes" ("name", "desc", "race_id", "main_attr", "desc_short", "type", "hidden", "tag_id") VALUES
('Воин',	'Человеческий воин, сражающийся на ближней дистанции и использующий физическую силу.',	1,	'strength',	'',	'Воин',	'f',	7),
('Маг',	'Человеческий воин, сражающийся на дальней дистанции и использующий магическую силу.',	1,	'intelligence',	'',	'Маг',	'f',	8),
('Лучник',	'Человеческий воин, сражающийся на средне-дальней дистанции и использующий физическую силу.',	1,	'dexterity',	'',	'Лучник',	'f',	9);

CREATE TABLE IF NOT EXISTS classes_effects (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "type" text NOT NULL,
  "attribute" text NOT NULL,
  "value" double precision NOT NULL,
  "class_id" integer NOT NULL,
  FOREIGN KEY (class_id) REFERENCES CLASSES(id) ON DELETE CASCADE
);


INSERT INTO "classes_effects" ("class_id", "type", "attribute", "value", "name") VALUES
(1,  'number',   'hp_modify',    3,  'Мод. Хп'),
(1,	'number',	'qi_modify',	2,	'Мод. Ки'),
(1,	'number',	'mana_modify',	1,	'Мод. Маны'),
(1,	'percent',	'strength',	0.05,	'Сила'),
(1,	'percent',	'health',	0.025,	'Здоровье'),
(1,	'percent',	'soul',	0.025,	'Дух'),

(2,	'number',	'hp_modify',	2,	'Мод. Хп'),
(2,	'number',	'qi_modify',	1,	'Мод. Ки'),
(2,  'number',   'mana_modify',  3,  'Мод. Маны'),
(2,	'percent',	'intelligence',	0.05,	'Интеллект'),
(2,	'percent',	'speed',	0.025,	'Скорость'),
(2,	'percent',	'soul',	0.025,	'Дух'),

(3,	'number',	'hp_modify',	3,	'Мод. Хп'),
(3,	'number',	'qi_modify',	1,	'Мод. Ки'),
(3,	'number',	'mana_modify',	2,	'Мод. Маны'),
(3,	'percent',	'dexterity',	0.05,	'Ловкость'),
(3,	'percent',	'accuracy',	0.025,	'Меткость'),
(3,	'percent',	'soul',	0.025,	'Дух');

COMMIT;
