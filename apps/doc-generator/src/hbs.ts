import { create } from "express-handlebars";

// TODO after moving to monorepo try to create single source of truth
// for this enum

enum Sexes {
  MALE = "male",
  FEMALE = "female",
  UNKNOWN = "unknown",
}

export const hbs = create({
  extname: ".hbs",
  helpers: {
    formatDate(date: Date) {
      // TODO just for Polish dates - might need to extend that
      return date.toLocaleDateString("pl-PL");
    },
    translateSex(sex: Sexes) {
      switch (sex) {
        case Sexes.MALE:
          return "samiec";
        case Sexes.FEMALE:
          return "samica";
        case Sexes.UNKNOWN:
          return "nieznana";
      }
    },
    inc(i: number) {
      return i + 1;
    },
    firstPairIdx(i: number) {
      return i * 2 + 1;
    },
    secondPairIdx(i: number) {
      return (i + 1) * 2;
    },
  },
});
