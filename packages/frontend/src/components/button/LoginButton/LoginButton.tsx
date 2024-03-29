import React from "react";

import { SubmitButton } from "@/components/button/SubmitButton";

type LoginButtonProps = {
  readonly onClick?: () => void;
};

/**
 * login button
 */
export const LoginButton: React.FC<LoginButtonProps> = ({ onClick }) => {
  return (
    <SubmitButton className="w-[8em]" onClick={onClick}>
      Login
    </SubmitButton>
  );
};
