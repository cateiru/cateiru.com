import type { NextApiRequest, NextApiResponse } from "next";
import type { Form } from "../../utils/form";

const handler = async (req: NextApiRequest, res: NextApiResponse) => {
  console.log("OK");
  let status = 400;
  if (req.method === "POST") {
    const form = req.body as Form;
    const ip = req.headers["x-forwarded-for"] || req.socket.remoteAddress || "";
    status = await sendDiscord(form, ip);
  }

  res.status(status);
  res.end();
};

const sendDiscord = async (
  form: Form,
  ip: string[] | string,
): Promise<number> => {
  const token = process.env.DISCORD_TOKEN || "";
  if (typeof ip !== "string") {
    ip = ip.join(" ");
  }

  let status = 200;

  const text = `【新着問い合わせ】
* お名前: ${form.name}
* メールアドレス: ${form.mail}
* URL: ${form.url || "Null"}
* 送信日時: ${form.date}
* IPアドレス: ${ip}
* 言語: ${form.lang}
------
${form.subject}
------
${form.body}
`;

  try {
    const r = await fetch(token, {
      method: "POST",
      body: JSON.stringify({ content: text }),
      headers: { "Content-Type": "application/json" },
    });
    if (!r.ok) {
      throw new Error(r.statusText);
    }
  } catch (error) {
    console.log(error);
    status = 500;
  }

  return status;
};

export default handler;
