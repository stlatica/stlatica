import React from "react";
import { SubmitButton } from "@/components/button/SubmitButton";
import { CancelButton } from "@/components/button/CancelButton";

type FollowButtonProps = {
  isFollow: boolean;
  onClick: () => void;
};

export const FollowButton: React.FC<FollowButtonProps> = ({ isFollow, onClick }) => {
  if (isFollow) {
    return (
      <CancelButton style={{ width: "8em" }} onClick={onClick}>
        フォロー中
      </CancelButton>
    );
  }

  return (
    <SubmitButton style={{ width: "8em" }} onClick={onClick}>
      フォロー
    </SubmitButton>
  );
};
