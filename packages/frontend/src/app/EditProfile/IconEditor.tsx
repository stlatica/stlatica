import React from "react";

export const IconEditor: React.FC = () => {
  return (
    <div className="rounded border border-white p-3">
      <div className="text-gray-500">Icon</div>
      <div>
        <div className="h-[100px] w-[100px] bg-white text-black">tmp</div>
      </div>
      <div>zoom bar</div>
      <div>file path</div>
    </div>
  );
};
