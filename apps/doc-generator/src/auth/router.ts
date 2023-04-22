import { Router } from "express";
import { getLoginForm, postLoginForm, logout } from "./controller"

const authRouter = Router();

authRouter.get("/login", getLoginForm);
authRouter.post("/login", postLoginForm);
authRouter.get("/logout", logout);

export { authRouter };
