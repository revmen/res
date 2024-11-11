<script lang="ts">
	import * as Popover from '$lib/components/ui/popover';
	import { buttonVariants } from '$lib/components/ui/button';
	import Button from '../ui/button/button.svelte';
	import Input from '../ui/input/input.svelte';

	type Props = {
		value: string | null;
	};

	let { value = $bindable() }: Props = $props();
</script>

<Popover.Root>
    
	<Popover.Trigger class={buttonVariants({ variant: 'ghost' })}>
		<Button variant="ghost">{value}</Button>
	</Popover.Trigger>

	<Popover.Content>
		<div class="flex flex-row space-x-4">
			<Input bind:value />
			<Popover.Close class={buttonVariants({ variant: 'outline' })}>OK</Popover.Close>
		</div>
	</Popover.Content>

</Popover.Root>
