import type { Species, Aviary, BreedingProject } from '@prisma/client';
export const loadFalconFormData: () => Promise<
	| { status: number; error: Error }
	| { species: Species[]; aviaries: Aviary[]; breedingProjects: BreedingProject[] }
> = async () => {
	// todo
	// refactor using Promise.all
	const speciesUrl = `/api/species`;
	const speciesRes = await fetch(speciesUrl);
	if (!speciesRes.ok) {
		return {
			status: speciesRes.status,
			error: new Error(`Could not load ${speciesUrl}`)
		};
	}

	const breedingProjectsUrl = `/api/breeding-projects`;
	const breedingProjectsRes = await fetch(breedingProjectsUrl);

	if (!breedingProjectsRes.ok) {
		return {
			status: breedingProjectsRes.status,
			error: new Error(`Could not load ${breedingProjectsUrl}`)
		};
	}

	const aviariesUrl = `/api/aviaries`;
	const aviariesRes = await fetch(aviariesUrl);

	if (!aviariesRes.ok) {
		return {
			status: aviariesRes.status,
			error: new Error(`Could not load ${aviariesUrl}`)
		};
	}

	const species = await speciesRes.json();
	const breedingProjects = await breedingProjectsRes.json();
	const aviaries = await aviariesRes.json();

	return { species, breedingProjects, aviaries };
};
