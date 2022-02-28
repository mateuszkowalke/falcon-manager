import { prisma } from '$lib/prisma';

export async function get({ locals, params }) {
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
