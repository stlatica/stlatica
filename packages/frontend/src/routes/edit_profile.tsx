import { CancelButton } from "@/components/button/CancelButton";
import { SubmitButton } from "@/components/button/SubmitButton";
import { TextEditor } from "@/components/common/TextEditor";

export const IconEditor: React.FC = () => {
  return (
    <div className="relative rounded-lg border p-3 dark:border-gray-600">
      <div className="absolute -translate-x-3.5 -translate-y-6 scale-75 bg-gray-800 px-2 text-gray-400">
        Icon
      </div>
      <div className="size-[100px] bg-white text-black">tmp</div>
      <div>zoom bar</div>
      <div>file path</div>
    </div>
  );
};

export default function Home() {
  return (
    <main>
      <div className="flex bg-gray-800 p-12 text-white">
        <div className="w-1/2">
          <IconEditor />
        </div>
        <div className="ml-12 w-1/2">
          <div>
            <TextEditor defaultValue="" label="Name" maxLength={16} />
          </div>
          <div className="mt-12">
            <TextEditor defaultValue="" label="Self-Intorduction" maxLength={512} />
          </div>
          <div className="mt-12 flex justify-around">
            {/* Assuming SubmitButton and CancelButton are styled components, you would need to pass the class names instead of the style prop. */}
            <SubmitButton className="w-[85px]">Save</SubmitButton>
            {/* Adjust the width as needed */}
            <CancelButton className="w-[85px]">Cancel</CancelButton>
            {/* Adjust the width as needed */}
          </div>
        </div>
      </div>
    </main>
  );
}
