<script lang="ts">
	import * as Accordion from '$lib/components/ui/accordion/index.js';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';
	import * as Card from '$lib/components/ui/card/index.js';
	import Input from '$lib/components/ui/input/input.svelte';
	import { models } from '$lib/wailsjs/go/models';

	interface CircleObject {
		x: number;
		y: number;
		radius: number;
	}

	type Props = {
		experiment: models.Experiment;
		selectedObject: CircleObject | null;
	};

	let { experiment = $bindable(), selectedObject = $bindable() }: Props = $props();

	// Which sections should be open by default
	const initialValues: string[] = [];
	if (experiment.candidateGens.length < 2) initialValues.push('candidates');
	if (experiment.voterGens.length < 1) initialValues.push('voters');
	if (experiment.notes) initialValues.push('notes');

	const cardClass = (obj: CircleObject) => {
		let cc = '';
		if (obj === selectedObject) {
			cc += ' border-2 border-slate-50';
		}
		return cc;
	};

	const selectObject = (obj: CircleObject) => {
		selectedObject = obj;
	};
</script>

<Accordion.Root type="multiple" value={initialValues}>
	<Accordion.Item value="candidates">
		<Accordion.Trigger>Candidates ({experiment.candidateGens.length})</Accordion.Trigger>

		<Accordion.Content>
			{#each experiment.candidateGens as cgen}
				<Card.Root onclick={() => selectObject(cgen)} class={cardClass(cgen)}>
					<Card.Content>
						<div>
							<input type="color" bind:value={cgen.color} />
							{cgen.name}
						</div>
					</Card.Content>
				</Card.Root>
			{/each}
		</Accordion.Content>
	</Accordion.Item>

	<Accordion.Item value="voters">
		<Accordion.Trigger>Voters ({experiment.voterGens.length})</Accordion.Trigger>
		<Accordion.Content>
			{#each experiment.voterGens as vgen}
				<Card.Root onclick={() => selectObject(vgen)} class={cardClass(vgen)}>
					<Card.Content>
						<div>
							<input type="color" bind:value={vgen.color} />
							{vgen.name}
						</div>
					</Card.Content>
				</Card.Root>
			{/each}
		</Accordion.Content>
	</Accordion.Item>

	<Accordion.Item value="notes">
		<Accordion.Trigger>Notes</Accordion.Trigger>
		<Accordion.Content>
			<div class="p-2">
				<Textarea class="h-60" bind:value={experiment.notes} />
			</div>
		</Accordion.Content>
	</Accordion.Item>
</Accordion.Root>
