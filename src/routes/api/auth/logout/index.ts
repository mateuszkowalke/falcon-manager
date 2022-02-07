export function post() {
	return {
		headers: {
			'set-cookie': 'token=deleted; max-age=0;'
		}
	};
}
