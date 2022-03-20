import { prisma } from '$lib/prisma';
import type { User } from '@prisma/client';

export async function get({ locals }: { locals: { user: User } }) {
	const user = locals.user;
	if (!user) return { status: 401 };

	const breedingProjects = await prisma.breedingProject.findMany({
		where: {
			ownerId: user.id
		}
	});
	const aviaries = await prisma.aviary.findMany({
		where: {
			breedingProjectId: { in: breedingProjects.map((el) => el.id) }
		}
	});

	return {
		body: [...aviaries]
	};
}
