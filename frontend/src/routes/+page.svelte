<script lang="ts">
	import { resetMode, setMode } from 'mode-watcher';
	import { Button, buttonVariants } from '$lib/components/ui/button/index.js';
	import * as Resizable from '$lib/components/ui/resizable/index.js';
	import * as Sheet from '$lib/components/ui/sheet/index.js';
	import * as Select from '$lib/components/ui/select/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import Settings from 'lucide-svelte/icons/settings';
	import Label from '$lib/components/ui/label/label.svelte';
	import { toast } from 'svelte-sonner';
	import { invalidateAll } from '$app/navigation';

	import { models } from '$lib/wailsjs/go/models';
	import { SaveCurrentExperiment } from '$lib/wailsjs/go/main/App'

	const { data } = $props();

	let currentExperiment: models.Experiment = $state(data.currentExperiment);

	$effect(() => {
		currentExperiment = data.currentExperiment;
	});

	$effect(() => {
		toast.success(`Loaded ${currentExperiment.name}`);
	});

	$effect(() => {
		const result = SaveCurrentExperiment(currentExperiment).then((result: Boolean) => {
			if (result) {
				toast.success(`Saved ${currentExperiment.name}`);
			} else {
				toast.error(`Failed to save ${currentExperiment.name}`);
			}
		});
	})

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

<div class="flex items-center justify-between">
	<div class="mx-auto">{currentExperiment?.name ?? 'No experiment loaded'}</div>

	<Sheet.Root>
		<Sheet.Trigger class={buttonVariants({ variant: 'ghost' })}>
			<Settings />
		</Sheet.Trigger>

		<Sheet.Content>
			<div class="grid w-full items-center gap-16">
				<Sheet.Header>
					<Sheet.Description>Rev's Election Sim v0.1</Sheet.Description>
				</Sheet.Header>

				<div class="flex flex-col space-y-2">
					<Label for="experimentName">Experiment Name</Label>
					<Input bind:value={currentExperiment.name} id="experimentName" />
				</div>

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
		</Sheet.Content>
	</Sheet.Root>
</div>

<Resizable.PaneGroup direction="horizontal" class="border">
	<Resizable.Pane defaultSize={25}>
		<div class="flex h-full items-center justify-center p-6">
			<span class="font-semibold">Menu</span>
		</div>
	</Resizable.Pane>
	<Resizable.Handle />
	<Resizable.Pane defaultSize={75}>
		<div class="flex h-full items-center justify-center p-6">
			<span class="font-semibold">Content</span>
		</div>
	</Resizable.Pane>
</Resizable.PaneGroup>
