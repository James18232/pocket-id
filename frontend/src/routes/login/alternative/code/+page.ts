import type { PageLoad } from './$types';

export const load: PageLoad = async ({ url }) => {
	const redirectParam = url.searchParams.get('redirect') || '/settings';

	const isExternalClient = redirectParam.startsWith('http') && !redirectParam.startsWith(url.origin);

	return {
	code: url.searchParams.get('code'),
	redirect: redirectParam,
	isExternalClient
	};
};