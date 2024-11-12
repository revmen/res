<script lang="ts">
	import { models } from '$lib/wailsjs/go/models';

	type Props = {
		experiment: models.Experiment;
		selectedObject: {type: string, index: number} | null;
	};

	let { experiment = $bindable(), selectedObject = $bindable() }: Props = $props();

	const size = 500;
	const maxCoord = 100;
	const minCoord = -100;

	
	const modelToMapCoord = (coord: number) => {
		return ((coord - minCoord) / (maxCoord - minCoord)) * size
	}
	
	const mapToModelCoord = (coord: number) => {
		return (coord * (maxCoord - minCoord)) / size + minCoord;
	}

	const modelToMapRadius = (r: number) => {
		return r * size / (maxCoord - minCoord)
	}
	
	const boundedModelCoord = (coord: number, r: number) => {
		if (coord > maxCoord - r) return maxCoord - r
		if (coord < minCoord + r) return minCoord + r
		return coord
	}

	if (experiment.candidateGens.length < 1) {

		experiment.candidateGens = [];
		
		const cgColors = [
			"#c01c28",
			"#1c71d8",
			"#2ec27e",
			"#f8e45c",
			"#9141ac"
		]
		
		for (let i = 0; i < 4; i++) {
			const radius = Math.random() * 25
			const x = boundedModelCoord(Math.random() * 200 - 100, radius)
			const y = boundedModelCoord(Math.random() * 200 - 100, radius)
			experiment.candidateGens.push({
				name: `Candidate ${i}`,
				x,
				y,
				radius,
				distribution: 'uniform',
				power: 100,
				color: cgColors[i],
			})
		}
	}

	const circleStroke = (type: string, index: number) => {
    if (type == selectedObject?.type && index == selectedObject.index) {
        return "red";
    }
    return "black";

}

let selectedCircle: models.CandidateGen | models.VoterGen | null = $state(null)
$effect(()=>{
	if (selectedObject?.type == "cgen"){
		selectedCircle = experiment.candidateGens[selectedObject.index]
	} else if (selectedObject?.type == "vgen") {
		selectedCircle = experiment.voterGens[selectedObject.index]
	} else {
		selectedCircle = null;
	}
});

let dragging = $state(false)

const handleMouseDown = (e: MouseEvent) => {

}

const handleMouseMove = (e: MouseEvent) => {
	if (selectedCircle && dragging) {
		const svg = e.currentTarget as SVGSVGElement;
		const rect = svg.getBoundingClientRect();
		
		// Adjust the client coordinates by the SVG's position
		const x = e.clientX - rect.left;
		const y = e.clientY - rect.top;

		// Map the coordinates and update the selectedCircle's position
		selectedCircle.x = boundedModelCoord(mapToModelCoord(x), selectedCircle.radius);
		selectedCircle.y = boundedModelCoord(mapToModelCoord(y), selectedCircle.radius);
	}
};


const handleMouseUp = (e: MouseEvent) => {
	dragging = false;
}

const handleMouseLeave = (e: MouseEvent) => {
	dragging = false;
}



const circleCenterMouseDown = (e: MouseEvent) => {
	if (!selectedCircle) return;
	dragging = true;
}

const handleCGenClick = (i: number) => {
	selectedObject = {type: "cgen", index: i}
}

const handleSvgClick = (e: MouseEvent) => {
	if (!(e.target instanceof SVGCircleElement)) {
		selectedObject = {type: "", index: 0}
	}
}
	
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_click_events_have_key_events -->
<svg 
	width={size} 
	height={size} 
	style="border: 2px solid black; background-color: #f4f4f4; cursor: {dragging ? 'none': 'arrow'};"
	onmousedown={handleMouseDown}
	onmousemove={handleMouseMove}
	onmouseup={handleMouseUp}
	onmouseleave={handleMouseLeave}
	onclick={handleSvgClick}
	>

	<!-- Candidate Generator Circles -->
	{#each experiment.candidateGens as cgen, i}
		<circle
			cx={modelToMapCoord(cgen.x)} 
			cy={modelToMapCoord(cgen.y)} 
			r={modelToMapRadius(cgen.radius)}
			stroke={circleStroke("cgen", i)}
			stroke-width="2"
			fill={cgen.color}
			fill-opacity="0.25"
			onclick={()=>handleCGenClick(i)}
		/>
	{/each}

	<!-- Vertical Axis -->
	<line
		x1={size/2}
		y1="0"
		x2={size/2}
		y2={size}
		stroke="grey"
		stroke-width="2"
	/>

	<line
		x1={size/2 - 10}
		y1={size/4}
		x2={size/2 + 10}
		y2={size/4}
		stroke="grey"
		stroke-width="2"
	/>

	<line
		x1={size/2 - 10}
		y1={3*size/4}
		x2={size/2 + 10}
		y2={3*size/4}
		stroke="grey"
		stroke-width="2"
	/>


	<!-- Horizontal Axis -->
	<line
		x1="0"
		y1={size/2}
		x2={size}
		y2={size/2}
		stroke="grey"
		stroke-width="2"
	/>

	<line
		x1={size/4}
		y1={size/2 - 10}
		x2={size/4}
		y2={size/2 + 10}
		stroke="grey"
		stroke-width="2"
	/>

	<line
		x1={3*size/4}
		y1={size/2 - 10}
		x2={3*size/4}
		y2={size/2 + 10}
		stroke="grey"
		stroke-width="2"
	/>



	<!-- Selection Handle -->
	{#if selectedCircle}
		<circle
			cx={modelToMapCoord(selectedCircle.x)}
			cy={modelToMapCoord(selectedCircle.y)}
			r=5
			fill="red"
			onmousedown={circleCenterMouseDown}
			style="cursor: {dragging ? 'none' : 'pointer'};"
		/>
	{/if}

</svg>


