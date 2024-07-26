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
    <main className={mainContainer}>
      <div className={innerContainer}>
        <div className={title}>Stlatica</div>
        <div className={textEditorContainer}>
          <TextEditor defaultValue="" maxLength={32} title="mail address" />
        </div>
        <div className={textEditorContainer}>
          <TextEditor defaultValue="" maxLength={32} title="password" />
        </div>
        <div className={buttonContainer}>
          <SubmitButton>Login</SubmitButton>
        </div>
      </div>
    </main>
  );
}
