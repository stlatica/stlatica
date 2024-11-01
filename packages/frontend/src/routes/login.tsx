import { SubmitButton } from "@/components/button/SubmitButton";
import { TextEditor } from "@/components/common/TextEditor";
import {
  mainContainer,
  innerContainer,
  title,
  textEditorContainer,
  buttonContainer,
} from "@/styles/routes/login.css";

export default function LoginScene() {
  return (
    <main class={mainContainer}>
      <div class={innerContainer}>
        <div class={title}>Stlatica</div>
        <div class={textEditorContainer}>
          <TextEditor defaultValue="" maxLength={32} title="mail address" />
        </div>
        <div class={textEditorContainer}>
          <TextEditor defaultValue="" maxLength={32} title="password" />
        </div>
        <div class={buttonContainer}>
          <SubmitButton>Login</SubmitButton>
        </div>
      </div>
    </main>
  );
}
