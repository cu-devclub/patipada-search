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
} from '@chakra-ui/react'
import { Request } from '../../models/request'
import { useNavigate } from 'react-router-dom'
interface Props {
    data : Request[]      
}

function RequestTable({ data }: Props) {
    const navigate = useNavigate()
    return (
      <TableContainer>
        <Table variant="simple" layout="fixed">
          <TableCaption>คำขอแก้ไข</TableCaption>
          <Thead>
            <Tr>
              <Th width="5%">เลขที่</Th>
              <Th width="10%">เวลาเริ่มต้น</Th>
              <Th width="10%">เวลาสิ้นสุด</Th>
              <Th width="15%">คำถาม</Th>
              <Th width="15%">คำตอบ</Th>
              <Th width="10%">สร้างเมื่อ</Th>
              <Th width="10%">โดย</Th>
              <Th width="10%">สถานะ</Th>
            </Tr>
          </Thead>
          <Tbody >
            {data?.map((item) => (
                <Tooltip label="กดเพื่อตรวจสอบ" >

              <Tr key={item.requestID}  onClick={() => navigate("/admin/edit-request/"+item.requestID)}>
                <Td width="5%" >
                  <Text >{item.requestID}</Text>
                </Td>
                <Td width="10%" >
                  <Text >{item.startTime}</Text>
                </Td>
                <Td width="10%" >
                  <Text >{item.endTime}</Text>
                </Td>

                <Td width="15%" >
                    <Box overflow={"hidden"}>
                        <Text>{item.question}</Text>
                    </Box>
                </Td>
                <Td width="15%">
                  <Box overflow={"hidden"}>
                    <Text >{item.answer}</Text>
                  </Box>
                </Td>
                <Td width="10%">
                  <Text >{item.createdAt}</Text>
                </Td>
                <Td width="10%">
                  <Text >{item.by}</Text>
                </Td>
                <Td width="10%">
                  <Text >{item.status}</Text>
                </Td>
              </Tr>
              </Tooltip>

            ))}
          </Tbody>
        </Table>
      </TableContainer>

    );
  }

export default RequestTable