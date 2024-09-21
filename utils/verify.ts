import Cookies from "js-cookie";
import { NextRouter } from "next/router";
import { api } from "./api";
import { User, UserSchema } from "./types";

export class Verify {
  private isRedirect: boolean;
  private setUser: (u: User | null) => void;
  private router: NextRouter;
  private redirectPath: string | undefined;
  private failed: boolean;

  constructor(
    isRedirect: boolean,
    setUser: (u: User | null) => void,
    router: NextRouter,
    redirectPath?: string,
  ) {
    this.isRedirect = isRedirect;
    this.setUser = setUser;
    this.router = router;
    this.redirectPath = redirectPath;
    this.failed = false;
  }

  private async me(): Promise<User | null> {
    const res = await fetch(api("/user/me"), {
      credentials: "include",
      mode: "cors",
    });

    if (res.ok) {
      return UserSchema.parse(await res.json());
    }
    this.failed = true;

    return null;
  }

  private checkLoginFromCookie(): boolean {
    const isSession = Cookies.get("cateirucom-issession");

    return typeof isSession === "string";
  }

  private after(handler: () => void) {
    if (this.failed && this.isRedirect) {
      this.router.push(this.redirectPath ?? "/");
    }
    handler();
  }

  public async exec(handler: () => void) {
    if (!this.checkLoginFromCookie()) {
      this.failed = true;
      this.setUser(null);
      this.after(handler);
      return;
    }

    const res = await this.me();
    this.setUser(res);
    this.after(handler);
  }
}
