import React from "react";

export const IconEditor: React.FC = () => {
  return (
    <div className="border border-white rounded p-3">
      <div className="text-gray-500">Icon</div>
      <div>
        <div className="w-[100px] h-[100px] text-black bg-white">tmp</div>
      </div>
      <div>zoom bar</div>
      <div>file path</div>
    </div>
  );
};
