import { z } from "zod";

const address = z.object({
  street: z.string(),
  no: z.string(),
  city: z.string(),
  zipCode: z.string(),
});

const species = z.object({
  id: z.string(),
  name: z.string(),
  latin: z.string(),
});

const falconSpecies = z.array(
  z.object({
    species,
    count: z.number(),
  })
);

const falconSexes = z.object({
  male: z.number(),
  female: z.number(),
  unknown: z.number(),
  total: z.number(),
});

const falconBirths = z.object({
  latest: z.date(),
  earliest: z.date(),
});

const parent = z.object({
  accquiredDate: z.date(),
  source: z.string(),
  citesNo: z.string(),
});

const parentPair = z.object({
  id: z.string(),
  female: parent,
  male: parent,
});

const falconParentPairs = z.array(
  z.object({
    id: z.string(),
    female: parent,
    male: parent,
    firstIdx: z.number(),
    lastIdx: z.number(),
  })
);

export const generateDocData = z.object({
  falcons: z.array(
    z.object({
      ring: z.string(),
      species,
      sex: z.enum(["male", "female", "unknown"]),
      birthDate: z.date(),
      parentPair,
    })
  ),
  falconSpecies,
  falconSexes,
  falconBirths,
  falconParentPairs,
  office: z.object({
    name: z.string(),
    officeType: z.object({
      name: z.string(),
    }),
    address,
    breedingProject: z.object({
      vetRegNo: z.string(),
      address,
      owner: z.object({
        firstName: z.string(),
        lastName: z.string(),
        personalAddress: address,
      }),
    }),
  }),
});

export type FalconBirths = z.infer<typeof falconBirths>;
export type FalconSexes = z.infer<typeof falconSexes>;
export type FalconSpecies = z.infer<typeof falconSpecies>;
export type FalconParentPairs = z.infer<typeof falconParentPairs>;
export type GenerateDocData = z.infer<typeof generateDocData>;
