import { Flex, Text,Heading } from "@chakra-ui/react";
import { Logo,AnimatedBackground } from "../components";

function NotFoundPage() {
  return (
    <AnimatedBackground>
      <Flex
        w="100%"
        minH="100svh"
        justify={"center"}
        align={"center"}
        direction={"column"}
        pt={12}
      >
        <Logo size="2xs" />
        <Heading color="white" fontSize={["3xl", "6xl"]} fontWeight={"normal"}>
          ไม่พบหน้าที่ต้องการค้นหา
        </Heading>
        <Text fontSize={["xl", "2xl"]} color="white">
          หน้าที่คุณค้นหาไม่มีอยู่ในระบบ
        </Text>
      </Flex>
    </AnimatedBackground>
  );
}

export default NotFoundPage;