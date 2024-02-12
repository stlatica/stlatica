import Button, { ButtonProps } from "@mui/material/Button";

export type SubmitButtonProps = ButtonProps;

/**
 *
 */
export const SubmitButton: React.FC<SubmitButtonProps> = (props) => {
  const { children, ...others } = props;
  return (
    <Button type="button" variant="contained" {...others}>
      {children}
    </Button>
  );
};
