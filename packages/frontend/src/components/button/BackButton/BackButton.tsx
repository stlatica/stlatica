import React from "react";
import { MdArrowBack } from "react-icons/md";

type BackButtonProps = {
  // children: React.ReactNode;
};

/**
 *
 */
export const BackButton: React.FC<BackButtonProps> = () => {
  return (
    <div className="flex justify-center">
      <button className="border-2 border-black" type="button">
        <MdArrowBack size="3em" />
      </button>
    </div>
  );
};
