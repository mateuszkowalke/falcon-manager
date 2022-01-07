import cookie from 'cookie';
import type { Handle } from '@sveltejs/kit';
import * as jwt from 'jsonwebtoken';
import { JWT_SECRET } from '$lib/config';

export const handle: Handle = async ({ request, resolve }) => {
	const cookies = cookie.parse(request.headers.cookie || '');

	if (cookies.token) request.locals.authenticated = true;

	const response = await resolve(request);

	return response;
};

export const getSession = async (request) => {
	const cookies = cookie.parse(request.headers.cookie || '');
	if (!cookies.token) return;
	const token = jwt.verify(cookies.token, JWT_SECRET);
	const user = token;

	return {
		user
	};
};
