import { JTDDataType } from "ajv/dist/core";

const favoriteSchema = {
  type: "object",
  properties: {
    data: {
      type: "object",
      properties: {
        favorite_count: { type: "int32" },
        is_favorite: { type: "boolean" },
      },
    },
  },
} as const;

export type FavoriteResponse = JTDDataType<typeof favoriteSchema>;

export type FavoriteCreate = {
  sayka_id: string;
  is_favorite: boolean;
};
