import { list } from "@keystone-6/core";
import { text, relationship, timestamp, file } from "@keystone-6/core/fields";

import {
  defaultAccess,
  attachSessionUser,
  sharedResourceAccess,
} from "../auth/auth";

export const DocumentType = list({
  access: sharedResourceAccess,

  fields: {
    name: text({ validation: { isRequired: true } }),
  },

  ui: {
    labelField: "name",
  },
});

export const Document = list({
  access: defaultAccess,

  fields: {
    falcon: relationship({ ref: "Falcon.documents" }),
    documentType: relationship({ ref: "DocumentType" }),
    documentNumber: text(),
    scanFile: file({ storage: "documentsStorage" }),
    rawFile: file({ storage: "documentsStorage" }),
    owner: relationship({
      ref: "User.documents",
      ui: { hideCreate: true, createView: { fieldMode: "hidden" } },
      many: false,
    }),
    createdAt: timestamp({
      defaultValue: { kind: "now" },
    }),
    updatedAt: timestamp({
      defaultValue: { kind: "now" },
    }),
  },

  ui: {
    labelField: "documentNumber",
  },

  hooks: {
    resolveInput: attachSessionUser,
  },
});
