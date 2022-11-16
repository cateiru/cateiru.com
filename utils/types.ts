import {number, z} from 'zod';

export const UserSchema = z.object({
  id: z.number(),
  given_name: z.string(),
  family_name: z.string(),
  given_name_ja: z.string(),
  family_name_ja: z.string(),
  user_id: z.string(),
  mail: z.string(),
  birth_date: z.string(),
  location: z.string(),
  location_ja: z.string(),
  sso_token: z.string(),
  avatar_url: z.optional(z.string().url()),
  selected: z.optional(z.boolean()),
  created: z.string(),
  modified: z.string(),
});

export const UsersListSchema = z.array(UserSchema);

export const BioSchema = z.object({
  id: z.number(),
  user_id: z.number(),
  is_public: z.optional(z.boolean()),
  location_id: z.number(),
  position: z.string(),
  position_ja: z.string(),
  join: z.string(),
  leave: z.string(),
  created: z.string(),
  modified: z.string(),
});

export const LocationSchema = z.object({
  id: z.number(),
  type: z.string(),
  name: z.string(),
  name_ja: z.string(),
  address: z.string(),
  address_ja: z.string(),
  created: z.string(),
  modified: z.string(),
});

export const BioLocSchema = z.object({
  biography: BioSchema,
  location: z.optional(LocationSchema),
});
export const BioLocArraySchema = z.array(BioLocSchema);

export const ProductSchema = z.object({});
export const LocationArraySchema = z.array(LocationSchema);

export type User = typeof UserSchema._type;
export type Bio = typeof BioSchema._type;
export type AllUsers = typeof UsersListSchema._type;
export type Product = typeof ProductSchema._type;
export type BioLoc = typeof BioLocSchema._type;
export type BioLocArray = typeof BioLocArraySchema._type;
export type Location = typeof LocationSchema._type;
export type LocationArray = typeof LocationArraySchema._type;
