import { prisma } from '$lib/prisma';
import type { RequestHandler } from '@sveltejs/kit';

export const get: RequestHandler = async function ({ locals, params }) {
	const user = locals.user;
	if (!user) return { status: 401 };

	const aviary = await prisma.aviary.findUnique({
		where: {
			id: parseInt(params.id)
		},
		include: {
		    falcons: true
		}
	});

	const breedingProject = await prisma.breedingProject.findUnique({
		where: {
            id: aviary.breedingProjectId
		}
	});

	if (user.id !== breedingProject.ownerId) return { status: 401 }

	return {
		body: aviary
	};
}

