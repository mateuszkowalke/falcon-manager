import { prisma } from '$lib/prisma';
import type { RequestHandler } from '@sveltejs/kit';

export const get: RequestHandler = async function ({ locals, params }) {
	const user = locals.user;
	if (!user) return { status: 401 };

	const falcon = await prisma.falcon.findUnique({
		where: {
			id: parseInt(params.id)
		},
		include: {
			species: true,
			aviaries: true
		}
	});

	const breedingProject = await prisma.breedingProject.findUnique({
		where: {
            id: falcon.breedingProjectId
		}
	});

	if (user.id !== breedingProject.ownerId) return { status: 401 }

	return {
		body: falcon
	};
}

export const patch: RequestHandler = async function ({ request, locals, params }) {
	const user = locals.user;
	if (!user) return { status: 401 };

	const falcon = await prisma.falcon.findUnique({
		where: {
			id: parseInt(params.id)
		},
	});

	const breedingProject = await prisma.breedingProject.findUnique({
		where: {
            id: falcon.breedingProjectId
		}
	});

	if (user.id !== breedingProject.ownerId) return { status: 401 }

    const updatedFalcon = await request.json()
	await prisma.falcon.update({where: {id: falcon.id}, data: {...updatedFalcon}})

	return {
		body: falcon
	};
}
