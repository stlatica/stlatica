import React from "react";

type UserIconProps = {
  // children: React.ReactNode;
  /**
   * bese64
   */
  readonly iconImage: string;
};

/**
 * user icon
 */
export const UserIcon: React.FC<UserIconProps> = (props) => {
  const { iconImage } = props;
  return (
    <div>
      <img alt="icon" className="rounded-[50%]" src={iconImage} />
    </div>
  );
};
