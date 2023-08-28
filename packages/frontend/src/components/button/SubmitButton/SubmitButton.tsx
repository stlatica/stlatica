import React from "react";
import Button from "@mui/material/Button";
import type { ButtonProps } from "@mui/material/Button";

/**
 *
 */
export const SubmitButton: React.FC<ButtonProps> = ({ children, ...others }) => {
  return (
    <Button variant="contained" {...others}>
      {children}
    </Button>
  );
};
