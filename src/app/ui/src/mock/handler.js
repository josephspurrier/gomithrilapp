import { rest } from "msw";

export const handlers = [
  // Healthcheck URL.
  rest.get("/api/v1", (req, res, ctx) => {
    return res(
      ctx.status(200),
      ctx.json({
        status: "OK",
        message: "ready",
      })
    );
  }),
];
