import { CancelButton } from "@/components/button/CancelButton";
import { SubmitButton } from "@/components/button/SubmitButton";
import { TextEditor } from "@/components/common/TextEditor";
import {
  mainContainer,
  halfWidth,
  marginLeft,
  marginTop,
  flexContainer,
  iconEditorContainer,
  iconEditorLabel,
  iconEditorContent,
} from "@/styles/routes/edit_profile.css";

export const IconEditor: React.FC = () => {
  return (
    <div class={iconEditorContainer}>
      <div class={iconEditorLabel}>Icon</div>
      <div class={iconEditorContent}>tmp</div>
      <div>zoom bar</div>
      <div>file path</div>
    </div>
  );
};

export default function Home() {
  return (
    <main>
      <div class={mainContainer}>
        <div class={halfWidth}>
          <IconEditor />
        </div>
        <div class={`${halfWidth} ${marginLeft}`}>
          <div>
            <TextEditor defaultValue="" label="Name" maxLength={16} />
          </div>
          <div class={marginTop}>
            <TextEditor defaultValue="" label="Self-Intorduction" maxLength={512} />
          </div>
          <div class={`${marginTop} ${flexContainer}`}>
            <SubmitButton>Save</SubmitButton>
            <CancelButton>Cancel</CancelButton>
          </div>
        </div>
      </div>
    </main>
  );
}
