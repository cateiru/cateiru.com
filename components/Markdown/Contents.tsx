/**********************************************************
 * markdown components.
 *
 * @author Yuto Watanabe <yuto.w51942@gmail.com>
 * @version 1.0.0
 *
 * Copyright (C) 2021 hello-slide
 **********************************************************/
import type { Components } from "react-markdown";
import * as headers from "./Contents/Heading";
import * as lists from "./Contents/Lists";
import * as texts from "./Contents/Texts";

export const components: Components = {
  h1: headers.H1,
  h2: headers.H2,
  h3: headers.H3,
  h4: headers.H4,
  h5: headers.H5,
  h6: headers.H6,
  code: texts.Code,
  ul: lists.Ul,
  li: lists.Li,
  ol: lists.Ol,
};
