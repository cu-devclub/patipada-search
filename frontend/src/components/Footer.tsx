import { Text,Flex, Button, Modal, ModalBody, ModalCloseButton, ModalContent, ModalFooter, ModalHeader, ModalOverlay, useDisclosure} from "@chakra-ui/react";
import { useState } from "react";
function Footer() {
  const [modalBody, setModalBody] = useState("")
  const onOpenModal = (state : string) => {
    if (state=="develop") {
      setModalBody("ขณะนี้ระบบกำลังอยู่ระหว่างการพัฒนาอาจมีข้อผิดพลาดเกี่ยวกับการสะกดคำหรือปัญหาอื่น ๆ ที่เกิดขึ้นได้")
    } else {
      setModalBody("จัดทำโดย นายวาริส หลักทอง 6334460723 ภาควิชาคณิตศาสตร์และวิทยาการคอมพิวเตอร์ คณะวิทยาศาสตร์ จุฬาลงกรณ์มหาวิทยาลัย")
    }
    onOpen();
  }
  const { isOpen, onOpen, onClose } = useDisclosure();
  return (
    <Flex
      w="100%"
      h="8xs"
      bg="brand_gray.500"
      direction="row"
      justify={{ xl: "flex-end", md: "center", base: "center" }}
      align="center"
      px="4"
      gap={4}
    >
      <Text onClick={() => onOpenModal("about")} cursor={"pointer"}>
        เกี่ยวกับ
      </Text>
      <Text onClick={() => onOpenModal("develop")} cursor={"pointer"}>
        อยู่ระหว่างการพัฒนา
      </Text>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent mx={4}>
          <ModalHeader />
          <ModalCloseButton />
          <ModalBody>{modalBody}</ModalBody>
          <ModalFooter>
            <Button colorScheme="yellow" mr={3} onClick={onClose} color="white">
              รับทราบ
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </Flex>
  );
}

export default Footer;
