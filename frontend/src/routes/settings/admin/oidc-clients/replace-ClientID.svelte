<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import Label from '$lib/components/ui/label/label.svelte';
	import { m } from '$lib/paraglide/messages';
	import type { OidcClient } from '$lib/types/oidc.type';
	import FormInput from '$lib/components/form/form-input.svelte';
	import type { HTMLAttributes } from 'svelte/elements';

	let {
		client,
		...restProps
	}: HTMLAttributes<HTMLDivElement> & {
		client?: OidcClient;
	} = $props();

	let newClientIdInput = '';
	let showClientInput = $state(false);

</script>

<div {...restProps}>
	<Button class="mt-3" variant="secondary" size="sm" onclick={() => (showClientInput = !showClientInput)} type="button">
		Replace Client ID or Client secret
	</Button>

	<!-- {#if showClientInput}
		<div class="mt-4">
			<Input
				id="newClientIdInput"
				bind:value={newClientIdInput}
				placeholder={`${client?.id}`}
			/>
			<Button class="mt-2">Replace Client ID</Button>
		</div>
	{/if} -->

	{#if showClientInput}
		<div class="space-y-3 rounded-lg border p-4">
			<div class="flex items-center justify-between gap-4">
				<Input
					id="newClientIdInput"
					bind:value={newClientIdInput}
					placeholder={`${client?.id}`}
					class="flex-grow"
				/>
				<Button class="mt-0 whitespace-nowrap">Replace Client ID</Button>
			</div>
		</div>
	{/if}
</div>
