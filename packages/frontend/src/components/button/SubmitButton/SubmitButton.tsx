import Button from "@mui/material/Button";
import React from "react";

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
