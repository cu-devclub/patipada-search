import {
  TableContainer,
  Table,
  Thead,
  Tr,
  Th,
  Tbody,
  Td,
  Text,
  Button,
} from "@chakra-ui/react";
import { User } from "../../models/user";
import { getCookie } from "typescript-cookie";
import { CheckRoleDisplayDeleteButton } from "../../functions/user";

interface UserTablesProps {
  users: User[];
  removeUser: (id: string) => void;
}

function UserTables({ users, removeUser }: UserTablesProps) {
  const role = getCookie("role");

  return (
    <TableContainer w="full">
      <Table variant="simple">
        <Thead>
          <Tr>
            <Th>username</Th>
            <Th>email</Th>
            <Th>role</Th>
            <Th>action</Th>
          </Tr>
        </Thead>
        <Tbody>
          {users?.map((user) => (
            <Tr key={user.id}>
              <Td>
                <Text>{user.username}</Text>
              </Td>
              <Td>
                <Text>{user.email}</Text>
              </Td>
              <Td>
                <Text>{user.role}</Text>
              </Td>
              <Td>
                {CheckRoleDisplayDeleteButton(role || "", user.role) && (
                  <Button
                    variant={"cancel"}
                    onClick={() => removeUser(user.id)}
                  >
                    Remove user{" "}
                  </Button>
                )}
              </Td>
            </Tr>
          ))}
        </Tbody>
      </Table>
    </TableContainer>
  );
}

export default UserTables;
