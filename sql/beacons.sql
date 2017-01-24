CREATE TABLE "beacons" (
  "id"        BIGSERIAL PRIMARY KEY,
  "hex"       TEXT   NOT NULL,
  "timestamp" BIGINT NOT NULL,
  "flight"    TEXT,
  "altitude"  INTEGER,
  "speed"     INTEGER,
  "heading"   INTEGER,
  "lat"       REAL,
  "lon"       REAL,
  "src"       TEXT   NOT NULL UNIQUE
);

CREATE INDEX ON "beacons" ("hex");
CREATE INDEX ON "beacons" ("timestamp");
CREATE INDEX ON "beacons" ("flight");
