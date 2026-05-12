import { type ActionFunctionArgs, json, redirect } from "@remix-run/node";
import { Form, useActionData, useNavigation } from "@remix-run/react";

import { SubmitButton } from "@/components/button/SubmitButton";
import { TextEditor } from "@/components/common/TextEditor";
import {
  buttonContainer,
  innerContainer,
  mainContainer,
  textEditorContainer,
  title,
} from "@/styles/routes/login.css";

const backendLoginURL = "http://localhost:8080/internal/v1/login";

type ActionData = {
  error?: string;
  mailAddress?: string;
};

export const action = async ({ request }: ActionFunctionArgs) => {
  const formData = await request.formData();
  const mailAddress = formData.get("mail_address")?.toString().trim() ?? "";
  const password = formData.get("password")?.toString() ?? "";

  if (mailAddress === "" || password === "") {
    return json<ActionData>(
      {
        error: "mail address と password は必須です。",
        mailAddress,
      },
      { status: 400 },
    );
  }

  const response = await fetch(backendLoginURL, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
    },
    body: JSON.stringify({
      mail_address: mailAddress,
      password,
    }),
  });

  if (!response.ok) {
    const errorMessage = await readErrorMessage(response);
    return json<ActionData>(
      {
        error: errorMessage,
        mailAddress,
      },
      { status: response.status },
    );
  }

  const responseBody = (await response.json()) as { user_id?: string };
  if (responseBody.user_id === undefined || responseBody.user_id === "") {
    return json<ActionData>(
      {
        error: "ログイン応答が不正です。",
        mailAddress,
      },
      { status: 500 },
    );
  }

  const headers = new Headers();
  appendSetCookieHeaders(headers, response.headers);
  return redirect(`/user/${responseBody.user_id}/timeline`, { headers });
};

export default function LoginScene() {
  const actionData = useActionData<typeof action>();
  const navigation = useNavigation();
  const isSubmitting = navigation.state === "submitting";

  return (
    <main className={mainContainer}>
      <Form method="post" style={{ width: "100%" }}>
        <div className={innerContainer}>
          <div className={title}>Stlatica</div>
          <div className={textEditorContainer}>
            <TextEditor
              autoComplete="email"
              defaultValue={actionData?.mailAddress ?? ""}
              maxLength={256}
              name="mail_address"
              required
              title="mail address"
            />
          </div>
          <div className={textEditorContainer}>
            <TextEditor
              autoComplete="current-password"
              defaultValue=""
              maxLength={72}
              name="password"
              required
              title="password"
              type="password"
            />
          </div>
          {actionData?.error ? (
            <div style={{ color: "#fca5a5", textAlign: "center" }}>{actionData.error}</div>
          ) : null}
          <div className={buttonContainer}>
            <SubmitButton loading={isSubmitting} type="submit">
              Login
            </SubmitButton>
          </div>
        </div>
      </Form>
    </main>
  );
}

const readErrorMessage = async (response: Response): Promise<string> => {
  try {
    const body = (await response.json()) as { message?: string };
    if (body.message) {
      return body.message;
    }
  } catch {
    // noop
  }

  if (response.status === 401) {
    return "mail address もしくは password が正しくありません。";
  }
  return "ログインに失敗しました。";
};

const appendSetCookieHeaders = (targetHeaders: Headers, sourceHeaders: Headers) => {
  const getSetCookie = (
    sourceHeaders as Headers & {
      getSetCookie?: () => string[];
    }
  ).getSetCookie;

  const setCookies =
    getSetCookie?.call(sourceHeaders) ??
    (sourceHeaders.get("set-cookie") ? [sourceHeaders.get("set-cookie") as string] : []);

  for (const setCookie of setCookies) {
    targetHeaders.append("Set-Cookie", setCookie);
  }
};
