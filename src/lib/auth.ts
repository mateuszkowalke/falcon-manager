import cookie from 'cookie';
import * as jwt from 'jsonwebtoken';
import { JWT_SECRET } from '$lib/config';

export const checkUser = async (request) => {
	const cookies = cookie.parse(request.headers.get('cookie') || '');
	if (!cookies.token) return;
	const token = jwt.verify(cookies.token, JWT_SECRET);
	const user = token;

	return {
		user
	};
}
