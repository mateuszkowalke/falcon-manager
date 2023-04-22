import { LoginData } from "./dto/loginData.dto";

class AuthService {
  constructor(private keyestoneGraphQLUrl: string) {}

  async login({ username, password }: LoginData) {
    let keystoneResp: Response;
    try {
      keystoneResp = await fetch(this.keyestoneGraphQLUrl, {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: JSON.stringify({
          query: `mutation ($email: String!, $password: String!) {
                  authenticateUserWithPassword(
                    email: $email,
                    password: $password
                  ) {
                   ... on UserAuthenticationWithPasswordSuccess {
                    item {
                      id
                      email
                    }
                  }
                    ... on UserAuthenticationWithPasswordFailure {
                      message
                    }
                  }
                }`,
          variables: { email: username, password },
        }),
      });
    } catch (err) {
      throw new Error("Error fetching data: " + err);
    }

    try {
      const body = await keystoneResp.json();
      if (body?.data?.authenticateUserWithPassword?.item) {
        const cookies = keystoneResp.headers.get("set-cookie");
        const sessionCookie = cookies
          ?.split(";")
          .map((el) => el.trim().split("="))
          .filter((el) => el[0] === "keystonejs-session")
          .flat()[1];
        return sessionCookie
          ? { sessionCookie, email: username }
          : { sessionCookie: "", email: "" };
      }
    } catch (err) {
      throw new Error("Error parsing response: " + err);
    }

    return { sessionCookie: "", email: "" };
  }

  async getUser(sessionCookie: string, email: string) {
    let keystoneResp: Response;
    try {
      keystoneResp = await fetch(this.keyestoneGraphQLUrl, {
        method: "POST",
        headers: {
          "content-type": "application/json",
          cookie: sessionCookie,
        },
        body: JSON.stringify({
          query: `query ($where: UserWhereUniqueInput!) {
                  user(
                    where: $where
                  ) {
                    firstName,
                    lastName,
                    email
                  }
                }`,
          variables: { where: { email } },
        }),
      });
    } catch (err) {
      throw new Error("Error fetching data: " + err);
    }

    try {
      const body = await keystoneResp.json();
      return body.data?.user;
    } catch (err) {
      throw new Error("Error parsing response: " + err);
    }
  }
}

export const authService = new AuthService(process.env.KEYSTONE_GRAPHQL_URL!);
