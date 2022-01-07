import { prisma } from '$lib/prisma';
import * as bcrypt from 'bcrypt';
import cookie from 'cookie';
import { JWT_SECRET } from '$lib/config';
import * as jwt from 'jsonwebtoken';

export async function post({ body }) {
	const user = await prisma.user.findFirst({
		where: { name: body.name }
	});
	if (user) return { status: 409 };
	const hashedPass = await bcrypt.hash(body.password, 8).catch((err) => console.error(err));
	if (hashedPass) {
		const role = await prisma.role.findUnique({
			where: {
				name: body.role
			}
		});

		const address = await prisma.address.create({
			data: { ...body.address }
		});

		const user = await prisma.user.create({
			data: {
				name: body.name,
				password: hashedPass,
				email: body.email,
				roleId: role.id,
				addressId: address.id
			},
			include: {
				role: {
					include: {
						permissions: true
					}
				}
			}
		});

		await prisma.address.update({
			where: {
				id: address.id
			},
			data: {
				ownerId: user.id
			}
		});

		const { email, password, createdAt, updatedAt, ...tokenData } = user;

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
}
