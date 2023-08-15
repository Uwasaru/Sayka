import { JTDDataType } from "ajv/dist/core";

const loggedInUserSchema = {
  type: "object",
  properties: {
    user: {
      type: "object",
      properties: {
        Id: { type: "number" },
        Name: { type: "string" },
        Img: { type: "string" },
      },
    },
  },
} as const;

export type AuthUserResponse = {
  user: {
    Id: number;
    Name: string;
    Img: string;
  };
};
