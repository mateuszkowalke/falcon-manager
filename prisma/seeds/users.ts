import { PrismaClient } from '@prisma/client'
import * as bcrypt from 'bcrypt';

const prisma = new PrismaClient();

export async function seedUsers() {
    const superAdminRole = await prisma.role.findUnique({
        where: {
            name: 'superadmin'
        }
    })
	const hashedPass = await bcrypt.hash('superadmin', 8)
    const superAdmin = await prisma.user.create({
        data: {
            email: 'superadmin@falcon-manager.com',
            name: 'superadmin',
            password: hashedPass,
            roleId: superAdminRole.id,
        }
    })
}
