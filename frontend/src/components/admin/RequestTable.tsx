import {
  Table,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
  TableCaption,
  TableContainer,
  Text,
  Box,
  Tooltip,
  Button,
} from "@chakra-ui/react";
import { Request } from "../../models/request";
import { useNavigate } from "react-router-dom";
import { REQUEST_STATUS } from "../../constant";
import { convertStatusWord } from "../../functions/request";

interface Props {
  data: Request[];
  syncRequestRecord: (id: string) => void;
}

function RequestTable({ data, syncRequestRecord }: Props) {
  const navigate = useNavigate();

  return (
    <TableContainer>
      <Table variant="simple" layout="fixed">
        <TableCaption>คำขอแก้ไข</TableCaption>
        <Thead>
          <Tr>
            <Th>เลขที่</Th>
            <Th>เวลาเริ่มต้น</Th>
            <Th>เวลาสิ้นสุด</Th>
            <Th>คำถาม</Th>
            <Th>คำตอบ</Th>
            <Th>สร้างเมื่อ</Th>
            <Th>โดย</Th>
            <Th>สถานะ</Th>
            <Th>Action</Th>
          </Tr>
        </Thead>
        <Tbody>
          {data?.map((item) => (
            <Tooltip label="กดเพื่อตรวจสอบ">
              <Tr
                key={item.requestID}
                _hover={{ bg: "gray.100" }}
                onClick={() =>
                  navigate("/admin/edit-request/" + item.requestID)
                }
              >
                <Td>
                  <Text>{item.requestID}</Text>
                </Td>
                <Td>
                  <Text>{item.startTime}</Text>
                </Td>
                <Td>
                  <Text>{item.endTime}</Text>
                </Td>

                <Td>
                  <Box overflow={"hidden"}>
                    <Text>{item.question}</Text>
                  </Box>
                </Td>
                <Td>
                  <Box overflow={"hidden"}>
                    <Text>{item.answer}</Text>
                  </Box>
                </Td>
                <Td>
                  <Text>{item.createdAt}</Text>
                </Td>
                <Td>
                  <Text>{item.by}</Text>
                </Td>
                <Td>
                  <Text>{item.status}</Text>
                </Td>

                <Td>
                  {item.status ==
                    convertStatusWord(REQUEST_STATUS.REVIEWED) && (
                    <Button
                      colorScheme="blue"
                      onClick={(e) => {
                        e.stopPropagation();
                        syncRequestRecord(item.requestID);
                      }}
                    >
                      Sync
                    </Button>
                  )}
                </Td>
              </Tr>
            </Tooltip>
          ))}
        </Tbody>
      </Table>
    </TableContainer>
  );
}

export default RequestTable;
