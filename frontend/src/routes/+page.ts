import type { PageLoad } from './$types';
import { GetCurrentExperiment } from '$lib/wailsjs/go/main/App';
import { models } from '$lib/wailsjs/go/models';
import { get } from 'svelte/store';

export const load = (async ({ depends }) => {
    depends('app:currentExperiment');

    const getCurrentExperiment = async () => {
        let exp: models.Experiment;
        exp = await GetCurrentExperiment();
        return exp;
    }

    return {currentExperiment: await getCurrentExperiment()};
}) satisfies PageLoad;