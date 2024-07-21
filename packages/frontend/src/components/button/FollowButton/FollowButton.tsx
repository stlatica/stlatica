import React from "react";

import { CancelButton } from "@/components/button/CancelButton";
import { SubmitButton } from "@/components/button/SubmitButton";

type FollowButtonProps = {
  readonly isFollow: boolean;
  readonly onClick?: () => void;
};

export const FollowButton: React.FC<FollowButtonProps> = ({ isFollow, onClick }) => {
  if (isFollow) {
    return <CancelButton onClick={onClick}>フォロー中</CancelButton>;
  }

  return <SubmitButton onClick={onClick}>フォロー</SubmitButton>;
};
