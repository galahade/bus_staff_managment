ALTER TABLE `bus_system`.`staff`
  ADD COLUMN `is_resign` TINYINT(4) UNSIGNED NULL DEFAULT 0 AFTER `deleted_at`,
  ADD COLUMN `resign_date` DATE NULL AFTER `is_resign`;

ALTER TABLE bus_basic ADD CONSTRAINT uc_engine_no UNIQUE (engine_no);
