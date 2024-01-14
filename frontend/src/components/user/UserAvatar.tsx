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
import { getCookie } from "typescript-cookie";
import { useLocation } from "react-router-dom";

interface UserAvatarProps {
  username: string;
}

function UserAvatar({ username }: UserAvatarProps) {
  const { addToast } = MessageToast();
  const location = useLocation();

  const navigate = useNavigate();
  const role = getCookie("role");
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
        {role == "admin" || role == "super-admin" ? (
          <MenuItem
            onClick={() =>
              navigate(
                location.pathname.startsWith("/admin")
                  ? "/"
                  : "/admin/dashboard"
              )
            }
          >
            {location.pathname.startsWith("/admin")
              ? "หน้าค้นหา"
              : "หน้าแอดมิน"}
          </MenuItem>
        ) : null}
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
