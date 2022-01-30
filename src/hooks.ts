import cookie from 'cookie';
import type { Handle } from '@sveltejs/kit';
import { checkUser } from '$lib/auth';

export const handle: Handle = async ({ event, resolve }) => {
	const cookies = cookie.parse(event.request.headers.get('cookie') || '');

	if (cookies.token) event.locals.authenticated = true;

	const response = await resolve(event);

	return response;
};

export const getSession = async (event) => {
	return checkUser(event.request);
};
