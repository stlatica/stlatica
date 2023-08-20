import React from "react";
import Button from "@mui/material/Button";
import type { ButtonProps } from "@mui/material/Button";

export type CancelButtonProps = {
  children: React.ReactNode;
};

/**
 *
 */
export const CancelButton: React.FC<ButtonProps> = ({ children, ...others }) => {
  return (
    <Button variant="outlined" {...others}>
      {children}
    </Button>
  );
};
