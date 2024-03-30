import { Text, Flex } from "@chakra-ui/react";
import { useState } from "react";
import { InformModal, BaseModal } from "../../modal";
import { RatingForm } from "../../rating";

function Footer() {
  const [selectInformModal, setSelectInformModal] = useState(false);
  const [modalBody, setModalBody] = useState<React.ReactNode | string>();
  const [openModal, setopenModal] = useState(false);

  const onOpenModal = (state: "develop" | "about" | "recommend") => {
    if (state == "develop") {
      setModalBody(
        "ขณะนี้ระบบกำลังอยู่ระหว่างการพัฒนาอาจมีข้อผิดพลาดเกี่ยวกับการสะกดคำหรือปัญหาอื่น ๆ ที่เกิดขึ้นได้"
      );
    } else if (state == "about") {
      setModalBody(
        "ระบบการค้นหาคำถาม คำตอบ จัดทำโดย นายวาริส หลักทอง 6334460723 ภาควิชาคณิตศาสตร์และวิทยาการคอมพิวเตอร์ คณะวิทยาศาสตร์ จุฬาลงกรณ์มหาวิทยาลัย"
      );
    } else if (state == "recommend") {
      setModalBody(<RatingForm closeModal={() => setopenModal(false)} />);
    }
    setSelectInformModal(state == "recommend" ? false : true);
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
      <Text onClick={() => onOpenModal("recommend")} cursor={"pointer"}>
        ข้อเสนอแนะ
      </Text>
      <Text onClick={() => onOpenModal("about")} cursor={"pointer"}>
        ค้นธรรม ณ ธรรมนาวา (v0.2)
      </Text>
      <Text onClick={() => onOpenModal("develop")} cursor={"pointer"}>
        อยู่ระหว่างการพัฒนา
      </Text>
      {selectInformModal ? (
        <InformModal
          modalBody={modalBody}
          modalButtonText="รับทราบ"
          openModal={openModal}
          closeModal={() => setopenModal(false)}
        />
      ) : (
        <BaseModal
          modalBody={modalBody}
          openModal={openModal}
          closeModal={() => setopenModal(false)}
        /> // Replace this with the component you want to render when selectInformModal is false
      )}
    </Flex>
  );
}

export default Footer;
