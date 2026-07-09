UPDATE "parser_building" SET "district" = '1' WHERE "district" IS NULL;

-- Добавление связи поля district с таблицей district для таблицы building
ALTER TABLE "parser_building" ALTER COLUMN "district" TYPE BIGINT USING "district"::BIGINT;

ALTER TABLE "parser_building" ALTER COLUMN "district" SET DEFAULT '1';

ALTER TABLE "parser_building" ALTER COLUMN "district" SET NOT NULL;

ALTER TABLE "parser_building" ADD CONSTRAINT fk_building_district FOREIGN KEY ("district") REFERENCES "parser_district"("id");

CREATE INDEX idx_building_district ON "parser_building" ("district");

-- Дроп связи с таблицей district
ALTER TABLE "parser_apartment" DROP CONSTRAINT parser_apartment_district_id_34f87205_fk_parser_district_id;
