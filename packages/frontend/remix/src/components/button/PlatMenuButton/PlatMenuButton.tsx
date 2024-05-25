import { Menu } from "@mantine/core";
import { MdMoreHoriz } from "react-icons/md";

type PlatMenuButtonProps = {
  readonly onClick?: () => void;
};
/**
 *
 */
export const PlatMenuButton: React.FC<PlatMenuButtonProps> = () => {
  return (
    <Menu position="bottom">
      <Menu.Target>
        {/* ここをdivで囲わないとメニューの表示位置がおかしくなる */}
        <div>
          <MdMoreHoriz />
        </div>
      </Menu.Target>
      <Menu.Dropdown>
        <Menu.Item>Delete this plat</Menu.Item>
        <Menu.Item>Mute this user</Menu.Item>
      </Menu.Dropdown>
    </Menu>
  );
};
