import { Button, type ButtonProps } from "@mantine/core";
import type { MouseEventHandler } from "react";

export type SubmitButtonProps = {
  onClick?: MouseEventHandler<HTMLButtonElement> | undefined;
} & ButtonProps;

/**
 *
 */
export const SubmitButton: React.FC<SubmitButtonProps> = ({ children, ...others }) => {
  return (
    <Button type="button" variant="contained" {...others}>
      {children}
    </Button>
  );
};
