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
    <div className={iconEditorContainer}>
      <div className={iconEditorLabel}>Icon</div>
      <div className={iconEditorContent}>tmp</div>
      <div>zoom bar</div>
      <div>file path</div>
    </div>
  );
};

export default function Home() {
  return (
    <main>
      <div className={mainContainer}>
        <div className={halfWidth}>
          <IconEditor />
        </div>
        <div className={`${halfWidth} ${marginLeft}`}>
          <div>
            <TextEditor defaultValue="" label="Name" maxLength={16} />
          </div>
          <div className={marginTop}>
            <TextEditor defaultValue="" label="Self-Intorduction" maxLength={512} />
          </div>
          <div className={`${marginTop} ${flexContainer}`}>
            <SubmitButton>Save</SubmitButton>
            <CancelButton>Cancel</CancelButton>
          </div>
        </div>
      </div>
    </main>
  );
}
