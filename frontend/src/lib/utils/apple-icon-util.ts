import { mode } from 'mode-watcher';
import { cachedApplicationAppleIcon } from '$lib/utils/cached-image-util';

export function getAppleIconUrl() {
	const isLightMode = mode.current === 'light';
	return cachedApplicationAppleIcon.getUrl(isLightMode);
}
