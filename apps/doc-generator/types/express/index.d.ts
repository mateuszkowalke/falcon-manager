import { KeystoneUser } from "../../src/auth/dto/keystoneUser.dto";
// we need export to mark file as an external module
export {};

declare global {
  namespace Express {
    export interface Request {
      user?: KeystoneUser & { sessionCookie: string};
    }
  }
}
