import { Request, Response } from "express";
import { loginDataSchema, LoginData } from "./dto/loginData.dto";
import { authService } from "./service";

export function getLoginForm(req: Request, res: Response) {
  res.render("auth/login-form", { title: "Log in", user: req.user });
}

export async function postLoginForm(
  req: Request<any, any, LoginData>,
  res: Response
) {
  try {
    const loginData = loginDataSchema.parse(req.body);
    const cookies = await authService.login(loginData);
    res.cookie("keystonejs-session", cookies.sessionCookie, {
      httpOnly: true,
      secure: true,
    });
    res.cookie("email", cookies.email, {
      httpOnly: true,
      secure: true,
    });
    res.redirect("/");
  } catch (err) {
    console.error({ err });
  }
}

export function logout(req: Request, res: Response) {
  res.cookie("keystonejs-session", "");
  res.cookie("email", "");
  res.redirect("/");
}
