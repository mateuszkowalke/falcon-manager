import { prisma } from '$lib/prisma';
import type { User } from '@prisma/client';

export async function get({ locals }: { locals: { user: User } }) {
	const user = locals.user;
	if (!user) return { status: 401 };

	const species = await prisma.species.findMany();

	return {
		body: [...species]
	};
}
