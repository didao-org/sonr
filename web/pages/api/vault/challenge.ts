import type { NextRequest } from "next/server";

export const config = {
  runtime: "experimental-edge",
};

export default async function handler(req: NextRequest) {
  // Get API URL
  let domain = new URL(req.url).searchParams.get("domain");
  let apiUrl = "https://api.sonr.network";
  if (process && process.env.NODE_ENV === "development") {
    apiUrl = "http://localhost:1317";
  }

  const requestOptions = {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  };

  const resp = await fetch(
    apiUrl + "/sonr-io/highway/vault/challenge/" + domain,
    requestOptions
  );
  const data = await resp.json();
  console.log(data);
  return new Response(JSON.stringify(data), {
    status: 200,
    headers: {
      "content-type": "application/json",
    },
  });
}
