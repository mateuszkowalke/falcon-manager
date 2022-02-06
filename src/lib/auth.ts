import cookie from 'cookie';
import * as jwt from 'jsonwebtoken';
import { JWT_SECRET } from '$lib/config';

export const checkUser = (request) => {
	const cookies = cookie.parse(request.headers.get('cookie') || '');
	if (!cookies.token) return null;
	const token = jwt.verify(cookies.token, JWT_SECRET);
	if (!token) return null;
	const user = token;

	return {
		user
	};
}
