import { JTDDataType } from "ajv/dist/core";

const loggedInUserSchema = {
  type: "object",
  properties: {
    data: {
      type: "object",
      properties: {
        id: { type: "string" },
        name: { type: "string" },
        img: { type: "string" },
      },
    },
  },
} as const;

export type AuthUserResponse = JTDDataType<typeof loggedInUserSchema>;
