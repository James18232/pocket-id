<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { updateClientId } from '$lib/services/oidcService';
	import { Input } from '$lib/components/ui/input';
	import { m } from '$lib/paraglide/messages';
	import type { OidcClient, OidcClientSecretInput, OidcClientMetaData } from '$lib/types/oidc.type';
	import { axiosErrorToast } from '$lib/utils/error-util';

	import type { HTMLAttributes } from 'svelte/elements';
	import { TextCursorInput } from '@lucide/svelte';

	let {
		client,
		...restProps
	}: HTMLAttributes<HTMLDivElement> & {
		client: OidcClient;
	} = $props();

	let newClientIdInput: OidcClientMetaData['id'] = '';
	let newClientSecretInput: OidcClientSecretInput = '';
	let expandUpdateClientIdentifiers = $state(false);

	async function handleUpdateClientId() {
		try {
			const updatedClient = await updateClientId(client.id, newClientIdInput);
			toast.success('Client ID updated successfully');
			newClientIdInput = '';
		} catch (e) {
			axiosErrorToast(e);
		}
	}

</script>

<div {...restProps}>
	<Button class="mt-0" variant="secondary" size="sm" onclick={() => (expandUpdateClientIdentifiers = !expandUpdateClientIdentifiers)} type="button">
	<TextCursorInput class="mr-1 size-4" />
		{m.update()} {m.client_id()} / {m.client_secret()}
	</Button>

	{#if expandUpdateClientIdentifiers}
		<div class="mt-4 space-y-3 rounded-lg border p-4">
			<div class="flex items-center justify-between gap-4">
				<Input
					id="newClientIdInput"
					bind:value={newClientIdInput}
					placeholder={`${client.id}`}
					class="flex-grow"
				/>
				<Button class="mt-0 whitespace-nowrap" variant="secondary" on:click={handleUpdateClientId}>{m.update()} {m.client_id()}</Button>
			</div>
			<div class="flex items-center justify-between gap-4">
				<Input
					id="newClientSecretInput"
					bind:value={newClientSecretInput}
					placeholder={"••••••••••••••••••••••••••••••••"}
					class="flex-grow"
				/>
				<Button class="mt-0 whitespace-nowrap" variant="secondary">{m.update()} {m.client_secret()}</Button>
			</div>
		</div>
	{/if}
</div>
