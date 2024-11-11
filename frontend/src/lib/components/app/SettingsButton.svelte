<script lang="ts">
	import Settings from 'lucide-svelte/icons/settings';
	import * as Popover from '$lib/components/ui/popover';
	import { buttonVariants } from '$lib/components/ui/button';
	import * as Select from '$lib/components/ui/select';
	import { Label } from '$lib/components/ui/label';
	import { resetMode, setMode } from 'mode-watcher';

	/////////// THEME  ////////////////
	const themeModes = [
		{ value: 'light', label: 'Light' },
		{ value: 'dark', label: 'Dark' },
		{ value: 'system', label: 'System' }
	];

	let themeMode = $state('system');

	let themeLabel = $derived(() => {
		const mode = themeModes.find((mode) => mode.value === themeMode);
		return mode?.label || 'Select';
	});

	$effect(() => {
		switch (themeMode) {
			case 'light':
				setMode('light');
				break;
			case 'dark':
				setMode('dark');
				break;
			case 'system':
				resetMode();
				break;
		}
	});
</script>


		<Popover.Root>
			<Popover.Trigger class={buttonVariants({ variant: 'ghost' })}>
				<Settings />
			</Popover.Trigger>

			<Popover.Content>
				<div class="grid w-full items-center gap-16">
					<div class="flex flex-col space-y-2">
						<Label for="themeModeSelect">Theme</Label>
						<Select.Root type="single" bind:value={themeMode} name="themeModeSelect">
							<Select.Trigger>
								{themeLabel()}
							</Select.Trigger>
							<Select.Content>
								{#each themeModes as mode}
									<Select.Item value={mode.value}>{mode.label}</Select.Item>
								{/each}
							</Select.Content>
						</Select.Root>
					</div>
				</div>
			</Popover.Content>
		</Popover.Root>
