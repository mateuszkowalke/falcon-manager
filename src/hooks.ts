import { checkUser } from "$lib/auth";

export async function handle({ event, resolve }) {
    event.locals.user = checkUser(event.request)
    return await resolve(event);
}

export function getSession({ locals }) {
    return {
        user: locals.user && {
            ...locals.user
        }
    };
}
