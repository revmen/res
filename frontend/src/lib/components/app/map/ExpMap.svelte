<script lang="ts">
	import { models } from '$lib/wailsjs/go/models';
	import { toast } from 'svelte-sonner';

	interface CircleObject {
		x: number,
		y: number,
		radius: number
	}

	type Props = {
		experiment: models.Experiment;
		selectedObject: CircleObject | null;
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

	const mapToModelRadius = (r: number) => {
    return r * (maxCoord - minCoord) / size;
}
	
	const boundedModelCoord = (coord: number, r: number) => {
		if (coord > maxCoord - r) return maxCoord - r
		if (coord < minCoord + r) return minCoord + r
		return coord
	}

	const boundedModelRadius = (x: number, y: number, r: number) => {
		r = Math.min(r, maxCoord - x, maxCoord - y)
		return Math.min(r, x - minCoord, y - minCoord)
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
			const radius = 25
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

	if (experiment.voterGens.length < 1) {
		experiment.voterGens = [];
		const cgColors = [
			"#c01c28",
			"#1c71d8",
			"#2ec27e",
			"#f8e45c",
			"#9141ac"
		]
		
		for (let i = 0; i < 3; i++) {
			const radius = 25
			const x = boundedModelCoord(Math.random() * 200 - 100, radius)
			const y = boundedModelCoord(Math.random() * 200 - 100, radius)
			experiment.voterGens.push({
				name: `Voter Bloc ${i}`,
				x,
				y,
				radius,
				distribution: 'uniform',
				strategic: 100,
				color: cgColors[i],
			})
		}
	}

	const circleStroke = (obj: CircleObject) => {
    if (obj === selectedObject) {
        return "red";
    }
    return "black";

}

let dragging = $state(false)
let resizing = $state(false)

const handleMouseDown = (e: MouseEvent) => {

}

const handleMouseMove = (e: MouseEvent) => {
	if (!selectedObject) return;
	
	const svg = e.currentTarget as SVGSVGElement;
	const rect = svg.getBoundingClientRect();
	
	// Adjust the client coordinates by the SVG's position
	const x = e.clientX - rect.left;
	const y = e.clientY - rect.top;
	
	// Dragging to Move
	if (dragging) {
		// Map the coordinates and update the selectedCircle's position
		selectedObject.x = boundedModelCoord(mapToModelCoord(x), selectedObject.radius);
		selectedObject.y = boundedModelCoord(mapToModelCoord(y), selectedObject.radius);
	}

	// Dragging right to resize
	if (resizing) {
		const mapDistance = Math.max(Math.abs(x - modelToMapCoord(selectedObject.x)), Math.abs(y - modelToMapCoord(selectedObject.y)))
		selectedObject.radius = boundedModelRadius(selectedObject.x, selectedObject.y, mapToModelRadius(mapDistance))
	}
};


const handleMouseUp = (e: MouseEvent) => {
	if (!dragging && !resizing) {
		if (!(e.target instanceof SVGCircleElement)) {
			selectedObject = null;
		}
	}
	dragging = false;
	resizing = false;
}

const handleMouseLeave = (e: MouseEvent) => {
	dragging = false;
	resizing = false;
}

const circleCenterMouseDown = (e: MouseEvent) => {
	if (!selectedObject) return;
	dragging = true;
}

const circleEdgeMouseDown = (e: MouseEvent) => {
	if (!selectedObject) return;
	resizing = true;
}

const handleGenClick = (obj: CircleObject) => {
	selectedObject = obj
}
	
</script>

<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_click_events_have_key_events -->
<svg 
	width={size} 
	height={size} 
	style="border: 2px solid black; background-color: #f4f4f4; cursor: {dragging || resizing ? 'none': 'arrow'};"
	onmousedown={handleMouseDown}
	onmousemove={handleMouseMove}
	onmouseup={handleMouseUp}
	onmouseleave={handleMouseLeave}
	>

	<!-- Candidate Generator Circles -->
	{#each experiment.candidateGens as circ}
		<circle
			cx={modelToMapCoord(circ.x)} 
			cy={modelToMapCoord(circ.y)} 
			r={modelToMapRadius(circ.radius)}
			stroke={circleStroke(circ)}
			stroke-width="2"
			fill={circ.color}
			fill-opacity="0.25"
			onclick={()=>handleGenClick(circ)}
		/>
	{/each}

	<!-- Voter Generator Circles -->
	{#each experiment.voterGens as circ (circ)}
		<circle
			cx={modelToMapCoord(circ.x)} 
			cy={modelToMapCoord(circ.y)} 
			r={modelToMapRadius(circ.radius)}
			stroke={circleStroke(circ)}
			stroke-width="2"
			stroke-dasharray="5,5"
			fill={circ.color}
			fill-opacity="0.25"
			onclick={()=>handleGenClick(circ)}
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



	{#if selectedObject}
	<!-- Center Handle for Moving-->
		<circle
			cx={modelToMapCoord(selectedObject.x)}
			cy={modelToMapCoord(selectedObject.y)}
			r=6
			fill="red"
			onmousedown={circleCenterMouseDown}
			style="cursor: {dragging ? 'none' : 'pointer'};"
		/>

		<!-- Edge Handles for Resizing -->
		<circle
			cx={modelToMapCoord(selectedObject.x + selectedObject.radius)}
			cy={modelToMapCoord(selectedObject.y)}
			r=4
			fill="red"
			onmousedown={circleEdgeMouseDown}
			style="cursor: {resizing ? 'none' : 'pointer'};"
		/>

		<circle
			cx={modelToMapCoord(selectedObject.x - selectedObject.radius)}
			cy={modelToMapCoord(selectedObject.y)}
			r=4
			fill="red"
			onmousedown={circleEdgeMouseDown}
			style="cursor: {resizing ? 'none' : 'pointer'};"
		/>

		<circle
			cx={modelToMapCoord(selectedObject.x)}
			cy={modelToMapCoord(selectedObject.y + selectedObject.radius)}
			r=4
			fill="red"
			onmousedown={circleEdgeMouseDown}
			style="cursor: {resizing ? 'none' : 'pointer'};"
		/>

		<circle
			cx={modelToMapCoord(selectedObject.x)}
			cy={modelToMapCoord(selectedObject.y - selectedObject.radius)}
			r=4
			fill="red"
			onmousedown={circleEdgeMouseDown}
			style="cursor: {resizing ? 'none' : 'pointer'};"
		/>
	{/if}

</svg>


