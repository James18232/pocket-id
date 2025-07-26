import type { RequestHandler } from '@sveltejs/kit';

export const GET: RequestHandler = ({ url }) => {
	const isLight = url.searchParams.get('light') !== 'false';

	const manifest = {
		name: 'PocketID',
		icons: [
			{
        			src: `/api/application-configuration/AppleIcon${isLight ? '' : '?light=false'}`,
				sizes: '180x180',
				type: 'image/png'
			}
		],
		display: 'browser'
	};

	return new Response(JSON.stringify(manifest), {
		headers: {
			'Content-Type': 'application/manifest+json'
		}
	});
};
