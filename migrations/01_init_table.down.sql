ALTER TABLE "queues" DROP FOREIGN KEY ("customer_id");
ALTER TABLE "queues" DROP FOREIGN KEY ("recipient_id");
ALTER TABLE "users" DROP FOREIGN KEY ("role_id");

DROP TABLE IF EXISTS "queues";
DROP TABLE IF EXISTS "roles";
DROP TABLE IF EXISTS "users";
