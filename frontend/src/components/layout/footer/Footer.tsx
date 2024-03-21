import { Text, Flex } from "@chakra-ui/react";
import { useState } from "react";
import { InformModal } from "../../modal";

function Footer() {
  const [modalBody, setModalBody] = useState("");
  const [openModal, setopenModal] = useState(false);
  const onOpenModal = (state: string) => {
    if (state == "develop") {
      setModalBody(
        "ขณะนี้ระบบกำลังอยู่ระหว่างการพัฒนาอาจมีข้อผิดพลาดเกี่ยวกับการสะกดคำหรือปัญหาอื่น ๆ ที่เกิดขึ้นได้"
      );
    } else {
      setModalBody(
        "จัดทำโดย นายวาริส หลักทอง 6334460723 ภาควิชาคณิตศาสตร์และวิทยาการคอมพิวเตอร์ คณะวิทยาศาสตร์ จุฬาลงกรณ์มหาวิทยาลัย"
      );
    }
    setopenModal(true);
  };

  return (
    <Flex
      w="100%"
      h="100%"
      bg="brand_gray.400"
      direction="row"
      justify={{ lg: "space-between", md: "center", base: "center" }}
      align="center"
      px="4"
      gap={4}
    >
      <Text onClick={() => onOpenModal("about")} cursor={"pointer"}>
        เกี่ยวกับ
      </Text>
      <Text onClick={() => onOpenModal("develop")} cursor={"pointer"}>
        ค้นธรรม ณ ธรรมนาวา (v0.2)
      </Text>
      <Text onClick={() => onOpenModal("develop")} cursor={"pointer"}>
        อยู่ระหว่างการพัฒนา
      </Text>
      <InformModal
        modalBody={modalBody}
        modalButtonText="รับทราบ"
        openModal={openModal}
        closeModal={() => setopenModal(false)}
      />
    </Flex>
  );
}

export default Footer;
