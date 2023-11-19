import React from "react";

type UserIconProps = {
  // children: React.ReactNode;
  /**
   * bese64
   */
  readonly iconImage: string;
  readonly size: number;
};

/**
 * user icon
 */
export const UserIcon: React.FC<UserIconProps> = (props) => {
  const { iconImage, size } = props;
  return (
    <div style={{ height: size, width: size }}>
      <img alt="icon" className="rounded-full" src={iconImage} />
    </div>
  );
};
