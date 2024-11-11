<script lang="ts">
	import SettingsButton from '$lib/components/app/SettingsButton.svelte';
	import App from '$lib/components/app/App.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { Textarea } from '$lib/components/ui/textarea/index.js';
	import { toast } from 'svelte-sonner';

	import { models } from '$lib/wailsjs/go/models';
	import { NewExperiment } from '$lib/wailsjs/go/main/App';

	let experiment: models.Experiment | null = $state(null);

	let showJSON = $state(false);

	const newExperiment = async () => {
		experiment = await NewExperiment();
	};

	const closeExperiment = () => {
		experiment = null;
	};

	let textAreaContents = $state('');

	const jsonOK = () => {
		let exp: models.Experiment;
		try {
			exp = JSON.parse(textAreaContents) as models.Experiment;
			toast.success(textAreaContents);
		} catch (e: any) {
			toast.error('Failed to parse JSON');
			toast.error(e.message);
			return;
		}
		experiment = exp;
		showJSON = false;
	};
</script>

{#if experiment}
	<App bind:experiment {closeExperiment}/>
{:else}
	<div>
		<div class="flex items-center justify-between">
			<div class="ml-auto">
				<SettingsButton />
			</div>
		</div>
	</div>
	<div class="flex w-full flex-col items-center">
		<h1>Create or load an experiment to save democracy.</h1>
		<Button onclick={newExperiment} class="mt-8">New</Button>
		<Button class="mt-8">Load from File</Button>
		<Button class="mt-8" onclick={() => (showJSON = !showJSON)}>Load from JSON</Button>
		{#if showJSON}
			<div class="mt-8 w-1/2">
				<Textarea bind:value={textAreaContents} class="h-60" placeholder="Paste experiment JSON" />
				<div class="flex justify-end">
					<Button onclick={jsonOK} class="ml-auto">OK</Button>
				</div>
			</div>
		{/if}
	</div>
{/if}
