import { Request, Response, NextFunction } from "express";
import { KeystoneUser } from "./dto/keystoneUser.dto";
import { authService } from "./service";

export async function authMiddleware(
  req: Request,
  res: Response,
  next: NextFunction
) {
  if (!req.user) {
    return res.status(401).redirect("/");
  }
  next();
}

export async function extractUser(
  req: Request,
  res: Response,
  next: NextFunction
) {
  const sessionCookie = 'keystonejs-session=' + req.cookies["keystonejs-session"];
  const email = req.cookies["email"];
  try {
    const user: KeystoneUser | null = await authService.getUser(sessionCookie, email);
    if (user) {
      req.user = { ...user, sessionCookie };
    }
    next();
  } catch (err) {
    next(err);
  }
}
