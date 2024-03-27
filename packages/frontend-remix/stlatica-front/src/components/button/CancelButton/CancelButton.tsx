import { Button, ButtonProps } from "@mantine/core";
import { MouseEventHandler } from "react";

export type CancelButtonProps = {
  onClick?: MouseEventHandler<HTMLButtonElement> | undefined;
} & ButtonProps;

/**
 *
 */
export const CancelButton: React.FC<CancelButtonProps> = ({
  children,
  ...others
}) => {
  return (
    <Button type="button" variant="outline" {...others}>
      {children}
    </Button>
  );
};
