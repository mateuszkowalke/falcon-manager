import prismaClient from '@prisma/client';

const { PrismaClient, Sex } = prismaClient;
const prisma = new PrismaClient();

export async function seedFalcons() {
	const sakerSpecies = await prisma.species.create({
		data: {
			name: 'Saker',
			latin: 'Falco cherrug'
		}
	});
	const peregrineSpecies = await prisma.species.create({
		data: {
			name: 'Peregrine',
			latin: 'Falco peregrinus'
		}
	});
	const pietrzwaldBreedingProject = await prisma.breedingProject.findFirst({
		where: {
			name: 'Pietrzwałd'
		}
	});
	const firstAviary = await prisma.aviary.create({
		data: {
			name: 'first',
			capacity: 1,
			lastCleaned: new Date(),
			breedingProjectId: pietrzwaldBreedingProject.id
		}
	});
	const secondAviary = await prisma.aviary.create({
		data: {
			name: 'second',
			capacity: 1,
			lastCleaned: new Date(),
			breedingProjectId: pietrzwaldBreedingProject.id
		}
	});
	await prisma.falcon.createMany({
		data: [
			{
				name: 'Prima',
				ring: 'todo',
				speciesId: peregrineSpecies.id,
				sex: Sex.FEMALE,
				birthDate: new Date('2019-04'),
				source: 'from Krafek',
				aviaryId: firstAviary.id,
				breedingProjectId: pietrzwaldBreedingProject.id
			},
			{
				name: 'Vera',
				ring: 'todo',
				speciesId: peregrineSpecies.id,
				sex: Sex.FEMALE,
				birthDate: new Date('2019-04'),
				source: 'from Krafek',
				aviaryId: secondAviary.id,
				breedingProjectId: pietrzwaldBreedingProject.id
			}
		]
	});
}
