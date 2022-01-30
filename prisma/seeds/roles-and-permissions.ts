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
	DeleteUser = 'DELETE_USER',

	GetBreedingProjects = 'GET_BREEDING_PROJECTS',
	CreateBreedingProject = 'CREATE_BREEDING_PROJECT',
	UpdateBreedingProject = 'UPDATE_BREEDING_PROJECT',
	DeleteBreedingProject = 'DELETE_BREEDING_PROJECT',

	GetFalcons = 'GET_FALCONS',
	CreateFalcon = 'CREATE_FALCON',
	UpdateFalcon = 'UPDATE_FALCON',
	DeleteFalcon = 'DELETE_FALCON'
}

const prisma = new PrismaClient();

export async function seedRolesAndPermissions() {
	// create role if not exists
	const adminRole = await createRoleIfNotExist(RolesEnum.SUPER_ADMIN);
	const breederRole = await createRoleIfNotExist(RolesEnum.BREEDER);
	const workerRole = await createRoleIfNotExist(RolesEnum.WORKER);

	const getUsersPerm = await createPermissionIfNotExist(PermissionEnum.GetUsers);
	const createUserPerm = await createPermissionIfNotExist(PermissionEnum.CreateUser);
	const updateUserPerm = await createPermissionIfNotExist(PermissionEnum.UpdateUser);
	const deleteUserPerm = await createPermissionIfNotExist(PermissionEnum.DeleteUser);

	const getBreedingProjectsPerm = await createPermissionIfNotExist(PermissionEnum.GetBreedingProjects);
	const createBreedingProjectPerm = await createPermissionIfNotExist(PermissionEnum.CreateBreedingProject);
	const updateBreedingProjectPerm = await createPermissionIfNotExist(PermissionEnum.UpdateBreedingProject);
	const deleteBreedingProjectPerm = await createPermissionIfNotExist(PermissionEnum.DeleteBreedingProject);

	const getFalconsPerm = await createPermissionIfNotExist(PermissionEnum.GetFalcons);
	const createFalconPerm = await createPermissionIfNotExist(PermissionEnum.CreateFalcon);
	const updateFalconPerm = await createPermissionIfNotExist(PermissionEnum.UpdateFalcon);
	const deleteFalconPerm = await createPermissionIfNotExist(PermissionEnum.DeleteFalcon);

	// connect permissions with role
	// Super Admin
	await connectRoleWithPermissions(
		adminRole,
		getUsersPerm,
		createUserPerm,
		updateUserPerm,
		deleteUserPerm,
		getBreedingProjectsPerm,
		createBreedingProjectPerm,
		updateBreedingProjectPerm,
		deleteBreedingProjectPerm,
		getFalconsPerm,
		createFalconPerm,
		updateFalconPerm,
		deleteFalconPerm
	);

	// Breeder
	await connectRoleWithPermissions(
		breederRole,
		getUsersPerm,
		createUserPerm,
		updateUserPerm,
		deleteUserPerm,
		getBreedingProjectsPerm,
		createBreedingProjectPerm,
		updateBreedingProjectPerm,
		deleteBreedingProjectPerm,
		getFalconsPerm,
		createFalconPerm,
		updateFalconPerm,
		deleteFalconPerm
	);

	// Worker
	await connectRoleWithPermissions(
		workerRole,
		getUsersPerm,
		createUserPerm,
		updateUserPerm,
		deleteUserPerm,
		getBreedingProjectsPerm,
		getFalconsPerm,
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
