-- Удаляем индекс и внешний ключ связи таблиц district и building
DROP INDEX IF EXISTS idx_building_district;

ALTER TABLE "parser_building" DROP CONSTRAINT IF EXISTS fk_building_district;

ALTER TABLE "parser_building" ALTER COLUMN "district" DROP NOT NULL;

ALTER TABLE "parser_building" ALTER COLUMN "district" DROP DEFAULT;

ALTER TABLE "parser_building" ALTER COLUMN "district" TYPE VARCHAR(255) USING "district"::VARCHAR;

-- Восстанавливаем связь между district и building
ALTER TABLE "parser_apartment" ADD CONSTRAINT "parser_apartment_district_id_34f87205_fk_parser_district_id" 
FOREIGN KEY ("district") REFERENCES "parser_district"("id");