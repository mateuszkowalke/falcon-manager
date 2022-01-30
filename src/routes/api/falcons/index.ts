import { prisma } from '$lib/prisma';
import { checkUser } from '$lib/auth';

export async function get({ request }) {
    const user = request.headers.get('cookie')

    if (!user) return { status: 401 }

    const breedingProjects = await prisma.breedingProject.findMany({
        where: {
            ownerId: user.id
        }
    })
    const falcons = await prisma.falcon.findMany({
        where: {
            breedingProjectId: { in: breedingProjects.map(el => el.id)}
        }
    })

	return {
		body: {
			falcons
		}
	};
}
