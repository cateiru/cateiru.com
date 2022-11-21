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
  type: z.enum(['univ', 'corp']),
  name: z.string(),
  name_ja: z.string(),
  address: z.string(),
  address_ja: z.string(),
  created: z.string(),
  modified: z.string(),
});
export const LocationArraySchema = z.array(LocationSchema);

export const BioLocSchema = z.object({
  biography: BioSchema,
  location: z.optional(LocationSchema),
});
export const BioLocArraySchema = z.array(BioLocSchema);

export const ProductSchema = z.object({
  id: z.number(),
  user_id: z.number(),
  name: z.string(),
  name_ja: z.string(),
  detail: z.string(),
  detail_ja: z.string(),
  site_url: z.string(),
  github_url: z.optional(z.string()),
  dev_time: z.string(),
  thumbnail: z.optional(z.string()),
  created: z.string(),
  modified: z.string(),
});
export const ProductArraySchema = z.array(ProductSchema);

export const LinkSchema = z.object({
  id: z.number(),
  user_id: z.number(),
  name: z.string(),
  name_ja: z.string(),
  site_url: z.string().url(),
  category_id: z.number(),
  created: z.string(),
  modified: z.string(),
});
export const CategorySchema = z.object({
  id: z.number(),
  name: z.string(),
  name_ja: z.string(),
  emoji: z.string(),
  created: z.string(),
  modified: z.string(),
});
export const LinkCategorySchema = z.object({
  category: CategorySchema,
  link: LinkSchema,
});

export const NoticeSchema = z.object({
  id: z.number(),
  user_id: z.number(),
  discord_webhook: z.optional(z.string().url()),
  slack_webhook: z.optional(z.string().url()),
  mail: z.optional(z.string().email()),
  created: z.string(),
  modified: z.string(),
});

export const ContactSchema = z.object({
  id: z.number(),
  to_user_id: z.number(),
  name: z.string(),
  title: z.string(),
  detail: z.string(),
  mail: z.string().email(),
  ip: z.string(),
  lang: z.string(),
  url: z.optional(z.string().url()),
  category: z.optional(z.string()),
  custom_title: z.optional(z.string()),
  custom_value: z.optional(z.string()),
  device_name: z.optional(z.string()),
  os: z.optional(z.string()),
  browser_name: z.optional(z.string()),
  is_mobile: z.optional(z.boolean()),
  created: z.string(),
  modified: z.string(),
});

export const PublicBioSchema = z.array(
  z.object({
    position: z.string(),
    position_ja: z.string(),
    join: z.string(),
    leave: z.string(),
    type: z.enum(['univ', 'corp']),
    name: z.string(),
    name_ja: z.string(),
    address: z.string(),
    address_ja: z.string(),
  })
);

export const PublicLinkSchema = z.object({
  category_id: z.number(),
  category_name: z.string(),
  category_name_ja: z.string(),
  emoji: z.string(),

  links: z.array(
    z.object({
      name: z.string(),
      name_ja: z.string(),
      site_url: z.string(),
      favicon_url: z.string().optional(),
    })
  ),
});

export const PublicSchema = z.object({
  given_name: z.string(),
  family_name: z.string(),
  given_name_ja: z.string(),
  family_name_ja: z.string(),
  user_id: z.string(),
  birth_date: z.string(),
  location: z.string(),
  location_ja: z.string(),
  avatar_url: z.string().url(),
  created: z.string(),
  modified: z.string(),

  biographies: PublicBioSchema,
  products: z.array(
    z.object({
      id: z.number(),
      name: z.string(),
      name_ja: z.string(),
      detail: z.string(),
      detail_ja: z.string(),
      dev_time: z.string(),
      thumbnail: z.optional(z.string().url()),
      github_url: z.optional(z.string().url()),
    })
  ),
  links: z.array(PublicLinkSchema),
});

export const PublicProductSchema = z.object({
  id: z.number(),
  name: z.string(),
  name_ja: z.string(),
  detail: z.string(),
  detail_ja: z.string(),
  site_url: z.string(),
  github_url: z.optional(z.string()),
  dev_time: z.string(),
  thumbnail: z.optional(z.string()),
});

export type User = typeof UserSchema._type;
export type Bio = typeof BioSchema._type;
export type AllUsers = typeof UsersListSchema._type;
export type Product = typeof ProductSchema._type;
export type ProductArray = typeof ProductArraySchema._type;
export type BioLoc = typeof BioLocSchema._type;
export type BioLocArray = typeof BioLocArraySchema._type;
export type Location = typeof LocationSchema._type;
export type LocationArray = typeof LocationArraySchema._type;
export type Link = typeof LinkSchema._type;
export type Category = typeof CategorySchema._type;
export type LinkCategory = typeof LinkCategorySchema._type;
export type Notice = typeof NoticeSchema._type;
export type Contact = typeof ContactSchema._type;
export type Public = typeof PublicSchema._type;
export type PublicProduct = typeof PublicProductSchema._type;
export type PublicBio = typeof PublicBioSchema._type;
export type PublicLink = typeof PublicLinkSchema._type;
