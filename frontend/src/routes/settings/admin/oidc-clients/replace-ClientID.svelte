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
</script>
<div> test2 </div>
<div {...restProps}>
	<div>
		<label class="block mb-2 font-medium">test</label>
	</div>
	<!-- <div>
		<label for="newClientIdInput" class="block mb-2 font-medium">{m.client_id()} ({m.new()})</label>
		<Input
			id="newClientIdInput"
			value={newClientIdInput | ""}
			placeholder="Enter new Client ID"
		/>
		<Button class="mt-2">Replace Client ID</Button>
	</div> -->
</div>
