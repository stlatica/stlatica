import React from "react";

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
