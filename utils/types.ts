import {z} from 'zod';

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

export const ProductSchema = z.object({});

export type User = typeof UserSchema._type;
export type AllUsers = typeof UsersListSchema._type;
export type Product = typeof ProductSchema._type;
