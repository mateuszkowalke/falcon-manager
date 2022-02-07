import { prisma } from '$lib/prisma';

export async function get({ locals }) {
    const user = await prisma.user.findFirst({
        where: { name: locals.user.name }
    });

    if (!user) return { status: 401 };

    const breedingProjects = await prisma.breedingProject.findMany({
        where: {
            ownerId: user.id
        }
    });
    const falcons = await prisma.falcon.findMany({
        where: {
            breedingProjectId: { in: breedingProjects.map((el) => el.id) }
        }
    });

    return {
        body: [
            ...falcons
        ]
    };
}
