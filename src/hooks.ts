import { checkUser } from "$lib/auth";

export async function handle({ event, resolve }) {
	event.locals.user = checkUser(event.request)
	return await resolve(event);
}

export function getSession({ locals }) {
	return {
		user: locals.user && {
			username: locals.user.username,
			email: locals.user.email,
			image: locals.user.image,
			bio: locals.user.bio
		}
	};
}
