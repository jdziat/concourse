BEGIN;
  ALTER TABLE resource_caches ADD COLUMN version_md5 text;

  UPDATE resource_caches SET version_md5 = md5(version::text);

  DROP INDEX resource_caches_resource_config_id_version_params_hash_uniq;

  ALTER TABLE resource_caches DROP COLUMN version;

  CREATE UNIQUE INDEX resource_caches_resource_config_id_version_md5_params_hash_uniq
  ON resource_caches (resource_config_id, version_md5, params_hash);
COMMIT;
