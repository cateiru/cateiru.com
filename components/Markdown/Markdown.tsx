/**********************************************************
 * Markdown effects.
 *
 * @author Yuto Watanabe <yuto.w51942@gmail.com>
 * @version 1.0.0
 *
 * Copyright (C) 2021 hello-slide
 **********************************************************/
import ReactMarkdown from "react-markdown";
import { components } from "./Contents";

const Markdown: React.FC<{ children: string }> = ({ children }) => {
  return <ReactMarkdown components={components}>{children}</ReactMarkdown>;
};

export default Markdown;
