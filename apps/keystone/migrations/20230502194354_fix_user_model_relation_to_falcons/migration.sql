-- DropIndex
DROP INDEX "Falcon_owner_key";

-- CreateIndex
CREATE INDEX "Falcon_owner_idx" ON "Falcon"("owner");
