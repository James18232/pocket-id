import { mode } from 'mode-watcher';

export function getLightDark() {
	const isLightMode = mode.current === 'light';
	return isLightMode;
}
