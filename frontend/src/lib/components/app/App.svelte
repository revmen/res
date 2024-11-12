<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Resizable from '$lib/components/ui/resizable/index.js';
	import * as Select from '$lib/components/ui/select/index.js';

	import { Input } from '$lib/components/ui/input/index.js';
	import Label from '$lib/components/ui/label/label.svelte';
	import { toast } from 'svelte-sonner';
	import SettingsButton from './SettingsButton.svelte';
	import TextPopover from './TextPopover.svelte';
	import ExpMenu from './ExpMenu.svelte';
	import ExpMap from './map/ExpMap.svelte';

	import { models } from '$lib/wailsjs/go/models';
	import { ClipboardSetText } from '$lib/wailsjs/runtime/runtime';

	type Props = {
		experiment: models.Experiment;
		closeExperiment: () => void;
	};

	let { experiment = $bindable(), closeExperiment }: Props = $props();

	let selectedObject: models.CandidateGen | models.VoterGen | null = $state(null)

	// Copy the experiment to the clipboard as JSON
	const jsonToClipboard = async () => {
		const result = await ClipboardSetText(JSON.stringify(experiment, null, 2));
		if (result) {
			toast.success('JSON copied to clipboard');
		} else {
			toast.error('Failed to copy JSON to clipboard');
		}
	};

	// Track whether changes have been made to the experiment
	let changed = false;
	let cleanCopy = { ...experiment };

	$effect(() => {
		if (changed) return;
		changed = JSON.stringify(experiment) !== JSON.stringify(cleanCopy);
	});

	const handleClose = () => {
		if (changed) {
			if (confirm('Close without saving?')) {
				closeExperiment();
			}
		} else {
			closeExperiment();
		}
	};

</script>

<div class="flex flex-row justify-between">
	<div class="justify-start">
		<Button onclick={handleClose} variant="ghost">Close</Button>
		<Button onclick={jsonToClipboard} variant="ghost">to Clipboard</Button>
	</div>
	<TextPopover bind:value={experiment.name} />
	<div class="justify-end">
		<SettingsButton />
	</div>
</div>

<Resizable.PaneGroup direction="horizontal" class="border">
	<Resizable.Pane defaultSize={35}>
		<div class="mx-4">
			<ExpMenu bind:experiment bind:selectedObject/>
		</div>
	</Resizable.Pane>
	<Resizable.Handle />
	<Resizable.Pane defaultSize={65}>
		<div class="px-4">
			<ExpMap bind:experiment bind:selectedObject/>
		</div>
	</Resizable.Pane>
</Resizable.PaneGroup>
