import { prisma } from '$lib/prisma';
import type { User } from '@prisma/client';
import type { RequestHandler } from '@sveltejs/kit';

export async function get({ locals }: { locals: { user: User } }) {
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
            aviary: true
        }
    });

    return {
        body: [...falcons]
    };
}

export const post: RequestHandler = async function({ request, locals }) {
    const user = locals.user;
    if (!user) return { status: 401 };

    const falcon = await request.json()
    await prisma.falcon.create({
        data: { ...falcon }
    })

    return {
        body: falcon
    };
}
