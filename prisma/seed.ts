import { PrismaClient, Role, Permission, RolesOnPermissions } from '@prisma/client';

enum RolesEnum {
	SUPER_ADMIN = 'Super Admin',
	BREEDER = 'Breeder',
	WORKER = 'Worker'
}

enum PermissionEnum {
	GetUsers = 'GET_USERS',
	CreateUser = 'CREATE_USER',
	UpdateUser = 'UPDATE_USER',
	DeleteUser = 'DELETE_USER'
}

const prisma = new PrismaClient();

async function main() {
	seedRolesAndPermissions();
}

main()
	.catch((e) => {
		console.error(e);
		process.exit(1);
	})
	.finally(async () => {
		await prisma.$disconnect();
	});

export async function seedRolesAndPermissions() {
	// create role if not exists
	const adminRole = await createRoleIfNotExist(RolesEnum.SUPER_ADMIN);
	const breederRole = await createRoleIfNotExist(RolesEnum.BREEDER);
	const workerRole = await createRoleIfNotExist(RolesEnum.WORKER);

	const getUsersPerm = await createPermissionIfNotExist(PermissionEnum.GetUsers);
	const createUserPerm = await createPermissionIfNotExist(PermissionEnum.CreateUser);
	const updateUserPerm = await createPermissionIfNotExist(PermissionEnum.UpdateUser);
	const deleteUserPerm = await createPermissionIfNotExist(PermissionEnum.DeleteUser);

	// connect permissions with role
	// Super Admin
	await connectRoleWithPermissions(
		adminRole,
		getUsersPerm,
		createUserPerm,
		updateUserPerm,
		deleteUserPerm
	);

	// Breeder
	await connectRoleWithPermissions(
		breederRole,
		getUsersPerm,
		createUserPerm,
		updateUserPerm,
		deleteUserPerm
	);

	// Worker
	await connectRoleWithPermissions(
		workerRole,
		getUsersPerm,
		createUserPerm,
		updateUserPerm,
		deleteUserPerm
	);
}

async function createRoleIfNotExist(name: string) {
	let role = await prisma.role.findFirst({
		where: { name },
		include: { permissions: true }
	});

	if (!role) {
		role = await prisma.role.create({
			data: { name },
			include: { permissions: true }
		});
	}
	return role;
}

async function createPermissionIfNotExist(name: string) {
	let permission = await prisma.permission.findFirst({
		where: { name },
		include: { roles: true }
	});

	if (!permission) {
		permission = await prisma.permission.create({
			data: { name },
			include: { roles: true }
		});
	}
	return permission;
}

async function connectRoleWithPermissions(
	role: Role & {
		permissions: RolesOnPermissions[];
	},
	...permissions: (Permission & {
		roles: RolesOnPermissions[];
	})[]
) {
	for (const permission of permissions) {
		if (!role.permissions.map(({ permissionId }) => permissionId).includes(permission.id)) {
			await prisma.rolesOnPermissions.create({
				data: {
					roleId: role.id,
					permissionId: permission.id
				}
			});
		}
	}
}
