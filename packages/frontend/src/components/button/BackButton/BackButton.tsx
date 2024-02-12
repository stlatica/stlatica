import IconButton from "@mui/material/IconButton";
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
    <IconButton>
      <MdArrowBack size="3em" />
    </IconButton>
  );
};
