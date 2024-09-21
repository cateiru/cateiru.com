import type { MultiLang } from "./lang";

export interface Product {
  name: MultiLang;
  createDate: Date;
  link: string;
  description: MultiLang;
}
