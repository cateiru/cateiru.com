import { createGzip, Gzip } from "zlib";
import {
  SitemapItem,
  SitemapStream,
  EnumChangefreq,
  streamToPromise,
} from "sitemap";
import { getPublicProfile } from "../utils/public";

const defaultLinks: SitemapItem[] = [
  {
    url: "/",
    changefreq: EnumChangefreq.NEVER,
    priority: 1.0,
    img: [],
    video: [],
    links: [],
  },
  {
    url: "/contact",
    changefreq: EnumChangefreq.NEVER,
    priority: 1.0,
    img: [],
    video: [],
    links: [],
  },
  {
    url: "/brand_resources",
    changefreq: EnumChangefreq.NEVER,
    priority: 1.0,
    img: [],
    video: [],
    links: [],
  },
  {
    url: "/brand_resources/icon_generator",
    changefreq: EnumChangefreq.NEVER,
    priority: 1.0,
    img: [],
    video: [],
    links: [],
  },
];

/**
 * サイトマップをAPIからのデータを使用して作成する
 *
 * @returns {Promise<Buffer>} - サイトマップのバイナリ
 */
export async function generateSitemap(): Promise<Gzip> {
  const d = await getPublicProfile();

  const smStream = new SitemapStream({ hostname: "https://cateiru.com/" });
  const pipeline = smStream.pipe(createGzip());

  for (const link of defaultLinks) {
    smStream.write(link);
  }

  for (const product of d.products) {
    smStream.write({
      url: `/product/${product.id}`,
      changefreq: EnumChangefreq.MONTHLY,
      priority: 0.8,
      img: [],
      video: [],
      links: [],
    });
  }

  smStream.end();

  return pipeline;
}
