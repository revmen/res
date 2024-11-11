<script lang="ts">
	import { Button } from '$lib/components/ui/button/index.js';
	import * as Resizable from '$lib/components/ui/resizable/index.js';
	import * as Select from '$lib/components/ui/select/index.js';
	import * as Accordion from '$lib/components/ui/accordion/index.js';
	import { Input } from '$lib/components/ui/input/index.js';
	import Label from '$lib/components/ui/label/label.svelte';
	import { toast } from 'svelte-sonner';
	import SettingsButton from './SettingsButton.svelte';
	import TextPopover from './TextPopover.svelte';

	import { models } from '$lib/wailsjs/go/models';
	import { ClipboardSetText } from '$lib/wailsjs/runtime/runtime';
	import Textarea from '../ui/textarea/textarea.svelte';

	type Props = {
		experiment: models.Experiment;
		closeExperiment: () => void;
	};

	let { experiment = $bindable(), closeExperiment }: Props = $props();

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

	const jsonToClipboard = async () => {
		const result = await ClipboardSetText(JSON.stringify(experiment, null, 2));
		if (result) {
			toast.success('JSON copied to clipboard');
		} else {
			toast.error('Failed to copy JSON to clipboard');
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
			<div class="px-4">
			<Accordion.Root type="multiple">
				<Accordion.Item value="candidates">
					<Accordion.Trigger>Candidates</Accordion.Trigger>
					<Accordion.Content>Yes. It adheres to the WAI-ARIA design pattern.</Accordion.Content>
				</Accordion.Item>
				<Accordion.Item value="voters">
					<Accordion.Trigger>Voters</Accordion.Trigger>
					<Accordion.Content>Yes. It adheres to the WAI-ARIA design pattern.</Accordion.Content>
				</Accordion.Item>
				<Accordion.Item value="notes">
					<Accordion.Trigger>Notes</Accordion.Trigger>
					<Accordion.Content>
						<Textarea class="h-60"bind:value={experiment.notes} />
					</Accordion.Content>
				</Accordion.Item>
			</Accordion.Root>
		</div>
		</Resizable.Pane>
		<Resizable.Handle />
		<Resizable.Pane defaultSize={65}>
			<div class="flex h-96 items-center justify-center p-6">
				<span class="font-semibold">Content</span>
			</div>
		</Resizable.Pane>
	</Resizable.PaneGroup>
