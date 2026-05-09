import { afterEach, describe, expect, test, vi } from "vitest";

vi.mock("@/components/button/SubmitButton", () => ({
  SubmitButton: () => null,
}));

vi.mock("@/components/common/TextEditor", () => ({
  TextEditor: () => null,
}));

vi.mock("@/styles/routes/login.css", () => ({
  buttonContainer: "buttonContainer",
  innerContainer: "innerContainer",
  mainContainer: "mainContainer",
  textEditorContainer: "textEditorContainer",
  title: "title",
}));

import { action } from "./login";

afterEach(() => {
  vi.restoreAllMocks();
});

describe("login action", () => {
  test("backend の Set-Cookie を redirect response に引き継ぐ", async () => {
    const fetchMock = vi.fn().mockResolvedValue(
      createBackendResponse(
        {
          user_id: "user-123",
        },
        {
          setCookies: [
            "access_token=access-token-value; Path=/; HttpOnly",
            "refresh_token=refresh-token-value; Path=/; HttpOnly",
          ],
        },
      ),
    );
    vi.stubGlobal("fetch", fetchMock);

    const response = await action({
      request: createLoginRequest({
        mailAddress: "user@example.com",
        password: "password",
      }),
      context: {},
      params: {},
    });

    expect(fetchMock).toHaveBeenCalledOnce();
    expect(response.status).toBe(302);
    expect(response.headers.get("Location")).toBe("/user/user-123/timeline");
    expect(getSetCookieHeaders(response.headers)).toEqual([
      "access_token=access-token-value; Path=/; HttpOnly",
      "refresh_token=refresh-token-value; Path=/; HttpOnly",
    ]);
  });

  test("backend が 401 を返したら error json を返す", async () => {
    vi.stubGlobal(
      "fetch",
      vi.fn().mockResolvedValue(
        createBackendResponse(
          {
            message: "unauthorized",
          },
          {
            status: 401,
          },
        ),
      ),
    );

    const response = await action({
      request: createLoginRequest({
        mailAddress: "user@example.com",
        password: "wrong-password",
      }),
      context: {},
      params: {},
    });

    expect(response.status).toBe(401);
    await expect(response.json()).resolves.toEqual({
      error: "unauthorized",
      mailAddress: "user@example.com",
    });
  });
});

const createLoginRequest = ({
  mailAddress,
  password,
}: {
  mailAddress: string;
  password: string;
}) => {
  const formData = new FormData();
  formData.set("mail_address", mailAddress);
  formData.set("password", password);

  return new Request("http://localhost/login", {
    method: "POST",
    body: formData,
  });
};

const createBackendResponse = (
  body: unknown,
  options?: {
    status?: number;
    setCookies?: string[];
  },
) => {
  const headers = new Headers({
    "Content-Type": "application/json",
  });

  for (const setCookie of options?.setCookies ?? []) {
    headers.append("Set-Cookie", setCookie);
  }

  return new Response(JSON.stringify(body), {
    status: options?.status ?? 200,
    headers,
  });
};

const getSetCookieHeaders = (headers: Headers) => {
  const getSetCookie = (
    headers as Headers & {
      getSetCookie?: () => string[];
    }
  ).getSetCookie;

  return getSetCookie?.call(headers) ?? [];
};
