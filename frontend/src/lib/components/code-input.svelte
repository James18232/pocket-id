<script lang="ts">
	import { onMount } from 'svelte';

	let {
		value = $bindable(''),
		length = 8,
		disabled = false,
		autofocus = false,
		onsubmit
	}: {
		value: string;
		length?: number;
		disabled?: boolean;
		autofocus?: boolean;
		onsubmit?: () => void;
	} = $props();

	let inputEl: HTMLInputElement | null = $state(null);
	let focused = $state(false);

	// A single input holds the whole value; the boxes are display-only. Using one input instead of
	// one per character means focus never moves while typing, which is what keeps the on-screen
	// keyboard from dismissing after every character on mobile browsers.
	// dashes added in middle <6 chars or every 4 otherwise.

	const chars = $derived(Array.from({ length }, (_, i) => value[i] ?? ''));
	const activeIndex = $derived(Math.min(value.length, length - 1));

	function handleInput(e: Event) {
		const el = e.target as HTMLInputElement;
		value = el.value.slice(0, length);
	}

	function handleKeydown(e: KeyboardEvent) {
		if (e.key === 'Enter' && onsubmit) {
			e.preventDefault();
			onsubmit();
		}
	}

	onMount(() => {
		if (autofocus && !disabled) inputEl?.focus();
	});
</script>

<div
	class={length > 6
		? 'relative max-sm:grid max-sm:grid-cols[repeat(8,40px)] max-sm:justify-items-center gap-x-1 gap-y-3 sm:flex sm:flex-row sm:items-center sm:justify-center sm:gap-x-1.5'
		: 'relative flex flex-nowrap items-center justify-center gap-x-1.5'}
>
	{#each chars as char, i}
		{#if i > 0}
			{#if length > 6 && i % 4 === 0}
				<span class="text-muted-foreground hidden text-2xl font-light sm:inline" aria-hidden="true"
					>–</span
				>
			{:else if length <= 6 && i === Math.floor(length / 2)}
				<span class="text-muted-foreground text-2xl font-light" aria-hidden="true">–</span>
			{/if}
		{/if}

		<div
			class={`border-input bg-background dark:bg-input/30 flex h-12 w-full max-w-[40px] items-center justify-center rounded-lg border text-center text-lg font-bold shadow-xs transition-all ${!disabled && focused && i === activeIndex ? 'border-ring ring-ring/50 ring-[3px]' : ''} ${disabled ? 'opacity-50' : ''}`}
			aria-hidden="true"
		>
			{char}
		</div>
	{/each}

	{#if !disabled}
		<input
			bind:this={inputEl}
			{value}
			oninput={handleInput}
			name="auth_code"
			data-bwignore="true"
			data-lpignore="true"
			data-1p-ignore="true"
			onkeydown={handleKeydown}
			onfocus={() => (focused = true)}
			onblur={() => (focused = false)}
			inputmode="text"
			autocapitalize="none"
			autocomplete="one-time-code"
			autocorrect="off"
			spellcheck="false"
			maxlength={length}
			aria-label="Verification code"
			class="absolute inset-0 h-full w-full cursor-pointer opacity-0 outline-none"
		/>
	{/if}
</div>
