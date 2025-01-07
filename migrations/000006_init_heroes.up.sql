BEGIN;

CREATE TABLE IF NOT EXISTS heroes (
  "id" serial PRIMARY KEY,

  "user_id" integer not null,
  "race_id" integer not null,
  "class_id" integer not null,

  "name" text not null,
  "rank" integer not null,
  "level" integer not null,
  "exp" double precision not null,

  "is_hidden" boolean not null,

  FOREIGN KEY (user_id) REFERENCES USERS(id) ON DELETE CASCADE,
  FOREIGN KEY (race_id) REFERENCES RACES(id) ON DELETE CASCADE,
  FOREIGN KEY (class_id) REFERENCES CLASSES(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS heroes_specs (
  "id" serial PRIMARY KEY,

  "hero_id" integer NOT NULL,

  "accuracy" integer NOT NULL,
  "strength" integer NOT NULL,
  "health" integer NOT NULL,
  "speed" integer NOT NULL,
  "dexterity" integer NOT NULL,
  "soul" integer NOT NULL,
  "intelligence" integer NOT NULL,
  "submissions" integer NOT NULL,

  "critical_rate" double precision NOT NULL,
  "critical_damage" double precision NOT NULL,
  "resistance" double precision NOT NULL,

  "total_spec" integer NOT NULL,
  "free_spec" integer NOT NULL,

  FOREIGN KEY (hero_id) REFERENCES HEROES(id) ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS heroes_skills (
  "id" serial PRIMARY KEY,
  "hero_id" integer NOT NULL,
  "skill_id" integer NOT NULL,
  "level" integer NOT NULL,
  "exp" double precision NOT NULL,

  FOREIGN KEY (hero_id) REFERENCES HEROES(id) ON DELETE CASCADE,
  FOREIGN KEY (skill_id) REFERENCES SKILLS(id) ON DELETE CASCADE
);


COMMIT;