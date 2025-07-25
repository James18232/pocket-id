import type { RequestHandler } from './$types';
import { getAppleIconUrl } from '$lib/utils/apple-icon-util';
import type { RequestEvent } from '@sveltejs/kit';

export const prerender = true; // this should just prerender manifest.json and not treat it as a server route

export const GET: RequestHandler = async () => {
	const logoUrl = getAppleIconUrl();
	const manifest = {
		name: 'PocketID',
		icons: [
			{
				src: logoUrl,
				sizes: 'any'
			}
		],
		display: 'browser',
		background_color: '#000000',
		theme_color: '#000000'
	};

	return new Response(JSON.stringify(manifest), {
		headers: {
			'Content-Type': 'application/manifest+json'
		}
	});
};
