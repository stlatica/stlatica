import React from "react";

export const IconEditor: React.FC = () => {
  return (
    <div
      style={{
        border: "1px solid #FFFFFF",
        borderRadius: "5px",
        padding: "12px",
      }}
    >
      <div style={{ color: "#9F9F9F" }}>Icon</div>
      <div>
        <div
          style={{ width: "100px", height: "100px", color: "#000000", backgroundColor: "#FFFFFF" }}
        >
          tmp
        </div>
      </div>
      <div>zoom bar</div>
      <div>file path</div>
    </div>
  );
};
