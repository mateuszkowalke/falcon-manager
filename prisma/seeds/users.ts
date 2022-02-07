import prismaClient from '@prisma/client';
import * as bcrypt from 'bcrypt';

const { PrismaClient } = prismaClient;
const prisma = new PrismaClient();

export async function seedUsers() {
	const superAdminRole = await prisma.role.findUnique({
		where: {
			name: 'Super Admin'
		}
	});
	const hashedPass = await bcrypt.hash('superadmin', 8);
	const superAdmin = await prisma.user.create({
		data: {
			email: 'superadmin@falcon-manager.com',
			name: 'superadmin',
			password: hashedPass,
			roleId: superAdminRole.id
		}
	});
	const breedingProject1 = await prisma.breedingProject.create({
		data: {
			name: 'Pietrzwałd',
			vetRegNo: 'todo',
			ownerId: superAdmin.id
		}
	});
	const breedingProject2 = await prisma.breedingProject.create({
		data: {
			name: 'Kościerzyna',
			vetRegNo: 'todo',
			ownerId: superAdmin.id
		}
	});
	await prisma.usersBreedingProjects.createMany({
		data: [
			{
				userId: superAdmin.id,
				breedingProjectId: breedingProject1.id
			},
			{
				userId: superAdmin.id,
				breedingProjectId: breedingProject2.id
			}
		]
	});
}
