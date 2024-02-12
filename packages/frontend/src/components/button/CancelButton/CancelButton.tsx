import Button, { ButtonProps } from "@mui/material/Button";

export type CancelButtonProps = ButtonProps;

/**
 *
 */
export const CancelButton: React.FC<CancelButtonProps> = (props) => {
  const { children, ...others } = props;
  return (
    <Button type="button" variant="outlined" {...others}>
      {children}
    </Button>
  );
};

export type SubmitButtonProps = ButtonProps;
