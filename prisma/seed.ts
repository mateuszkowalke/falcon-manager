import { PrismaClient } from '@prisma/client';
import { seedRolesAndPermissions } from './seeds/roles-and-permissions';
import {seedUsers } from './seeds/users'

const prisma = new PrismaClient();

async function main() {
	await seedRolesAndPermissions();
	await seedUsers();
}

main()
	.catch((e) => {
		console.error(e);
		process.exit(1);
	})
	.finally(async () => {
		await prisma.$disconnect();
	});
