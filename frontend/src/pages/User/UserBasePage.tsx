import { ReactNode } from "react";
import {
  Flex,
  VStack,
  Box,
  IconButton,
  Tooltip,
  AlertDialog,
  AlertDialogBody,
  AlertDialogContent,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogOverlay,
  Button,
  useDisclosure,
  AlertDialogCloseButton,
} from "@chakra-ui/react";
import { Logo } from "../../components";
import { ArrowBackIcon } from "@chakra-ui/icons";
import { useNavigate } from "react-router-dom";
import { useRef } from "react";
interface BasePageProps {
  children: ReactNode;
}

export const UserBasePage: React.FC<BasePageProps> = ({ children }) => {
  const { isOpen, onOpen, onClose } = useDisclosure();
  const cancelRef = useRef<HTMLButtonElement | null>(null);
  const navigate = useNavigate();
  return (
    <Flex
      w="100%"
      minH="100svh"
      justify={"flex-start"}
      align={"center"}
      direction={"column"}
      pt={12}
      bg="gray.700"
    >
      <Tooltip label="กลับหน้าค้นหาข้อมูล">
        <Box position="absolute" top={["2", "4"]} left={["2", "4"]}>
          <IconButton
            colorScheme="orange"
            aria-label="back to search page"
            icon={<ArrowBackIcon />}
            onClick={onOpen}
          />
        </Box>
      </Tooltip>

      <Logo size="7xs" />
      <VStack spacing={0} pb={4}>
        {children}
      </VStack>

      <AlertDialog
        motionPreset="slideInBottom"
        leastDestructiveRef={cancelRef}
        onClose={onClose}
        isOpen={isOpen}
      >
        <AlertDialogOverlay />

        <AlertDialogContent>
          <AlertDialogHeader>กลับหน้าค้นหา</AlertDialogHeader>
          <AlertDialogCloseButton />
          <AlertDialogBody>ยืนยันกลับหน้าค้นหาข้อมูล</AlertDialogBody>
          <AlertDialogFooter>
            <Button ref={cancelRef} onClick={onClose}>
              ยกเลิก
            </Button>
            <Button colorScheme="orange" ml={3} onClick={() => navigate("/")}>
              ยืนยัน
            </Button>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </Flex>
  );
};
