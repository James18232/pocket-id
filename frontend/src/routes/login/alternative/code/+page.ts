import type { PageLoad } from './$types';
import AppConfigService from '$lib/services/app-config-service';

export const load: PageLoad = async ({ url }) => {
	const appConfigService = new AppConfigService();
	const appConfig = await appConfigService.list(true);
	return {
		code: url.searchParams.get('code'),
		redirect: url.searchParams.get('redirect') || '/settings',
		appConfig
	};
};