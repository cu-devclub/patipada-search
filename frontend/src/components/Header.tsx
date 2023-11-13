import { Center,Image } from "@chakra-ui/react";
function Header() {
  return (
    <Center w="100%" h="6xs" border="1px" borderColor="gray.300" shadow="md">
      <Image
        boxSize="7xs"
        objectFit="cover"
        src="Dhammanava.svg"
        alt="Dhammanava"
      />
    </Center>
  );
}

export default Header;
