import React from "react";
import Button from "@mui/material/Button";

export type CancelButtonProps = {
  children: React.ReactNode;
};

/**
 *
 */
export const CancelButton: React.FC<CancelButtonProps> = ({ children }) => {
  return <Button variant="outlined">{children}</Button>;
};
