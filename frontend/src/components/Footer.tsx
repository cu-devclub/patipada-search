import { Text,Flex, Button, Modal, ModalBody, ModalCloseButton, ModalContent, ModalFooter, ModalHeader, ModalOverlay, useDisclosure} from "@chakra-ui/react";
function Footer() {

  const { isOpen, onOpen, onClose } = useDisclosure();
  return (
    <Flex
      w="100%"
      h="6xs"
      border="1px"
      borderColor="gray.300"
      shadow="md"
      direction="row"
      justify={{ xl: "space-between", md: "center", base: "center" }}
      align="center"
      px="4"
    >
      <Text hideBelow="lg" color="white">
        ธรรมนาวา
      </Text>
      <Text hideBelow="lg">ธรรมนาวา</Text>
      <Text as="u" onClick={onOpen} cursor={"pointer"}>
        อยู่ระหว่างการพัฒนา
      </Text>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent mx={4}>
          <ModalHeader />
          <ModalCloseButton />
          <ModalBody>
            ขณะนี้ระบบกำลังอยู่ระหว่างการพัฒนา
            อาจมีข้อผิดพลาดเกี่ยวกับการสะกดคำหรือปัญหาอื่น ๆ ที่เกิดขึ้นได้
          </ModalBody>

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
