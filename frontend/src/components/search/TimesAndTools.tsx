import { RepeatIcon, EditIcon } from '@chakra-ui/icons';
import { HStack, Tooltip, IconButton,Text } from '@chakra-ui/react';
import { DataItem } from '../../models/qa';

interface TimesAndToolsProps {
    data: DataItem;
    token?: string;
    handleReplay: () => void;
    navigate: (path: string) => void;
}


function TimesAndTools({ data, token, handleReplay, navigate }: TimesAndToolsProps) {
  return (
    <HStack spacing={2} alignItems="center">
      <Text as="b" color="blue">
        เวลาเริ่มต้น {data.startTime} เวลาสิ้นสุด {data.endTime}
      </Text>
      <Tooltip
        hasArrow
        label="กดเพื่อเล่นวิดีโออีกครั้ง"
        bg="gray.300"
        color="black"
        placement="right"
      >
        <IconButton
          aria-label="Play Again"
          icon={<RepeatIcon />}
          onClick={handleReplay}
        />
      </Tooltip>
      {token && (
        <Tooltip
          hasArrow
          label="กดเพื่อเสนอข้อแก้ไข"
          bg="gray.300"
          color="black"
          placement="right"
        >
          <IconButton
            aria-label="Edit"
            icon={<EditIcon />}
            onClick={() => navigate(`/contributor/edit-record/${data.index}`)}
          />
        </Tooltip>
      )}
    </HStack>
  );
}

export default TimesAndTools