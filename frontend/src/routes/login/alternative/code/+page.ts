import type { PageLoad } from './$types';

export const load: PageLoad = async ({ url, fetch }) => {
    const response = await fetch(`${url.origin}/api/application-configuration`);
    const configArray = await response.json();
    
    const appConfigName = configArray.find((item: any) => item.key === 'appName').value;

    return {
        code: url.searchParams.get('code'),
        redirect: url.searchParams.get('redirect') || '/settings',
        appConfigName
    };
};