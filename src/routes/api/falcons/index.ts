import { prisma } from '$lib/prisma';

export async function get({ locals }) {
    const user = locals.user;
    if (!user) return { status: 401 };

    const breedingProjects = await prisma.breedingProject.findMany({
        where: {
            ownerId: user.id
        }
    });
    const falcons = await prisma.falcon.findMany({
        where: {
            breedingProjectId: { in: breedingProjects.map((el) => el.id) }
        },
        include: {
            species: true,
            aviaries: true
        }
    });

    return {
        body: [
            ...falcons
        ]
    };
}
