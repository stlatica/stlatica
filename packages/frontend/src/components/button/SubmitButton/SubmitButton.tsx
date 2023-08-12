import React from "react";
import Button from "@mui/material/Button";

export type SubmitButtonProps = {
  children: React.ReactNode;
};

/**
 *
 */
export const SubmitButton: React.FC<SubmitButtonProps> = ({ children }) => {
  return <Button variant="contained">{children}</Button>;
};
