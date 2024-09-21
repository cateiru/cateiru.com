import { backendApi } from "./api";
import {
  Public,
  PublicProduct,
  PublicProductSchema,
  PublicSchema,
} from "./types";

/**
 * get public data
 */
export async function getPublicProfile(): Promise<Public> {
  const res = await fetch(backendApi("/public/profile"), { method: "GET" });

  if (res.ok) {
    return PublicSchema.parse(await res.json());
  }
  const message = (await res.json()).message;
  throw new Error(`${res.status} : ${message}`);
}

/**
 * get public products
 * @param {number} id - product id
 */
export async function getPublicProduct(id: number): Promise<PublicProduct> {
  const res = await fetch(backendApi(`/public/product?product_id=${id}`), {
    method: "GET",
  });

  if (res.ok) {
    return PublicProductSchema.parse(await res.json());
  }
  const message = (await res.json()).message;
  throw new Error(`${res.status} : ${message}`);
}
