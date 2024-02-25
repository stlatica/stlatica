import { UnstyledButton } from "@mantine/core";
import React from "react";
import { MdArrowBack } from "react-icons/md";

type BackButtonProps = {
  // children: React.ReactNode;
};

/**
 *
 */
export const BackButton: React.FC<BackButtonProps> = () => {
  return (
    <UnstyledButton onClick={() => {}}>
      <MdArrowBack size="3em" />
    </UnstyledButton>
  );
};
