import { useEffect, useState } from "react";
import { UserTables } from "../../components/user";
import { AdminBasePage } from "./AdminBasePage";
import { useNavigate } from "react-router-dom";
import { MessageToast } from "../../components";
import {
  ERR_Messages_MAP,
  Role,
  SERVER_ERROR_MESSAGE,
  ToastStatus,
} from "../../constant";
import { AuthorizeAdmin } from "../../functions";
import { RegisterDTO, User } from "../../models/user";
import { ReturnError } from "../../service/error";
import {
  registerService,
  removeUserService,
  getAllUsersService,
} from "../../service/user";
import { Button, Flex, useDisclosure } from "@chakra-ui/react";
import { AddUserModal } from "../../components/modal";
import { AddUserForm } from "../../components/user/forms";

function AdminUserPage() {
  const navigate = useNavigate();
  const { addToast } = MessageToast();
  const { isOpen, onOpen, onClose } = useDisclosure();

  const [users, setUsers] = useState<User[]>([]);

  const [usernameError, setusernameError] = useState(false);
  const [emailError, setemailError] = useState(false);

  useEffect(() => {
    (async () => {
      const isAuthorize = await AuthorizeAdmin(Role.ADMIN);
      if (isAuthorize === false) {
        navigate("/user/login");
      }
    })();
    const getUsers = async () => {
      await getAllUsersService()
        .then((res: User[]) => {
          setUsers(res);
          addToast({
            description: "ดึงข้อมูลสำเร็จ",
            status: ToastStatus.SUCCESS,
          });
        })
        .catch(() => {
          addToast({
            description: "เกิดข้อผิดพลาดขณะทำการดึงข้อมูล",
            status: ToastStatus.ERROR,
          });
        });
    };
    getUsers();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const deleteUser = async (id: string) => {
    await removeUserService(id)
      .then(() => {
        addToast({
          description: "ลบผู้ใช้สำเร็จ",
          status: ToastStatus.SUCCESS,
        });
      })
      .catch((err: ReturnError) => {
        addToast({
          description: err.message,
          status: ToastStatus.ERROR,
        });
        return;
      });

    const newUsers = users.filter((user) => user.id !== id);
    setUsers(newUsers);
  };

  const submit = async (
    username: string,
    email: string,
    password: string,
    role: string
  ) => {
    const registerDTO: RegisterDTO = {
      username: username,
      email: email,
      password: password,
      role: role,
    };

    const newUser: User = {
      username: username,
      email: email,
      role: role,
      id: "",
    };

    await registerService(registerDTO)
      .then((res) => {
        newUser.id = res.user_id;
        addToast({
          description: "register successfully",
          status: ToastStatus.SUCCESS,
        });
      })
      .catch((error: ReturnError) => {
        if (error.status === 400) {
          if (error.message === SERVER_ERROR_MESSAGE.USERNAME_ALREADY_EXISTS) {
            setusernameError(true);
          } else if (
            error.message === SERVER_ERROR_MESSAGE.EMAIL_ALREADY_EXISTS
          ) {
            setemailError(true);
          }
          addToast({
            description: ERR_Messages_MAP[error.message],
            status: ToastStatus.WARNING,
          });
        } else {
          addToast({
            description: error.message,
            status: error.toastStatus,
          });
        }
        return;
      });

    setUsers([...users, newUser]);
    onClose();
  };

  const AddUserFormComponent = () => (
    <AddUserForm
      submit={submit}
      usernameError={usernameError}
      emailError={emailError}
      closeModal={onClose}
    />
  );

  return (
    <AdminBasePage activePage="Users">
      <Flex w="full" direction="column" gap={4}>
        <Flex alignSelf={"flex-end"} p={4}>
          <Button colorScheme="blue" onClick={onOpen}>
            Add user
          </Button>
        </Flex>
        <UserTables users={users} removeUser={deleteUser} />
      </Flex>
      <AddUserModal
        openModal={isOpen}
        closeModal={onClose}
        modalTitle={"ทดสอบ"}
        modalBody={<AddUserFormComponent />}
      />
    </AdminBasePage>
  );
}

export default AdminUserPage;
