import {
  Avatar,
  Menu,
  MenuButton,
  MenuItem,
  MenuList,
  Text,
  Box,
} from "@chakra-ui/react";

import { SignOut } from "../../functions";

import { MessageToast } from "..";
import { useNavigate } from "react-router-dom";

interface UserAvatarProps {
  username: string;
}

function UserAvatar({ username }: UserAvatarProps) {
  const { addToast } = MessageToast();
  const navigate = useNavigate();
  const signOut = () => {
    addToast({
      description: "ออกจากระบบสำเร็จ",
      status: "success",
    });
    SignOut();
  };
  return (
    <Menu>
      <MenuButton
        as={Avatar}
        aria-label="Options"
        icon={<Avatar bg="brand_orange.400" />}
        variant="outline"
        cursor="pointer"
      />
      <MenuList>
        <Box w="full" h="full" pl={2}>
          <Text fontWeight={"semibold"}> สวัสดี {username}</Text>
        </Box>
        <MenuItem onClick={() => navigate("/contributor/pending-request")}>
          ติดตามคำขอแก้ไข
        </MenuItem>
        <MenuItem onClick={() => navigate("/user/change-password")}>
          เปลี่ยนรหัสผ่าน
        </MenuItem>
        <MenuItem onClick={signOut}>ออกจากระบบ</MenuItem>
      </MenuList>
    </Menu>
  );
}

export default UserAvatar;
