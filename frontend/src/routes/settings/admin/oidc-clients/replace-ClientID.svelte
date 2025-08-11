<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { m } from '$lib/paraglide/messages';
	import type { OidcClient } from '$lib/types/oidc.type';
	import type { HTMLAttributes } from 'svelte/elements';
	import { TextCursorInput } from '@lucide/svelte';

	let {
		client,
		...restProps
	}: HTMLAttributes<HTMLDivElement> & {
		client?: OidcClient;
	} = $props();

	let newClientIdInput = '';
	let newClientSecretInput = '';
	let showClientInput = $state(false);

</script>

<div {...restProps}>
	<Button class="mt-0" variant="secondary" size="sm" onclick={() => (showClientInput = !showClientInput)} type="button">
	<TextCursorInput class="mr-1 size-4" />
		Replace Client ID / Client secret
	</Button>

	{#if showClientInput}
		<div class="mt-4 space-y-3 rounded-lg border p-4">
			<div class="flex items-center justify-between gap-4">
				<Input
					id="newClientIdInput"
					bind:value={newClientIdInput}
					placeholder={`${client?.id}`}
					class="flex-grow"
				/>
				<Button class="mt-0 whitespace-nowrap" variant="secondary" >Replace Client ID</Button>
			</div>
			<div class="flex items-center justify-between gap-4">
				<Input
					id="newClientSecretInput"
					bind:value={newClientSecretInput}
					placeholder={"••••••••••••••••••••••••••••••••"}
					class="flex-grow"
				/>
				<Button class="mt-0 whitespace-nowrap" variant="secondary">Replace Client secret</Button>
			</div>
		</div>
	{/if}
</div>
