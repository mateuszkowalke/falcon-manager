/*
  Warnings:

  - You are about to drop the column `falcons` on the `Aviary` table. All the data in the column will be lost.

*/
-- DropForeignKey
ALTER TABLE "Aviary" DROP CONSTRAINT "Aviary_falcons_fkey";

-- DropIndex
DROP INDEX "Aviary_falcons_key";

-- AlterTable
ALTER TABLE "Aviary" DROP COLUMN "falcons";

-- AlterTable
ALTER TABLE "Falcon" ADD COLUMN     "aviary" TEXT;

-- CreateIndex
CREATE INDEX "Falcon_aviary_idx" ON "Falcon"("aviary");

-- AddForeignKey
ALTER TABLE "Falcon" ADD CONSTRAINT "Falcon_aviary_fkey" FOREIGN KEY ("aviary") REFERENCES "Aviary"("id") ON DELETE SET NULL ON UPDATE CASCADE;
