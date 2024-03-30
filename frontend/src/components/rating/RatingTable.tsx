import {
  TableContainer,
  Table,
  Thead,
  Tr,
  Th,
  Tbody,
  Td,
  Text,
} from "@chakra-ui/react";
import { FullRating } from "../../models/ratings";

interface RatingTablesProps {
  ratings: FullRating[];
}

function RatingTable({ ratings }: RatingTablesProps) {
  return (
    <TableContainer w="full">
      <Table variant="simple">
        <Thead>
          <Tr>
            <Th>rating id</Th>
            <Th>stars</Th>
            <Th>comment</Th>
          </Tr>
        </Thead>
        <Tbody>
          {ratings?.map((rating) => (
            <Tr key={rating.rating_id}>
              <Td w="30%">
                <Text>{rating.rating_id}</Text>
              </Td>
              <Td w="10%">
                <Text>{rating.stars}</Text>
              </Td>
              <Td>
                <Text>{rating.comment}</Text>
              </Td>
            </Tr>
          ))}
        </Tbody>
      </Table>
    </TableContainer>
  );
}

export default RatingTable;
