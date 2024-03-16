import { UnstyledButton } from "@mantine/core";
import React from "react";
import { MdArrowBack } from "react-icons/md";

type BackButtonProps = {
  // children: React.ReactNode;
  readonly onClick?: () => void;
};

/**
 *
 */
export const BackButton: React.FC<BackButtonProps> = ({ onClick }) => {
  return (
    <UnstyledButton onClick={onClick}>
      <MdArrowBack size="3em" />
    </UnstyledButton>
  );
};
