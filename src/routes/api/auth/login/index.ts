import { prisma } from '$lib/prisma';
import * as bcrypt from 'bcrypt';
import cookie from 'cookie';
import * as jwt from 'jsonwebtoken';
import { JWT_SECRET } from '$lib/config';

export async function post({ request }) {
    const body = await request.json();
	const user = await prisma.user.findUnique({
		where: {
			name: body.name
		},
		include: {
			role: {
				include: {
					permissions: true
				}
			}
		}
	});

	if (!user) {
		return { status: 401 };
	}

	const passMatch = await bcrypt
		.compare(body.password, user.password)
		.catch((err) => console.error(err));

	if (!passMatch) {
		return {
			status: 401
		};
	}

	const { password, createdAt, updatedAt, ...tokenData } = user;

	const token = jwt.sign(tokenData, JWT_SECRET, { expiresIn: '1h' });

	const headers = {
		'Set-Cookie': cookie.serialize('token', token, {
			httpOnly: true,
			maxAge: 60 * 60,
			sameSite: true,
			path: '/'
		})
	};

	return {
		headers,
		body: {
			user
		}
	};
}
