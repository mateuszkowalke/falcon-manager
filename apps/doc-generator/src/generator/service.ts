import puppeteer from "puppeteer";
import { hbs } from "../hbs";
import {
  FalconParentPairs,
  FalconSexes,
  FalconSpecies,
  GenerateDocData,
} from "./dto/generateDocData.dto";

class GeneratorService {
  constructor(private keyestoneGraphQLUrl: string) {}

  async getFormData(sessionCookie: string) {
    const keystoneResp = await fetch(this.keyestoneGraphQLUrl, {
      method: "POST",
      headers: {
        "content-type": "application/json",
        cookie: sessionCookie,
      },
      body: JSON.stringify({
        query: `query {
                    breedingProjects {
                        name,
                        offices {
                            id,
                            name
                        }
                    }
                    pairs {
                        name,
                        children {
                            id,
                            name,
                            ring
                        }
                    }
                }`,
      }),
    });

    try {
      const body = await keystoneResp.json();
      return body.data;
    } catch (err) {
      throw new Error("Internal server error: " + err);
    }
  }

  async getDocData(
    sessionCookie: string,
    officeId: string,
    falconIds: string[]
  ): Promise<GenerateDocData> {
    const keystoneResp = await fetch(this.keyestoneGraphQLUrl, {
      method: "POST",
      headers: {
        "content-type": "application/json",
        cookie: sessionCookie,
      },
      body: JSON.stringify({
        query: `query ($falconIds: FalconWhereInput!, $officeId: OfficeWhereUniqueInput!){
                    falcons(where: $falconIds) {
                        ring,
                        species {
                            id,
                            name,
                            latin
                        },
                        sex,
                        birthDate,
                        parentPair {
                            id,
                            male {
                                accquiredDate,
                                source,
                                citesNo
                            }
                            female {
                                accquiredDate,
                                source,
                                citesNo
                            }
                        }
                    },
                    office(where: $officeId) {
                        name,
                        officeType {
                            name
                        },
                        address {
                            street,
                            no,
                            city,
                            zipCode
                        },
                        breedingProject {
                            address {
                                street,
                                no,
                                city,
                                zipCode
                            },
                            vetRegNo,
                            owner {
                                firstName,
                                lastName,
                                personalAddress {
                                    street,
                                    no,
                                    city,
                                    zipCode
                                }
                            }
                        }
                    }
                }`,
        variables: {
          falconIds: { id: { in: falconIds } },
          officeId: { id: officeId },
        },
      }),
    });

    try {
      const body = await keystoneResp.json();
      return this.processRawGenerateDocData(body.data as GenerateDocData);
    } catch (err) {
      console.error("Error processing keystoneResp.");
      throw err;
    }
  }

  async generateBredInCaptivity(docData: GenerateDocData) {
    const html = await hbs.render(
      "views/generator/doc-templates/bred-in-captivity.hbs",
      {
        docData,
      }
    );

    const browser = await puppeteer.launch({
      headless: true,
    });
    const page = await browser.newPage();
    await page.setContent(html);
    const pdf = await page.pdf({ format: "A4" });
    await browser.close();
    return pdf;
  }

  private processRawGenerateDocData(data: GenerateDocData) {
    data.falcons = data.falcons
      .map((falcon) => ({
        ...falcon,
        birthDate: new Date(falcon.birthDate),
        parentPair: {
          ...falcon.parentPair,
          male: {
            ...falcon.parentPair.male,
            accquiredDate: new Date(falcon.parentPair.male.accquiredDate),
          },
          female: {
            ...falcon.parentPair.female,
            accquiredDate: new Date(falcon.parentPair.female.accquiredDate),
          },
        },
        sex: <typeof falcon.sex>falcon.sex.toLowerCase(),
      }))
      .sort((a, b) =>
        a.parentPair.id > b.parentPair.id ? (a.ring > b.ring ? 1 : 1) : -1
      );

    data.falconSpecies = data.falcons.reduce((acc, curr) => {
      let species = acc.find((val) => val.species.id === curr.species.id);
      if (!species) {
        species = {
          species: curr.species,
          count: 0,
        };
        acc.push(species);
      }
      species.count++;
      return acc;
    }, [] as FalconSpecies);

    data.falconSexes = data.falcons.reduce(
      (acc, curr): FalconSexes => {
        acc[curr.sex]++;
        acc.total++;
        return acc;
      },
      { male: 0, female: 0, unknown: 0, total: 0 } as FalconSexes
    );

    data.falconBirths = {
      earliest: data.falcons.reduce((prev, curr) =>
        prev.birthDate.getTime() < curr.birthDate.getTime() ? prev : curr
      ).birthDate,
      latest: data.falcons.reduce((prev, curr) =>
        prev.birthDate.getTime() > curr.birthDate.getTime() ? prev : curr
      ).birthDate,
    };

    data.falconParentPairs = data.falcons.reduce((acc, curr, i) => {
      if (!acc.map((pair) => pair.id).includes(curr.parentPair.id)) {
        acc.push({ ...curr.parentPair, firstIdx: i, lastIdx: i });
      } else {
        acc[acc.length - 1]!.lastIdx++;
      }
      return acc;
    }, [] as FalconParentPairs);

    return data;
  }
}

export const generatorService = new GeneratorService(
  process.env.KEYSTONE_GRAPHQL_URL!
);
