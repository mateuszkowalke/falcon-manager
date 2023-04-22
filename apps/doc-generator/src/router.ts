import { Router, Request, Response } from "express";
import { authMiddleware, extractUser } from "./auth/middleware";
import { authRouter } from "./auth/router";
import { generatorRouter } from "./generator/router";

const mainRouter = Router();

mainRouter.use(extractUser);

mainRouter.use("/auth", authRouter);
mainRouter.use("/generator", authMiddleware, generatorRouter);

// main page
mainRouter.get("/", (req: Request, res: Response) => {
  res.render("home", { title: "Document generator", user: req.user });
});

// 404 page
mainRouter.use((req, res, next) => {
  res.status(404).render("404");
});

export { mainRouter };
