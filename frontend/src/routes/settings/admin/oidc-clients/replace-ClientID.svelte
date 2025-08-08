<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { m } from '$lib/paraglide/messages';

	let { data } = $props();
	let client = $state(data.client);

	let newClientIdInput = $state('');
</script>

<div class="space-y-6">
	{#if client}
		<div>
			<label class="block mb-2 font-medium">{m.client_id()} ({m.current()})</label>
			<Input value={client.id} readonly />
		</div>
		<div>
			<label for="newClientIdInput" class="block mb-2 font-medium">{m.client_id()} ({m.new()})</label>
			<Input
				id="newClientIdInput"
				bind:value={newClientIdInput}
				placeholder="Enter new Client ID"
			/>
			<Button class="mt-2">Replace Client ID</Button>
			<!-- Button is not wired yet, since API call is not a focus -->
		</div>
	{:else}
		<div>Loading client...</div>
	{/if}
</div>
